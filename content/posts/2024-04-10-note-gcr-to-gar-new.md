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
date: 2024-04-10T12:00:00+08:00
title: "[筆記] 一行指令搬遷 google container registry 到 google artifact registry"
url: /2024/04/10/note-gcr-to-gar-new/
images:
  - /images/2024-04-10/note-gcr-to-gar-new.png
---

> Container Registry will be replaced by Artifact Registry. Please upgrade your projects to Artifact Registry before March 18, 2025.

{{< figure src="/images/2024-04-10/20240410001.jpg" alt="announce" >}}

Link: [https://cloud.google.com/container-registry/docs/release-notes](https://cloud.google.com/container-registry/docs/release-notes)

官方文件連結：[Automatically migrate from Container Registry to Artifact Registry](https://cloud.google.com/artifact-registry/docs/transition/auto-migrate-gcr-ar)

之前的筆記：[[筆記] 設定 google container registry 轉址到 google artifact registry](/2024/10/29/note-gcr-to-gar/)

本文是紀錄使用 Google Cloud 新指令，使得設定更為便捷。

建議操作的帳號擁有 Project Owner Role。

## 執行指令進行搬遷

```bash
gcloud artifacts docker upgrade migrate \
    --projects=PROJECTS
```

其中 `PROJECTS` 可以是一個 Project ID 或是以逗號分隔的多組 Project ID。例如:

```bash
gcloud artifacts docker upgrade migrate \
    --projects=my-project-1,my-project-2,my-project-3
```

該指令會執行以下步驟：
- 在 Artifact Registry 中為對應區域中列出的每個 gcr.io 專案建立 gcr.io repositories。
- 為每個 repositories 建議 IAM Policy，並根據使用者選項套用該 Policy 或忽略。
- 將所有流量從 gcr.io 重新導向至 Artifact Registry。 Artifact Registry 透過 request-time copying 從 Container Registry 複製缺少的 Image 來暫時提供服務，直到所有容器映像都複製到 Artifact Registry。
- 將 gcr.io buckets 中儲存的所有 Container Image 複製到 Artifact Registry 上託管的新建立的 gcr.io repositories。
- 停用 request-time copying。ArtifactRegistry 上託管的 gcr.io repositories 不再依賴 Container Registry。

由於步驟中會複製 gcr.io repositories 中的 Container Image，所以如果有許多 Container Image 就會花上不少時間。可以加上參數 `--recent-images` 或 `--last-uploaded-versions` 來決定再複製的時候選擇的 Container Image：
- `--recent-images=NUM_DAY`: 只複製近期有 Pull 或 Push 的 Container Image，NUM_DAY 的範圍是 30 ~ 90。
- `--last-uploaded-versions=N`: 近期上傳的 Container Image，N 的範圍是整數。

要注意 `--recent-images` 與 `--last-uploaded-versions` 不能同時使用。

```bash
gcloud artifacts docker upgrade migrate \
    --recent-images=60 \
    --projects=my-project
```

或是

```bash
gcloud artifacts docker upgrade migrate \
    --last-uploaded-versions=5 \
    --projects=my-project
```

這個指令會以多執行緒來複製 Container Registry 的 Container Image，可以透過參數 `--max-threads` 調整，預設值為 `8`。

## 參考資料
- [Transition from Container Registry](https://cloud.google.com/artifact-registry/docs/transition/transition-from-gcr)
- [gcloud artifacts docker upgrade](https://cloud.google.com/sdk/gcloud/reference/artifacts/docker/upgrade)
