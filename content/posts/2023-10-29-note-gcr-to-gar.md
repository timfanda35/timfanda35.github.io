---
categories:
  - gcp
keywords:
  - note
  - gcp
  - google cloud
  - google container registry
  - gcr
  - google artifact registry
  - gar
comments: true
date: 2023-10-29T12:00:00+08:00
title: "[筆記] 設定 google container registry 到 google artifact registry"
url: /2023/10/29/note-gcr-to-gar/
---

**" Container Registry will be phased out, starting May 15, 2024**

Link: [https://cloud.google.com/container-registry/docs/release-notes#May_15_2023](https://cloud.google.com/container-registry/docs/release-notes#May_15_2023)

Google Cloud 即將要停止 google container registry 服務而改用 google artifact registry。但兩者的網域並不相同，例如：
- google container registry: `gcr.io`
- google artifact registry: `us-docker.pkg.dev`

現在可以透過轉址設定將現有 `grc.io` 的請求從 google container registry 轉導至 google artifact registry。
減少搜尋取代現有設定的工作量。本文是操作步驟的筆記，並假設操作者為 `Project Owner`。

官方文件連結：[https://cloud.google.com/artifact-registry/docs/transition/setup-gcr-repo](https://cloud.google.com/artifact-registry/docs/transition/setup-gcr-repo)

## 設定環境變數

我們可以於 Google Cloud Console 頁面，切換至欲設定的專案，並啟動 Cloud Shell

執行指令設定環境變數，方便後續使用

```bash
PROJECT_ID=$(gcloud config get-value core/project)
PROJECT_NUMBER=$(gcloud projects describe $PROJECT_ID --format="value(projectNumber)")
```

## 確認啟用需要的 Cloud API

```bash
gcloud services enable \
    cloudresourcemanager.googleapis.com \
    artifactregistry.googleapis.com
```

## 建立 Repository

因為我是使用 `gcr.io`，所以 location 要選擇 `us`，如果是使用其他 hostname，則需要對照表格：

| Container Registry hostname	| Artifact Registry repository name |
| --- | --- |
| gcr.io | 	gcr.io |
| asia.gcr.io | 	asia.gcr.io |
| eu.gcr.io | 	eu.gcr.io |
| us.gcr.io | 	us.gcr.io |


```bash
gcloud artifacts repositories create gcr.io \
    --repository-format=docker \
    --location=us
```

建立之後於 Cloud Console 上，於 Artifact Repository 新增權限給原本使用 container registry 的帳號
- 對應原來 Role：Storage Admin
- 新增 Role：Artifact Repository Administrator

![Artifact Repository permission](/images/2023-10-29/20231029001.jpg)

## 搬移 Container Images

在啟用轉址前，我們需要先搬移 Container Images 到 Artifacts Repository。

### 設定權限

新增權限給 Artifact Repository Service Account。

```bash
gcloud projects add-iam-policy-binding $PROJECT_ID \
    --member="serviceAccount:service-$PROJECT_NUMBER@gcp-sa-artifactregistry.iam.gserviceaccount.com" \
    --role='roles/storage.objectViewer'
```

### 搬移

下載 `gcrane` https://github.com/google/go-containerregistry/releases/tag/v0.16.1


執行指令搬移全部 Container Images(不同 location 需要個別執行指令)。執行之前建議先清理不再需要的 Container Images，不然可能會需要執行很久；或是指定要搬移的 Repository。

```bash
./gcrane cp -r gcr.io/$PROJECT_ID us-docker.pkg.dev/$PROJECT_ID/gcr.io
```

### 移除權限

當搬移完成後，就可以把 Artifact Repository Service Account 的權限移除。

```bash
gcloud projects remove-iam-policy-binding $PROJECT_ID \
    --member="serviceAccount:service-$PROJECT_NUMBER@gcp-sa-artifactregistry.iam.gserviceaccount.com" \
    --role='roles/storage.objectViewer'
```

## 驗證設定

使用 `dry-run` 在啟用轉址前來檢查設定。

```bash
gcloud beta artifacts settings enable-upgrade-redirection \
    --project=$PROJECT_ID --dry-run
```

## 啟用轉址

```bash
gcloud beta artifacts settings enable-upgrade-redirection \
    --project=$PROJECT_ID
```

## 停用轉址

```bash
gcloud beta artifacts settings disable-upgrade-redirection \
    --project=$PROJECT_ID
```

## 其他

原先的 GitLab 上的 docker config 可能會有權限錯誤，先在其他機器上執行以下指令，再把 `~/.docker/config.json` 的內容放到 GitLab 的 CI 變數上。

```bash
cat key.json | docker login -u _json_key --password-stdin https://gcr.io
```

## 參考資料
- [Authentication settings in the Docker configuration file](https://cloud.google.com/container-registry/docs/advanced-authentication#docker-config)
- [Using password authentication](https://cloud.google.com/artifact-registry/docs/transition/changes-docker#password)
