---
categories:
  - gem
keywords:
  - rails
  - gem
comments: true
date: 2024-01-02T00:00:00+08:00
title: "Rails 使用者驗證：Devise Gem Customize Views"
url: /2024/01/02/devise-gem-customize-views/
images:
  - /images/2024-01-02/devise-gem-customize-views.png
---

我們在 [Rails 使用者驗證：Devise Gem Getting Start][Rails 使用者驗證：Devise Gem Getting Start] 一文中使用 [Devise gem][Devise gem] 在 Rails 專案中快速實現最基礎的使用者驗證功能。本文進一步對其頁面樣式進行客製化。

## 簡介

我們可以透過 Devise 的 Generator 來產生 Devise 使用的 View Template，依照我們的需求來客製化頁面樣式。

本文我們會使用 [demo-devise][demo-devise] Rails 專案開始修改，此專案的進度如 [Rails 使用者驗證：Devise Gem Getting Start][Rails 使用者驗證：Devise Gem Getting Start] 一文。還會使用 [Tailwind CSS 的 Form Template][Tailwind CSS 的 Form Template] 作為登入與註冊頁面範本的基礎樣式。

## 環境

本文使用的環境：
- macOS
- Ruby 3.2.2
- Rails 7.1.2

取得專案程式碼：

```bash
git clone https://github.com/timfanda35/demo-devise.git
cd demo-devise
```

## 客製化頁面樣式

以下指令可以產生 Devise 提供的所有頁面範本：

```bash
bin/rails generate devise:views
```

{{< figure src="/images/2024-01-02/20240102001.jpg" alt="Generate Devise View" >}}

如果只需要客製化部分功能的範本，可以用 `-v` 指定，多個可以用空白分隔，例如：

```bash
bin/rails generate devise:views -v registrations confirmations
```

可以選擇的功能有：
- `confirmations`：驗證信箱
- `passwords`：忘記密碼
- `registrations`：註冊/更新
- `sessions`：登入
- `unlocks`：帳號解鎖
- `mailer`：所有信件範本

### 使用 Play CDN 引入 Tailwind CSS

引入 Tailwind CSS 最簡單的方式是使用 [Play CDN][Get started with Tailwind CSS]。(但要注意官方不建議在生產環境使用這種方式引入，這主要是用來測試與體驗)

修改 `app/views/layouts/application.html.erb`

```html
<!DOCTYPE html>
<html>
  <head>
    <title>DemoDevise</title>
    <meta name="viewport" content="width=device-width,initial-scale=1">
    <%= csrf_meta_tags %>
    <%= csp_meta_tag %>

    <%= stylesheet_link_tag "application", "data-turbo-track": "reload" %>
    <%= javascript_importmap_tags %>
    <script src="https://cdn.tailwindcss.com"></script>
  </head>

  <body>
    <header>
      <nav class="bg-gray-800">
        <div class="flex h-16 items-center justify-end gap-3">
          <%= link_to "Posts",
                      posts_path,
                      class: "text-gray-300 hover:bg-gray-700 hover:text-white rounded-md px-3 py-2 text-sm font-medium" %>
          <% if user_signed_in? %>
            <span class="text-gray-300 px-3 py-2 text-sm font-medium">
              <%= current_user.email %>
            </span>
            <%= link_to "Sign Out",
                        destroy_user_session_path,
                        data: {turbo_method: :delete },
                        class: "text-gray-300 hover:bg-gray-700 hover:text-white rounded-md px-3 py-2 text-sm font-medium" %>
          <% else %>
            <%= link_to "Sign In",
                        new_user_session_path,
                        class: "text-gray-300 hover:bg-gray-700 hover:text-white rounded-md px-3 py-2 text-sm font-medium" %>
          <% end %>
        <div>
      </nav>
    </header>

    <%- if controller_name != 'sessions' && controller_name != 'registrations' %>
      <main class="grid min-h-full place-items-center bg-white px-6 py-24 sm:py-32 lg:px-8">
    <% else %>
      <main>
    <% end %>
      <%= yield %>
    </main>
  </body>
</html>
```

啟動 Rails Server 來確認一下

```bash
bin/rails server
```

用瀏覽器開啟網址 [http://localhost:3000/][localhost]

{{< figure src="/images/2024-01-02/20240102002.jpg" alt="Import Tailwind CSS" >}}

其實當我們在建立 Rails 專案時，我們也能預先引入 Tailwind CSS，這樣透過 Scaffold 建立的頁面就會直接套用預設的 Tailwind CSS 樣式：

```bash
rails new --css=tailwind new-project
```

{{< figure src="/images/2024-01-02/20240102003.jpg"
    alt="Sample page with Tailwind CSS"
    caption="Rails Default Tailwind CSS Style" >}}

### 修改頁面

修改 `app/views/devise/sessions/new.html.erb` 套用樣式：

```html
<div class="flex min-h-full flex-col justify-center px-6 py-12 lg:px-8">
  <div class="sm:mx-auto sm:w-full sm:max-w-sm">
    <img class="mx-auto h-10 w-auto" src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600" alt="Your Company">
    <h2 class="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">Sign in to your account</h2>
  </div>

  <div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
    <%= form_for(resource, as: resource_name, url: session_path(resource_name), html: { class: "space-y-6" }) do |f| %>
      <div>
        <%= f.label :email, "Email address", class: "block text-sm font-medium leading-6 text-gray-900" %>
        <div class="mt-2">
          <%= f.email_field :email,
                            autofocus: true,
                            autocomplete: "email",
                            class: "block w-full rounded-md border-0 px-3 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6",
                            required: true %>
        </div>
      </div>

      <div>
        <div class="flex items-center justify-between">
          <%= f.label :password, class: "block text-sm font-medium leading-6 text-gray-900" %>
          <div class="text-sm">
            <%= link_to "Forgot password?", new_password_path(resource_name), class: "font-semibold text-indigo-600 hover:text-indigo-500" %>
          </div>
        </div>
        <div class="mt-2">
          <%= f.password_field :password,
                               autocomplete: "current-password",
                               class: "block w-full rounded-md border-0 px-3 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6",
                               required: true %>
        </div>
      </div>

      <div>
        <%= f.submit "Sign in",
                     class: "flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600" %>
      </div>
    <% end %>

    <p class="mt-10 text-center text-sm text-gray-500">
      Not a member?
      <%= link_to "Sign Up", new_registration_path(resource_name), class: "font-semibold text-indigo-600 hover:text-indigo-500" %>
    </p>
  </div>
</div>
```

{{< figure src="/images/2024-01-02/20240102004.jpg" alt="Sign In page with Tailwind CSS" >}}

修改 `app/views/devise/registrations/new.html.erb` 套用樣式：

```html
<div class="flex min-h-full flex-col justify-center px-6 py-12 lg:px-8">
  <div class="sm:mx-auto sm:w-full sm:max-w-sm">
    <img class="mx-auto h-10 w-auto" src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600" alt="Your Company">
    <h2 class="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">Sign up</h2>
  </div>

  <div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
    <%= form_for(resource, as: resource_name, url: registration_path(resource_name), html: { class: "space-y-6" }) do |f| %>
      <div>
        <%= f.label :email, "Email address", class: "block text-sm font-medium leading-6 text-gray-900" %>
        <div class="mt-2">
          <%= f.email_field :email,
                            autofocus: true,
                            autocomplete: "email",
                            class: "block w-full rounded-md border-0 px-3 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6",
                            required: true %>
        </div>
      </div>

      <div>
          <%= f.label :password, class: "block text-sm font-medium leading-6 text-gray-900" %>
          <% if @minimum_password_length %>
          <em>(<%= @minimum_password_length %> characters minimum)</em>
          <% end %>
          <div class="mt-2">
          <%= f.password_field :password,
                               autocomplete: "new-password",
                               class: "block w-full rounded-md border-0 px-3 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6",
                               required: true %>
        </div>
      </div>

      <div>
          <%= f.label :password_confirmation, class: "block text-sm font-medium leading-6 text-gray-900" %>
          <div class="mt-2">
          <%= f.password_field :password_confirmation,
                               autocomplete: "new-password",
                               class: "block w-full rounded-md border-0 px-3 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6",
                               required: true %>
        </div>
      </div>

      <div>
        <%= f.submit "Sign up",
                     class: "flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600" %>
      </div>
    <% end %>
  </div>
</div>
```

{{< figure src="/images/2024-01-02/20240102005.jpg" alt="Registration page with Tailwind CSS" >}}

## 註冊表單新增欄位

我們為 `User` Model 新增一個欄位為 `username`。

新增 Migration 檔案

```bash
bin/rails g migration add_username_to_users username:string
```

更新 DB Schema

```bash
bin/rails db:migrate
```

更新 `app/views/devise/registrations/new.html.erb` 在表單中加入 `username` 欄位：

```html
<div>
  <%= f.label :username, class: "block text-sm font-medium leading-6 text-gray-900" %>
  <div class="mt-2">
    <%= f.text_field :username,
                      class: "block w-full rounded-md border-0 px-3 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6",
                      required: true %>
  </div>
</div>
```

{{< figure src="/images/2024-01-02/20240102006.jpg" alt="Add username field to Registration page" >}}

雖然我們在註冊頁面新增了欄位，但實際註冊一個新使用者後，發現 `username` 的值並未存到資料庫中：

{{< figure src="/images/2024-01-02/20240102007.jpg" alt="Register a new user" >}}

執行指令查看：

```bash
bin/rails runner "puts JSON.pretty_generate(User.last.as_json)"
```

```json
{
  "id": 2,
  "email": "cindy@example.com",
  "created_at": "2024-01-02T00:20:07.719Z",
  "updated_at": "2024-01-02T00:20:07.719Z",
  "username": null
}
```

我們可以在 Rails 的 log 中發現，送出請求後，`username` 是 Unpermitted parameter：

```ruby
Unpermitted parameter: :username
```

這是因為 Devise 有使用 [Strong Parameters][Strong Parameters] 來限制傳入參數的關係。

### 修改 Strong Parameters

更新 `app/controllers/application_controller.rb`，我們告訴 Devise 如果是註冊頁面，就寫入資料庫時允許使用 `username` 參數：

```ruby
class ApplicationController < ActionController::Base
  before_action :configure_permitted_parameters, if: :devise_controller?

  protected

  def configure_permitted_parameters
    devise_parameter_sanitizer.permit(:sign_up, keys: [:username])
  end
end
```

再次註冊一個新的使用者：

{{< figure src="/images/2024-01-02/20240102008.jpg" alt="Register another new user" >}}

執行指令查看：

```bash
bin/rails runner "puts JSON.pretty_generate(User.last.as_json)"
```

```json
{
  "id": 3,
  "email": "dennis@example.com",
  "created_at": "2024-01-02T00:38:17.384Z",
  "updated_at": "2024-01-02T00:38:17.384Z",
  "username": "Dennis"
}
```

`username` 的值成功存到資料庫中。

## 總結

透過本文的步驟，我們為 Devise 提供的登入與註冊頁面客製化樣式，並在註冊頁面新增自訂欄位。

## 參考資料

- [Tailwind CSS 的 Form Template][Tailwind CSS 的 Form Template]
- [Get started with Tailwind CSS][Get started with Tailwind CSS]
- [Strong Parameters][Strong Parameters]

<!-- Links -->
[Rails 使用者驗證：Devise Gem Getting Start]: /2024/01/01/devise-gem-getting-start/
[demo-devise]: https://github.com/timfanda35/demo-devise
[Devise Gem]: https://github.com/heartcombo/devise
[Tailwind CSS 的 Form Template]: https://tailwindui.com/components/application-ui/forms/sign-in-forms
[Get started with Tailwind CSS]: https://tailwindcss.com/docs/installation/play-cdn
[Strong Parameters]: https://github.com/heartcombo/devise?tab=readme-ov-file#strong-parameters
[localhost]: http://localhost:3000/
