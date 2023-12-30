---
categories:
  - vue
keywords:
  - vue
  - vue3
  - vue-cli
  - vue-cli-service
  - frontend
  - development
  - https
comments: true
date: 2022-05-15T12:00:00+08:00
title: "如何設定 Vue3 開發環境使用 HTTPS 連線"
url: /2022/05/15/vue3-devserver-https/
images:
  - /images/2022-05-15/vue3-devserver-https.png
---

如果 Vue3 專案是使用 `vue-cli-service` 作為主要開發環境，可以透過編輯 `vue.config.js` 讓開發環境使用的 Server 支援 HTTPS。

## 設定

主要的設定如下：

```
module.exports = {
  devServer: {
    https: true,
    host: "localhost",
  }
}
```

- `devServer.https`: 啟用 HTTPS。
- `devServer.host`: 指定 Host，預設為本地端 IP。我們必須將其設定為 localhost，這樣才能在開發時自動更新頁面。否則會在網頁的 Console 出現 `sockjs-node/info` 相關的錯誤。

## 啟動 Server

啟動開發環境的 Server

```
yarn serve
```

確認輸出內容。Local 與 Network 的值將會同樣是 `https://localhost:8080/`。(除非 Port 8080 已被其他服務用走，才會更換成其他 Port)

```
 DONE  Compiled successfully in 1942ms                                                       11:27:13 PM


  App running at:
  - Local:   https://localhost:8080/
  - Network: https://localhost:8080/

  Note that the development build is not optimized.
  To create a production build, run yarn build.
```

因為是使用自簽憑證，用 Chrome 瀏覽頁面時會出現不安全的警告，要求訪問頁面即可。

## Ref

- https://cli.vuejs.org/config/#devserver
- https://www.shangmayuan.com/a/b7f639a1a6774570b0306f11.html
