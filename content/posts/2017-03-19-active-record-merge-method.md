---
categories:
  - rails
  - activerecord
comments: true
date: 2017-03-19T05:13:00+08:00
title: ActiveRecord merge method
url: /2017/03/19/active-record-merge-method/
---

![](http://i.imgur.com/EsiorQx.jpg)

[圖片來源](https://www.pakutaso.com/20160628181pc-11.html)

之前查資料找到這篇 [Rails where() clause on model association?](http://stackoverflow.com/questions/6334537/rails-where-clause-on-model-association)

原來在查詢條件包含關聯表欄位的時候還有 `merge` 的方法可以用

之前都在 `joins` 裡面硬刻 SQL query...

<!--more-->

`merge` 可以合併查詢子句，要做依賴關聯表條件的查詢時非常有用

以 [ActiveRecord::SpawnMethods#merge API Dock](http://apidock.com/rails/ActiveRecord/SpawnMethods/merge) 的例子來說：「想要找出至少有一筆非垃圾留言的文章」

首先建立一個新 rails 專案

```bash
$ rails new test_active_record_merge
$ cd test_active_record_merge
```

新增 Migragion file

```bash
$ bundle exec rails new start
```

編輯 Migragion file

```ruby
class Start < ActiveRecord::Migration[5.0]
  def change
    create_table :posts do |t|
      t.boolean :published, default: false
    end

    create_table :comments do |t|
      t.boolean :spam, default: false
      t.references :post, index: true, foreign_key: true
    end
  end
end
```

執行 `migrate`

```bash
$ bundle exec rake db:migrate
```

新增 Model file

```ruby
# app/models/post.rb
class Post < ApplicationRecord
  has_many :comments
end

# app/models/comment.rb
class Comment < ApplicationRecord
  belongs_to :post
end
```

在 Terminal 中執行 `rails c` 建立測試資料

```ruby
Comment.create(spam: true, post: Post.create(published: true))
Comment.create(spam: true, post: Post.create(published: false))
Comment.create(spam: false, post: Post.create(published: true)) #符合條件
Comment.create(spam: false, post: Post.create(published: false))
Post.create(published: true)
Post.create(published: false)
```

查詢「至少有一筆非垃圾留言的文章」

```ruby
Post.where(published: true).joins(:comments).merge(Comment.where(spam: false))

# 相當於
# Post.where(published: true).joins('INNER JOIN comments ON comments.post_id = posts.id').where('comments.spam = ?', false)
```

結果找到一篇文章

```ruby
  Post Load (0.7ms)  SELECT "posts".* FROM "posts" INNER JOIN "comments" ON "comments"."post_id" = "posts"."id" WHERE "posts"."published" = ? AND "comments"."spam" = ?  [["published", true], ["spam", false]]
=> #<ActiveRecord::Relation [#<Post id: 3, published: true>]>
```

但以上測試資料文章都只有一個留言，如果為符合條件的文章再加上一個非垃圾留言後再查詢

```ruby
Comment.create(spam: false, post: posts.first)
posts = Post.where(published: true).joins(:comments).merge( Comment.where(spam: false))
```

會發現找到兩筆結果，但卻是重複的文章

這是因為 `SQL INNER JOIN` 的關係，如果我們再為該文章加上一筆新的非垃圾留言，就會變成三筆重複的文章

```ruby
  Post Load (0.1ms)  SELECT "posts".* FROM "posts" INNER JOIN "comments" ON "comments"."post_id" = "posts"."id" WHERE "posts"."published" = ? AND "comments"."spam" = ?  [["published", true], ["spam", false]]
=> #<ActiveRecord::Relation [#<Post id: 3, published: true>, #<Post id: 3, published: true>]>
```

這時候可以加上 `distinct` 來排除重複的資料

```ruby
Post.where(published: true).joins(:comments).merge(Comment.where(spam: false)).distinct
```

結果

```ruby
  Post Load (0.2ms)  SELECT DISTINCT "posts".* FROM "posts" INNER JOIN "comments" ON "comments"."post_id" = "posts"."id" WHERE "posts"."published" = ? AND "comments"."spam" = ?  [["published", true], ["spam", false]]
=> #<ActiveRecord::Relation [#<Post id: 3, published: true>]>
```

## 參考

[Rails where() clause on model association?](http://stackoverflow.com/questions/6334537/rails-where-clause-on-model-association)
[ActiveRecord::SpawnMethods#merge API Dock](http://apidock.com/rails/ActiveRecord/SpawnMethods/merge)
[ActiveRecord::SpawnMethods#joins API Dock](http://apidock.com/rails/ActiveRecord/QueryMethods/joins)