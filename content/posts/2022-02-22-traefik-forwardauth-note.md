---
categories:
  - docker
keywords:
  - docker
  - docker compose
  - traefik
  - load balancer
  - load balancing
  - authentication
comments: true
date: 2022-02-22T12:00:00+08:00
title: Traefik ForwardAuth 筆記
url: /2022/02/22/traefik-forwardauth-note/
images:
  - /images/2022-02-22/traefik-forwardauth-note.png
---

想要在 Traefik 設定呼叫外部服務來驗證請求是否可以轉發至後端。

## 示意圖

當請求送往 Traefik 的時候，會先將請求轉發至 Auth Service，透過 Auth Service 回傳的 HTTP Status Code 來決定能不能將請求轉送至 Backend。

```
Client <--> Traefik <--> Backend
               ^
               | Forward Auth
               v
             Auth Service
```

## 使用 Docker Compose 建立測試環境

假設本地環境為 Linux-like 作業系統，並已安裝好 Docker、Docker Compose、curl 工具。

建立目錄

```bash
mkdir traefik-forward-demo
cd traefik-forward-demo
```

### 新增 traefik 設定檔

新增 `conf/traefik.yml`，設定可參考 [server configuration](https://doc.traefik.io/traefik/reference/static-configuration/file/)。

```yaml
api:
  insecure: true
providers:
  file:
    directory: "/etc/traefik"
    watch: true
```

新增 `conf/dynamic.yml`，設定可參考 [File provider configuration](https://doc.traefik.io/traefik/providers/file/)。為了測試方便，我們分別建立可以驗證成功與失敗的 routers 與 middlewares。

```yaml
http:
  routers:
    succ:
      rule: "Host(`succ-auth.docker.localhost`)"
      service: myip
      middlewares:
        - "succ-auth"
    fail:
      rule: "Host(`fail-auth.docker.localhost`)"
      service: myip
      middlewares:
        - "fail-auth"
  middlewares:
    succ-auth:
      forwardAuth:
        address: "http://httpbin/headers"
        trustForwardHeader: true
    fail-auth:
      forwardAuth:
        # Set a not found url to make auth fail
        address: "http://httpbin/headersx"
        trustForwardHeader: true
  services:
    myip:
      loadBalancer:
        passHostHeader: false
        servers:
        - url: "https://api.ipify.org"
```

新增 `docker-compose.yml`，建立 traefik container 與用來當作 Auth Service 的 httpbin container。

```yaml
version: '3'

services:
  reverse-proxy:
    # The official v2 Traefik docker image
    image: traefik:v2.5
    ports:
      # The HTTP port
      - "80:80"
      # The Web UI (enabled by --api.insecure=true)
      - "8080:8080"
    volumes:
      # So that Traefik can listen to the Docker events
      - /var/run/docker.sock:/var/run/docker.sock
      - ./conf:/etc/traefik
  httpbin:
    image: kennethreitz/httpbin
    environment: # Enable httpbin log
      - GUNICORN_CMD_ARGS=${GUNICORN_CMD_ARGS}
```

新增 `.env`，這樣可以讓 httpbin container 印出 access log。

```
GUNICORN_CMD_ARGS="--capture-output --error-logfile - --access-logfile - --access-logformat '%(h)s %(t)s %(r)s %(s)s Host: %({Host}i)s X-MY-ID: %({X-MY-ID}i)s'"
```

## 啟動測試環境

啟動 containers。

```bash
docker-compose up
```

在測試過程中可以觀察 httpbin 印出的 access log。

開啟另一個 terminal session 用來執行以下測試的指令。

## 測試驗證成功的請求

```
curl \
    -H Host:succ-auth.docker.localhost \
    -H X-MY-ID:hello@dream \
    "http://127.0.0.1"
```

如果成功，畫面將會輸出對外 IP。

```bash
172.25.0.1
```

由於 Auth Service 回傳 HTTP 狀態碼 200，所以請求會繼續往下送，轉發至 myip service。

![](/images/2022-02-22/traefik-forwardauth-note/001.jpg)

## 測試驗證失敗的請求

```
curl \
    -H Host:fail-auth.docker.localhost \
    -H X-MY-ID:hello@dream \
    "http://127.0.0.1"
```

如果成功，畫面將會輸出對外 IP。

```
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 3.2 Final//EN">
<title>404 Not Found</title>
<h1>Not Found</h1>
<p>The requested URL was not found on the server.  If you entered the URL manually please check your spelling and try again.</p>
```

由於我們將 fail-auth middleware 驗證的端點指向 Auth Service 不存在的路徑，所以驗證的時候會回傳 HTTP 狀態碼 404，請求就不會繼續往 myip service 送，而是直接把 Auth Service 的回應返回至 client。

![](/images/2022-02-22/traefik-forwardauth-note/002.jpg)

## 停止測試環境

於剛剛執行 `docker-compute up` 指令的 terminal session 輸入 `crtl` + `c` 停止 containers。

## Ref

- https://doc.traefik.io/traefik/
