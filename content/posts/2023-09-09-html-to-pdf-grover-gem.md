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
  - gem
  - PDF
  - HTML
  - PNG
  - JPEG
  - convert
  - print
comments: true
date: 2023-09-09T12:00:00+08:00
title: "HTML to PDF PDFkit 的替代方案 - Grover 介紹"
url: /2023/09/09/html-to-pdf-grover-gem/
images:
  - /images/2023-09-09/html-to-pdf-grover-gem.png
---

## 問題：PDFkit 的替代方案

為什麼要尋找 [PDFkit](https://github.com/pdfkit/pdfkit) 的替代方案？除了在包 Container 的時候有點困擾，其實 PDFkit 運作的還不錯。但今年發現 PDFkit 使用的 [wkhtmltopdf](https://github.com/wkhtmltopdf/wkhtmltopdf) 已經變成 archived。沒有在維護的專案未來會有不被新版本作業系統相容與安全性風險，於是就花了幾天找看看其他方案。

## 方案一：Grover

[Grover](https://github.com/Studiosity/grover) 是一個可以將 HTML 轉成 PDF、PNG 和 JPEG 的 gem。它包裝了一層，呼叫 node 指令執行 [Google Puppeteer](https://github.com/puppeteer/puppeteer) 透過 [Chromium](https://www.chromium.org/Home) 將 HTML 進行轉換。

### 安裝

本文使用的環境：
- macOS
- Ruby 3.2.2
- Rails 7.0.7.2
- Nodejs v20.5.0
- Docker Desktop 4.21.1

建立 Rails 專案

```bash
rails new demo
cd demo
```

執行指令安裝 Google Puppeteer

```bash
yarn add puppeteer
``````

執行指令安裝 Grover gem

```bash
bundle add grover
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

### 設定 Grover

我們可以設定 Grover 的預設值。

新增 `config/initializers/grover.rb`。`browser_ws_endpoint` 設定為遠端 Chrome 的位址。本文的 rails 假設是跑在本機環境，所以我們填入 `ws://localhost:9222/`。

```ruby
Grover.configure do |config|
  config.options = {
    format: 'Letter',
    margin: {
      top: '0.2in',
      right: '0.2in',
      bottom: '0.2in',
      left: '0.2in',
    },
    browser_ws_endpoint: 'ws://localhost:9222/'
  }
end
```

### 轉換方法一：直接呼叫 Grover 轉換

好處是比較容易在 Controller 中做錯誤處理。

建立範例 Controller

```bash
rails g controller string_pdfs index
```

編輯 `app/controllers/string_pdfs_controller.rb`

```ruby
class StringPdfsController < ApplicationController
  def index
    filename = 'string.pdf'
    pack_slip = render_to_string layout: false

    response.headers['content-disposition'] = "attachment; filename=#{filename}"
    render_pdf pack_slip, filename: filename
  rescue => e
    response.headers['content-disposition'] = ''
    Rails.logger.error "#{e.class}: #{e.message}"
    render plain: "#{e.class}: #{e.message}"
  end

  def render_pdf(html, filename:)
    pdf = Grover.new(html).to_pdf
    send_data pdf, filename: filename, type: "application/pdf"
  end
end
```

### 轉換方法二：透過 Middleware 轉換

Grover 提供了 Rails Middleware 整合，只要網址結尾為 `.pdf` 就會自動呼叫 Grover 進行轉換。[參考程式碼](https://github.com/Studiosity/grover/blob/a08dad78d6e63d28739a5db6fbf034f408b31dcd/lib/grover/middleware.rb#L48)

編輯 `config/application.rb`

```ruby
require_relative "boot"

require "rails/all"
require 'grover'

# Require the gems listed in Gemfile, including any gems
# you've limited to :test, :development, or :production.
Bundler.require(*Rails.groups)

module RailsPdfGenerator
  class Application < Rails::Application
    config.load_defaults 7.0
    config.middleware.use Grover::Middleware
  end
end
```

建立範例 Controller

```bash
rails g controller middleware_pdfs index
```

編輯 Controller `app/controllers/middleware_pdfs_controller.rb`

```ruby
class MiddlewarePdfsController < ApplicationController
  def index
    filename = 'middleware.pdf'

    response.headers['content-disposition'] = "attachment; filename=#{filename}"
    render layout: false
  end
end
```

### 範例首頁

建立範例 Controller

```bash
rails g controller home index
```

編輯 `app/views/home/index.html.erb`

```html
<div>
  <%= link_to 'string pdf', string_pdfs_path, target: "_blank", rel: "nofollow" %>
</div>
<div>
  <%= link_to 'middleware pdf', "#{middleware_pdfs_path}.pdf", target: "_blank", rel: "nofollow" %>
</div>
```

編輯 `config/routes.rb`

```ruby
Rails.application.routes.draw do
  resource :home, only:[:index]
  resources :string_pdfs, only:[:index]
  resources :middleware_pdfs, only:[:index]

  root "home#index"
end
```

### 測試

啟動 Rails server

```bash
bin/rails server
```

於瀏覽器開啟 [http://localhost:3000](http://localhost:3000)，點擊連結測試能否成功下載 PDF，PDF 內容是否正確。

![pdfs are fine](/images/2023-09-09/20230909001.jpg)

## CJK 問題

在以上測試，英文的轉換都沒問題，但是當 HTML 的內容中有中文的時候，會出現亂碼。

編輯 `app/views/string_pdfs/index.html.erb`

```html
<h1>Hello 世界！</h1>
```

編輯 `app/views/middleware_pdfs/index.html.erb`

```html
<h1>哈囉 World！</h1>
```

於瀏覽器開啟 http://localhost:3000，點擊連結測試。可以發現下載的 PDF 無法正確地顯示中文。

![pdfs has CJK problem](/images/2023-09-09/20230909002.jpg)

### 尋找原因

開啟瀏覽器訪問 [http://localhost:9222](http://localhost:9222)

可以看到 `browerless/chrome` container 的測試工具。

![debug tool of browserless/chrome](/images/2023-09-09/20230909003.png)

參考 [Grover 呼叫 puppeteer 的程式碼](
https://github.com/Studiosity/grover/blob/78f0695c01c4c01bb0672785964a9acfb28a9887/lib/grover/js/processor.cjs#L172)，在左側輸入以下程式碼進行測試：

```javascript
// For PDFs, let's take some API content and inject some simple styles
export default async ({ page }: { page: Page }) => {

  let urlOrHtml = `<div>A天氣真好B<div>`;
  urlOrHtml += `<div style="font-family: 'Noto Sans TC';">A天氣真好B<div>`;
  urlOrHtml += `<div style='font-family: BlinkMacSystemFont, -apple-system, "Segoe UI", "Roboto", "Oxygen", "Ubuntu", "Cantarell", "Fira Sans", "Droid Sans", "Helvetica Neue", "Helvetica", "Arial", sans-serif;'>A天氣真好B<div>`;

  await page.setRequestInterception(true);
  let htmlIntercepted = false;
  page.on('request', request => {
    // We only want to intercept the first request - ie our HTML
    if (htmlIntercepted)
      request.continue();
    else {
      htmlIntercepted = true
      request.respond({ body: urlOrHtml });
    }
  });
  await page.goto('http://example.com', {waitUntil: 'networkidle0'});

  // Return a PDF buffer to trigger the editor to download.
  return page.pdf();
};
```

可以發現就算在樣式中指定了字型，右側還是都顯示亂碼。

![it can not render CJK](/images/2023-09-09/20230909004.png)

如果用 `setContent` 取代 `Interception`，在左側輸入以下程式碼進行測試：

```javascript
// For PDFs, let's take some API content and inject some simple styles
export default async ({ page }: { page: Page }) => {

  let urlOrHtml = `<div>A天氣真好B<div>`;
  urlOrHtml += `<div style="font-family: 'Noto Sans TC';">A天氣真好B<div>`;
  urlOrHtml += `<div style='font-family: BlinkMacSystemFont, -apple-system, "Segoe UI", "Roboto", "Oxygen", "Ubuntu", "Cantarell", "Fira Sans", "Droid Sans", "Helvetica Neue", "Helvetica", "Arial", sans-serif;'>A天氣真好B<div>`;

  await page.setContent(urlOrHtml)

  // Return a PDF buffer to trigger the editor to download.
  return page.pdf();
};
```

即使樣式不指定字型，也能在右側顯示中文。

![it can render CJK](/images/2023-09-09/20230909005.png)


看起來主要是這一段程式碼導致中文沒有辦法轉換，如果不修改這一段似乎就沒有辦法解決顯示中文的問題。

## 結論

我需求是尋找可替代 PDFkit 的 HTML 轉換 PDF 方案，但是：
1. Grover 需要 nodejs 執行環境執行 node 指令
2. Grover 需要 Chrome/Chromium 進行轉換
3. 初步來看 Grover 需要修改原始碼才能解決中文顯示問題

所以我最後沒有在專案上使用 Grover。我會在之後介紹我在研究過程中所接觸到的其他方案。

## 參考資料
- [PDFkit](https://github.com/pdfkit/pdfkit)
- [wkhtmltopdf](https://github.com/wkhtmltopdf/wkhtmltopdf)
- [Grover](https://github.com/Studiosity/grover)
- [Google Puppeteer](https://github.com/puppeteer/puppeteer)
- [browserless](https://www.browserless.io/docs/docker-quickstart)
- [browserless/chrome](https://github.com/browserless/chrome)
- [Converting HTML to PDF using Rails](https://dev.to/ayushn21/converting-html-to-pdf-using-rails-54e7)
- [Fly.io Sample](https://github.com/fly-apps/dockerfile-rails/blob/main/DEMO.md#demo-6---grover--puppeteer--chromium)