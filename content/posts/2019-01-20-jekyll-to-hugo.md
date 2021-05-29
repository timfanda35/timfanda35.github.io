---
categories:
comments: true
date: 2019-01-20T01:20:26+08:00
title: 從 jekyll 搬到 hugo 筆記
url: /2019/01/20/jekyll-to-hugo/
---

![](/images/2019-01-20/jekyll-to-hugo/cover.jpg)

[圖片來源](https://www.pakutaso.com/20151102331web-14.html)
[フリー写真素材ぱくたそ](https://www.pakutaso.com/)

總之從 Jekyll 搬到了 Hugo，再慢慢調版面與新增文章。

<!--more-->

今天在公司剛好提到要打算用 GitHub Page 作為之後分享技術文章的平台。

再度搜尋了一下有哪些工具可用，就想到遙遠記憶中網路上有人推薦的 [Hugo](https://gohugo.io/)。

搬家的過程還算順利的，網路搜尋了一下就找到有人寫好的轉換工具。

使用 Hugo 本身的轉換功能產生新的目錄，但發現文章中的 Metadata 都消失了。

```bash
hugo import jekyll jekyll_root_path target_path
```

所以再使用 [ConvertToHugo](https://github.com/coderzh/ConvertToHugo) ，將轉好的文章放到剛剛產生的目錄下的 `content/post`。

```bash
python ConvertToHugo.py jekyll-post-dir output-dir
```

雖然 Metadata 也帶過來了，但時間有格式問題，我是直接用 VS Code 的取代全部來轉換格式。最後按照文件調整 `config.yaml`。

由於個人 Repo 的 GitHub Pages 只能用 Master branch 做為來源，依照 [Hugo 文件](https://gohugo.io/hosting-and-deployment/hosting-on-github/) 操作，再開一個新的 Repo，用加入 gitmodule 的方式更新個人 Repo。

其他參考：

* [使用hugo搭建个人博客站点](https://blog.coderzh.com/2015/08/29/hugo/)
