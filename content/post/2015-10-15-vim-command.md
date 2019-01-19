---
categories:
  - vim
comments: true
date: 2015-10-15T06:05:00+08:00
title: Vim 指令(游標移動)
url: /2015/10/15/vim-command/
---

最近重新學習 vim，紀錄一下覺得好用的指令。

# 移動

----

### 單字間移動

`W` 移動到下一個字首(以空格作分隔)

`w` 移動到下一個字首

`E` 移動到下一個字尾(以空格作分隔)

`e` 移動到下一個字尾

`B` 移動到上一個字首(以空格作分隔)

`b` 移動到上一個字首

<!--more-->

### 頁面移動

`Ctrl h` 移到螢幕最上面

`Ctrl l` 移到螢幕最下面

`Ctrl f` 下一頁

`Ctrl b` 上一頁

`Ctrl d` 上半頁

`Ctrl u` 下半頁

### 單字間跳躍

`%` 移動到對應括號

`*` 移動到下一個相同單字

`#` 移動到上一個相同單字

# 縮排

----

`>>` 增加縮排

`<<` 減少縮排

# 分頁

----

`:Texplore` 開啟新分頁

`gt` 下一頁

`gT` 上一頁

`{i} gt` {i} 是數字，到指定頁，例：`5gt` 就是到第五頁

`:tabs` 顯示所有分頁

`:tabclose {i}` {i} 是數字，關閉指定頁，例： `:tabclose 5` 就是關閉第五頁

# 其他

----

`cw` 刪除目前單字

`:! {command}` 執行 shell 指令 {command}，例： `:! pwd`

`q:` 顯示歷史指令

`Ctrl [` 相當於 `Esc`

----

# 參考：

[技巧：快速提高 Vi/Vim 使用效率的原则与途径](https://www.ibm.com/developerworks/cn/linux/l-cn-tip-vim/)

[简明 Vim 练级攻略](http://coolshell.cn/articles/5426.html)

[无插件Vim编程技巧](http://coolshell.cn/articles/11312.html)

[所需即所获：像 IDE 一样使用 vim](https://github.com/yangyangwithgnu/use_vim_as_ide/blob/master/README.md)

[高見龍 Category: Vim](http://blog.eddie.com.tw/category/vim/)

[學Vim的啟示](http://enginechang.logdown.com/posts/296611-learn-something-from-learning-vim)

[Learn Vimscript the Hard Way](http://learnvimscriptthehardway.stevelosh.com)

[Don’t tell people to use VIM (because) You’re Using It Wrong](http://antjanus.com/blog/thoughts-and-opinions/use-vim/)

[@be_vimmer_tw | Twitter](https://twitter.com/be_vimmer_tw)

[Vim Videos - Posted by Derek Wyatt on](http://derekwyatt.org/vim/tutorials/)