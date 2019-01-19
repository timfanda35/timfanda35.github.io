---
categories:
  - ruby
comments: true
date: 2015-10-14T08:26:00+08:00
title: Ruby 的 %Q, %q, %W, %w, %x, %r, %s, %i
url: /2015/10/14/ruby-percent-literals/
---

有時候在看其他人的 Ruby 程式碼時，常常發現有 `%w`, `%r` 等 % 開頭的程式碼。上網查了一下，才知道這些叫做 [Percent literals(或Percent Strings)](http://ruby-doc.org/core-2.2.2/doc/syntax/literals_rdoc.html#label-Percent+Strings)。

以下內容是改寫自 [%Q, %q, %W, %w, %x, %r, %s](https://simpleror.wordpress.com/2009/03/15/q-q-w-w-x-r-s/) 與 [Ruby 里的 %Q, %q, %W, %w, %x, %r, %s, %i (译)](https://ruby-china.org/topics/18512)。

執行環境：

- ruby: 2.2.2
- console: irb

## %Q
----

代替雙引號字串，如果在字串中有很多雙引號的時候很好用

```ruby
>> what_frank_said = 'Hello!'
=> "Hello!"

>> %Q(Joe said: "Frank said: "#{what_frank_said}"")
=> "Joe said: \"Frank said: \"Hello!\"\""

>> puts %Q(Joe said: "Frank said: "#{what_frank_said}"")
Joe said: "Frank said: "Hello!""
=> nil
```

<!--more-->

也可以省略 Q

在 [Ruby Style Guide](https://github.com/bbatsov/ruby-style-guide#percent-literals) 建議使用 `%()`，因為比 `%Q()` 短

```ruby
>> what_frank_said = 'Hello!'
=> "Hello!"

>> %(Joe said: "Frank said: "#{what_frank_said}"")
=> "Joe said: \"Frank said: \"Hello!\"\""

>> puts %(Joe said: "Frank said: "#{what_frank_said}"")
Joe said: "Frank said: "Hello!""
=> nil
```

## %q
----

代替單引號字串，如果在字串中有很多單引號的時候用

```ruby
>> %q(Joe said: 'Frank said: '#{what_frank_said} ' ')
=> "Joe said: 'Frank said: '\#{what_frank_said} ' '"

>> puts %q(Joe said: 'Frank said: '#{what_frank_said} ' ')
Joe said: 'Frank said: '#{what_frank_said} ' '
=> nil
```


## %W
----

產生用雙引號包起來的字串陣列(用空白做分隔，空白需要用 \ 符號跳過)

```ruby
>> %W(a b c\ d \#e #{1}f)
=> ["a", "b", "c d", "#e", "1f"]
```

## %w
----

產生用單引號包起來的字串陣列(用空白做分隔，空白需要用 \ 符號跳過)

可以觀察與使用 `%W` 出來的差異

```ruby
>> %w(a b c\ d \#e #{1}f)
=> ["a", "b", "c d", "\\#e", "\#{1}f"]
```

## %x
----

使用``方法執行一段 shell script 並回傳標準輸出的結果

```ruby
>> %x(echo foo:#{'Foo'})
=> "foo:Foo\n"
```

## %r
----

產生正規表示式

```ruby
>> %r(/home/#{'Foo'})
=> /\\/home\\/Foo/
```

## %s 
----

產生 symbol，但不會對表示式等內容作轉換

```ruby
>> %s(foo)
=> :foo
>> %s(foo bar)
=> :"foo bar"
>> %s(#{foo} bar)
=> :"\#{foo} bar"
```

## %i
----

產生 symbol 陣列

```ruby
>> %i(a b c)
=> [:a, :b, :c]
```

---

除了用 `()` 也可以用其他非數字字母的符號或是成對的符號代替：

```ruby
%Q(...) 
%Q[...]
%Q!...!
%Q+...+
%Q{...}
%Q-...-
```

參考： 

[%Q, %q, %W, %w, %x, %r, %s](https://simpleror.wordpress.com/2009/03/15/q-q-w-w-x-r-s/)

[Ruby 里的 %Q, %q, %W, %w, %x, %r, %s, %i (译)](https://ruby-china.org/topics/18512)

[Ruby : General Delimited Input](http://jokercatz.blogspot.tw/2015/05/ruby-general-delimited-input.html)

[Ruby Doc - Percent Strings](http://ruby-doc.org/core-2.2.2/doc/syntax/literals_rdoc.html#label-Percent+Strings)

[Ruby Style Guide Percent Literals](https://github.com/bbatsov/ruby-style-guide#percent-literals)

[Ruby Style Guide 百分比字面](https://github.com/JuanitoFatas/ruby-style-guide/blob/master/README-zhTW.md#百分比字面)
