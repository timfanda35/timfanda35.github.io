---
categories:
  - rails
keywords:
  - Rails
  - Ruby
  - PDFkit
  - wkhtmltopdf
  - grover
  - puppeteer
  - puppeteer-ruby
  - gem
  - PDF
  - HTML
  - PNG
  - JPEG
  - convert
  - print
comments: true
date: 2023-09-23T12:00:00+08:00
title: "HTML to PDF PDFkit 的替代方案 - puppeteer-ruby 介紹"
url: /2023/09/23/html-to-pdf-puppeteer-ruby-gem/
images:
  - /images/2023-09-23/html-to-pdf-puppeteer-ruby.png
---

本文接續之前文章 [HTML to PDF PDFkit 的替代方案 - Grover 介紹](/2023/09/09/html-to-pdf-grover-gem/)。

## 方案二：Puppeteer-ruby

在之前文章發現 [Grover](https://github.com/Studiosity/grover) 有以下問題是我想要改善的：
1. Grover 需要 nodejs 執行環境執行 node 指令
2. Grover 需要 Chrome/Chromium 進行轉換
3. 初步來看 Grover 需要修改原始碼才能解決中文顯示問題

其中第三點雖然目前還沒有需求，但未來也是個隱憂，在沒有能力修改原始碼的情況下，我打算先看看有沒有其他選擇。

以前研究爬蟲的時候，也多少了解到 Puppeteer 是如何去與 Chrome/Chromium 溝通的。Chrome/Chromium 提供了基於 WebSocket 技術實現的 [DevTools Protocol](https://chromedevtools.github.io/devtools-protocol/)，開發者也可以多花點功夫讀懂規格並去自己實作相關功能，不一定非得要使用 Puppeteer。

而我的需求是透過 Chrome/Chromium 將畫面轉成 PDF，且主要開發語言是 Ruby，為此要安裝 nodejs 執行環境並安裝 Puppeteer，似乎殺雞焉用牛刀。

於是我把念頭轉向了尋找 porting Puppeteer 功能的 Ruby gem 上，而我找到了 [puppeteer-ruby](https://github.com/YusukeIwaki/puppeteer-ruby)。

puppeteer-ruby 是純 Ruby 實作，雖然還沒有支援 puppeteer 所有的功能，但我所需要的 Remote Connect 與 PDF 都已經實現了。

### 安裝

本文使用的環境：
- macOS
- Ruby 3.2.2
- Rails 7.0.7.2
- Docker Desktop 4.21.1

建立 Rails 專案

```bash
rails new demo
cd demo
```

執行指令安裝 `puppeteer-ruby`

```bash
bundle add puppeteer-ruby
```

### 建立 Chrome container(可選)

Google Puppeteer 會自動下載 Chromium，但也可以透過 connect 方法去連線到現有的 Chromium 或是 Chrome。我們可以直接使用本地環境上的 Chromium 或是 Chrome，也可以使用網路上第三方包好的 Container Image。

由於我在線上環境都是使用 Container Image 來部署應用程式，所以在本文我們會使用 [browserless](https://www.browserless.io/docs/docker-quickstart) 所建置的 [Container Image](https://github.com/browserless/chrome) 來作為 Grover 遠端連接的 Chrome，而不是讓 Grover 自動下載。

注意：Chrome 其實蠻吃資源的，所以要留意部署機器上的空間。我曾因為 Docker VM 硬碟空間不足導致連線時 Container 中的 Chrome 無法開啟 Browser 而連線失敗。

我們透過 docker compose 來管理 Chrome container。新增 `docker-compose.yml`

```yaml
services:
  chrome:
    image: browserless/chrome
    ports:
      - 9222:3000
```

執行指令建立 Chrome container

```
# v2
docker compose up -d
```

### 轉換方法：直接呼叫 `puppeteer-ruby` 轉換

建立範例 Controller

```bash
rails g controller puppeteer_pdfs index
```

編輯 `app/controllers/puppeteer_pdfs_controller.rb`

```ruby
class PuppeteerPdfsController < ApplicationController
  def index
    filename = 'puppeteer-ruby.pdf'
    pack_slip = render_to_string layout: false

    response.headers['content-disposition'] = "attachment; filename=#{filename}"
    render_pdf pack_slip, filename: filename
  rescue => e
    response.headers['content-disposition'] = ''
    Rails.logger.error "#{e.class}: #{e.message}"
    render plain: "#{e.class}: #{e.message}"
  end

  def render_pdf(html, filename:)
    # 連到遠端的 Chrome/Chromium
    browser = Puppeteer.connect(browser_url: 'http://localhost:9222')

    # 取得 Puppeteer Page
    page = browser.pages.first

    # 設定 Puppeteer Page Content
    page.set_content(html)

    # 從 Puppeteer Page 產生 PDF
    pdf = page.pdf(format: 'Letter',
                   margin: {
                     top: '0.2in',
                     right: '0.2in',
                     bottom: '0.2in',
                     left: '0.2in',
                   })

    # 回傳 PDF 給使用者端
    send_data pdf, filename: filename, type: "application/pdf"

    # 關閉與遠端 Chrome/Chromium 的連線
    browser.disconnect
  end
end
```

編輯 `app/views/puppeteer_pdfs/index.html.erb`

```html
<h1>Hello 世界！</h1>
```

### 範例首頁

建立範例 Controller

```bash
rails g controller home index
```

編輯 `app/views/home/index.html.erb`

```html
<div>
  <%= link_to 'puppeteer pdf', puppeteer_pdfs_path, target: "_blank", rel: "nofollow" %>
</div>
```

編輯 `config/routes.rb`

```ruby
Rails.application.routes.draw do
  root 'home#index'

  resources :puppeteer_pdfs, only: [:index]
end
```


### 測試

啟動 Rails server

```bash
bin/rails server
```

於瀏覽器開啟 [http://localhost:3000](http://localhost:3000)，點擊連結測試能否成功下載 PDF，PDF 內容是否正確。

![CJK pdf is fine](/images/2023-09-23/20230923001.jpg)

## 結論

從測試中可以發現，我們解決了上一篇想要改善的問題 2 與問題 3：
- 我們在執行環境中不再需要依賴 nodejs 執行環境與 Puppeteer package
- 我們解決了 CJK 問題，產生的 PDF 可以正確地顯示中文

剩下最後一個想要改善的問題是 Chrome/Chromium 需要使用大量資源，光是上述測試使用的 Container Image size 就高達了 3GB。

目前我的需求只是偶爾需要產生 PDF，為此要在伺服器上長駐 Chrome/Chromium 好像有點奢侈了，但這已經是個還不錯的方案。

能省則省，我想了又想，既然 `wkhtmltopdf` 可以透過 Qt 來實現 HTML to PDF，說不定也有其他不依賴 Chrome/Chromium 也能做到 HTML to PDF 的方案。

我最後沒有在專案上使用 `puppeteer-ruby`，因為我找到了不依賴 Chrome/Chromium 也能做到 HTML to PDF 的方案，我會在另一篇介紹。

## 參考資料
- [PDFkit](https://github.com/pdfkit/pdfkit)
- [wkhtmltopdf](https://github.com/wkhtmltopdf/wkhtmltopdf)
- [Grover](https://github.com/Studiosity/grover)
- [Google Puppeteer](https://github.com/puppeteer/puppeteer)
- [puppeteer-ruby](https://github.com/YusukeIwaki/puppeteer-ruby)
- [browserless](https://www.browserless.io/docs/docker-quickstart)
- [browserless/chrome](https://github.com/browserless/chrome)
- [Converting HTML to PDF using Rails](https://dev.to/ayushn21/converting-html-to-pdf-using-rails-54e7)
- [Fly.io Sample](https://github.com/fly-apps/dockerfile-rails/blob/main/DEMO.md#demo-6---grover--puppeteer--chromium)
- [DevTools Protocol](https://chromedevtools.github.io/devtools-protocol/)
