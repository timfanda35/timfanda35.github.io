---
categories:
  - ruby
comments: true
date: 2017-01-05T12:00:00+08:00
title: Start ruby project with bundler
url: /2017/01/05/start-ruby-project-with-bundler/
---

有時候會想要寫個 side project

不是只有一個檔案的 script

因為會用到其他人寫好的 gem，想要用 [bundler][bundler] 來管理

在這裡紀錄一下建立用 [bundler][bundler] 管理依賴的 ruby 專案的起始步驟

<!--more-->

首先先建立一個目錄，並移動到該目錄下

```
$ mkdir project
$ cd project
```

執行指令初始化目錄

```
$ bundle init

# 如果沒有 bundle 指令，可能是沒有安裝 bundler 的關係
$ gem install bundler 
```

以上指令會在目錄下產生 `Gemfile`，當然這檔案你也可以自己打，不過打指令產生不是好記又快嗎？

```
# frozen_string_literal: true
# A sample Gemfile
source "https://rubygems.org"

# gem "rails"
```

記得編輯完 `Gemfile` 要執行 `bundle install` 安裝 gem

接下來新增一個作為專案起始檔案 `app.rb`，在檔案開頭加上

```
require 'rubygems'
require 'bundler/setup'
```

`require 'bundler/setup'` 會將 `Gemfile` 的 gem 加入到 `LOAD_PATH` 裡，這樣就可以用 `require` 引入`Gemfile` 裡面的 gem 了

但如果你要引入的 gem 很多，懶得一個個引入的話，可以用 `Bundler.require` 一次引入 `Gemfile` 裡面全部的 gem，如將 `app.rb` 修改如下：

```
require 'rubygems'
require 'bundler/setup'
Bundler.require
```

不過 `Bundler.require` 在 require 的時候是直接使用 gem 的名稱，如果有 gem 在 require 的時候要用不同的名稱的話，必須在 `Gemfile` 裡指名 require 時的名稱，如

```
# Gemfile
gem 'rack-cache', :require => 'rack/cache'
```

到這邊就已經可以用 `Gemfile` 來管理專案所以用到的 gem ，可以開始開發囉！

[bundler]: http://bundler.io/