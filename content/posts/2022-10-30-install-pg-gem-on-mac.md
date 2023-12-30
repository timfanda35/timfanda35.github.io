---
categories:
  - ruby
keywords:
  - ruby
  - gem
  - macOS
  - pg
  - postgresql
comments: true
date: 2022-10-30T12:00:00+08:00
title: "在 Mac 上安裝 pg gem"
url: /2022/10/30/install-pg-gem-on-mac/
images:
  - /images/2022-10-30/install-pg-gem-on-mac.png
---

想在 Mac OS 上安裝 pg gem 但不想要安裝整套 PostgreSQL。

`pg` gem 在 Mac OS 上需要 native build，依賴 `libpg` 套件。

## 用 Homebrew 安裝 libpq

```shell
brew install libpq
```

在安裝過程中 Summary 會顯示 libpg 安裝的位置。

```shell
==> Summary
🍺  /opt/homebrew/Cellar/libpq/15.0: 2,366 files, 28.5MB
```

我們將該位置紀錄成環境變數

```shell
LIBPG_PATH=/opt/homebrew/Cellar/libpq/15.0
```

## 自訂建置參數安裝 pg gem

```shell
gem install pg -- \
	--with-pg-include=$LIBPG_PATH/include \
	--with-pg-lib=$LIBPG_PATH/lib
```
