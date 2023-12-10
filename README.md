# Bear Su's Blog

Render by [Hugo](https://gohugo.io/).

Install Hugo

```
brew install hugo
```

Add Post

```bash
hugo new post/new-post.md
```

Local Test

```bash
make
```

## Generate OG Image

Install dependency

```bash
go mod tidy
```

Build

```bash
make build_og_gen
```

Generate image

```bash
og_gen <post-file>

# for example
og_gen content/posts/2023-10-14-dry-initializer-gem.md
```

Output

```bash
2023/12/11 00:44:35 postTitle: dry-initializer 介紹
2023/12/11 00:44:35 OG Image Save to: static/images/2023-10-14/dry-initializer-gem.png
```

We can set the image to the post meta

```yaml
images:
  - images/2023-10-14/dry-initializer-gem.png
```
