---
categories:
  - gcp
keywords:
  - note
  - gcp
  - google cloud
  - cloud logging
  - cloud monitoring
comments: true
date: 2024-09-07T08:00:00+08:00
title: "Google Cloud Logging Summary Truncate"
url: /2024/09/07/google-cloud-logging-summary-truncate/
images:
  - /images/2024-09-07/google-cloud-logging-summary-truncate.png
---

## 問題描述

Google Cloud Logging 是 Google Cloud 提供的日誌查詢工具。除了平台基礎建設的日誌之外，我們也能透過 API 或是函式庫將日誌發送到 Cloud Logging 儲存與分析。

Cloud Logging 每一筆資料稱為 Log Entry，是以 JSON 格式儲存。除了日誌內容之外，還有 Google 還會加上額外資訊，例如資源、接收日期等。

因為欄位很多，所以 Log Entry 預設情況是折疊起來，只顯示幾個欄位與主要日誌內容。

在我們需要透過瀏覽日誌中一些欄位去找出 Pattern 時，一一展開 Log Entry 確認欄位是很不現實的。

## Summary fields

為了能夠快速瀏覽多筆 Log Entry 中的日誌欄位，我們可以將關心的欄位設定為 Summary Field。

展開 Log Entry，點擊欄位，就選單中點擊「Add field to summary line」。

{{< figure src="/images/2024-09-07/20240907001.png" alt="add summary field" >}}

或是可以在上方工具列點擊 SUMMARY 旁的「Edit」，彈出「Manage summary fields」畫面。

在「Custom summary fields」欄位中填入想要加入到 Summary 的 Log Entry 欄位。

{{< figure src="/images/2024-09-07/20240907002.jpg" alt="add summary field" >}}

這樣我們所關心的欄位就會在 Log Entry 折疊的情況下以綠色亮底呈現。我們便可以快速瀏覽多筆 Log Entry 來找尋 Pattern 或是異常。

{{< figure src="/images/2024-09-07/20240907005.png" alt="show summary fields" caption="Source: [官方文件](https://cloud.google.com/logging/docs/view/logs-explorer-interface#add_summary_fields)" >}}

## Truncate summary fields

有時候應用程式的日誌內容很長，關心的內容剛好是在日誌內容結尾，無法在有限的螢幕寬度中顯示。

我們可以在「Manage summary fields」畫面中，啟用「Truncate summary fields」功能，並選擇從開頭或是結尾只顯示固定字數呈現。

{{< figure src="/images/2024-09-07/20240907003.png" alt="truncate summary field" >}}

設定之後，Summary Field 欄位就只會顯示限制字數了。這在需要查看長日誌結尾的資訊時十分有用。

{{< figure src="/images/2024-09-07/20240907004.png" alt="result" >}}

## 參考
- [Find patterns in your logs by using summary fields](https://cloud.google.com/logging/docs/view/logs-explorer-interface#add_summary_fields)
