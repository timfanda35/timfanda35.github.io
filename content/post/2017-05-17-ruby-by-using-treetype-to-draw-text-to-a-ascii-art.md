---
categories:
  - ruby
  - freetype
  - ascii-art
comments: true
date: 2017-05-17T13:28:00+08:00
title: Ruby 透過 FreeType 產生點陣文字
url: /2017/05/17/ruby-by-using-treetype-to-draw-text-to-a-ascii-art/
---

![](http://i.imgur.com/EsiorQx.jpg)

[圖片來源](https://www.pakutaso.com/20160628181pc-11.html)
[フリー写真素材ぱくたそ](https://www.pakutaso.com/)

最近因為工作上的需求，需要將使用者輸入的文字轉成 16 x 16 的點陣文字

在網路上找到 jserv 這篇 [透過 FreeType 繪製 Unicode ASCII Art](http://blog.linux.org.tw/~jserv/archives/002050.html) 而有了實作的方向

<!--more-->

本文將使用 [ft2-ruby](https://github.com/customink/ft2-ruby) 從 TTF 字型讀取資料，以繪製點陣文字

本文所使用的字型：[源真ゴシック (げんしんゴシック)](http://jikasei.me/font/genshin/#zip__10)

## 安裝 ft2-ruby

```bash
gem install ft2-ruby
```

## 下載字型

從 [源真ゴシック (げんしんゴシック)](http://jikasei.me/font/genshin/#zip__10) 下載後，任一挑選一個字型檔案使用，本文選用 `GenShinGothic-Monospace-Regular.ttf`

## 範例程式

```ruby
#!/usr/bin/env ruby

require "ft2"

class Ft2Example
  def initialize(ttf_path)
    @face = FT2::Face.new ttf_path
    @face.set_pixel_sizes 16, 0
  end

  def draw
    glyph = @face.glyph

    for row in 0..(glyph.bitmap.rows - 1)
      for pixel in 0..(glyph.bitmap_left - 1)
        print "  "
      end

      for pixel in 0..(glyph.bitmap.width - 1)
        # pitch 是指一個 row 有多少 bytes
        # 跟 c/c++ 不一樣的是，ruby 的條件式是只有 nil 跟 false 會是 false
        # 所以這邊要設定條件 > 0
        index = row * glyph.bitmap.pitch + pixel / 8
        if glyph.bitmap.buffer[index].ord & (0xC0 >> (pixel % 8)) > 0
          print "██"
        else
          print "  "
        end
      end
      puts
    end
  end

  def draw_ch(ch)
    # 先取得索引
    ch_index = @face.char_index ch.ord

    # 設定現在要渲染的文字
    @face.load_glyph ch_index, FT2::Load::DEFAULT

    # 設定渲染模式
    @face.glyph.render_glyph FT2::RenderMode::MONO

    draw
  end

  def draw_string(str)
    str.split("").each do |ch|
      draw_ch(ch)
      
      puts
    end
  end
end

if ARGV.size < 2
  puts "Usage:"
  puts "    ruby main.rb TTF_PATH \"WORD\""

  puts

  puts "Example:"
  puts "    ruby main.rb ./GenShinGothic-Monospace-Regular.ttf 安安您好"

  exit 1
end

ft2_example = Ft2Example.new(ARGV[0])
ft2_example.draw_string(ARGV[1])
```

## 執行範例

```bash
ruby main.rb ./GenShinGothic-Monospace-Regular.ttf 安安您好
```

結果：

![Imgur](http://i.imgur.com/f99ibym.png)


如果 `ch_index = @face.char_index ch.ord` 結果是 `0`，表示在該字型檔中並沒有這個字的資料

只能試著換個字型檔繪製

## 參考

[透過 FreeType 繪製 Unicode ASCII Art](http://blog.linux.org.tw/~jserv/archives/002050.html) 

[FreeType简易教程](http://mingplusplus.com/tech/2014/09/13/freetype/)

[FreeType Development](https://www.freetype.org/developer.html)

[ft2-ruby](https://github.com/customink/ft2-ruby)

[FT2 document](http://www.rubydoc.info/gems/ft2-ruby/FT2)

[思源黑體修改版「思源真黑體」字型免費下載](https://free.com.tw/genshin-font/)

[源真ゴシック (げんしんゴシック)](http://jikasei.me/font/genshin/#zip__10)