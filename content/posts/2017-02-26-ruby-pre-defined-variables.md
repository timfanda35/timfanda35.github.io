---
categories:
  - ruby
comments: true
date: 2017-02-26T07:41:00+08:00
title: Ruby Pre-defined variables
url: /2017/02/26/ruby-pre-defined-variables/
---

![](http://i.imgur.com/EsiorQx.jpg)

[圖片來源](https://www.pakutaso.com/20160628181pc-11.html)
[フリー写真素材ぱくたそ](https://www.pakutaso.com/)

ruby 本身有先定義了一些變數與常數
 
可以用在取得系統資訊，或是除錯上

有些變數可以使得在開發時少打一些程式碼，但相對的可讀性可能就會降低一些(不是所有人都熟悉符號的意義)
 
這裡紀錄一些我覺得值得一提的部分

<!--more-->

## 正規表示式比對

|變數名稱|說明|
|----|----|
|$~ |最後比對到的結果，MatchData 類別|
|$& |完整比對到的字串|
|$` |比對到的字串之前的字串|
|$' |比對到的字串之後的字串|
|$1, $2, ... |比對到的群組，這裡是從 1 開始|
|$+ |最後比對到的群組|

例子：

```ruby
# 比對字串
> /b(c)(d)e/ =~ "abcdef"
=> 1

# 最後比對到的結果
> $~
=> #<MatchData "bcde" 1:"c" 2:"d">

# 完整比對到的字串
> $&
=> "bcde"

# 比對到的字串之前的字串
> $`
=> "a"

# 比對到的字串之後的字串
> $'
=> "f"

# 比對到的群組
> $1
=> "c"
> $2
=> "d"
> $3
=> nil

# 最後比對到的群組
> $+
=> "d"
```

## 檔案相關

`$0` 是目前執行 ruby 的檔案名稱，假設用 `irb` 印出 `$0`

```ruby
> $0
=> "irb"
```
若是執行 `test.rb` 

```bash
bash$ echo 'puts $0' > test.rb
bash$ ruby test.rb
test.rb
```

`DATA` 這個比較常用到在 pure ruby script，它是用來取得在檔案 `__END__` 以下的字串

```ruby
# test.rb

# 注意 DATA 是 File 類別，所以可以用 `read` 取得內容
puts DATA.read

__END__
I have a pen.
I have an apple.
```

執行 `ruby test.rb`

```ruby
bash$ ruby test.rb
I have a pen.
I have an apple.
```

`ENV` 這個有用 `rails` 的開發者應該不陌生，可以用來取得環境變數

## 錯誤

`$!` 取得最後 raise 的 Exception 訊息

`$@` 取得最後的 Exception 錯誤推疊陣列

```ruby
# test.rb

begin
  raise StandardError.new "Mayday"
rescue => e
  puts $!
  puts $@
end
```

執行 `ruby test.rb`

```ruby
bash$ ruby test.rb
Mayday
test.rb:2:in `<main>'
```

## 參考

[ruby-doc globals_rdoc](https://ruby-doc.org/core-2.1.1/doc/globals_rdoc.html)
[ruby-doc Regexp](https://ruby-doc.org/core-2.1.1/Regexp.html#class-Regexp-label-Special+global+variables)