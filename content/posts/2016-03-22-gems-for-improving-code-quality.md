---
categories:
  - ruby
  - rails
  - brakeman
  - bullet
  - rubycritic
  - newrelic
  - gem
comments: true
date: 2016-03-22T11:28:00+08:00
title: 推薦協助提高程式碼品質的 Gem
url: /2016/03/22/gems-for-improving-code-quality/
---

用 Rails 開發網站時除了寫測試之外，還可以透過一些 gem 來協助維護 code 的品質

這篇推薦幾個 gem：

[Brakeman - 檢查安全性](https://github.com/presidentbeef/brakeman)

[Bullet - 檢查 N+1 Query](https://github.com/flyerhzm/bullet)

[RubyCritic - 靜態分析程式碼(已整合多套工具)](https://github.com/whitesmith/rubycritic)

[New Relic RPM - 監測網站效能](https://github.com/newrelic/rpm)

<!--more-->

## [Brakeman](https://github.com/presidentbeef/brakeman)

Brakeman 可以用來分析程式碼，以檢測網站的安全性問題提出報告作為改善參考

可以直接在終端機輸入指令安裝 gem：

```shell
gem install brakeman
```

或是修改專案中的 `Gemfile` 透過 `bundle` 安裝

```ruby
group :development do
  gem 'brakeman', :require => false
end
```

因為只有開發環境才會用到，所以通常都放在 `development` 群組中

這裡只是用來在 `bundle install` 一併安裝 Brakeman，在專案應用執行過程中並不會引用到 Brakman，所以將 `require` 設為 `false`

使用 Brakeman 很簡單，於終端機在專案目錄下執行指令：

```shell
brakeman
```

就會批哩啪啦地跑出檢測報告了

接著就可以依照檢測報告改善程式碼

## [Bullet](https://github.com/flyerhzm/bullet)

通常都會用 Bullet 來檢查 N + 1 Query 問題(對資料庫額外執行多個查詢子句而影響效能)

Bullet 會在執行查詢的時候，將發生 N + 1 Query 的部分輸出於 log 上

修改專案中的 `Gemfile` 透過 `bundle` 安裝

```ruby
group :test, :development do
  gem 'bullet'
end
```

將 Bullet 加入到 `development` 與 `test` 群組以看到 Bullet 所產生的 log 

修改 `config/environments/test.rb`，加入以下設定

```ruby
config.after_initialize do
  Bullet.enable = true
  Bullet.bullet_logger = true
  Bullet.raise = true # raise an error if n+1 query occurs
end
```

並在 `spec/spec_helper.rb` 或 `spec/rails_helper.rb` 中加入以下設定

```ruby
RSpec.configure do |config|

  # other configures...
  
  if Bullet.enable?
    config.before(:each) do
      Bullet.start_request
    end

    config.after(:each) do
      Bullet.perform_out_of_channel_notifications if Bullet.notification?
      Bullet.end_request
    end
  end
end
```

這樣在跑測試的時候，如果有 N + 1 Query 時就會無法通過測試

再依照出現的訊息去改善程式吧

## [RubyCritic](https://github.com/whitesmith/rubycritic)

RubyCritic 整合了多個靜態分析程式碼的工具像是 [Reek](https://github.com/troessner/reek), [Flay](https://github.com/seattlerb/flay) and [Flog](https://github.com/seattlerb/flog)，產生報表用以檢視專案的程式碼品質與改善建議

可以直接在終端機輸入指令安裝 gem：

```shell
gem install rubycritic
```

或是修改專案中的 `Gemfile` 透過 `bundle` 安裝

```ruby
group :development do
  gem 'rubycritic', :require => false
end
```

通常都是在 `development` 環境下，於專案目錄下執行指令：

```shell
rubycritic
```

檢測結束後，就會自動開啟瀏覽器顯示報告

## [New Relic RPM](https://github.com/newrelic/rpm)

New Relic 是廣為人知的網站效能檢測服務，有提供免費方案足以應付小型網站的需求

而且所提供的 `new_relic_rpm` gem 可以在`development` 環境中在本機端直接查看效能分析

修改專案中的 `Gemfile` 透過 `bundle` 安裝

```ruby
gem 'newrelic_rpm'
```

執行指令產生設定檔 `config/newrelic.yml`

```shell
newrelic install --license_key="YOUR_KEY" "My application"
```

**YOUR_KEY** 是在 New Relic 官網取得的 License Key

**My application** 是自訂的應用程式名稱

只是想在開發環境測試時，或還沒在 New Relic 官網建立，以上資料可以先隨意填就好

接著執行 `rails s` 啟動專案，在瀏覽器輸入 `http://localhost:3000/newrelic`

就可以看到測試環境用的除錯頁面了

執行專案網站中的幾項操作，回到 `http://localhost:3000/newrelic` 就能夠看到剛剛操作的執行時間

接下來就可以針對較慢的部分調整程式碼了！

----

以上的 gem 結合了眾人的智慧，提供許多改善程式碼的參考

相信能夠在精進技術的路上提供不少幫助，多多利用吧！