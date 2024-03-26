---
categories:
  - gcp
keywords:
  - note
  - gcp
  - google cloud
  - private google access
  - vpc
comments: true
date: 2024-03-26T00:00:00+08:00
title: "Private Google Access 的 IPv6 Range"
url: /2024/03/26/private-google-access-ipv6/
images:
  - /images/2024-03-26/private-google-access-ipv6.png
---

## 問題

在 Google Cloud 的專案中，沒有外部 IP 的 VM 使用設定 IP Address Restriction 的 API Key 呼叫 Google Maps API。出現了錯誤訊息：

XML 版本：

```xml
<?xml version="1.0" encoding="UTF-8"?>
<GeocodeResponse>
  <status>REQUEST_DENIED</status>
  <error_message>This IP, site or mobile application is not authorized to use this API key. Request received from IP address fda3:e722:ac3:10:c0:58ab:a8c:6, with empty referer</error_message>
</GeocodeResponse>
```

JSON 版本：

```json
{
   "error_message" : "This IP, site or mobile application is not authorized to use this API key. Request received from IP address fda3:e722:ac3:10:c0:58ab:a8c:6, with empty referer",
   "results" : [],
   "status" : "REQUEST_DENIED"
}
```

透過 curl 查看連接時的 Destination IP 是 IPv4。

{{< figure src="/images/2024-03-26/202403260007.jpg" alt="curl result" >}}

在 2020 年就有人發現相關的問題 [Restrict Google Maps API key for us in Private Google Access environment][Restrict Google Maps API key for us in Private Google Access environment]，但沒有很好的解釋。

## 環境分析

1. 這是 Google Cloud 新專案，可以當做都是預設設定。
2. VPC Network 與 Subnet 都是預設建立的 `default`。在預設情況下並不會啟用 [dual-stack][[IPv4/IPv6 dual-stack networking]]，所以在此 VPC Network 中的 VM 並不會被分配到 IPv6 的 IP。
3. Subnet 啟用了 Private Google Access，
4. 在此 Subnet 中的 VM 沒有給予 External IP。
5. 建立了一個 API Key 用於呼叫 Google Maps API，並且設定了 IP Address Restriction，限制只有該 Subnet 可以存取 Google Maps API。
6. 並沒有建立 Cloud NAT。所以 VM 一定會透過 Private Google Access 的功能去存取 Google Maps API。

{{< figure src="/images/2024-03-26/202403260003.jpg" alt="Initial API Key Setting" >}}

問題中的端點可以很簡單地分為：
1. VM
2. Private Google Access
3. Google Maps API

在 VM 上都沒有與 IPv6 相關的設定；Google Maps API 也沒有可以修改的地方，那麼看起來問題就是在 Private Google Access 上。

{{< figure src="/images/2024-03-26/202403260008.jpg" alt="Initial API Key Setting" >}}

## Private Google Access

由於 Google Maps API 為公開端點，當透過公開 DNS 解析 Google Maps API 的網域時，會得到外部 IP，這表示如果 Client 端沒有連上 Internet 能力的話，便無法存取 Google Maps API。

既然 Google Cloud VM 跟 Google Maps API 都在 Google 的網路中，有沒有辦法不繞過 Internet，讓只有 Google Cloud VM 可以使用內部 IP 存取 Google Maps API？

[Private Google Access][Private Google Access] 便是為了解決該問題的功能，我們可以在 VPC Network 中的 Subnet 去啟用。

{{< figure src="/images/2024-03-26/202403260002.jpg" alt="Enable Private Google Access" >}}

啟用後 Google Cloud VM 可以就可以使用內部 IP 去存取 Google Maps API，而不用再透過 Internet。

{{< figure src="/images/2024-03-26/202403260001.jpg" alt="Private Google Access Architecture" >}}

那麼 Private Google Access 是如何與只有外部 IP 的 Google Maps API 溝通的呢？

理論上 IPv4 與 IPv6 是不互通的，在 VM 呼叫 Google Maps API 的過程中應該需要做轉換。而且 IP 是有限的，Google Cloud 所有客戶的專案若直接用內部 IP 去存取 Google Maps API 勢必會遇到網段衝突。所以我們可以假設 Private Google Access 會替我們的 Subnet 做 SNAT，類似於 Cloud NAT 的功能將封包中的來源 IP 替換掉。

## 確認使用 Private Google Access 的請求來源 IP

我們可以在 VM 上透過以下指令取得 Google Maps API 看到的請求來源 IP：

```bash
curl https://toolbox.googleapps.com/apps/browserinfo/info/
```

輸出結果會像：

```json
{
 "reportUnixTime":1711447452.2128165,
 "charset":"",
 "userAgent": "curl/7.68.0,gzip(gfe)",
 "keepAlive": "Keep-Alive",
 "remoteAddr": "fda3:e722:ac3:10:c0:58ab:a8c:6",
 "reportTime": "Tue, 26 Mar 2024 10:04:12 +0000",
 "language": "",
 "encoding": "gzip(gfe)",
 "xForwardedFor": "",
 "httpVia": "",
 "dnt": ""
}
```

其中 `remoteAddr` 欄位便是 Google Maps API 所看到的請求來源 IP。

我們也可以發現這與錯誤訊息中提到 IP 是相同的。

```bash
curl https://maps.googleapis.com/maps/api/geocode/json?place_id=ChIJeRpOeF67j4AR9ydy_PIzPuM&key=<YOUR_API_KEY>
```

```json
{
   "error_message" : "This IP, site or mobile application is not authorized to use this API key. Request received from IP address fda3:e722:ac3:10:c0:58ab:a8c:6, with empty referer",
   "results" : [],
   "status" : "REQUEST_DENIED"
}
```

透過以上方式可以確認，使用 Private Google Access 會替換掉請求封包中的來源 IP 並且是 IPv6。所以當我們要設定 API Key 的 IP Address Restriction 時，需要設定的是 Private Google Access 的 IPv6 Range，而不是 Subnet 的 IPv4 Range。

## 取得 Private Google Access 的 IPv6 Range

那麼 Private Google Access 的 IPv6 Range 是什麼呢？

從網路上查看其他人的錯誤訊息交叉比對下發現是 `fda3:e722:ac3:10:0:0:0:0/64`，但這個範圍也太大了，明顯超出原本 Subnet `10.140.0.0/20` 太多。

在詢問 Google Support 後得知，Private Google Access 的 IPv6 Range 是有一個轉換規則的：

1. 前 64 bits 是固定的 Prefix
2. 接下來的 32 bits 是 VPC 的 GUID 轉換而來的
3. 最後 32 bits 是 Subnet 的 IP 轉換而來的

{{< figure src="/images/2024-03-26/202403260006.jpg" alt="Private Google Access IPv6 format" >}}

我們可以先透過 `https://toolbox.googleapps.com/apps/browserinfo/info/` 取得前 96 bits 的部分：`fda3:e722:ac3:10:c0:58ab`

後面 32 bits 我們再從 IPv4 10 進位換算 IPv6 16 進位。Subnet `10.140.0.0` 就變成 `a8c:0`。

遮罩的部分則是 `128 - (32 - 20)` = `116`。

所以最後會 Subnet `10.140.0.0/20` 對應到 Private Google Access 的 IPv6 Range 為: `fda3:e722:ac3:10:c0:58ab:a8c:0/116`。

我們修改 API Key 的 IP Address Restriction：

{{< figure src="/images/2024-03-26/202403260004.jpg" alt="Update API Key Setting" >}}

再做測試，並得到成功的回應：

{{< figure src="/images/2024-03-26/202403260005.jpg" alt="Success Response" >}}

## 總結

Private Google Access 會做 SNAT 的部分並未在文件上特別說明，花上了不少時間做猜想。最後基於排除所有可能選項才確定是 Private Google Access 會做 SNAT。最後也還好有向 Google Support 確認到這是有一定的轉換規則。希望這部分能夠寫在公開說明文件上面。

Subnet 停用 Private Google Access，並使用 Cloud NAT 固定來源 IP 也是一種解決方法；但有些服務例如 GKE Private Cluster 會強制啟用 Private Google Access 就不適用了。

## 參考資料
- [Restrict Google Maps API key for us in Private Google Access environment][Restrict Google Maps API key for us in Private Google Access environment]
- [IPv4/IPv6 dual-stack networking][IPv4/IPv6 dual-stack networking]
- [Private Google Access][Private Google Access]

<!-- Links -->
[Restrict Google Maps API key for us in Private Google Access environment]: https://stackoverflow.com/questions/63921749/restrict-google-maps-api-key-for-us-in-private-google-access-environment
[IPv4/IPv6 dual-stack networking]: https://cloud.google.com/anthos/clusters/docs/bare-metal/latest/how-to/dual-stack-networking
[Private Google Access]: https://cloud.google.com/vpc/docs/private-google-access
