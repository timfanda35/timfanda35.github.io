---
categories:
  - ruby
  - rails
comments: true
date: 2017-02-11T11:59:00+08:00
title: Ruby/Rails load setting file
url: /2017/02/11/ruby-rails-load-setting-file/
---

最近在網路上看到了這兩篇文章：

[Rails 最佳实践之配置管理](https://ruby-china.org/topics/32126)

[在 Rails 內建立類似 database.yml 的設定文件（global settings）](http://railsfun.tw/t/rails-database-yml-global-settings/996)

才知道原來 Rails 4.2 之後有好用的 `config_for` 方法啊

<!--more-->

## 環境變數

由於自己寫的都是些小專案，接的服務也不多

用環境變數來管理因應環境而要調整的設定就綽綽有餘了

這部分可以使用以下 gem 來協助：

* [dotenv](https://github.com/bkeepers/dotenv)

* [figaro](https://github.com/laserlemon/figaro)

## YAML

不過如果需要結構化設定，或是有大量要調整的設定項目

建議使用 yaml 來管理

在 ruby 專案中可以用以下方式讀取

```ruby
require 'yaml'
env = ENV['RUBY_ENV']
$settings = YAML.load(File.open("#{__dir__}/settings.yml"))[env]
```

在 Rails 4.2 以後，可以修改 `config/application.yml` 

用以下方式讀取 `config/settings.yml`

```ruby
module MyApplication
  class Application < Rails::Application
  	# ...
    
    # add this line
    $settings = config_for(:settings)
  end
end
```

## Gem

如果你有很複雜的想法，可以試試這個 gem： [config](https://github.com/railsconfig/config)

支援：
* Ruby 2.x
* Rails >= 3.1, 4 and 5
* Padrino
* Sinatra

Rails 專案想要可以在線上修改設定？ 可以試試這個 gem： [rails-settings-cached](https://github.com/huacnlee/rails-settings-cached)

說明寫得很詳細，包含設定頁面的程式碼範例





