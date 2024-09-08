---
categories:
  - gcp
keywords:
  - note
  - gcp
  - google cloud
  - cloud logging
  - cloud monitoring
  - regex
comments: true
date: 2024-09-08T08:00:00+08:00
title: "Cloud Logging 欄位正規表示式查詢"
url: /2024/09/08/google-cloud-logging-regexp/
images:
  - /images/2024-09-08/google-cloud-logging-regexp.png
---

## 問題描述

我們想要在 Cloud logging 中找出延遲大於 3 秒的日誌記錄，在日誌中 `jsonPayload.message` 查到的格式如下：

```
2024-09-08T05:05:00.929Z pid=1 tid=nip class=Amz::SpApi::ScheduleCreateReportJob jid=dc516a464c733d9dc9d29f04 elapsed=0.675 INFO: done
```

我們可以確認應用程式中是以純文字的方式紀錄執行時間，單位是秒。

如果我們搜尋 `elapsed=` 只能找到包含該文字的所有日誌，但無法只過濾出大於 3 秒的日誌。

{{< figure src="/images/2024-09-08/20240908001.jpg" alt="search match" >}}

即使使用 Pattern Match，也是無法進行比較。

{{< figure src="/images/2024-09-08/20240908002.jpg" alt="search pattern match" >}}

Cloud Logging 的 Query Language 有提供 function，我們可以使用 function 來從日誌抽取欄位並轉成可以比較的數值。

## REGEXP_EXTRACT function

使用 `REGEXP_EXTRACT` function 來抽取延遲時間。

```
REGEXP_EXTRACT(jsonPayload.message, "elapsed=([0-9.]+)")
```

在第二個參數中，我們需要在正規表示式中用 Match Group，將我們需要的部分用括號包起來，這樣 function 的執行結果就會是括號包起來中間的部分。

但要注意回傳的資料型態是文字，所以是不能跟數字比較大小的。

以下分別是比較 1 秒與 3 秒的結果，可以發現直接用 `REGEXP_EXTRACT` 抽取出來的值並沒有正確地與數字比較大小，而是用字串比較的方式。所以在試圖搜尋大於 3 秒的結果中會漏掉執行時間為 16 秒的資料。

{{< figure src="/images/2024-09-08/20240908003.jpg" alt="REGEXP_EXTRACT function > 1" >}}

{{< figure src="/images/2024-09-08/20240908004.jpg" alt="REGEXP_EXTRACT function > 3" >}}

## Cast function

為了能夠過濾出執行時間大於 3 秒的日誌，我們需要將日誌中的文字轉成數字型態。

```
CAST(REGEXP_EXTRACT(jsonPayload.message, "elapsed=([0-9.]+)"), FLOAT64) > 3
```

我們將 `REGEXP_EXTRACT(jsonPayload.message, "elapsed=([0-9.]+)")` 作為 Cast function 的第一個參數，將其執行結果轉換成 Float 數字型態，這樣就可以跟數字比較大小了。

以下就可以看到我們得到了符合預期的搜尋結果：

{{< figure src="/images/2024-09-08/20240908005.jpg" alt="Cast function" >}}

## 總結

透過 `REGEXP_EXTRACT` function 從日誌中找出符合正規表示式 `elapsed=([0-9.]+)` 的字串，並從正規表示式的 Match Group 括號取得數字格式的字串 `[0-9.]+`，並透過 `CAST` function 將字串轉換數字型態，來與數字進行比較。

## 參考
- [Logging query language#Functions](https://cloud.google.com/logging/docs/view/logging-query-language#functions)
