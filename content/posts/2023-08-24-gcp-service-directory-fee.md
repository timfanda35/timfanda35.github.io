---
categories:
  - google cloud
keywords:
  - google cloud platform
  - google cloud
  - gcp
  - billing
  - networking
  - service directory
comments: true
date: 2023-08-24T12:00:00+08:00
title: "GCP Networking Service Directory Registered Resource 費用來源"
url: /2023/08/24/gcp-service-directory-fee/
images:
  - /images/2023-08-24/gcp-service-directory-fee.png
---

## 費用 SKU

在檢查 Google Cloud Platform 帳單時，發現一筆奇怪的費用：

```
Networking -> Networking Service Directory Registered Resource
```

![](/images/2023-08-24/001.png)

## 確認資源

這個專案是我用來測試 GKE 功能的測試專案，有可能是在做一些測試的時候，Google 自動產生未預期的資源，測試結束時沒有清除乾淨。

從 Cloud Console 查看「Network Services -> Service Directory」，但該專案並未啟用 Service Directory API，理論上應該是沒有資源才對的，所以我往其他的方向查。但後來發現我是錯的。
![](/images/2023-08-24/002.png)

從「IAM & Admin -> Asset Inventory」查看目前還有資源的 Resource Type，發現 `servicedirectory.Namespaces` 底下的確有資源存在。

![](/images/2023-08-24/003.png)

`goog-psc-default` 是 Google 自動產生的 Namespace，依照[官方文件說明](https://cloud.google.com/vpc/docs/configure-private-service-connect-services#create-endpoint)：

" When you create an endpoint, it is automatically registered with Service Directory, using a namespace that you choose, or the default namespace, `goog-psc-default`. "

而這也與網路上搜尋到的文章[GKEのクラスタを削除してもネットワークサービスで課金が止まらない場合の対処法](https://zenn.dev/ikechan0829/articles/gcp_stop_billing_deleted_gke_cluster)中的敘述相同。

點擊名稱確認相關資訊。
![](/images/2023-08-24/004.png)

確定 Service Directory 有資源存在後，就回到「Network Services -> Service Directory」，啟用 Service Directory API。啟用後重新整理頁面應該就可以看到該資源：

![](/images/2023-08-24/005.png)

## 清除資源

由於該專案並不再需要使用該資源，所以我決定直接刪除它。

一開始我想從 Cloud Console 刪除，卻跳出 `Only standard namespaces can be deleted` 的錯誤訊息：

![](/images/2023-08-24/006.png)

不過我們可以開啟 Cloud Shell 使用 `gcloud` 刪除。請記得將指令中的 `REGION` 改為實際的位置：

```bash
REGION=asia-east1
gcloud service-directory namespaces delete goog-psc-default --location=$REGION
```

成功畫面：

![](/images/2023-08-24/007.png)

## 心得

透過「IAM & Admin -> Asset Inventory」可以尋找專案中已建立的資源，對於尋找未預期、Google 自動建立的資源很有幫助。

## 參考資料

- [Is billing in Google Cloud strange?](https://www.reddit.com/r/googlecloud/comments/14oejrc/is_billing_in_google_cloud_strange/)
- [GKEのクラスタを削除してもネットワークサービスで課金が止まらない場合の対処法](https://zenn.dev/ikechan0829/articles/gcp_stop_billing_deleted_gke_cluster)
- [Access published services through endpoints#Create an endpoint](https://cloud.google.com/vpc/docs/configure-private-service-connect-services#create-endpoint)