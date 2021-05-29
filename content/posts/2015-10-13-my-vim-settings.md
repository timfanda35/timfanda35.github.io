---
categories:
  - vim
comments: true
date: 2015-10-13T08:01:00+08:00
title: My Vim 設定
url: /2015/10/13/my-vim-settings/
---

紀錄一下我的 vim 設定，還在學習中，所以會不定期更新。

----

# 安裝 Vundle

## install dependency in ubuntu

```shell
sudo apt-get install git curl
```

## Install Vundle

```shell
git clone https://github.com/gmarik/vundle.git ~/.vim/bundle/vundle
```

[Vundle：Vim Plugin 自動下載、安裝、更新與管理工具（Vim Bundle)][link1]

<!--more-->

----

# Configure Vim

編輯 `~/.vimrc`

```config
set nocompatible              " be iMproved, required
set noswapfile
set encoding=utf-8

filetype off                  " required

" set the runtime path to include Vundle and initialize
set rtp+=~/.vim/bundle/vundle/
call vundle#rc()

" let Vundle manage Vundle, required
Plugin 'gmarik/vundle'

Plugin 'scrooloose/nerdtree'
Plugin 'vim-ruby/vim-ruby'
Plugin 'tpope/vim-rails'
Plugin 'tpope/vim-bundler'
Plugin 'tpope/vim-dispatch'
Plugin 'thoughtbot/vim-rspec'
Plugin 'vim-scripts/tComment'
Plugin 'ap/vim-css-color'
Plugin 'MarcWeber/vim-addon-mw-utils'
Plugin 'tomtom/tlib_vim'
Plugin 'garbas/vim-snipmate'
Plugin 'honza/vim-snippets'
Plugin 'ntpeters/vim-better-whitespace'

filetype plugin indent on     " required

" RSpec.vim mappings
map <Leader>t :call RunCurrentSpecFile()<CR>
map <Leader>s :call RunNearestSpec()<CR>
map <Leader>l :call RunLastSpec()<CR>
map <Leader>a :call RunAllSpecs()<CR>

" common setting
syntax on
set expandtab
set shiftwidth=2
set softtabstop=2
set tabstop=2
set number
set cindent
set autoindent

" 在所有 window 顯示狀態列
set laststatus=2

" 提示游標所在的行
set cursorline
" set cursorcolumn

" 高光搜尋
set hlsearch

" default show hidden files
let NERDTreeShowHidden=1
" use F3 toggle
noremap <F3> <ESC>:NERDTreeToggle<CR>
" user F5 to run rake
noremap <F5> <ESC>:!rake<CR>

" auto remove unwant space:w
autocmd BufWritePre * StripWhitespace

```

----

## Plugin list

## [nerdtree](https://github.com/scrooloose/nerdtree)

> 檔案導覽，[文件]()

## [vim-ruby](https://github.com/vim-ruby/vim-ruby)

> Ruby 支援，[文件](https://github.com/vim-ruby/vim-ruby/blob/master/doc/vim-ruby.txt)

## [vim-rails](https://github.com/tpope/vim-rails)

> Rails 支援，[文件](https://github.com/tpope/vim-rails/blob/master/doc/rails.txt)

## [vim-bundler](https://github.com/tpope/vim-bundler)

> Bundler 指令支援，[文件](https://github.com/tpope/vim-bundler/blob/master/doc/bundler.txt)

## [vim-dispatch](https://github.com/tpope/vim-dispatch)

> 非同步建置與測試，[文件](https://github.com/tpope/vim-dispatch/blob/master/doc/dispatch.txt)

## [vim-rspec](https://github.com/thoughtbot/vim-rspec)

> Rspec 支援，[文件](https://github.com/thoughtbot/vim-rspec)

## [tComment](https://github.com/vim-scripts/tComment)

> 提供註解的熱鍵，[文件](https://github.com/vim-scripts/tComment/blob/master/doc/tcomment.txt)

## [vim-css-color](https://github.com/ap/vim-css-color)

> 可以在 vim 中看到色碼的顏色

## [vim-addon-mw-utils](https://github.com/MarcWeber/vim-addon-mw-utils)

> For vim-snipmate

## [vim-tlib_vim](https://github.com/tomtom/tlib_vim)

> For vim-snipmate

## [vim-snipmate](https://github.com/garbas/vim-snipmate)

> 提供一些常用的程式碼片段，按  tab 補齊，讓寫 code 稍微快一點，[文件](https://github.com/garbas/vim-snipmate/blob/master/doc/SnipMate.txt)

## [vim-snippets](https://github.com/honza/vim-snippets)

> 支援更多程式語言的程式碼片段

## [vim-better-whitespace](https://github.com/ntpeters/vim-better-whitespace)

> 顯示行尾多餘空白，並提供移除功能

----

## Install Plugins

Run command in terminal:

```shell
vim +PluginInstall +qall
```

Or run command in vim

```shell
:PluginInstall
```
----

# FAQ

Q: Occur error when edit git commit in MAC OSX

A: run `git config --global core.editor /usr/bin/vim` to fix.

[Fixing "There was a problem with the editor 'vi'" for Git on Mac OS X Snow Leopard][link2]

----

# 參考

- [Vim for Rubyist: 10 useful vim plugins for Ruby][link3]
- [Awesome Vim Plugins][link4]
- [十個必用的 Vim Plugin][link5]
- [定制你的 Vim](https://www.darkshell.me/2015/05/06/ding-zhi-ni-de-vim/)
- [手把手将 Vim 打造成开发 Ruby 和 Rails 的强大 IDE](https://ruby-china.org/topics/25295)

[link1]: http://blogger.gtwang.org/2014/04/vundle-vim-bundle-plugin-manager.html
[link2]: http://tooky.co.uk/there-was-a-problem-with-the-editor-vi-git-on-mac-os-x/
[link3]: https://masteruby.github.io/productivity-booster/2014/05/02/vim-plugins-for-ruby.html#.VhyteLwkzMU
[link4]: https://reinteractive.net/posts/166-awesome-vim-plugins
[link5]: http://www.openfoundry.org/en/tech-column/2244--vim-plugin