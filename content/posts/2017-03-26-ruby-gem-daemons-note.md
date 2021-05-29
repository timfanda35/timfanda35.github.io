---
categories:
  - ruby
comments: true
date: 2017-03-26T08:04:00+08:00
title: Ruby gem daemons 筆記
url: /2017/03/26/ruby-gem-daemons-note/
---

![](http://i.imgur.com/ZjslPbL.jpg)

[圖片來源](https://www.pakutaso.com/20170137024post-10119.html)
[フリー写真素材ぱくたそ](https://www.pakutaso.com/)

[daemons gem][daemons] 是一個能夠讓 ruby 程式以行程(process)獨立執行的 gem

當你想要有一個程式能夠在背景持續執行時，就可以可慮使用 [daemons gem][daemons] 來協助開發

<!--more-->

什麼是 daemon？可以參考鳥哥的文章 [17.1 什麼是 daemon 與服務 (service)](http://linux.vbird.org/linux_basic/0560daemons.php#daemon) 

簡單來說可以當作是在背景一直執行的服務

透過 [daemons gem][daemons] 你可以：
1. 將其他 script 作為 daemon 執行 
2. 將 proc 丟到 daemon 執行 
3. 將目前執行的程式變成 daemon 執行

## 寫個範例

範例程式(`main.rb`)是一個可以用在程式中建立 **每 5 秒將 Process 名稱寫入 Log** 的 daemon：

```ruby
# main.rb
require 'daemons'
require "logger"

# 在執行的時候先取得當前目錄的絕對路徑
DIR = File.absolute_path(File.dirname(__FILE__))

class TittleDaemon
  def run(name)
    # 修改行程 (Process) 顯示的名稱
    Process.setproctitle(name)

    logger = Logger.new(log_path(name), 'daily')

    loop do
      logger.info "daemon #{name} is running"
      sleep(5)
    end
  end

  def log_path(name)
    File.join(DIR, "#{name}-running.log")
  end
end

class Main
  def start(name)
    options = {
      :ARGV                 => ['start'], # 啟動 daemon
      :backtrace            => true # 如果有發生例外會記錄到 log 中
    }
    Daemons.run_proc(name, options) { TittleDaemon.new.run(name) }
  end

  def stop(name)
    options = {
      :ARGV               => ['stop'], # 停止 daemon
      :backtrace          => true # 如果有發生例外會記錄到 log 中
    }
    Daemons.run_proc(name, options) { TittleDaemon.new.run(name) }
  end
end
```

操作：

首先在 Terminal 中開啟 `irb`，並加上 `-r` 參數在啟動 irb 的同時 `require ./main.rb`

```bash
$ irb -r ./main.rb 
```

在 irb 中建立 `Main` 的實體，並啟動一個名為 `little-daemon` 的 daemon

```ruby
> main = Main.new
> main.start("little-daemon")
```

![Imgur](http://i.imgur.com/zjbm1MX.png)

開啟另一個 Terminal 執行

```bash
$ ps aux | grep little
```

![Imgur](http://i.imgur.com/MiADDw1.png)

發現有一個名為 `little-daemon` 的行程正在執行


回到第一個 Terminal，這次我們請啟動另一個名為 `little-daemon2` 的 daemon

```ruby
> main.start("little-daemon2")
```

![Imgur](http://i.imgur.com/6Lmdl7o.png)


在第二個 Terminal 執行以下指令可以觀察到 `little-daemon2` 也在執行了

```bash
$ ps aux | grep little
```

![Imgur](http://i.imgur.com/xUn2Gzm.png)

看看目錄底下有什麼：

```bash
$ ls -lh
```

![Imgur](http://i.imgur.com/FLyCAyh.png)

可以發現除了預計會產生的 log 之外，還多了兩個結尾為 `.pid` 的檔案

用文字編輯器可以打開，裡面是記錄 daemon 的 PID

透過 PID 我們可以傳送訊號給 daemon ，像是立即停止，或是做其他運用

可以參考鳥哥文章 [16.1 什麼是程序 (process)](http://linux.vbird.org/linux_basic/0440processcontrol.php#whatis)

最後我們回到第一個 Terminal 停止剛剛啟動的 daemon

```ruby
> main = Main.new
> main.stop("little-daemon")
> main.stop("little-daemon2")
```

![Imgur](http://i.imgur.com/pk8hCA3.png)

在第二個 Terminal 執行以下指令會發現 daemon 已經停止，結尾為 `.pid` 的檔案也已經被刪除了

```bash
$ ps aux | grep little
$ ls -lh
```

![Imgur](http://i.imgur.com/G5IycHI.png)

## 尾聲

想要更了解 [daemons][daemons] 的話，推薦玩玩 GitHub 上的 [範例](https://github.com/thuehlinger/daemons/tree/master/examples)

還有看看程式碼裡的[註解](https://github.com/thuehlinger/daemons/blob/master/lib/daemons.rb)

了解一下傳入的 `options` 有哪些選項可以使用

要注意的是程式碼一直都有在更新，記得在閱讀的時候要將 branch 調成目前所使用的版本才不會錯亂喲

![Imgur](http://i.imgur.com/gTTPWT2.png)

## 參考

[Ruby daemons gem official repository](https://github.com/thuehlinger/daemons)

[歡迎光臨鳥哥的 Linux 私房菜](http://linux.vbird.org/)

[daemons]: https://github.com/thuehlinger/daemons