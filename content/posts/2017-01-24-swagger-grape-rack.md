---
categories:
  - ruby
comments: true
date: 2017-01-24T06:43:00+08:00
title: swagger + grape + rack
url: /2017/01/24/swagger-grape-rack/
---

常常聽到 [swagger](http://swagger.io/) 跟 [grape](https://github.com/ruby-grape/grape) 的名字

所以用一個很陽春的[專案](https://github.com/timfanda35/test-grape-api)練習一下要如何使用

<!--more-->

## 使用到的 Gem:

* [grape](https://github.com/ruby-grape/grape) API 框架
* [grape-swagger](https://github.com/ruby-grape/grape-swagger) API 文件產生器
* [vienna](https://github.com/kmikael/vienna) 簡單設定讓 Rack 處理public 靜態文件

## 其他使用到的 Repo

* [swagger-ui](https://github.com/swagger-api/swagger-ui) API 文件瀏覽前端

## 開發

用 [grape](https://github.com/ruby-grape/grape) 開發 API，以及加上 swagger 文件支援，這些都照著官方文件走就可以了

由於要用 rack 同時跑 API 與 [swagger-ui](https://github.com/swagger-api/swagger-ui)，所以在 `config.ru` 用了 [vienna](https://github.com/kmikael/vienna)

[vienna](https://github.com/kmikael/vienna) 可以讓 rack 套用 `public` 目錄下的靜態檔案成為網站的一部份

`config.ru 範例`

```ruby
require 'bundler/setup'
Bundler.require

require "vienna"
require "grape"
require "grape-swagger"
require_relative "twitter/api.rb"

module API
  class Root < Grape::API
    format :json

    # CORS
    before do
      header['Access-Control-Allow-Origin'] = '*'
      header['Access-Control-Request-Method'] = '*'
    end

    mount Twitter::API
    add_swagger_documentation
  end
end

# 除了 API 還要跑 public 下的靜態檔案
run Rack::Cascade.new [API::Root, Vienna]
```

## Swagger UI

從 [swagger-ui](https://github.com/swagger-api/swagger-ui) 下載壓縮檔

把 `swagger-ui/dist` 下的檔案複製到專案的 `public` 目錄 下

打開 `index.html` 將

```js
if (url && url.length > 1) {
  url = decodeURIComponent(url[1]);
} else {
  url = "http://petstore.swagger.io/v2/swagger.json";
}
```

修改成

```js
if (url && url.length > 1) {
  url = decodeURIComponent(url[1]);
} else {
  url = "http://localhost:9292/swagger_doc.json";
}
```

## Execute

Run command:
```
bundle exec rackup
```

打開瀏覽器瀏覽 `http://localhost:9292/` 就可以看到 swagger 風格的 API 文件了！