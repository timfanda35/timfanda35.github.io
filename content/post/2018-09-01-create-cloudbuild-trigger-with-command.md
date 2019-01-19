---
categories:
  - gcp
  - cloudbuild
comments: true
date: 2018-09-01T16:10:00+08:00
title: 用指令建立 Cloud Build Trigger
url: /2018/09/01/create-cloudbuild-trigger-with-command/
---


![](/images/2018-09-01/create-cloudbuild-trigger-with-command/cover.jpg)

[圖片來源](https://www.pakutaso.com/20151102331web-14.html)
[フリー写真素材ぱくたそ](https://www.pakutaso.com/)

目前還無法使用 `gcloud` 指令直接建立 Cloud Build Trigger，必須要透過 REST API

在編寫 Shell Script 的時候對於認證的部分就比較麻煩，要使用 REST API 就必須要有 Token 來驗證身份

現在我們可以透過 `gcloud auth application-default print-access-token` 輕鬆取得所需的 Token

<!--more-->

首先確定專案有啟用 Cloud Build API
```
$ gcloud services enable cloudbuild.googleapis.com
```

接下來會需要用到 `application-default` 憑證，用來產生 Token

我們可以使用 `gcloud` 指令產生

```
$ gcloud auth application-default login
```

或是透過設定 `GOOGLE_APPLICATION_CREDENTIALS` 環境變數使用 Service Account 的憑證(需要給予 Service Account 角色 Cloud Build Editor 權限)

```
$ export GOOGLE_APPLICATION_CREDENTIALS=/path/credentials.json
```

最後執行以下指令建立 Cloud Build Trigger，請將 `[[YOUR_REPO_NAME]]` 改成 Source Repositories 上的名稱：

```
$ export REPO_NAME=[[YOUR_REPO_NAME]]
$ export PROJECT_ID=$(gcloud config get-value core/project)
$ export OAUTH_TOKEN=$(gcloud auth application-default print-access-token)
$ curl -X POST -H "Authorization: Bearer ${OAUTH_TOKEN}" -H "Content-Type: application/json" \
https://cloudbuild.googleapis.com/v1/projects/${PROJECT_ID}/triggers -d \
'
{
  "description": "Build master for GAE",
  "triggerTemplate": {
    "projectId": "'${PROJECT_ID}'",
    "repoName": "'${REPO_NAME}'",
    "branchName": "master"
  },
  "substitutions": {},
  "filename": "cloudbuild.yaml"
}
'
```

其他參考：

* [REST API Document](https://cloud.google.com/cloud-build/docs/api/reference/rest/v1/projects.triggers/create)