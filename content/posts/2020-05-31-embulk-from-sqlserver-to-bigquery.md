
---
categories:
  - gcp
  - bigquery
  - sqlserver
  - embulk
comments: true
date: 2020-05-31T12:00:00+08:00
title: 使用 Embulk 從 SQL Server 匯入資料至 BigQuery
url: /2020/05/31/embulk-from-sqlserver-to-bigquery/
---

![](/images/2020-05-31/embulk-from-sqlserver-to-bigquery/cover.jpg)

<a style="background-color:black;color:white;text-decoration:none;padding:4px 6px;font-family:-apple-system, BlinkMacSystemFont, &quot;San Francisco&quot;, &quot;Helvetica Neue&quot;, Helvetica, Ubuntu, Roboto, Noto, &quot;Segoe UI&quot;, Arial, sans-serif;font-size:12px;font-weight:bold;line-height:1.2;display:inline-block;border-radius:3px" href="https://unsplash.com/@shadejay?utm_medium=referral&amp;utm_campaign=photographer-credit&amp;utm_content=creditBadge" target="_blank" rel="noopener noreferrer" title="Download free do whatever you want high-resolution photos from Antony Xia"><span style="display:inline-block;padding:2px 3px"><svg xmlns="http://www.w3.org/2000/svg" style="height:12px;width:auto;position:relative;vertical-align:middle;top:-2px;fill:white" viewBox="0 0 32 32"><title>unsplash-logo</title><path d="M10 9V0h12v9H10zm12 5h10v18H0V14h10v9h12v-9z"></path></svg></span><span style="display:inline-block;padding:2px 3px">Antony Xia</span></a>

本文將使用開源軟體 Embulk  SQL Server 匯入資料至 BigQuery。

一般批次匯入 BigQuery 皆會先從 SQL Server 匯出資料表成 CSV 格式後，再透過 ETL 工具或是直接執行 BigQuery Load Job 匯入。在網路上也有商業的 ETL 軟體與服務能夠直接串接 SQL Server 至 BigQuery，省去自己匯出 CSV 的麻煩。

SQL Server 匯出成 BigQuery 可以接受的 CSV 格式，比較麻煩的主要有兩點：

1. 文字欄位的分隔。若欄位值有包含半形逗號(,)或是包含半形逗號(,)與雙引號(")，若沒有正確地處理這兩個符號，在解析 CSV 欄位的時候就會難以判斷分隔欄位導致匯入失敗。
2. NULL 值。SQL Server 匯出 CSV 時，會在欄位值中填入 NULL ，但在匯入 BigQeury 時會被解析成字串，所以必須要先將欄位值為 NULL 的欄位修改為空值再匯入。

<!--more-->

## Embulk 簡介

[Embulk]([https://github.com/embulk/embulk](https://github.com/embulk/embulk)) 是一個使用 JRuby 開發的開源 Bulk Data Loader，支援多種資料來源與目的地。 

> Embulk is a parallel bulk data loader that helps data transfer between various storages, databases, NoSQL and cloud services.

透過安裝不同的 input/output plugin 即可支援不同的資料庫或儲存媒介；設定檔為 YAML 格式，也能做成 liquid 格式的範本，動態改變設定檔中的值而不必重新再寫一份類似內容只為了少數變化。

## Embulk 安裝

本文環境以 Google Compute Engine，作業系統為 Linux Centos7 為例。

Google Compute Engine 上的 Service Account 權限為：

- Storage: 讀寫
- BigQuery: 已啟用

安裝 Open JDK 1.8

```bash
sudo yum update -y
sudo yum install -y java-1.8.0-openjdk-devel
```

安裝 Embulk 最新版於使用者目錄

```bash
curl --create-dirs -o ~/.embulk/bin/embulk -L "https://dl.embulk.org/embulk-latest.jar"
chmod +x ~/.embulk/bin/embulk
echo 'export PATH="$HOME/.embulk/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

測試範例，可以觀察在 `try1` 目錄下的檔案，與 `config.yml` 的內容。

```bash
# 產生測試來源
embulk example ./try1

# 猜測欄位並產生設定檔
embulk guess   ./try1/seed.yml -o config.yml

# 預覽載入作業
embulk preview config.yml

# 執行載入作業
embulk run     config.yml
```

若沒有任何錯誤就可以往下一步邁進。

`./try1` 目錄：

```bash
./try1/
├── csv
│   └── sample_01.csv.gz
└── seed.yml
```

`./try1/seed.yml` 內容，這裡只有 input/output plugin 等基本設定：

```yaml
in:
  type: file
  path_prefix: '/home/USER_NAME/./try1/csv/sample_'
out:
  type: stdout
```

`config.yml`  內容，透過 `guess` 指令，從 `seed.yml` 於原始資料產生內容更加詳盡的設定檔：

```yaml
in:
  type: file
  path_prefix: /home/USER_NAME/./try1/csv/sample_
  decoders:
  - {type: gzip}
  parser:
    charset: UTF-8
    newline: LF
    type: csv
    delimiter: ','
    quote: '"'
    escape: '"'
    null_string: 'NULL'
    trim_if_not_quoted: false
    skip_header_lines: 1
    allow_extra_columns: false
    allow_optional_columns: false
    columns:
    - {name: id, type: long}
    - {name: account, type: long}
    - {name: time, type: timestamp, format: '%Y-%m-%d %H:%M:%S'}
    - {name: purchase, type: timestamp, format: '%Y%m%d'}
    - {name: comment, type: string}
out: {type: stdout}
```

從 `config.yml` 中可以得知，將會從 File 類型的資料來源，匯入至標準輸出，也就是顯示在指令輸出中。

## Embulk Plugin

本文我們將安裝 [SQL Server Input Plugin(embulk-input-jdbc)]([https://github.com/embulk/embulk-input-jdbc](https://github.com/embulk/embulk-input-jdbc)) 與 [BigQuery Output Plugin]([https://github.com/embulk/embulk-output-bigquery](https://github.com/embulk/embulk-output-bigquery))。

```bash
embulk gem install embulk-input-sqlserver
embulk gem install embulk-output-bigquery
```

## 撰寫設定檔

假設要使用的 SQL Server 資料表 Schema 如下：

```sql
CREATE TABLE PurchaseUsers (
    id         INT IDENTITY,
    name       VARCHAR(255) NOT NULL,
    create_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
```

設定檔可以寫成這樣，並命名為 `config.yml.liquid`：

```yaml
in:
  type: sqlserver
  host: "{{ env.IN_SOURCE_HOST }}"
  user: "{{ env.IN_SOURCE_DB_USER }}"
  password: "{{ env.IN_SOURCE_DB_PASS }}"
  database: "{{ env.IN_SOURCE_DB }}"
  table: "PurchaseUsers"
  select: "create_at, id, name"
  column_options:
    create_at: {type: timestamp}
    id: {type: long}
    name: {type: string}
out:
  type: bigquery
  auth_method: application_default
  project: "{{ env.GCP_PROJECT_ID }}"
  dataset: "{{ env.BQ_DATASET }}"
  table:   "PurchaseUsers"
  compression: GZIP
  gcs_bucket: "{{ env.GCP_PROJECT_ID }}-embulk"
  auto_create_gcs_bucket: true
  column_options:
  - {name: create_at, type: DATETIME, format: '%Y-%m-%d %H:%M:%S'}
  - {name: id, type: INTEGER}
  - {name: name, type: STRING}
```

其中 `{{ env.variable_name }}` 格式的字串是 [liquid]([https://shopify.github.io/liquid/](https://shopify.github.io/liquid/)) 的語法，可以代入環境變數。要注意的是，設定檔要使用 liquid 格式範本，檔名就必須以 liquid 結尾，如 `config.yml.liquid`。

`output.gcs_bucket` 欄位是必填的，BigQuery Output Plugin 背後其實也是先將資料變成 CSV 上傳到 Google Cloud Storage Bucket 後再執行 Load Job 載入至 BigQuery。 `output.auto_create_gcs_bucket` 則是設定成自動建立。

`output.auth_method` 的值為 `application_default`，由於本文環境於 Google Compute Engine 上執行，所以 Embulk 會直接使用 Google Compute Engine 預設的 Service Account 權限呼叫 Google Cloud Storage 與 BigQuery API。其他方式可參考 [BigQuery Output Plugin Authentication 章節]([https://github.com/embulk/embulk-output-bigquery#authentication](https://github.com/embulk/embulk-output-bigquery#authentication))。

我們可以很簡單地寫一個 Shell Script 使用以上的設定去執行載入任務，填寫要代入設定檔的機敏資訊：

```bash
#!/bin/bash

export GCP_PROJECT_ID="<PROJECT_ID>"
export BQ_DATASET="<DATASET>"

export IN_SOURCE_HOST="<SQLServer HOST>"
export IN_SOURCE_DB_USER="<LOGIN USER>"
export IN_SOURCE_DB_PASS="<LOGIN PASSWORD>"
export IN_SOURCE_DB="<DATABASE>"

embulk run config.yml.liquid
```

執行 Shell Script：

```bash
chmod u+x run.sh
./run.sh
```
