---
categories:
  - docker
keywords:
  - docker
  - google cloud platform
  - google cloud
  - gcp
  - google compute engine
  - compute engine
  - gce
  - cloud logging
comments: true
date: 2022-02-20T12:00:00+08:00
title: Docker log on Google Compute Engine
url: /2022/02/20/docker-log-on-gce/
images:
  - /images/2022-02-20/docker-log-on-gce.png
---

在 Google Compute Engine 上安裝 Docker，想要將 Container 的 log 匯出至 Google Cloud Logging，可以將 Docker Log Driver 改成 `gcplogs` 後重新啟動 Docker service。

編輯 `/etc/docker/daemon.json`，在設定檔中加入 `log-driver` 設定：

```json
{
  "log-driver":"gcplogs"
}
```

記得重新啟動 Docker Service 載入變更後的設定：

```
sudo systemctl restart docker
```

這樣 Container 寫入 stdout 的 log 就會被送至 Google Cloud Logging。

其他相關的參數設定可參考文末連結。

## Ref

- https://docs.docker.com/config/containers/logging/gcplogs/
