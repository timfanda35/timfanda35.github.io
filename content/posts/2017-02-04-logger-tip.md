---
categories:
  - ruby
  - rails
  - logger
comments: true
date: 2017-02-04T07:05:00+08:00
title: logger tip
url: /2017/02/04/logger-tip/
---

所謂有記有保庇

log 能夠幫助瞭解程式在執行的時候做了什麼事

log 記得好在除錯的時候能幫上很大的忙

這邊紀錄跟 logger 的一些相關技巧

<!--more-->

## Rotate

為了避免過多的 log 佔據硬碟空間

在 Linux 上可以透過 logrotate 保留部分的 log

參考資料：

[使用 logrotate 定期整理 Rails Log 檔案] (https://ihower.tw/blog/archives/3565)

[Rotating Rails Production Logs with LogRotate] (https://gorails.com/guides/rotating-rails-production-logs-with-logrotate)

範例：

```
sudo vi /etc/logrotate.conf
```

```
/home/deploy/app/current/log/*.log {
  daily
  missingok
  rotate 7
  compress
  delaycompress
  notifempty
  copytruncate
}
```

## lograge

[lograge](https://github.com/roidrage/lograge) 是可以精簡 Rails log 的一個 gem

將原本輸出的 log

```
Started GET "/" for 127.0.0.1 at 2012-03-10 14:28:14 +0100
Processing by HomeController#index as HTML
  Rendered text template within layouts/application (0.0ms)
  Rendered layouts/_assets.html.erb (2.0ms)
  Rendered layouts/_top.html.erb (2.6ms)
  Rendered layouts/_about.html.erb (0.3ms)
  Rendered layouts/_google_analytics.html.erb (0.4ms)
Completed 200 OK in 79ms (Views: 78.8ms | ActiveRecord: 0.0ms)
```

變成

```
method=GET path=/jobs/833552.json format=json controller=JobsController  action=show status=200 duration=58.33 view=40.43 db=15.26
```

設定方式很簡單

在 `Gemfile` 中加上以下設定，並執行 `bundle install` 安裝

```
gem "lograge"
```

再將以下設定加到 `config/initializers/lograge.rb` 或是 `config/environments/production.rb`

```
## MyApp 改成你專案的名稱
MyApp::Application.configure do
  config.lograge.enabled = true
end
```

## 自訂 Rails Logger

要將 log 導向其他系統儲存，推薦以下兩篇閱讀：

[Rails Log Process](https://ruby-china.org/topics/27523)

[Rails Log 格式化 + Kids日志收集](https://ruby-china.org/topics/29417)

## Logger

Ruby 本身就有 `Logger` 類別，只要在程式中 `require` 就可以用了

[document](http://ruby-doc.org/stdlib-2.1.0/libdoc/logger/rdoc/Logger.html)

預設格式如下：

```
SeverityID, [DateTime #pid] SeverityLabel -- ProgName: message
```

輸出會長得像這樣：

```
I, [1999-03-03T02:34:24.895701 #19074]  INFO -- Main: info.
```

可以修改時間的顯示格式：

```
logger.datetime_format = '%Y-%m-%d %H:%M:%S'
# e.g. "2004-01-03 00:54:26"
```

也能自訂 log 顯示格式

```
logger.formatter = proc do |severity, datetime, progname, msg|
  "#{datetime}: #{msg}\n"
end
# e.g. "2005-09-22 08:51:08 +0900: hello world"
```

另外可以透過初始化傳入的參數，指定 log 檔案的儲存限制

```
logger = Logger.new('foo.log', 10, 1024000) # 保留 10 份日誌檔，每個檔案最大 1024000 bytes
logger = Logger.new('log.log', 'daily')     # 按天產生
logger = Logger.new('log.log', 'weekly')    # 按周產生
logger = Logger.new('log.log', 'monthly')   # 按月產生
```

如果你想紀錄錯誤的 stack 可以用以下方式：

```
logger.error err.backtrace.join("\n")
```

將日誌同時輸出到不同位置：

參考： [How can I have ruby logger log output to stdout as well as file?](http://stackoverflow.com/questions/6407141/how-can-i-have-ruby-logger-log-output-to-stdout-as-well-as-file)

```
class MultiIO
  def initialize(*targets)
     @targets = targets
  end

  def write(*args)
    @targets.each {|t| t.write(*args)}
  end

  def close
    @targets.each(&:close)
  end
end

log_file = File.open("log/debug.log", "a")
Logger.new MultiIO.new(STDOUT, log_file)
```