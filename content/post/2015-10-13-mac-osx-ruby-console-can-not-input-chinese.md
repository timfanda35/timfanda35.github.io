---
categories:
  - ruby
  - rbenv
  - shell
comments: true
date: 2015-10-13T08:23:00+08:00
title: MAC OSX ruby console 無法輸入中文
url: /2015/10/13/mac-osx-ruby-console-can-not-input-chinese/
---

我是在 MAC OSX 上使用 rbenv 安裝 ruby，在換了新電腦重新安裝環境後突然發現在不管在 irb 還是 pry 中都無法輸入中文，而且操作上怪怪的。

在網上搜尋後找到了這篇 [mac 无法在irb中输入中文](http://blog.58share.com/?p=404) 

首先用 [Homebrew](http://brew.sh/index_zh-tw.html) 安裝 readline

```shell
brew install readline
```

<!--more-->

再輸入以下指令用 rbenv 重新安裝 ruby 

```shell
CONFIGURE_OPTS="--disable-install-doc --with-readline-dir=$(brew --prefix readline)" rbenv install 2.2.2
```

最後執行 irb 測試輸入中文，一切正常！