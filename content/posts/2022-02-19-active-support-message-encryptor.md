---
categories:
  - rails
comments: true
date: 2022-02-19T12:00:00+08:00
title: Rails ActiveSupport::MessageEncryptor
url: /2022/02/19/active-support-message-encryptor/
---

啟動 Rails Server 時，如果沒有 `config/master.key` ，或是 `config/master.key` 無法解密 `config/credentials.yml.enc`，就會出現 `ActiveSupport::MessageEncryptor` 的錯誤

為 Production 環境準備一個 `config/master.key` 跟 `config/credentials.yml.enc` 在管理與部屬上不太方便

在 Rails 6 之後，可以依據環境變數產生並套用不同的 credentials

## Rails credentials

執行指令查看 rails credentials 相關說明

```
rails credentials:help
```

輸出(節錄部分)

```
...


=== Environment Specific Credentials

The `credentials` command supports passing an `--environment` option to create an
environment specific override. That override will take precedence over the
global `config/credentials.yml.enc` file when running in that environment. So:

   bin/rails credentials:edit --environment development

will create `config/credentials/development.yml.enc` with the corresponding
encryption key in `config/credentials/development.key` if the credentials file
doesn't exist.

The encryption key can also be put in `ENV["RAILS_MASTER_KEY"]`, which takes
precedence over the file encryption key.

In addition to that, the default credentials lookup paths can be overridden through
`config.credentials.content_path` and `config.credentials.key_path`.
```

## 新增/編輯 credentials

建立給 production environment 用的 credentials

```
rails credentials:edit --environment production
```

credentials 與 key 放在 `config/credentials` 目錄下，檔案為：

- `config/credentials/production.key`
- `config/credentials/production.yml.enc`

記得在 `.gitignore` 中新增規則排除 key

```
/config/credentials/*.key
```

## 查看 credentials

```
rails credentials:show --environment production
```

## 使用 credentials

在 Rails 程式中，我們可以這樣使用 credentials

```ruby
# 取得全部 config hash
Rails.application.credentials.config
#=> {:aws=>{:access_key_id=>"123", :secret_access_key=>"456"}}

# 只取部分值
Rails.application.credentials.aws[:access_key_id]
#=> "123"
```

## 使用環境變數

我們也可以將 Master key 設為環境變數而不儲存在檔案中，以 Heroku 為例：

```
heroku config:set RAILS_MASTER_KEY=`cat config/credentials/production.key`
```

這樣就不用把 key 上傳 git，或是想辦法塞入檔案了

## Ref
- https://blog.saeloun.com/2019/10/10/rails-6-adds-support-for-multi-environment-credentials.html
