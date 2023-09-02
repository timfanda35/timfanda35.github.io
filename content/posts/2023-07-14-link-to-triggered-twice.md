---
categories:
  - rails
keywords:
  - ruby
  - rails
  - turbo
  - turbolink
comments: true
date: 2023-07-14T12:00:00+08:00
title: "Rails 7 link_to 觸發兩次"
url: /2023/07/14/link-to-triggered-twice/
---

最近在 Rails 7 遇到 `link_to` 會觸發兩次的情況

Rails Version: `7.0.6`

## 正常情況

建立測試專案

```bash
rails new demo
cd demo
```

新增 Controller

```bash
rails g controller Articles index
```

編輯 `config/routes.rb`

```ruby
Rails.application.routes.draw do
  resources :articles, only: [:index]

  root "articles#index"
end
```

編輯 `app/views/articles/index.html.erb`

```html
<h1>Articles#index</h1>
<%= link_to 'link', 'https://www.google.com' %>
```

啟動 Rails Server

```bash
bin/rails server
```

使用 Chrome 瀏覽器開啟 http://localhost:3000。啟動開發者工具，切換到 Network 頁籤，勾選 Preserve log。

![](/images/2023-07-14/Xnip2023-07-14_15-57-13.jpg)

點擊 link 觀察請求。使用 filter 過濾 google，請求正常。

![](/images/2023-07-14/Xnip2023-07-14_15-58-35.jpg)

## 若為跳轉連結

這裡為求簡單，直接使用 `show` 來作為重現的 controller method。

編輯 `config/routes.rb`

```ruby
Rails.application.routes.draw do
  resources :articles, only: [:index, :show]

  root "articles#index"
end
```

編輯 `app/controllers/articles_controller.rb`

```ruby
class ArticlesController < ApplicationController
  def index
  end

  def show
    redirect_to('https://www.google.com', allow_other_host: true)
  end
end
```

編輯 `app/views/articles/index.html.erb`

```html
<h1>Articles#index</h1>
<%= link_to 'link', article_path(0) %>
```

測試後，會發現觸發了兩次請求，一次是 AJAX，一次是直接請求。

![](/images/2023-07-14/Xnip2023-07-14_16-05-14.jpg)

## 修正

由於該連結是直接挑轉到外部網站，所以我們可以在網頁元素加上 `data-turbo=false` 停用 `turbo` 來修正。

編輯 `app/views/articles/index.html.erb`

```html
<h1>Articles#index</h1>
<%= link_to 'link', article_path(0), data: { turbo: "false" } %>
```

測試結果正常

![](/images/2023-07-14/Xnip2023-07-14_16-09-21.jpg)

## 參考資料

- [data-turbo="false" doesn't seem to work](https://github.com/hotwired/turbo/issues/119#issuecomment-765708124)
