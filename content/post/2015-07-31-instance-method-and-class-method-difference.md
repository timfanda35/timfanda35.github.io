---
categories:
  - ruby
comments: true
date: 2015-07-31T06:02:00+08:00
title: instance method 與 class method 的差別
url: /2015/07/31/instance-method-and-class-method-difference/
---

簡單來說，instance method 需要一個實體物件才能呼叫，而 class method 則是可以直接用類別呼叫

```ruby
class Foo
  # 宣告 class method 需要在 method 名稱前加上 self.
  def self.class_bar
    # ...
  end
  
  def instance_bar
    # ...
  end
end

# class method 的呼叫方式
Foo.class_bar

# instance_bar 的呼叫方式
Foo.new.instance_bar
```

由以上範例程式可以看到，`instance_bar` 方法需要 `Foo.new` 先建立一個實體物件，才能夠呼叫到 `instance_bar` 這個 instance method。換句話說類別是無法呼叫 instance method，而實體變數也無法呼叫 class method

另一種宣告 class method 的方法

```ruby
class Foo
  class << self
    def class_bar1
      #...
    end
    
    def class_bar2
      #...
    end
  end
end
```

Monkey patch instance method 與 class method 的方法

```ruby
class Foo; end

Foo.class_eval do
  def self.class_bar_patch
    # ...
  end
  
  def instance_bar_patch
    # ...
  end
end	

# class method 的呼叫方式
Foo.class_bar_patch

# instance_bar 的呼叫方式
Foo.new.instance_bar_patch
```

## 參考資料

[Rails 102](http://rocodev.gitbooks.io/rails-102/content/chapter3-ruby/instance_method_class_method.html)

[RailsTips](http://www.railstips.org/blog/archives/2009/05/11/class-and-instance-methods-in-ruby/)