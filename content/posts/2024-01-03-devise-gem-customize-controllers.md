---
categories:
  - gem
keywords:
  - rails
  - gem
comments: true
date: 2024-01-03T00:00:00+08:00
title: "Rails 使用者驗證：Devise Gem Customize Controllers"
url: /2024/01/03/devise-gem-customize-controllers/
images:
  - /images/2024-01-03/devise-gem-customize-controllers.png
---

我們 [Rails 使用者驗證：Devise Gem Getting Start][Rails 使用者驗證：Devise Gem Getting Start] 一文中使用 [Devise gem][Devise gem] 在 Rails 專案中快速實現最基礎的使用者驗證功能。本文進一步對其路由與功能進行客製化。

## 簡介

我們可以透過 Devise 的 Generator 來產生 Devise 使用的 Controller，並依照我們的需求來修改功能的行為。

本文我們會使用 [demo-devise][demo-devise] Rails 專案，從 [Rails 使用者驗證：Devise Gem Customize Views][Rails 使用者驗證：Devise Gem Customize Views] 一文的進度開始修改。

## 環境

本文使用的環境：
- macOS
- Ruby 3.2.2
- Rails 7.1.2

取得專案程式碼：

```bash
git clone https://github.com/timfanda35/demo-devise.git
cd demo-devise

git checkout customize_views
```

## 產生 Controller 檔案

以下指令可以產生 Devise 提供的所有 Controllers，我們需要指定 `SCOPE` 也就是驗證 Model。在 `demo-devise` 中我們使用的是慣例的 `users`：

```bash
bin/rails generate devise:controllers users
```

{{< figure src="/images/2024-01-03/20240103001.jpg" alt="Generate Devise Controllers" >}}

在訊息中可以看到，我們將需要修改 `config/routes.rb`，應用程式運行的時候才能將請求送往我們客製化後的 Controller 進行處理。

如果只需要客製化部分 Controllers，可以用 `-c` 指定，多個可以用空白分隔，例如：

```bash
bin/rails generate devise:controllers users -c sessions passwords
```

可以選擇的功能有：
- `confirmations`：驗證信箱
- `passwords`：忘記密碼
- `registrations`：註冊/更新
- `sessions`：登入
- `unlocks`：帳號解鎖
- `omniauth_callbacks`：OmniAuth Callback

一但產生 Controller 檔案後，[官方文件][Devise Gem] 建議將產生的 Devise 頁面範本從專案目錄 `app/views/devise` 移到 `app/views/users` 下。這並不是一個必要的步驟，但會讓我們的專案結構更為一致。

## 修改路由

執行指令確認目前 `SCOPE` 為 `users` 的路由：

```bash
bin/rails routes --grep="users"
```

{{< figure src="/images/2024-01-03/20240103002.jpg" alt="Default Routes" >}}

可以發現目前 URI 在 `/users` 底下的路徑，都會將請求送到 `devise` 底下的 Controllers，而不是 `users` 底下的 Controllers。所以我們才需要依照提示修改 `config/routes` 修改路由將指定路徑的請求送到我們客製化的 Controller。

我們可以啟動 Rails Server 來觀察一下目前請求所使用的 Controller

```bash
bin/rails server
```

用瀏覽器開啟網址 [http://localhost:3000/][localhost]，點擊 Sign In，確認一下 Log：

{{< figure src="/images/2024-01-03/20240103003.jpg" alt="Default Routes for sign in" >}}

可以發現送往 `/users/sign_in` 的請求是由 `Devise::SessionsController` 處理。

讓我們修改 `config/routes.rb`，改成由上面步驟產生的 Controller 來處理：

```ruby
Rails.application.routes.draw do
  devise_for :users, controllers: {
    sessions: 'users/sessions'
  }

  resources :posts
  root "posts#index"

  get "up" => "rails/health#show", as: :rails_health_check
end
```

重新啟動 Rails Server

```bash
bin/rails server
```

用瀏覽器開啟網址 [http://localhost:3000/][localhost]，點擊 Sign In，確認一下 Log：

{{< figure src="/images/2024-01-03/20240103004.jpg" alt="Customize Routes for sign in" >}}

可以發現送往 `/users/sign_in` 的請求變成由 `Users::SessionsController` 來處理。

## 客製化 Controller

我們在前述步驟產生的 Controller 都繼承了 `Devise::Controller`，而且檔案中已經註解好方法：

{{< figure src="/images/2024-01-03/20240103005.jpg" alt="Commented Action" >}}

註解中的方法使用 `super` 去執行父物件的方法，我們可以查看 [Devise::SessionsController 的原始碼][Devise::SessionsController Source Code] 其中一段：

```ruby
  # POST /resource/sign_in
  def create
    self.resource = warden.authenticate!(auth_options)
    set_flash_message!(:notice, :signed_in)
    sign_in(resource_name, resource)
    yield resource if block_given?
    respond_with resource, location: after_sign_in_path_for(resource)
  end
```

可以看到中間有一段

```ruby
yield resource if block_given?
```

我們只需要傳入 `block`，就能夠在登入成功後接著進行我們自訂的額外行為，最後再繼續依照原本流程渲染回應頁面：

```ruby
class Users::SessionsController < Devise::SessionsController
  def create
    super do |resource|
      # do something...
    end
  end
end
```

我們移除 `Users::SessionsController` 中的 `create` 方法註解，並進行修改，使其能夠在使用者成功登入後，印出登入時來源 IP 的 Log。

```ruby
class Users::SessionsController < Devise::SessionsController
  def create
    super do |resource|
      Rails.logger.warn "\n\nUser ##{resource.id} has signed in from #{request.remote_ip}\n\n"
    end
  end
end
```

我們啟動 Rails Server，登入後確認 Log：

{{< figure src="/images/2024-01-03/20240103006.jpg" alt="Show Log" >}}

登入成功會顯示 `User #1 has signed in from 127.0.0.1`。

{{< figure src="/images/2024-01-03/20240103007.jpg" alt="Show Log" >}}

登入失敗則不會顯示。

## 總結

我們可以透過 Devise 提供的 Generator 產生可供客製化 Controller 檔案，在該 Controller 中我們將 block 傳入父物件的方法來在該方法的 Happy Path 中去執行我們所想要添加的額外行為。

## 參考資料
- [Rails 使用者驗證：Devise Gem Getting Start][Rails 使用者驗證：Devise Gem Getting Start]
- [Rails 使用者驗證：Devise Gem Customize Views][Rails 使用者驗證：Devise Gem Customize Views]
- [Devise Gem][Devise Gem]
- [Devise::SessionsController Source Code][Devise::SessionsController Source Code]

<!-- Links -->
[Rails 使用者驗證：Devise Gem Getting Start]: /2024/01/01/devise-gem-getting-start/
[Rails 使用者驗證：Devise Gem Customize Views]: /2024/01/02/devise-gem-customize-views/
[demo-devise]: https://github.com/timfanda35/demo-devise
[Devise Gem]: https://github.com/heartcombo/devise
[Devise::SessionsController Source Code]: https://github.com/heartcombo/devise/blob/e2242a95f3bb2e68ec0e9a064238ff7af6429545/app/controllers/devise/sessions_controller.rb#L17-L24
[localhost]: http://localhost:3000/
