---
categories:
  - ruby
keywords:
  - ruby
  - gem
  - dry
  - dry-rb
  - dry-initializer
comments: true
date: 2023-10-14T12:00:00+08:00
title: "dry-initializer 介紹"
url: /2023/10/14/dry-initializer-gem/
---

最近在研究 [Turbo](https://github.com/hotwired/turbo-rails) 與 [View Component](https://viewcomponent.org/)，在尋找資料的過程中從這部影片看到 [dry-initializer](https://dry-rb.org/gems/dry-initializer/) 的介紹。看起來能夠幫助在定義物件的時候少寫許多程式碼。

{{< youtube xBlw_vI8aPU >}}

## 簡介

[dry-rb](https://dry-rb.org/) 提供了許多方便的 gem 用來減少開發時重工的部分，而 dry-initializer 便是其中之一，透過 class methods 來簡化物件的 `initialize` method。

以往我們在定義物件的時候，如果需要給予初始值，可能會這樣定義：

```ruby
class Foo
  attr_reader :bar

  def initialize(bar)
    @bar = bar
  end
end
```

如果使用 dry-initializer，我們則可以改寫成：

```ruby
require 'dry-initializer'

class Foo
  extend Dry::Initializer

  param :bar
end
```

外部呼叫的程式碼不變。

```ruby
foo = Foo.new('BAR')
foo.bar # => BAR
```

## 用法

### Plain Argument

未使用 dry-initializer：

```ruby
class Foo
  attr_reader :bar, :zip

  def initialize(bar, zip)
    @bar = bar
    @zip = zip
  end
end
```

使用 dry-initializer：

```ruby
require 'dry-initializer'

class Foo
  extend Dry::Initializer

  param :bar
  param :zip
end
```

### Named (hash) Argument

未使用 dry-initializer：

```ruby
class Foo
  attr_reader :bar, :zip

  def initialize(bar:, zip:)
    @bar = bar
    @zip = zip
  end
end
```

使用 dry-initializer：

```ruby
require 'dry-initializer'

class Foo
  extend Dry::Initializer

  option :bar
  option :zip
end
```

### Default value

未使用 dry-initializer：

```ruby
class Foo
  attr_reader :bar, :zip

  def initialize(bar = 'BAR', zip: 'ZIP')
    @bar = bar
    @zip = zip
  end
end
```

使用 dry-initializer：

```ruby
require 'dry-initializer'

class Foo
  extend Dry::Initializer

  param :bar,  default: proc { 'BAR' }
  option :zip,  default: proc { 'ZIP' }
end
```

```ruby
Foo.new # => #<Foo:0x0000000106751208 @bar="BAR", @zip="ZIP">
```

可以指定 Named (hash) Argument 為 Optional，如果未給值則為 `Dry::Initializer::UNDEFINED`。

```ruby
require 'dry-initializer'

class Foo
  extend Dry::Initializer

  option :bar
  option :zip, optional: true
end
```

```ruby
foo1 = Foo.new(bar: 1, zip: 2)
foo1.bar # => 1
foo1.zip # => 2

foo2 = Foo.new(bar: 1)
foo2.bar # => 1
foo2.zip # => nil
foo2 # => #<Foo:0x0000000105d781f0 @bar=1, @zip=Dry::Initializer::UNDEFINED>
```

### Reader 與 Writer

我們可以定義 Argument 對應的 reader methods scope，`false` 表示不產生 reader method。

```ruby
require 'dry-initializer'

class Foo
  extend Dry::Initializer

  param :bar, reader: false
  param :zip, reader: :private   # the same as adding `private :name`
  param :app, reader: :protected # the same as adding `protected :email`
end
```

```ruby
foo = Foo.new('BAR', 'ZIP', 'APP')
foo.bar # => undefined method
foo.zip # => private method
foo.app # => protected method
```

dry-initializer 只有定義為 Argument 產生對應的 reader methods，我們需要自行依需要定義 attr_writer。

```ruby
require 'dry-initializer'

class Foo
  extend Dry::Initializer

  attr_writer :zip

  param :bar
  param :zip
end
```

```ruby
foo = Foo.new('BAR', 'ZIP')
foo.bar # => BAR
foo.zip # => ZIP

foo.bar = 'bar' # => undefined method
foo.zip = 'zip' # => 'zip'

foo.bar # => BAR
foo.zip # => zip
```

## 參考資料
- [Turbo](https://github.com/hotwired/turbo-rails)
- [View Component](https://viewcomponent.org/)
- [dry-initializer](https://dry-rb.org/gems/dry-initializer/)
