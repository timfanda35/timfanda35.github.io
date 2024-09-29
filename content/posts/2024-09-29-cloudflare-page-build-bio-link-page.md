---
categories:
  - note
keywords:
  - note
  - tailwind
  - bio
  - cloudflare
  - cloudflare page
  - build
  - github
comments: true
date: 2024-09-29T08:00:00+08:00
title: "[筆記] 使用 Cloudflare Page 建置 Bio Link 頁面"
url: /2024/09/29/cloudflare-page-build-bio-link-page/
images:
  - /images/2024-09-29/cloudflare-page-build-bio-link-page.png
---

## 緣由

看到推友建立了一個 Bio Link 頁面，想想主網域也一直空著，起心動念自己也來建置一個頁面。

<blockquote class="twitter-tweet"><p lang="zh" dir="ltr"><a href="https://twitter.com/hashtag/%E9%9A%8F%E4%BE%BF%E5%86%99%E5%86%99?src=hash&amp;ref_src=twsrc%5Etfw">#随便写写</a> 随手整了一个 Bio Link 自部署版本，用于 Twitter 的个人站点，甚至可以一键唤起 GitHub，假如你需要可右键显示源码自己改一下自用。<br>🤖 <a href="https://t.co/gFcSs6YTUK">https://t.co/gFcSs6YTUK</a> <a href="https://t.co/kkGdoPOA7g">pic.twitter.com/kkGdoPOA7g</a></p>&mdash; Tw93 (@HiTw93) <a href="https://twitter.com/HiTw93/status/1839821826383360137?ref_src=twsrc%5Etfw">September 28, 2024</a></blockquote>

## 靜態 HTML 專案

### 建立專案目錄

```bash
mkdir bio
mkdir -p bio/src bio/public
cd bio
```

- src: 原 css style
- public: 靜態資源

### 安裝 Tailwind CSS

原本是打算手刻 CSS 就好，但想想還是用 Tailwind CSS 好了，畢竟想 CSS Class 名字或是寫落落長的 style 也是很燒腦筋。

```bash
npm install -D tailwindcss
npx tailwindcss init
```

新增檔案 `src/input.css`。

```css
@tailwind base;
@tailwind components;
@tailwind utilities;
```

編輯檔案 `tailwind.config.js`。

```javascript
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./public/**/*.{html,js}"],
  theme: {
    extend: {},
  },
  plugins: [],
}
```

寫一個 Makefile 使用 Tailwind CLI 進行建置。

```makefile
dev:
	npx tailwindcss -i ./src/input.css -o ./public/style.css --watch

build:
	npx tailwindcss -i ./src/input.css -o public/style.css --minify
```

這樣我們在開發的時候就可以在終端機執行 `make` 或是 `make dev`，讓 Tailwind 即時建置。

部署的時候則是執行 `make build` 產生最小化的版本。

從指令可以看出來最後建置出的 css file 會是 `public/style.css`。

### 編寫 HTML

新增檔案 `public/index.html`，這一個步驟就只是編寫 HTML 頁面而已。

我們在 HTML 頁面中會引用 Tailwind CSS 建置出的 `public/style.css` 檔案。

```html
<!doctype html>
<html lang="zh-Hant">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link href="./style.css" rel="stylesheet">
...
```

圖示的部分我是使用 [Self-Hosted Dashboard Icons](https://selfh.st/icons/)。

## 部署到 Cloudflare Page

Cloudflare Page 可以與 GitHub Repo 連動，當有新的變更的時候就會自動建置部署。

步驟按照文件 [Cloudflare Page Git integration](https://developers.cloudflare.com/pages/get-started/git-integration/)，很快就能設定完成。

### 連結 GitHub 帳號

於 Cloudflare 控制台前往「Workers & Pages -> Overview」，點擊「Create」。

{{< figure src="/images/2024-09-29/cf-page-01.jpg" alt="manage cloudflare page" >}}

切換至「Pages」，點擊「Connect to Git」。

{{< figure src="/images/2024-09-29/cf-page-02.jpg" alt="connect github repo" >}}

我是使用 GitHub，所以連接 GitHub 帳號，並授權存取對應的 Git Repository。

點擊 Git Repository 後，會出現綠色 Check 的圖示，就可以點擊「Begin setup」進行下一步。

{{< figure src="/images/2024-09-29/cf-page-03.jpg" alt="choose repo" >}}

### 建置設定

在這邊可以輸入自訂的專案名稱，以及讓 Cloudflare 監聽的 Git Branch，當該 Git Branch 有更新時就會開始自動建置部署。

- Framework preset：因為我沒有用到靜態網站產生器或是 Javascript 框架，所以不用選擇。
- Build command： 使用我們自訂的 `make build`。
- Build output directory： `public`。

點擊「Save and Deploy」就會開始進行建置部署。

{{< figure src="/images/2024-09-29/cf-page-04.jpg" alt="setup build process" >}}

### 綁定自訂網域

前往 Cloudflare Page 專案詳情。

{{< figure src="/images/2024-09-29/cf-page-05.jpg" alt="page detail" >}}

切換至「Custom domains」，點擊「Set up a custom domain」。

{{< figure src="/images/2024-09-29/cf-page-06.jpg" alt="set up a custom domain" >}}

輸入想要綁定的網域。我使用的網域也是託管在 Cloudflare 上，所以這步驟會自動新增一筆 CNAME Record。

{{< figure src="/images/2024-09-29/cf-page-07.jpg" alt="enter custom domain" >}}

## 成果

{{< figure src="/images/2024-09-29/bio.jpg" alt="bio screenshot" >}}

目標網址在這裡：[https://bear-su.dev](https://bear-su.dev)

程式碼在這裡：[https://github.com/timfanda35/bio](https://github.com/timfanda35/bio)

## 注意事項

1. 由於我使用固定檔名，且 Cloudflare 預設會快取靜態資源，所以部署新版後，想要盡快使用最新的版本，可以自行在 [Cloudflare 控制台清除快取](https://developers.cloudflare.com/cache/how-to/purge-cache/)。
2. 由於我是使用 Cloudflare 免費方案，所以沒辦法使用台灣節點。從瀏覽器的開發者工具可以看到 Response Header [Cf-Ray](https://developers.cloudflare.com/fundamentals/reference/http-request-headers/#cf-ray) 是 LAX 結尾，從 [Cloudflare System Status](https://www.cloudflarestatus.com/) 查詢為 LAX 代號是 Los Angeles, CA, United States，從台灣訪問的回應速度大約 1 ~ 200ms。

{{< figure src="/images/2024-09-29/cf-region.jpg" alt="cf lax" >}}

## 參考
- [Get started with Tailwind CSS](https://tailwindcss.com/docs/installation)
