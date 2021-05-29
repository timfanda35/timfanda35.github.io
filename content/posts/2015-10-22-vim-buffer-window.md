---
categories:
  - vim
comments: true
date: 2015-10-22T13:29:00+08:00
title: Vim 指令( buffer 與 window )
url: /2015/10/22/vim-buffer-window/
---

最近在看 [Vim Videos - Posted by Derek Wyatt on](http://derekwyatt.org/vim/tutorials/) 來學 Vim 的操作，這篇是我看完 [Vim Novice Video Tutorials](http://derekwyatt.org/vim/tutorials/novice/) 章節中所做的筆記

一般游標移動的指令，我在 [Vim 指令(游標移動)](posts/304816-vim-command) 這篇有紀錄過就不再贅述了

以下主要紀錄有關在 Vim 中操作 buffer 與 window 的指令

有任何錯誤或建議都歡迎告知:)

<!--more-->

----

# buffer

buffer 簡單來說就是目前 Vim 所開啟的檔案的編輯區

當我們在用 Vim 開啟檔案時，Vim 會將檔案內容讀入 buffer 之中，我們平常就是在編輯 buffer 中的資料

最後存檔時，執行 `w` ，就是將 buffer 的內容寫回去檔案中

以下指令可以透過執行 `:h buffer` 查看說明文件

# 顯示與切換 buffer

`:ls` 列出所有 buffer，`:ls!` 為連系統隱藏的 buffer 都列出來

`:b [N]` 切換到編號 `[N]` 的 buffer

`:bn` 切換到下一個 buffer，`:bn!` 為強制切換

`:bN` 切換到上一個 buffer，`:bN!` 為強制切換

`:bf` 切換到第一個 buffer，`:bf!` 為強制切換

`:bl` 切換到最後一個 buffer，`:bl!` 為強制切換

# 新增與刪除 buffer

`:badd [filename]` 將 `[filename]` 加入 buffer 列表中

`:bd [N]`、`:bd [filename]` 刪除編號為 `[N]` 或是檔名為 `[filename]` 的 buffer，`:bd!` 為強制刪除，若不帶任何參數則刪除目前正在編輯的 buffer

# 對所有 buffer 同時操作

`:bufdo [cmd]` 對所有 buffer 執行 `[cmd]` 指令，`:bufdo! [cmd]` 為強制執行

`:wall` 將所有 buffer 寫入變更(存檔)，`:wall!` 為強制寫入

`:qall` 關閉所有 buffer，`:qall!` 為強制關閉

# 切換 buffer 時可能會遇到的問題

一般如果從未存檔的 buffer 切換到另一個 buffer 時，會出現以下訊息

```shell
E37: No write since last change (add ! to override)
```

這是提醒你要儲存變更，以免遺失　buffer 中的資料，你在指令後方加上 `!` 強制切換

你也可以執行 `:set hidden` 或在 `~/.vimrc` 中設置這個設定，將非顯示的 buffer 隱藏起來

這樣就不會再跳出提醒，順利切換了

----

# window

可以在畫面中切出許多編輯視窗，用來查看、比對不同檔案，不用為開多個終端機(Terminal)排列傷腦筋

以下指令可以透過執行 `:h window` 查看說明文件

# 新增分割視窗

`:sp [filename]` 往上新增水平分割視窗，如未輸入 `[filename]`，新視窗的內容跟原視窗相同

`Ctrl+w s` 往上新增水平分割視窗，新視窗的內容跟原視窗相同

`:vs [filename]` 往左新增垂直分割視窗，如未輸入 `[filename]`，新視窗的內容跟原視窗相同

`Ctrl+w v` 往左新增垂直分割視窗，新視窗的內容跟原視窗相同

`:new` 新增空檔案的新水平分割視窗(往上新增)

`:vnew` 新增空檔案的新垂直分割視窗(往左新增)

# 在視窗之間移動

`Ctrl+w h` 游標移動到左方視窗

`Ctrl+w j` 游標移動到下方視窗

`Ctrl+w k` 游標移動到上方視窗

`Ctrl+w l` 游標移動到右方視窗

`Ctrl+w w` 游標移動到下一個視窗

`:close` 或 `Ctrl+w c` 關閉目前視窗

# 移動視窗位置

`Ctrl+w H` 將目前視窗移到左方

`Ctrl+w J` 將目前視窗移到下方

`Ctrl+w K` 將目前視窗移到上方

`Ctrl+w L` 將目前視窗移到右方

`Ctrl+w x` 與下一個視窗互換位置

`Ctrl+w Ctrl+r` 輪轉視窗，每個視窗往下一個位置移動，最後的視窗會到第一個位置

# 其他

`Ctrl+w o` 畫面只保留目前視窗，關閉其他視窗。如果有設定 `:set hidden`，則其他視窗會隱藏起來

----

# window 與 buffer 的操作

`:sb [N]` 新增視窗，新視窗的內容為 buffer 編號 `[N]`

`:sbn` 新增視窗，新視窗的內容為下一個 buffer

`:sball` 將所有 buffer 的內容變為視窗顯示在畫面上

以上指令的結果都是水平分割，視窗上下排列，若要垂直分割的話，可以在指令前面加上 `:vert`

例如 `:vert ball` 就是將所有 buffer 的內容變為垂直分割視窗，左右排列顯示在畫面上

----

# 參考

[Vim Videos - Posted by Derek Wyatt on](http://derekwyatt.org/vim/tutorials/)

[Vim - 編輯緩衝區以及編輯視窗 (Buffers and Windows)](http://www.openfoundry.org/tw/tech-column/2383-vim--buffers-and-windows)

[Vim 101: Set Hidden](http://usevim.com/2012/10/19/vim101-set-hidden/)