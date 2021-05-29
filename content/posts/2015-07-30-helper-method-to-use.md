---
categories:
  - ruby
  - rails
comments: true
date: 2015-07-30T13:37:00+08:00
title: helper 的使用方法
url: /2015/07/30/helper-method-to-use/
---

Helper 是協助在 View 呈現資料的方法，通常都用來包裝程式邏輯，或是簡化撰寫 View 的作業。

Helper 檔案都會放在 `app/helpers` 底下，預設就有一個 `application_helper.rb` 的檔案。

一份 Helper 檔案範例如下，要注意是定義一個 **Module**，而非 **Class**：

```ruby
module FooHelper
  def bar
    ...
  end
end
```

之後就能在 View 中使用 `bar` 方法了。

如果想將 Controller 的方法變成 Helper 方法讓 View 呼叫，可以用 `helper_method` 方法達成：

```ruby
class ApplicationController < ActionController::Base
  # define helper method
  helper_method :foo

  protected

  def foo
    ...
  end
end
```

至於想要在 rails console 測試 helper ：

> 因為 Helper 預設只能在 Template 中使用，如果想在 rails console 中呼叫，必須加上 helper，例如 `helper.link_to`。另外，雖然機會不多，如果真的要在 Rails Controller 或 Model 程式中呼叫 Helper，則可以加上 `ApplicationController.helpers` 前置詞。 - __*Ruby on Rails 實戰聖經*__

最後推薦瀏覽參考資料 **如何運用 / 設計 Rails Helper** 系列，文章提到很多使用 Helper 的經驗與技巧。

## 參考資料

[Rails102](http://rocodev.gitbooks.io/rails-102/content/chapter1-mvc/v/helper.html)
[Ruby on Rails 實戰聖經](https://ihower.tw/rails4/actionview-helpers.html)
[BLOG.XDITE.NET 如何運用 / 設計 Rails Helper (1)](http://blog.xdite.net/posts/2011/12/09/how-to-design-helpers)
[BLOG.XDITE.NET 如何運用 / 設計 Rails Helper (2)](http://blog.xdite.net/posts/2011/12/09/how-to-design-helpers)
[BLOG.XDITE.NET 如何運用 / 設計 Rails Helper (3)](http://blog.xdite.net/posts/2011/12/09/how-to-design-helpers)
