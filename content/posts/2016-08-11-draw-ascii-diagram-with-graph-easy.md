---
categories:
  - shell
comments: true
date: 2016-08-11T09:15:00+08:00
title: Draw ASCII Diagram with Graph-Easy
url: /2016/08/11/draw-ascii-diagram-with-graph-easy/
---

因為想要用 ASCII 畫圖，上網搜尋了一番

需求是可以用指令產生，不依賴 java

最後找到了 [graph-easy](https://github.com/ironcamel/Graph-Easy)

這是 perl 的模組，可以用來繪製 ASCII Diagram，也能匯出成許多格式

<!--more-->

# Install

這模組依賴 [graphviz](http://www.graphviz.org/) 所以需要先安裝，在 OSX 可以用 Homebrew 安裝

```bash
brew install graphviz
```

再來要透過 cpan 安裝此模組，於指令列輸入

```bash
cpan
```
cpan 是 perl 的模組安裝工具， OSX 或是 Linux 上都有預設安裝

如果還沒有用過 cpan 的話，輸入後就一直 YES 下去，直到再次跳出輸入提示

```bash
cpan[1]> install Graph::Easy
```

如果沒有權限的話，就先輸入 exit 退出 cpan，用 sudo 轉成 root 身份執行 cpan 安裝

```
sudo cpan
cpan[1]> install Graph::Easy
```

# Usage

詳細文件請參考：[http://bloodgate.com/perl/graph/manual/](http://bloodgate.com/perl/graph/manual/)

在 [https://github.com/ironcamel/Graph-Easy/tree/master/t/txt](https://github.com/ironcamel/Graph-Easy/tree/master/t/txt) 有許多範例

主要透過 `graph-easy source.txt` 指令來產生圖形

註解使用 `#` ，但記得要接個一個空白再寫內容

```
#不OK
# OK
```

node 用 `[]` 表示，預設裡面的內容是 label 的值

```
[A]

# 產生

+---+
| A |
+---+
```

如果要產生兩個同個標籤的 node，就需要用到 attribute 來指定標籤

用 `{}` 包起來的就是 attribute，跟在 node 與 edge 後面就能夠設定

node 是用 `[]` 包起來的內容來區分不同的 node

```
[0] { label: A; }
[1] { label: B; }

[0] -- [1]

# 產生

+---+     +---+
| A | --- | B |
+---+     +---+
```

你也可以讓 node 看起來長一點，不過這需要多一點關聯才會看到效果

```
[0] { label: A; rows: 3; }
[1] { label: B; }

[0] --> [1]
[1] --> [0]

# 產生

+---+     +---+
|   | --> | B |
| A |     +---+
|   |       |
|   | <-----+
+---+
```

edge 有很多種樣式

```
->              
=>              
.>              
~>              
- >             
.->             
..->            
= >             
```

選擇圖形的方向，預設是 east

```
[0] { label: A; flow: south; }
[1] { label: B; }

[0] -- [1]

# 產生

+---+
| A |
+---+
  |
  |
  |
+---+
| B |
+---+
```

如果想要預設所有 node 或是 edge 的屬性，可以用 `graph {}` 跟 `edge {}` 設定

```
graph {
  flow: south;
}

edge {
  label: go;
}

[0] { label: A; }
[1] { label: B; }

[0] -- [1]

# 產生

+-----+
|  A  |
+-----+
  |
  | go
  |
+-----+
|  B  |
+-----+
```

groups

```
( group name
  [A]
)

# 產生

+ - - - - - - -+
' group name   '
'              '
' +----------+ '
' |    A     | '
' +----------+ '
'              '
+ - - - - - - -+
```

# Reference

[ASCII Art：使用纯文本流程图](http://weishu.me/2016/01/03/use-pure-ascii-present-graph/)

[Graph-Easy](https://github.com/ironcamel/Graph-Easy)

[Graph::Easy - Manual](http://bloodgate.com/perl/graph/manual/)

