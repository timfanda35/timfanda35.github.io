---
categories:
  - gitlab
  - docker
  - ruby
comments: true
date: 2022-02-16T12:00:00+08:00
title: gem install fail with docker in docker
url: /2022/02/16/gem-install-fail-with-dind/
---

打算照著 [GitLab](https://docs.gitlab.com/ee/user/packages/container_registry/#container-registry-examples-with-gitlab-cicd) 文件，用 dind 的方式建置 Container Image，但在執行 CI Job 時，收到了錯誤訊息 

```
You don't have write permission for the /usr/local/bundle
```

![](/images/2022-02-16/gem-install-fail-with-dind/001.png)

`.gitlab-ci.yml` 內容(參考官方文件)如下：

```
build:
  image: docker:19.03.12
  stage: build
  services:
    - docker:19.03.12-dind
  script:
    - docker login -u $CI_REGISTRY_USER -p $CI_REGISTRY_PASSWORD $CI_REGISTRY
    - docker build -t $CI_REGISTRY/group/project/image:latest .
    - docker push $CI_REGISTRY/group/project/image:latest
```

專案使用的 `Dockerfile` 內容如下：

```
FROM ruby:2.7.4-alpine

RUN apk add --no-cache --update build-base tzdata yarn mysql-dev libc6-compat
RUN gem install bundler --no-document
```

查了一下，發現因為 [Alpine 3.14 images can fail on Docker versions older than 20.10](https://github.com/docker-library/ruby/issues/351)，所以只要升級 docker 版本就可以了。將 `.gitlab-ci.yml` 修改如下：

```
stages:
  - build

build:
  image: docker:20.10.12
  stage: build
  services:
    - docker:20.10.12-dind
  script:
    - docker login -u $CI_REGISTRY_USER -p $CI_REGISTRY_PASSWORD $CI_REGISTRY
    - docker build -t $CI_REGISTRY/group/project/image:latest .
    - docker push $CI_REGISTRY/group/project/image:latest
```
