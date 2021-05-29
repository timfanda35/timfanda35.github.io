---
categories:
  - ruby
comments: true
date: 2017-03-05T13:33:00+08:00
title: Ruby Time and Date
url: /2017/03/05/ruby-time-and-date/
---

![](http://i.imgur.com/EsiorQx.jpg)

[圖片來源](https://www.pakutaso.com/20160628181pc-11.html)
[フリー写真素材ぱくたそ](https://www.pakutaso.com/)

內容是我在準備認證考試時，看文件所留下的一些紀錄

<!--more-->

## Time

`Time.new` 跟 `Time.now` 一樣，都是回傳現在時間

`Time` 相減會是 **Float** 類型

`Time + 1`  是加一秒

### 取得時間欄位

```ruby
> time = Time.new
=> 2017-03-05 21:36:21 +0800

> time.year
=> 2017

> time.month
=> 3

> time.day
=> 5
```

### 建立 Time

```ruby
> Time.mktime(2010)
=> 2010-01-01 00:00:00 +0800

> Time.mktime(2010, 7)
=> 2010-07-01 00:00:00 +0800

> Time.mktime(2010, 7, 7)
=> 2010-07-07 00:00:00 +0800
```

### 從字串轉成 Time

```ruby
> require "time"
=> true

> Time.parse("2010-7-7")
=> 2010-07-07 00:00:00 +0800
```

### 格式化時間

```ruby
> time = Time.new
=> 2017-03-05 21:36:21 +0800

> time.strftime("%Y-%m-%d")
=> "2017-03-05"

> require "time"
=> true

> time.rfc2822 # 符合 email header 的 format
=> "Sun, 05 Mar 2017 21:45:25 +0800"

> time.ios8601

> time.utc # 轉成 utc 時間，破壞性方法
=> 2017-03-05 13:45:25 UTC

> time.localtime # 轉成當地時間，破壞性方法
=> 2017-03-05 21:45:25 +0800
```

## Date

使用 Date 需要 `require "date"`

`Date` 相減會是 **Rational** 類型

`Date + 1`  是加一天

`Date >> 1` 是加一個月，`Date << 1` 是減一個月

### 建立 Date

```ruby
> require "date"
=> true

> Date.today
=> #<Date: 2017-03-05 ((2457818j,0s,0n),+0s,2299161j)>

> Date.new(2013, 2, -1) # 2013/2/28
=> #<Date: 2013-02-28 ((2456352j,0s,0n),+0s,2299161j)>

> Date.new(2016, 2, -1) # 2013/2/29
=> #<Date: 2016-02-29 ((2457448j,0s,0n),+0s,2299161j)>

> Date.parse("2010-7-7")
=> #<Date: 2010-07-07 ((2455385j,0s,0n),+0s,2299161j)>
```

### 日期遞增與遞減

```ruby
> require "date"
=> true

> date = Date.today
=> #<Date: 2017-03-05 ((2457818j,0s,0n),+0s,2299161j)>

> date >> 1 # 下個月
=> #<Date: 2017-04-05 ((2457849j,0s,0n),+0s,2299161j)>

> date << 1 # 上個月
=> #<Date: 2017-02-05 ((2457790j,0s,0n),+0s,2299161j)>
```

### 格式化時間

```ruby
> require "date"
=> true

> date = Date.today
=> #<Date: 2017-03-05 ((2457818j,0s,0n),+0s,2299161j)>

> date.strftime("%Y-%m-%d %H:%M:%S") # 小時以下的時間皆為 0
=> "2017-03-05 00:00:00"
```

## DateTime

基本上與 `Date` 大同小異，但會紀錄到小時以下的時間

```ruby
> require "date"
=> true

> datetime = DateTime.now # 用 now 方法取得現在時間
=> #<DateTime: 2017-03-05T22:00:33+08:00 ((2457818j,50433s,289029000n),+28800s,2299161j)>

> datetime.strftime("%Y-%m-%d %H:%M:%S")
=> "2017-03-05 22:00:33"
```

## 格式化參考

|符號|說明|
|----|----|
|%A	| Sunday, Monday, ... |
|%a	| Sun, Mon, ... |
|%B	| January, February, ... |
|%b	| Jan, Feb, ... |
|%c	| 時間 |
|%d	| 日 (01 - 31) |
|%H	| 24 小時 (00 - 23) |
|%l	| 12 小時 (01 - 12) |
|%j	| 一年中的第幾天 (001 - 366) |
|%M	| 分 (00 - 59) |
|%m	| 月 (01 - 12) |
|%p	| AM, PM |
|%S	| 秒 (00 - 60) |
|%U	| 第幾週，從星期日開始 (00 - 53) |
|%W	| 第幾週，從星期一開始 (00 - 53) |
|%w	| 星期幾，星期日開始 (0 - 6) |
|%X	| 24時：分：秒 |
|%x	| 月/日/年 |
|%Y	| 西元年 |
|%y	| 西元年後兩位 |
|%Z	| 時區 |
|%%	| % |

## 參考

因為我準備的是 ruby 2.1 認證，所以以下連結都是 2.1 版的文件

[Time Ruby document(core)](http://ruby-doc.org/core-2.1.0/Time.html)
[Time Ruby document(stdlib)](https://ruby-doc.org/stdlib-2.1.0/libdoc/time/rdoc/Time.html)
[Date Ruby document(stdlib)](https://ruby-doc.org/stdlib-2.1.0/libdoc/date/rdoc/Date.html)
[DateTime Ruby document(stdlib)](https://ruby-doc.org/stdlib-2.1.0/libdoc/date/rdoc/DateTime.html)
