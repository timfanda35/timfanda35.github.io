---
categories:
  - gcp
  - cloudbuild
  - container
comments: true
date: 2018-09-01T14:40:00+08:00
title: 用快取加速 Cloud Build 建置容器
url: /2018/09/01/cloudbuild-with-cache/
---


![](http://i.imgur.com/EsiorQx.jpg)

[圖片來源](https://www.pakutaso.com/20160628181pc-11.html)
[フリー写真素材ぱくたそ](https://www.pakutaso.com/)

參考至 Google Cloud Platform Cloud Build 官方文件：[Speeding up your Builds](https://cloud.google.com/cloud-build/docs/speeding-up-builds)

可以使用之前建置過的 Docker Image 來當作快取，來加快建置 Docker Image 的速度

這方法需要使用 `cloudbuild.yaml`

使用方式：

```shell
gcloud builds submit --config cloudbuild.yaml .
```

或是在設定 cloud build trigger 時指定使用 `cloudbuild.yaml`

<!--more-->

官方文件上的設定如下，不過這是在 Container Registry 已經有 `gcr.io/$PROJECT_ID/[IMAGE_NAME]:latest` 的前提下才能順利完成建置步驟。

若是從一開始就使用以下 `cloudbuild.yaml` 的話，在 Step 0 就會因為無法 pull `gcr.io/$PROJECT_ID/[IMAGE_NAME]:latest` 而停止建構

```yaml
steps:
- name: 'gcr.io/cloud-builders/docker'
  args: ['pull', 'gcr.io/$PROJECT_ID/[IMAGE_NAME]:latest']
- name: 'gcr.io/cloud-builders/docker'
  args: [
            'build',
            '-t', 'gcr.io/$PROJECT_ID/[IMAGE_NAME]:latest',
            '--cache-from', 'gcr.io/$PROJECT_ID/[IMAGE_NAME]:latest',
            '.'
        ]
images: ['gcr.io/$PROJECT_ID/[IMAGE_NAME]:latest']
```

因為不想要每次都要先自己手動建置一版推上 Container Registry 再修改 `cloudbuild.yaml`

所以將 Step 0 改成不管有沒有存在 `gcr.io/$PROJECT_ID/[IMAGE_NAME]:latest`，都讓 Step 0 通過建置

透過 `|| true` 可以讓該行指令的回傳值為 `0`，從而表示通過建置

這樣就可以在直接從專案初期就使用同一份 `cloudbuild.yaml` 了

```yaml
steps:
- name: 'gcr.io/cloud-builders/docker'
  entrypoint: 'bash'
  args: [
    '-c',
    '/usr/bin/docker pull gcr.io/${PROJECT_ID}/${REPO_NAME}:latest || true'
  ]
- name: 'gcr.io/cloud-builders/docker'
  args: [
    'build',
    '-t', 'gcr.io/${PROJECT_ID}/${REPO_NAME}:latest',
    '-t', 'gcr.io/${PROJECT_ID}/${REPO_NAME}:${SHORT_SHA}',
    '--cache-from', 'gcr.io/${PROJECT_ID}/${REPO_NAME}:latest',
    '.'
    ]
images: [
  'gcr.io/${PROJECT_ID}/${REPO_NAME}:${SHORT_SHA}', 
  'gcr.io/${PROJECT_ID}/${REPO_NAME}:latest'
  ]
```

其他參考：

* [Supported builder images for Google Cloud Build](https://cloud.google.com/cloud-build/)
* [Community-contributed images for Google Cloud Build](https://github.com/GoogleCloudPlatform/cloud-builders-community)