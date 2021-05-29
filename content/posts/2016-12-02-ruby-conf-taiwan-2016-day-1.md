---
categories:
  - ruby
comments: true
date: 2016-12-02T15:51:00+08:00
title: RubyConf Taiwan 2016 第一天
url: /2016/12/02/ruby-conf-taiwan-2016-day-1/
---

今年第一次參加 [RubyConf Taiwan 2016](https://2016.rubyconf.tw/)

紀錄今天聽到的議程：

- [Beware of Alpha Syndrome](https://2016.rubyconf.tw/#Yukihiro (Matz) Matsumoto)
- [Breaking Bad Habits with GitLab CI](https://2016.rubyconf.tw/#Ivan Nemytchenko)
- [Optimizing ruby core](https://2016.rubyconf.tw/#Shyouhei Urabe)
- [How to write complex data pipelines in Ruby](https://2016.rubyconf.tw/#Kazuyuki Honda) | [簡報](https://speakerdeck.com/hakobera/how-to-write-complex-data-pipeline-in-ruby)
- [It's More Fun to Compute](https://2016.rubyconf.tw/#Julian Cheal)
- [How I made a pure-Ruby word2vec program more than 3x faster](https://2016.rubyconf.tw/#Kei Sawada) | [簡報](https://speakerdeck.com/remore/how-i-made-a-pure-ruby-word2vec-program-more-than-3x-faster)
- [How to Write Middleware in Ruby](https://2016.rubyconf.tw/#Satoshi Tagomori) | [簡報](http://www.slideshare.net/tagomoris/how-to-write-middleware-in-ruby)

之後如有提供投影片再補充連結

<!--more-->

## [Beware of Alpha Syndrome](https://2016.rubyconf.tw/#Yukihiro (Matz) Matsumoto)

以人為本

過辛福的生活

讓世界更美好

期待更快更強大的 Ruby 3 

## [Breaking Bad Habits with GitLab CI](https://2016.rubyconf.tw/#Ivan Nemytchenko)

介紹 [GitLab CI](https://about.gitlab.com/gitlab-ci/) 除了測試還可以做到哪些事情

像是執行多個指令、打包檔案提供下載、部署應用程式

而且這些設定都很簡單

## [Optimizing ruby core](https://2016.rubyconf.tw/#Shyouhei Urabe)

這一場比較硬，我沒能好好聽懂 Orz

## [How to write complex data pipelines in Ruby](https://2016.rubyconf.tw/#Kazuyuki Honda)

[簡報](https://speakerdeck.com/hakobera/how-to-write-complex-data-pipeline-in-ruby)

在處理大量資料時，常常會需要很多步驟去做資料處理

在開發時，每個步驟的邏輯又常常有很多重複的事情要做，譬如：retry, logging...等等

講者因此開發了 [tumugi](https://github.com/tumugi/tumugi) gem 透過 DSL 的方式來減少這些重複的開發工作

在這場看到有趣的 gem:

[retriable](https://github.com/kamui/retriable) 透過 DSL 的方式為你的程式加入 retry

[tumugi](https://github.com/tumugi/tumugi) 透過 DSL 的方式開發複雜的 workflows

## [It's More Fun to Compute](https://2016.rubyconf.tw/#Julian Cheal)

不同的波型能夠有不同的聲音效果，透過程式來產生波型訊號，就能產生各種不同的聲音

在這場看到有趣的軟體：

[sonic pi](http://sonic-pi.net/) 寫 code 來玩音樂，酷吧！

## [How I made a pure-Ruby word2vec program more than 3x faster](https://2016.rubyconf.tw/#Kei Sawada)

[簡報](https://speakerdeck.com/remore/how-i-made-a-pure-ruby-word2vec-program-more-than-3x-faster)

Ruby 在處理大迴圈運算時，執行時間會變得很長

可以透過 *VirtualModule* 執行其他程式語言來解決

在這場看到有趣的語言：

[julia](http://julialang.org/) 一種動態程式語言

## [How to Write Middleware in Ruby](https://2016.rubyconf.tw/#Satoshi Tagomori)

[簡報](http://www.slideshare.net/tagomoris/how-to-write-middleware-in-ruby)

這場因為臨時有 issue 要處理，中途就沒有好好聽了 >"<