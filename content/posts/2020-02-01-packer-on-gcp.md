---
categories:
  - gcp
  - gce
  - hashicorp
  - packer
comments: true
date: 2020-02-01T12:00:00+08:00
title: Packer on GCP
url: /2020/02/01/packer-on-gcp/
---

![](/images/2020-02-01/packer-on-gcp/cover.jpg)

<a style="background-color:black;color:white;text-decoration:none;padding:4px 6px;font-family:-apple-system, BlinkMacSystemFont, &quot;San Francisco&quot;, &quot;Helvetica Neue&quot;, Helvetica, Ubuntu, Roboto, Noto, &quot;Segoe UI&quot;, Arial, sans-serif;font-size:12px;font-weight:bold;line-height:1.2;display:inline-block;border-radius:3px" href="https://unsplash.com/@shadejay?utm_medium=referral&amp;utm_campaign=photographer-credit&amp;utm_content=creditBadge" target="_blank" rel="noopener noreferrer" title="Download free do whatever you want high-resolution photos from Antony Xia"><span style="display:inline-block;padding:2px 3px"><svg xmlns="http://www.w3.org/2000/svg" style="height:12px;width:auto;position:relative;vertical-align:middle;top:-2px;fill:white" viewBox="0 0 32 32"><title>unsplash-logo</title><path d="M10 9V0h12v9H10zm12 5h10v18H0V14h10v9h12v-9z"></path></svg></span><span style="display:inline-block;padding:2px 3px">Antony Xia</span></a>

想要將在 GCP 建置 GCE VM Image 的過程自動化，又不想寫一堆針對 GCP Service  操作的 Script。可以使用 Hashicorp Packer 來協助一鍵建置。

<!--more-->

Packer 將會自動化以下步驟：

1. 建立 VM
2. 執行操作
3. 保留硬碟，刪除 VM
4. 用硬碟建立 Image
5. 刪除硬碟

***本文假設使用者已安裝並熟悉 Google Cloud SDK(gcloud)***

本文將以建置一個 Nginx 的 GCE Image 為例。

首先先到 Hashicorp  下載對應平台的 Packer 執行檔：[https://www.packer.io/](https://www.packer.io/)

我使用的平台是 macOS。

給予執行權限並放置到常用的路徑底下，以 Linux 來說如： `/usr/local/bin/`

建立專案目錄

    mkdir gce-image-nginx-demo
    cd gce-image-nginx-demo

## 設定

編輯 `template.json`

    {
      "variables": {
        "project_id": "{{env `GCP_PROJECT_ID`}}",
        "region": "asia-east1",
        "zone": "asia-east1-b"
      },
      "builders": [
        {
          "type": "googlecompute",
          "project_id": "{{user `project_id`}}",
          "image_name": "nginx-demo-{{timestamp}}",
          "source_image": "centos-7-v20191210",
          "ssh_username": "packer",
          "zone": "{{user `zone`}}"
        }
      ],
      "provisioners": []
    }

`template.json` 分為三個部分：

1. variables：定義可重複利用的變數
2. builders：要建置 Image 的平台設定，Packer 支援：[https://www.packer.io/docs/builders/index.html](https://www.packer.io/docs/builders/index.html)
3. provisioners：定義如何設定 VM 的方法

以下表示取得環境變數 `GCP_PROJECT_ID` 的值

    "{{env `GCP_PROJECT_ID`}}"

以下表示取得在 `variables` 區塊中定義的變數的值

    "{{user `project_id`}}"
    "{{user `zone`}}"

`timestamp` 是 Packer 預設的 Function，可以參考 [https://www.packer.io/docs/templates/engine.html#functions](https://www.packer.io/docs/templates/engine.html#functions)

    "prometheus-{{timestamp}}"

在 `variables` 區塊中，我定義了 GCP 常用的三個變數：

1. project id
2. region
3. zone

在 `builders` 區塊中，建置 GCP Linux 的 VM 時，需要設定如下：

    "builders": [
      {
        "type": "googlecompute",
        "project_id": "{{user `project_id`}}",
        "image_name": "prometheus-{{timestamp}}",
        "source_image": "centos-7-v20191210",
        "ssh_username": "packer",
        "zone": "{{user `zone`}}"
      }
    ]

- `type`: 固定為 `googlecompute`
- `project_id`: 要建置 GCE VM Image 的 GCP 專案 ID
- `image_name`: 要建置 GCE VM Image 的名稱，為避免名稱衝突，我通常會加上 timestamp
- `source_image`: 要建置 GCE VM Image 的原始 GCE VM Image，將以此 Image 作為基底進行建置，名稱可以從 GCE Image 的頁面查看
- `ssh_username`: 讓 Packer 遠端登入操作的使用者名稱
- `zone`: GCE VM 是 zone 層級的資源，所以需要指定 VM 要建立在哪個 zone

預設 Packer 會使用 Machine type 為 `n1-standard-1` 的 VM 進行建置，可以參考文件調整：

[https://www.packer.io/docs/builders/googlecompute.html#optional-](https://www.packer.io/docs/builders/googlecompute.html#optional-)

## 建置

因為在 `template.json` 中有使用到環境變數，所以我們可以先宣告：

    export GCP_PROJECT_ID="$(gcloud config get-value core/project)"

然後執行以下指令檢查格式

    packer validate template.json

如果沒問題，會看到以下輸出結果：

    Template validated successfully.

在正式使用時，建議使用 Service Account 執行 Packer。本文為了方便，改採用以下指令讓 Packer 取得授權操作：

    gcloud auth application-default login

這個指令將產生一份 `$HOME/.config/gcloud/application_default_credentials.json`，不需要的時候就可以刪除。

使用 Service Account 請參考：

[https://www.packer.io/docs/builders/googlecompute.html#precedence-of-authentication-methods](https://www.packer.io/docs/builders/googlecompute.html#precedence-of-authentication-methods)

執行以下指令進行建置：

    packer build template.json

輸出結果如下：

    googlecompute: output will be in this color.
    
    ==> googlecompute: Checking image does not exist...
    ==> googlecompute: Creating temporary SSH key for instance...
    ==> googlecompute: Using image: centos-7-v20191210
    ==> googlecompute: Creating instance...
        googlecompute: Loading zone: asia-east1-b
        googlecompute: Loading machine type: n1-standard-1
        googlecompute: Requesting instance creation...
        googlecompute: Waiting for creation operation to complete...
        googlecompute: Instance has been created!
    ==> googlecompute: Waiting for the instance to become running...
        googlecompute: IP: xxx.xxx.xxx.xxx
    ==> googlecompute: Using ssh communicator to connect: xxx.xxx.xxx.xxx
    ==> googlecompute: Waiting for SSH to become available...
    ==> googlecompute: Connected to SSH!
    ==> googlecompute: Deleting instance...
        googlecompute: Instance has been deleted!
    ==> googlecompute: Creating image...
    ==> googlecompute: Deleting disk...
        googlecompute: Disk has been deleted!
    Build 'googlecompute' finished.
    
    ==> Builds finished. The artifacts of successful builds are:
    --> googlecompute: A disk image was created: nginx-demo-1580527798

接著就能在 GCE Image 頁面中找到剛剛建置的 Image。

為了少打字與方便重複利用，我會用 `Makefile` 來執行建置。編輯 `Makefile`

    build:
    	GCP_PROJECT_ID=$(shell gcloud config get-value core/project) \
    	packer build template.json

之後只要執行 `make` 就好。

但我們並沒有安裝 Nginx，所以只是一個只有基底的 Image。

## 補完設定

編輯 `index.html`

    <h1>Hello, Packer!</h1>

編輯 `template.json`

    {
      "variables": {
        "project_id": "{{env `GCP_PROJECT_ID`}}",
        "region": "asia-east1",
        "zone": "asia-east1-b"
      },
      "builders": [
        {
          "type": "googlecompute",
          "project_id": "{{user `project_id`}}",
          "image_name": "nginx-demo-{{timestamp}}",
          "source_image": "centos-7-v20191210",
          "ssh_username": "packer",
          "zone": "{{user `zone`}}"
        }
      ],
      "provisioners": [
        {
          "type": "shell",
          "inline": ["sudo yum install -y nginx && sudo systemctl enable nginx"]
        },
        {
          "type": "file",
          "source": "./index.html",
          "destination": "/tmp/index.html"
        },
        {
          "type": "shell",
          "inline": ["sudo mv /tmp/index.html /usr/share/nginx/html/index.html"]
        }
      ]
    }

在 `provisioners` 區塊執行 3 個步驟：

1. 安裝 Nginx
2. 上傳 `index.html`
3. 將 `index.html` 移至指定目錄下(由於權限關係，所以才分兩步驟做)

再次執行驗證與建置

    packer validate template.json
    packer build template.json

用建置好的 Image 建立 VM，用瀏覽器訪問就能看到 "Hello, Packer!" 了。

Reference: [https://github.com/timfanda35/packer-on-gcp-demo/tree/master/nginx](https://github.com/timfanda35/packer-on-gcp-demo/tree/master/nginx)
