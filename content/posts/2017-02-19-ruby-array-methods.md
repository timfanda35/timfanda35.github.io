---
categories:
  - ruby
comments: true
date: 2017-02-19T14:16:00+08:00
title: Ruby Array 的一些筆記
url: /2017/02/19/ruby-array-methods/
---

![](http://i.imgur.com/EsiorQx.jpg)

[圖片來源](https://www.pakutaso.com/20160628181pc-11.html)
[フリー写真素材ぱくたそ](https://www.pakutaso.com/)

紀錄一下我覺得值得一提的那些陣列兩三事

<!--more-->

## 給予陣列預設值

```ruby
> Array.new(3, "str")
=> ["str", "str", "str"]

> Array.new([1, 2, 3])
=> [1, 2, 3]

> Array.new(3) {|i| i * 3}
=> [0, 3, 6]

# 要注意的是用預設值，所有元素都會指向同一個參考位址
> arr = Array.new(3, "a")
=> ["a", "a", "a"]

> arr[0].next!
=> "b"

# 全部的預設值元素都變成 "b" 了
> arr
=> ["b", "b", "b"]
```

## 新增元素

```ruby
> arr = []
=> []

# 在結尾插入
> arr.push("Bear")
=> ["Bear"]

# 在結尾插入
> arr << "Candy"
=> ["Bear", "Candy"]

# 在開頭插入
> arr.unshift("Andy")
=> ["Andy", "Bear", "Candy"]

# 合併陣列
> arr.concat ["Other", "Another"]
=> ["Andy", "Alex", "Bear", "Dennis", "Emily", "Other", "Another"]

# 以上方法都會改變原本陣列的值

# 合併陣列
> arr = arr + ["Dennis", "Emily"]
=> ["Andy", "Alex", "Bear", "Dennis", "Emily"]
```

## 填滿某個值

```ruby
arr = [1, 2, 3, 4, 5]
=> [1, 2, 3, 4, 5]

# 全部歸零
arr.fill(0)
=> [0, 0, 0, 0, 0]

# 談入對應的索引
arr.fill { |index| index }
=> [0, 1, 2, 3, 4]
```

## 一次取得多個不同索引的元素

```ruby
> [1, 2, 3, 4, 5].values_at(1, 2, 3)
=> [2, 3, 4]
```

## 清除陣列中全部的 `nil`

```ruby
> arr = [1, nil, 2, nil, 3]
=> [1, nil, 2, nil, 3]

# 只是回傳清除 nil 後的結果，但不會改變陣列的值
> arr.compact
=> [1, 2, 3]
> arr
=> [1, nil, 2, nil, 3]

# 要用 compact! 才會改變陣列的值
> arr.compact!
=> [1, 2, 3]
> arr
=> [1, 2, 3]
```

## 一邊取出陣列的值一邊運算的寫法

`arr.pop` 也能改成 `app.shift`

```ruby
> arr = [1, 2, 3]
> while item = arr.pop
>   puts item
> end
3
2
1
=> nil

> arr
=> []
```

## 根據型態，呼叫 `*` 方法會有不同的結果

```ruby
> [ 1, 2, 3 ] * 3
=> [ 1, 2, 3, 1, 2, 3, 1, 2, 3 ]

# * 字串會跟 join 方法一樣結果
> [ 1, 2, 3 ] * ","
=> "1,2,3"
```

## [轉置矩陣](https://zh.wikipedia.org/wiki/%E8%BD%AC%E7%BD%AE%E7%9F%A9%E9%98%B5)

```ruby
# 矩陣
# | 1, 2, 3|
# | 4, 5, 6|
arr = [[1, 2, 3], [4, 5, 6]]
=> [[1, 2, 3], [4, 5, 6]]

# 旋轉後
# | 1, 4|
# | 2, 5|
# | 3, 6|
> arr.transpose
=> [[1, 4], [2, 5], [3, 6]

# 不會改變陣列的值
> arr
=> [[1, 2, 3], [4, 5, 6]]
```

## 求陣列長度

```
# 一般
> [1, nil, nil, 3, nil, 4].length
=> 6

# 取得非 nil  的元素個數
> [1, nil, nil, 3, nil, 4].nitems
=> 3
```

## 參考：

[Ruby Document: Array Class](https://ruby-doc.org/core-2.2.0/Array.html)

另外關於操作集合的技巧可以參考：

[集合的使用技巧(一) | Ruby](http://www.10tiao.com/html/568/201609/2650829566/1.html)

[集合使用的技巧(二) | Ruby](http://www.10tiao.com/html/568/201609/2650829568/1.html)

[集合使用的技巧(三) | Ruby](http://www.10tiao.com/html/568/201609/2650829570/1.html)