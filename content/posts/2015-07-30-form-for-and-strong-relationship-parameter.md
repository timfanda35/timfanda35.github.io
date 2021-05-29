---
categories:
  - ruby
  - rails
comments: true
date: 2015-07-30T13:01:00+08:00
title: form_for 與 strong parameter 的關係
url: /2015/07/30/form-for-and-strong-relationship-parameter/
---

`form_for` 是用 Rails 做網頁很常用的 Helper 方法，他的方便之處在於很輕鬆的產生與 Model 對應的 HTML 表單。

大部份的時候都會將使用者輸入的資料，用 Hash 表示 Model 的方式傳送到伺服器上。

但是有兩個問題：

* 伺服器直接依照傳入的參數更新資料，如果有心人士偷塞了 Model 有的欄位，資料就會被改掉。
* 親自寫判斷去過濾傳入的參數，如果 Model 的種類一多，欄位一多就會變得很麻煩。

這邊可以用 `strong parameter` 來很簡單地過濾掉，傳入的 Model 參數中含有非預期的參數。

假設 Foo 定義如下，其中 `security_field` 欄位並不想被隨意修改，所以沒有呈現在表單上。

```ruby
# schema
create_table "foos", force: :cascade do |t|
  t.string   "name"
  t.integer  "bar"
  t.string   "security_field"
  t.datetime "created_at", null: false
  t.datetime "updated_at", null: false
end
```

但有心人士仍然可以發送這樣的請求給伺服器，試圖去修改資料庫的值。

```json
{
  "id": 1,
  "foo": {
    "name": "foo",
    "bar": 1,
    "security_field": "bad value"
  }
}
```

在這邊最簡單的方法就是用 `strong parameter` 來過濾，在 `Controller` 中寫一個 private method，要更新 Modle 時就使用該 method 來取得過濾後的參數來更新。

```ruby
class FoosController < ApplicationController
  ...
  
  private
  
  def foo_params
    # 只允許通過 :name, :bar 參數
  	params.require(:foo).permit(:name, :bar)
  end
end
```

如果有參數包含關聯 Modle，即巢狀表單，會寫成如下：

```ruby
# 只允許通過 :name, :bar 以及關聯的 member :id :name 參數
# 記得要允許關聯 Model 的 :id 參數，不然會無法正確更新
params.require(:foo).permit(:name, :bar, member_attributes:[:id, :name])
```

## 參考資料

[Rails102](http://rocodev.gitbooks.io/rails-102/content/chapter1-mvc/c/strong_parameters.html)
[Rails Guide](http://guides.rubyonrails.org/action_controller_overview.html#strong-parameters)