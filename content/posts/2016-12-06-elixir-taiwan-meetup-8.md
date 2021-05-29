---
categories:
  - elixir
comments: true
date: 2016-12-06T13:50:00+08:00
title: 'Elixir Taiwan Meetup #8'
url: /2016/12/06/elixir-taiwan-meetup-8/
---

最近正在上 [The Complete Elixir and Phoenix Bootcamp](https://www.udemy.com/the-complete-elixir-and-phoenix-bootcamp-and-tutorial/learn/v4/overview)

對我這初學者而言還是有蠻多觀念上的收穫

說到這才想到資料夾裡還有之前買的電子書還沒看


[Elixir.tw][elixirtw]  雖然還在建構中，但可以加入[Slack](https://slack.com/) 聊天了唷

<!--more-->

## [Exercism](http://www.exercism.io/)

這次開場練習的是 [Wordy](http://exercism.io/exercises/elixir/wordy/readme)

我還沒有寫完，努力中 XD

## [Routing around Plug](https://christopheradams.gitlab.io/plug_1_3/#/)

這次 [Plug][plug] Contributor [Christopher Adam](https://christopheradams.io/) 分享了 [Plug][plug] Routing 與新版 [Plug][plug] 的新功能

介紹了 [Plug][plug] 三大部分

- Plug.Conn: 用來描述請求與回應的資料結構
- Plug.Builder: 用來建立 Pipeline
- Plug.Router: 提供透過 DSL 的方式來建立路由

你可以像是在寫 [Phoenix](http://www.phoenixframework.org/) 的 router 一樣，在[Plug][plug] 1.3 你可以定義這樣的路由：

```
get "/world", to: MyApp.HelloPlug, init_opts: []
```

另外講者還有分享：

[plug_rest](https://github.com/christopheradams/plug_rest)

[The Elixir Style Guide](https://github.com/christopheradams/elixir_style_guide)

## [Building Elixir App Release with Distillery and Docker](http://www.slideshare.net/mickey1985/building-elixir-app-release-with-distillery-and-docker)

Micky 分享了他使用 [distillery][distillery] deploy [Elixir.tw][elixirtw] 時所遇到的坑：

在過程中出現 `Syntax on (` 這樣的錯誤，原因是建置與執行的平台不同的關係

在建置發佈的檔案時是用 Mac OSX，但 Server 的 OS 是跑 Ubuntu

所以無法正常執行

解決方法是使用 [Docker](https://www.docker.com/) 來建置發佈的檔案

Dockfile 在這： [Dockfile](https://github.com/elixirtw/elixir_tw/blob/master/Dockerfile)

另外分享了 podcast [The Elixir Fountain](http://elixirfountain.com/)

[plug]: https://github.com/elixir-lang/plug
[elixirtw]: http://elixir.tw/
[distillery]: https://github.com/bitwalker/distillery