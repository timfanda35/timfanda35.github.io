---
categories:
  - ruby
comments: true
date: 2017-02-05T05:58:00+08:00
title: ruby-refinement
url: /2017/02/05/ruby-refinement/
---

今天整理收集的文章時看到了這篇 [Ruby Best Practices and Tips by Toptal Developers](https://www.toptal.com/ruby/tips-and-practices?utm_medium=sponsored_post&utm_campaign=ruby_tips&utm_source=facebook.com)

發現裡面的範例有到用 refinement

雖然之前知道有這個語法，但看了範例才知道該怎麼用

<!--more-->

當想要為已存在的類別新增或覆寫方法時

常常會用 monkey patch 的方式

但這樣該類別在執行環境中就整個就被改掉了

monkey patch 若沒有規範好，會難以追蹤除錯

為了將在 `String` 擴充的方法限制在 `MyLogClass` scope 中

範例使用了 refinement

在 `MyLogClass` scope `裡面，`String` 才有 `red` 跟 `green` 方法可以使用

以外的 scope 中，`String` 還是原來沒有被改過的類別

這樣就可以控制擴充功能會出現的地方

```ruby
module ConsoleColorsStringRefinement
  refine String do
    def red
      "\e[31m#{self}\e[0m"
    end

    def green
      "\e[32m#{self}\e[0m"
    end
  end
end

class MyLogClass
  using ConsoleColorsStringRefinement

  def log(message, level)
    puts "%s: %s" % [level.to_s.red, message.to_s.green]
  end
end

MyLogClass.new.log('My Message', 'DEBUG') 
# DEBUG: My Message

'My Message'.to_s.red
# NoMethodError: undefined method `red' for "My Message":String
```

<blockquote class="imgur-embed-pub" lang="en" data-id="8qq6lOM"><a href="//imgur.com/8qq6lOM">View post on imgur.com</a></blockquote><script async src="//s.imgur.com/min/embed.js" charset="utf-8"></script>