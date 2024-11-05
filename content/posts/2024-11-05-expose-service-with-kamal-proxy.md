---
categories:
  - note
keywords:
  - note
  - kamal
  - kamal-proxy
  - uptime kuma
  - docker
  - container
  - deploy
comments: true
date: 2024-11-05T08:00:00+08:00
title: "[筆記] 使用 kamal-proxy 暴露服務"
url: /2024/11/05/expose-service-with-kamal-proxy/
images:
  - /images/2024-11-05/expose-service-with-kamal-proxy.png
---

## 緣由

之前在 [Hetzner](https://www.hetzner.com/) 上買了一個小 VPS，也用 Rails 8 + [Kamal2](https://kamal-deploy.org/) 部署了一個小專案。

正好最近 uptimerobot.com 宣布免費方案條款變更，打算自行架設一個 [Uptime Kuma](https://uptime.kuma.pet/) 來替代。

{{< figure src="/images/2024-11-05/uptimerobot-announce.jpg" alt="UptimeRobot Announce" >}}

## 問題

因為 Uptime Kuma 是一個獨立的服務，並不是跟之前部署的專案有依賴關係。所以我原本是打算另外開一個 Git Repo 來存放用 Kamal 部署 Uptime Kuma 的相關設定。

但是在我反覆嘗試與翻找資料後，我遇到了幾個問題：

### 不必要的設定值

由於 Uptime Kuma 已經提供了公開的 Container Image，所以其實我們就可以跳過 build 跟 push 的部分，但 Kamal 還是會要求要填寫 build 跟 registry 的區塊。

其中一部分是可以填預設值跟 Dummy data 來 workaround。

### 都會先跑一遍 docker login

不管是 `kamal deploy` 還是 `kamal accessory boot` 指令，都會在遠端主機上執行 `docker login` 指令且不能跳過。這就只能填真實資料才能通過。

### Image 要有 Label

使用 `kamal deploy` 好不容易到了最後一步，卻發現 kamal 會在遠端主機執行：

```bash
Running docker inspect -f '{{ .Config.Labels.service }}' louislam/uptime-kuma:1 | grep -x uptime_tburl_tw || (echo "Image louislam/uptime-kuma:1 is missing the 'service' label" && exit 1)
```

Uptime Kuma 的 image 當然沒有這個 label，就理所當然地得到了錯誤：

```bash
docker stdout: Image louislam/uptime-kuma:1 is missing the 'service' label
```

那怎麼辦呢？只好再寫一個 Dockfile，執行完整的 `kamal deploy`。

部署是部署成功了，但最後想想實在是沒必要。就執行 `kamal app remove` 移除，放棄用 kamal 部署了。

## 最後方案

由於遠端主機上是透過 kamal-proxy 對 Cloudflare 開放 443 port。或許我自己可以在遠端主機上自行啟動 container，然後透過 kamal-proxy 對外呢？

在參考文件與主機上現行設定後，以下是我最後在遠端主機上執行的指令：

```bash
KUMA_SERVICE=<容器名稱，例如 uptime_kuma>
KUMA_HOST=<網域名稱，例如 uptime.example.com>
```

建立 Volume 供 Uptime Kuma 掛載使用。

```bash
docker volume create ${KUMA_SERVICE}_storage
```

啟動 Uptime Kuma Container，並等待狀態變為 Healthy。

```bash
docker run -d \
  --name ${KUMA_SERVICE} \
  --restart unless-stopped \
  --network kamal \
  --volume ${KUMA_SERVICE}_storage:/app/data \
  --label destination="" \
  --label role=web \
  --label service=${KUMA_SERVICE} \
  louislam/uptime-kuma:1
```

向 kamal-deploy 註冊 Uptime Kuma(開啟 tls，與 Cloudflare 之間就要設定成 Full TLS)。

```bash
docker exec kamal-proxy \
  kamal-proxy deploy ${KUMA_SERVICE} \
  --target ${KUMA_SERVICE}:3001 \
  --host ${KUMA_HOST} \
  --tls
```

再到 Cloudflare 設定一下 DNS，我的 Uptime Kuma 就成功運作了！

因為 Git Repo 就沒有需要了，所以趕快寫這篇筆記記錄一下。
