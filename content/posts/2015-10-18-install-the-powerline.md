---
categories:
  - shell
  - vim
comments: true
date: 2015-10-18T13:30:00+08:00
title: 安裝 Powerline
url: /2015/10/18/install-the-powerline/
---

因為注意到 [vim-powerling](https://github.com/Lokaltog/vim-powerline) 已經不再維護，而建議改用 [Powerline][powerline] 時，就想說來試試看。

在設定 Vim 的狀態列時沒什麼問題，但在設定 Shell 時就出現奇怪的情況了，好在後來有解決。

這篇紀錄我在 Mac OSX 安裝 [Powerline][powerline] 的步驟。

有任何錯誤或建議都歡迎告知:)

環境：
- OS: OSX EI Capitan version 10.11
- Shell: bash

<!--more-->

----

# 安裝 Powerline

## [Powerline][powerline]

> 美觀的狀態列

確認系統 python 版本是 `2.7`

```shell
$ python -V
Python 2.7.10
```

由於我還沒有安裝過 [pip](https://pypi.python.org/pypi/pip) ， python 安裝套件的指令，執行以下指令安裝：

```shell
$ sudo easy_install pip
```

用 `pip` 安裝 Powerline (由於權限問題，我無法安裝到系統目錄，所以我加上了 `--user` 選項，選擇安裝到使用者目錄底下)

```shell
$ pip install --user powerline-status
```

查看 Powerline 被安裝到哪個路徑

```shell
$ pip show powerline-status
--
Metadata-Version: 1.1
Name: powerline-status
Version: 2.2
Summary: The ultimate statusline/prompt utility.
Home-page: https://github.com/powerline/powerline
Author: Kim Silkebaekken
Author-email: kim.silkebaekken+vim@gmail.com
License: MIT
Location: /Users/bear/Library/Python/2.7/lib/python/site-packages
Requires:

# 以我的例子，我的使用者名稱是 bear 
# 所以會安裝到 /Users/bear/Library/Python/2.7/lib/python/site-packages 目錄底下
```

----

# 安裝 Powerline 字型

這樣狀態列才不會變成亂碼

```shell
$ git clone https://github.com/Lokaltog/powerline-fonts.git 
$ cd powerline-fonts/
$ ./install.sh
```

----

# 設定 Vim 狀態列

編輯 `~/.vimrc`，加入以下設定，(請將 {path} 替換成 `pip show powerline-status` 所顯示的路徑)

```
# 在所有 window 中都顯示狀態列
set laststatus=2

# powerline 的相關設定
set rtp+={path}/powerline/bindings/vim
let g:minBufExplForceSyntaxEnable = 1
python from powerline.vim import setup as powerline_setup
python powerline_setup()
python del powerline_setup
```

螢幕截圖

![Vim 狀態列](http://user-image.logdown.io/user/6114/blog/6123/post/304497/RNT48fxTVyacWvppbiLA_Screen%20Shot%202015-10-18%20at%206.27.38%20PM.png)

---

# 設定 Shell 狀態列

在這邊我遇到了一個詭異的情況，原本應該是 `scripts` 目錄下的檔案卻都跑到 Python 的 `bin` 目錄底下

```shell
# 原本應該有這個目錄，底下會有 powerline 相關的檔案
/Users/bear/Library/Python/2.7/lib/python/site-packages/scripts

# 但都跑到這目錄底下了
/Users/bear/Library/Python/2.7/bin
```

所以需要搬動檔案到正確的目錄，以我的例子：

以下 `{path}` 要換成執行`pip show powerline-status` 顯示的路徑

```shell
# 移動到安裝目錄
$ cd {path}

# 建立 scripts 目錄
$ mkdir scripts

# 將檔案移到 scripts 目錄下
$ mv ../../../bin/* ./scripts/
```

最後在 `~./.bash_profile` 中加入以下設定，重新開啟終端機就會出現了絢麗的提示了(以下 `{path}` 要換成執行`pip show powerline-status` 顯示的路徑)

```shell
# Prompt
POWERLINE_INSTALL_PATH={path}
POWERLINE_SCRIPT=$POWERLINE_INSTALL_PATH/powerline/bindings/bash/powerline.sh
if [ -f $POWERLINE_SCRIPT ]; then
  $POWERLINE_INSTALL_PATH/scripts/powerline-daemon -q
  source $POWERLINE_SCRIPT
fi
```

另外客製化 Shell 狀態列的樣式，可以參考[文件](http://powerline.readthedocs.org/en/latest/configuration.html)
(以下 `{path}` 要換成執行`pip show powerline-status` 顯示的路徑)

```shell
# 建立目錄
$ mkdir -p ~/.config/powerline/themes/shell

# 複製設定檔
$ cp {path}/powerline/config_files/themes/shell/default.json ~/.config/powerline/themes/shell/default.json

# 修改設定檔
$ vi ~/.config/powerline/themes/shell/default.json
```

以下是我的設定檔:

- 使用者名稱前加上符號
- 加入時間
- 加上 git branch 提示

```json
{
  "segments": {
    "left": [
      {
        "function": "powerline.segments.shell.mode"
      },
      {
        "function": "powerline.segments.common.net.hostname",
        "priority": 10
      },
      {
        "function": "powerline.segments.common.env.user",
        "before": "ⓑ ",
        "priority": 30
      },
      {
        "function": "powerline.segments.common.env.virtualenv",
        "priority": 50
      },
      {
        "function": "powerline.segments.common.time.date",
        "args": {
          "format": "%H:%M"
        },
        "priority": 10
      },
      {
        "function": "powerline.segments.shell.cwd",
        "priority": 10
      },
      {
        "function": "powerline.segments.shell.jobnum",
        "priority": 20
      },
      {
        "function": "powerline.segments.shell.last_pipe_status",
        "priority": 10
      },
      {
        "function": "powerline.segments.common.vcs.branch",
        "priority": 40
      }
    ]
  }
}
```

螢幕截圖

![我的 Shell 狀態列](http://user-image.logdown.io/user/6114/blog/6123/post/304497/4O354gUfS3ih9M1f06Tz_Screen%20Shot%202015-10-18%20at%206.12.04%20PM.png)

----

# Vim-airline

後來有在網路上查到另外一個 [vim-airline][airline]，與 [Powerline][powerline] 不同的是 [vim-airline][airline] 是單純的 vim-plugin，如果只有 Vim 需要絢麗的狀態列的話，[vim-airline][airline] 應該會是比較好的選擇。

- [vim-airline - 狀態列顯示的好幫手](http://terrychen.logdown.com/posts/2014/10/18/vim-airline-status-bar-helper)
- [Vim改变生活，插件改变Vim (2) — Vim插件详解](http://fancyseeker.com/?p=592)
- [vim-airline](https://blog.othree.net/log/2013/07/04/vim-airline/)

----

參考：

- [为Bash和VIM配置一个美观奢华的状态提示栏](http://cenalulu.github.io/linux/mac-powerline/)
- [Powerline Document](https://powerline.readthedocs.org/en/latest/installation/osx.html#vim-installation)
- [Powerline：漂亮的 Vim 狀態列與 Bash Shell 命令提示字串外掛](http://blog.gtwang.org/linux/powerline-adds-powerful-statuslines-and-prompts-to-vim-and-bash/)

[powerline]: https://github.com/powerline/powerline
[airline]: https://github.com/bling/vim-airline