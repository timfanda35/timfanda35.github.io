---
categories:
  - elixir
comments: true
date: 2016-07-27T14:23:00+08:00
title: Elixir Char List
url: /2016/07/27/elixir-char-list/
---

最近除了 Ruby 以外，偶爾也有看看其他的語言

像是 [Elixir](http://elixir-lang.org/)

寫寫函數式語言轉轉腦筋也是挺不錯的

<!--more-->

## Char List

在 Elixir 中用雙引號 `"` 與用單引號 `'` 所包起來的文字會是不同的型態

不像 Ruby 最後還是 `String` 字串

用雙引號 `"` 包起來的文字，型態會是 `String` 字串

```erlang
$ iex
> is_bitstring "Bear"
true
> is_binary "Bear"
true
> is_list "Bear"
false
```

用單引號 `'` 包起來的文字，型態會是 `Char List` 字元列表

```erlang
$ iex
> is_bitstring 'Bear'
false
> is_binary 'Bear'
false
> is_list 'Bear'
true
```

有趣的是如果定義一個 `List`，其中元素是 32 ~ 126 之間的整數的話，就會被當作是 `Char List`，能夠直接印出字元，而不是列出 `List` 中的各個整數

```erlang
$ iex
> [66, 101, 97, 114]
'Bear'
> [31]
[31]
> [127]
[127]
```

另外如果想要知道字元對應的數字的話，可以透過 `?` 取得

```erlang
$ iex
> ?B
66
> [66]
'B'
```

經過以上的說明，就能夠輕鬆解出以下的題目：

[RNA Transcription in Elixir](http://exercism.io/exercises/elixir/rna-transcription/readme)


## 參考：

[Char Lists](http://elixir-lang.org/getting-started/binaries-strings-and-char-lists.html#char-lists)