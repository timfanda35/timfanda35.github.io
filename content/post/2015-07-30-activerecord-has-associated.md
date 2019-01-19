---
categories:
  - ruby
  - rails
comments: true
date: 2015-07-30T11:59:00+08:00
title: ActiveRecord has_many :through 關聯
url: /2015/07/30/activerecord-has-associated/
---

在 Rails 中 Model 常常使用到一對多的關係，像是一台車有四個輪子，但每個輪子都只屬於一台車

我們會用 `has_many` ， `belongs_to` 來定義 Model 之間的關係，在 `Tire` 的資料表需要有 `car_id` 欄位來儲存所屬車子的 ID。

```ruby
class Car
  has_many :tires
end

class Tire
  belongs_to :car
end
```

而多對多的關係就會用到 `has_many :through`，取個例子：在學校裡，一個社團可以有多名成員，每位成員也可以參加多個社團，這樣的關係就是「多對多關係」。

```plain
# 一個社團可以有多名成員
group
|-- member
|-- member
`-- member

# 每位成員可以參加多個社團
member
|-- group
|-- group
`-- group
```

我們無法只靠兩張表來表現多對多的關係，所以這時候需要再新增一個 Model，透過(:through)這一個 Model 來建立多得多的關係。

```plain
# group1 的成員有 member1, member2, member3
# group2 的成員有 member3
# group3 的成員有 member3

group1 -- group_member -- member1
       |- group_member -- member2
       |
       `- group_member -- member3
                        |
group2 -- group_member -|
group3 -- group_member -'

# 反過來說
# member1 參加了 group1
# member2 參加了 group1
# member3 參加了 group1, group2, group3
```

以上圖為例，Model 定義如下：

```ruby
class Group
  has_many :group_members
  has_many :members, through: :group_members
end

class GroupMember
  belongs_to :group
  belongs_to :members
end

class Member
  has_many :group_members
  has_many :groups, through: :group_members
end
```

`Group` 與 `Member` 要先與 `GroupMember` 建立好一對多的關係，才能透過 `GroupMember` 建立彼此多對多的關係。

在這邊可以看到 `GroupMember` 對於其他兩個 Model 的關聯是用 `belongs_to` 來定義，所以要記得在資料表中建立 `group_id` 與 `member_id` 欄位。

參考指令：

```shell
rails g migration CreateGroup name:string
rails g migration CreateGroupMember group_id:integer member_id:integer
rails g migration CreateMember name:string
```

## 參考資料

[Rails102](http://rocodev.gitbooks.io/rails-102/content/chapter1-mvc/m/has-many-through.html)
[Rails Guide](http://guides.rubyonrails.org/association_basics.html#the-has-many-through-association)