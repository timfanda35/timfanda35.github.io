---
categories:
  - rubymine
keywords:
  - rubymine
  - IDEA
  - JetBrains
  - IDE
  - Auto Completion
comments: true
date: 2023-09-01T12:00:00+08:00
title: "RubyMine Auto Completion Freeze"
url: /2023/09/01/rubymine-auto-completion-freeze/
images:
  - /images/2023-09-01/rubymine-auto-completion-freeze.png
---

## 問題

[RubyMine](https://www.jetbrains.com/ruby/features/) 是 [JetBrains](https://www.jetbrains.com/) 旗下提供 Ruby 開發環境的 IDE。

在使用 RubyMine 進行開發時，IDE 會根據輸入的程式碼提供 Auto Completion 的備選清單。

我的作業系統是 macOS，但今天不幸遇到，當 IDE 彈出備選清單後，整個視窗便卡住了。只能叫出 Activity Monitor 強制關閉。即使我重開了好幾次，只要彈出備選清單就會出現一樣卡住的情況。

## 解決過程

### 停用 ML completion 或 EditorConfig support (失敗)

網路搜尋到類似標題的問題 [IDEA sometimes complete hangs indefinitely when opening autocomplete modal](https://youtrack.jetbrains.com/issue/IDEA-280095/IDEA-sometimes-complete-hangs-indefinitely-when-opening-autocomplete-modal)

停用 ML completion
![](/images/2023-09-01/202309010001.png)

停用 EditorConfig support

![](/images/2023-09-01/202309010002.png)

但結果是失敗的

### 清除 RubyMine 快取 (成功)

網路搜尋到類似標題的問題 [WebStorm 2023.2 hangs](https://intellij-support.jetbrains.com/hc/en-us/community/posts/12843522069778-WebStorm-2023-2-hangs)

依照文中的回覆，我刪除了 RubyMine 的 Logs，並重開 RubyMine 但沒有作用。

注意：
- 以下路徑 `<macOS 使用者名稱>` 需要換成您 macOS 使用者名稱
- 以下路徑 `RubyMine2023.2` 需要換成您目前使用的版本

```
/Users/<macOS 使用者名稱>/Library/Logs/JetBrains/RubyMine2023.2/
```

我再往下看到另一個回覆說可以刪除 RubyMine 的 `caches`。

因為我也不了解是哪一份檔案有問題，所以我是刪除 RubyMine `caches` 目錄下的所有檔案

注意：
- 以下路徑 `<macOS 使用者名稱>` 需要換成您 macOS 使用者名稱
- 以下路徑 `RubyMine2023.2` 需要換成您目前使用的版本

```
/Users/<macOS 使用者名稱>/Library/Caches/JetBrains/RubyMine2023.2/caches/*
```

重新啟動 RubyMine 就恢復正常運作了

## 參考資料

- [IDEA sometimes complete hangs indefinitely when opening autocomplete modal](https://youtrack.jetbrains.com/issue/IDEA-280095/IDEA-sometimes-complete-hangs-indefinitely-when-opening-autocomplete-modal)
- [WebStorm 2023.2 hangs](https://intellij-support.jetbrains.com/hc/en-us/community/posts/12843522069778-WebStorm-2023-2-hangs)