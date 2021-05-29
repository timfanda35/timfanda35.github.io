---
categories:
  - ruby
  - heroku
comments: true
date: 2016-11-26T06:53:00+08:00
title: 'Side Project: Pocket to mail'
url: /2016/11/26/pocket-to-mail/
---

最近寫的小專案 [pocket_to_mail](https://github.com/timfanda35/pocket_to_mail)

主要原因是之前都是使用 [Pocket][pocket] 在收藏網路文章

然後在有空的時候再打開 [Pocket][pocket] 整理

不過每次要整理的時候看到一大堆文章就...懶了

所以決定寫一個小專案來解決這個問題

<!--more-->

## 起頭

這個專案很小，就是呼叫 API ，然後將連結寄到自己的信箱

所以我寫了一個簡單的 Ruby Script，使用 [pocket-ruby][pocket-ruby] gem 存取 [Pocket][pocket] API 然後透過 SMTP 將連結寄到自己的 Gmail

為了能夠每天排程執行，我將專案放在 [Heroku][heroku] 上面，並加上 [Heroku Scheduler][heroku_scheduler] Add-on，設定好每天固定時間執行

## 重構

有一天我發現信裡面的連結都不能點了，原來是某個 APP 分享過來的網址有透過轉址，而那個轉址的網站被 Gmail 認為不安全

於是趁這個機會我順便把網址後面的追蹤參數給刪掉，不然每次在將連結貼到筆記本時看到那些參數就很阿咂

剛開始我使用 URI 與 CGI 處理，但後來在寫測試的時候出現警告，看來是所使用的方法已經不建議再使用了。後來搜尋了一下改用 [addressable][addressable] gem，程式碼看起來也好讀很多

原本只是個 Ruby Script ，但最近剛好有一些時間，就來重構一下專案練練手，把方法拆一拆分給幾個類別處理

然後加上部分單元測試，用 Rack 包裝執行與測試指令

## 服務

掛上了 [Travis CI][travis_ci]、[Code Climate][code_climate]、[Hound CI][hound_ci]、[Deppbot][deppbot] 幾個服務，在 README 加上 Badges ~~滿足一下虛榮心~~

最後在 [Heroku][heroku] 專案的 Deploy 設定自動 deploy Github master branch。以後只要將程式碼 push 到 master 就會自動部署了！

## 結語

因為我每天都會清理信箱，所以每天處理五個連結還蠻輕鬆，日子過著過著就把 [Pocket][pocket] 裡累積的文章清空了

~~但自從我為了手機上的儲存空間把 app 給砍了之後，就幾乎沒再用到 [Pocket][pocket]了~~

[pocket]: https://getpocket.com/
[pocket-ruby]: https://github.com/turadg/pocket-ruby
[heroku]: https://www.heroku.com/
[heroku_scheduler]: https://devcenter.heroku.com/articles/scheduler
[addressable]: https://github.com/sporkmonger/addressable
[travis_ci]: https://travis-ci.org/
[code_climate]:https://codeclimate.com/dashboard
[hound_ci]: https://houndci.com/
[deppbot]: https://www.deppbot.com/