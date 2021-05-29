---
categories:
  - ruby
comments: true
date: 2017-01-04T09:46:00+08:00
title: Remove RVM
url: /2017/01/04/remove-rvm/
---

最近在整理筆記，慢慢會把一些東西搬上來

一開始學 Rails 的時候是用 [RVM][rvm] 在管理 ruby 版本，後來則改用了 [rbenv][rbenv]

兩者如果同時安裝在同一台電腦上的話就會有奇怪的問題發生

而且實際上兩者的主要功能都差不多

所以如果要使用 [rbenv][rbenv] 的話，記得要把[RVM][rvm] 砍掉

移除 RVM

```bash
$ rvm implode
$ rm -rf ~/.rvm
```

移除 `.bashrc` 與 `.bash_profile` 內 rvm 相關的東西 

[rvm]: https://rvm.io/
[rbenv]: https://github.com/rbenv/rbenv