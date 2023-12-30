---
categories:
  - rails
keywords:
  - rails
  - ruby
  - test
  - cucumber
  - gem
comments: true
date: 2023-10-07T12:00:00+08:00
title: "Cucumber Getting Start Note"
url: /2023/10/07/cucumber-getting-start-note/
images:
  - /images/2023-10-07/cucumber-getting-start-note.png
---

## 建立 Rails 專案

本文使用的環境：
- macOS
- Ruby 3.2.2
- Rails 7.0.8

執行指令建立 Rails 專案

```bash
rails new demo
```

設定工作目錄，或是使用編輯器開啟 Rails 專案

```shell
cd demo
```

## 安裝 Cucumber Rails

編輯 `Gemfile`，於 `test` group 中新增 `cucumber-rails` 與 `database_cleaner`。

```shell
group :test do
  # Use system testing [https://guides.rubyonrails.org/testing.html#system-testing]
  gem "capybara"
  gem "selenium-webdriver"
  gem 'cucumber-rails', require: false
  gem 'database_cleaner'
end
```

執行指令安裝

```shell
bundle install
```

執行 Cucumber Generator

```shell
rails g cucumber:install
```

Cucumber Generator 在 Rails 專案中新增或修改的檔案如下：

![rails files diff](/images/2023-10-07/20231007001.png)

## 執行 cucumber

```shell
rails cucumber
```

![execute cucumber](/images/2023-10-07/20231007002.jpg)

因為我們什麼都還沒做，所以結果會是空的，而且會出現一個提示框。

## 其他設定

預設 Cucumber Rails 會執行 `DatabaseCleaner.start` 與 `DatabaseCleaner.clean`，可以在 `features/support/env.rb` 停用。

```ruby
Cucumber::Rails::Database.autorun_database_cleaner = false
```

可以在 `config/cucumber.yml` 停用 `publish` 不再顯示提示框。

```ruby
default: <%= std_opts %> --publish-quiet features
```

![execute cucumber](/images/2023-10-07/20231007003.jpg)

## Demo

我們用一個範例來體驗一下使用 Cucumber 的流程。

測試案例：
1. 資料庫有兩筆測試資料
2. 當訪問 Article 首頁時，頁面標題為 "Article"
3. 且頁面上資料的顯示順序為倒序排列

### 新增 Scenario

新增檔案 `features/article.feature`，依照 [gherkin](https://cucumber.io/docs/gherkin/reference/) 語法撰寫。

```gherkin
Feature: Article
  Scenario: When user access index then get article list
    Given there are some articles
      | title | content |
      | post1 | Hello   |
      | post2 | Hi      |
    When I make a GET request to "/articles"
    Then the body has title "Articles"
    And the body has the articles sort by desc
      | title | content |
      | post2 | Hi      |
      | post1 | Hello   |
```

### 新增 Steps

Rails 預設包含 Capybara，所以我們可以使用 [Capybara DSL](https://github.com/teamcapybara/capybara#the-dsl) 來與頁面互動。我們會用 `Given`、`When`、`Then` 等方法去從 `.feature` 檔案中的敘述解析出需要的值，在測試過程中使用。

新增檔案 `features/step_definitions/article_steps.rb`。

```ruby
Given(/^there are some articles$/) do |table|
  # table is a table.hashes.keys # => [:title, :content]
  table.hashes.each do |hash|
    Article.create!(hash)
  end
end

When(/^I make a GET request to "(.+)"$/) do |path|
  visit path
end

Then(/^the body has title "([^"]*)"$/) do |title|
  assert_equal title, page.find('h1').text
end

And(/^the body has the articles sort by desc$/) do |table|
  expect_titles = table.hashes.map { |hash| hash['title'] }
  titles = page.all('div:not([id="articles"]) > p:first').map { |elm| elm.text.gsub('Title:', '').strip }
  assert_equal expect_titles, titles
end
```

測試

```shell
rails cucumber
```

![execute cucumber](/images/2023-10-07/20231007004.jpg)

可以看到顯示找不到 Article 的錯誤，這是因為我們還沒實作任何程式邏輯。

### 實作程式邏輯

執行指令建立 Article Controller、Model 與 View。

```shell
rails g scaffold article title:string content:string
```

執行指令更新資料庫 Schema。

```shell
rails db:migrate
```

測試

```
rails cucumber
```

![execute cucumber](/images/2023-10-07/20231007005.jpg)

從結果可以發現 Scenario 最後一個條件資料倒序的部分尚未被滿足，我們再修改一下程式。

編輯 `app/controllers/articles_controller.rb`，修改 `index` 方法。

```ruby
  def index
    @articles = Article.all.order(created_at: :desc)
  end
```

再次執行測試

```
rails cucumber
```

![execute cucumber](/images/2023-10-07/20231007006.jpg)

成功滿足 Scenario 的需求。

## 參考資料
- [Install Cucumber](https://cucumber.io/docs/installation/ruby/)
- [cucumber-rails](https://github.com/cucumber/cucumber-rails)
- [gherkin reference](https://cucumber.io/docs/gherkin/reference/)
- [Capybara DSL](https://github.com/teamcapybara/capybara#the-dsl)
- [可以測試的規格 - Rails 開發實踐](https://blog.aotoki.me/posts/2023/07/28/rails-in-practice-testable-specification/)
