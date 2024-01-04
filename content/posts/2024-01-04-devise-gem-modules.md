---
categories:
  - gem
keywords:
  - rails
  - gem
comments: true
date: 2024-01-04T00:00:00+08:00
title: "Rails 使用者驗證：Devise Gem Modules"
url: /2024/01/04/devise-gem-modules/
images:
  - /images/2024-01-04/devise-gem-modules.png
---

[Devise gem][Devise gem] 將主要功能分為 10 個模組，我們可以依據需求選擇啟用，並利用模組提供的方法來進行客製化開發。

## 簡介

從 [Devise gem][Devise gem] 的 GitHub README 上可以看到，這 10 個 Module 分別為：

1. `Database Authenticatable`：最主要的功能模組，用來對密碼作雜湊後儲存到資料庫中，並提供方法做使用者身份驗證。
2.  `Omniauthable`：用以支援 [OmniAuth Gem][OmniAuth Gem] 的功能模組。
3.  `Confirmable`：用以寄送登入指示與驗證使用者是否能登入系統的功能模組。
4.  `Recoverable`：用以重設密碼的功能模組。
5.  `Registerable`：提供註冊新使用者、編輯與刪除帳號的功能模組。
6.  `Rememberable`：記住登入的功能模組。
7.  `Trackable`：紀錄使用者登入次數、時間、與登入 IP 的功能模組。
8.  `Timeoutable`：自動登出閒置使用者的功能模組。
9.  `Validatable`：提供電子信箱與密碼驗證的功能模組。
10. `Lockable`：登入失敗即鎖定帳號的功能模組。

## Database Authenticatable

這是核心的功能模組，用來驗證使用者身份。定義了 `password=` 方法將密碼雜湊後存到資料表的 `encrypted_password` 欄位。

我們可以自行用 `valid_password?` 方法來確認密碼是否相符：

```ruby
User.find(1).valid_password?('password123')
```

  在 `config/initializers/devise.rb` 中調整密碼雜湊的強度：

```ruby
  # ==> Configuration for :database_authenticatable
  # For bcrypt, this is the cost for hashing the password and defaults to 12. If
  # using other algorithms, it sets how many times you want the password to be hashed.
  # The number of stretches used for generating the hashed password are stored
  # with the hashed password. This allows you to change the stretches without
  # invalidating existing passwords.
  #
  # Limiting the stretches to just one in testing will increase the performance of
  # your test suite dramatically. However, it is STRONGLY RECOMMENDED to not use
  # a value less than 10 in other environments. Note that, for bcrypt (the default
  # algorithm), the cost increases exponentially with the number of stretches (e.g.
  # a value of 20 is already extremely slow: approx. 60 seconds for 1 calculation).
  config.stretches = Rails.env.test? ? 1 : 12

  # Set up a pepper to generate the hashed password.
  # config.pepper = '4accb7eccfc24dea2a36548457b0746d655b8444911fd1e8dd03172f3c2c9ac4014d7e26e3283e171d4ed6428869524a4b71d0ff7a137527aff8a94641ff4eea'

  # Send a notification to the original email when the user's email is changed.
  # config.send_email_changed_notification = false

  # Send a notification email when the user's password is changed.
  # config.send_password_change_notification = false
```

修改設定，當使用者變更密碼成功後會寄送通知信件：

```ruby
# Send a notification email when the user's password is changed.
config.send_password_change_notification = false
```

## Omniauthable

這是用以支援 [OmniAuth Gem][OmniAuth Gem] 的功能模組，透過這個模組可以讓 Devise 整合第三方登入如 Google、GitHub 等等。

我們可以參考 [OmniAuth Facebook example][OmniAuth Facebook example] 進行設定。

## Confirmable

這個功能模提供確認使用者能否登入系統，以及寄送電子信箱驗證信的功能。

需要在驗證 Model 的資料表上新增欄位，可以在 `rails generate devise User` 產生的 Migration 檔案中移除註解：

```ruby
def change
  create_table :users do |t|
    ...

    ## Confirmable
    t.string   :confirmation_token
    t.datetime :confirmed_at
    t.datetime :confirmation_sent_at
    t.string   :unconfirmed_email # Only if using reconfirmable

    ...
  end
end
```

或是自行新增 Migration 檔案，為現有的 Model 的資料表加上欄位：

```ruby
def change
  ## Confirmable
  add_column :users, :confirmation_token, :string
  add_column :users, :confirmed_at, :datetime
  add_column :users, :confirmation_sent_at, :datetime
  add_column :users, :unconfirmed_email, :string
end
```

在驗證 Model 中啟用，編輯 `app/models/user.rb`

```ruby
class User < ApplicationRecord
  # Include default devise modules. Others available are:
  # :confirmable, :lockable, :timeoutable, :trackable and :omniauthable
  devise :database_authenticatable, :registerable,
         :recoverable, :rememberable, :validatable,
         :confirmable
end
```

啟用後，當我們註冊新使用者時，Devise 就會寄送一封 Confirm instructions 到使用者的電子信箱。使用者必須點擊信件中的連結才能夠登入系統。

{{< figure src="/images/2024-01-04/20240104001.jpg" alt="Confirm instructions" >}}

我們可以在 `config/initializers/devise.rb` 調整相關設定：

```ruby
  # ==> Configuration for :confirmable
  # A period that the user is allowed to access the website even without
  # confirming their account. For instance, if set to 2.days, the user will be
  # able to access the website for two days without confirming their account,
  # access will be blocked just in the third day.
  # You can also set it to nil, which will allow the user to access the website
  # without confirming their account.
  # Default is 0.days, meaning the user cannot access the website without
  # confirming their account.
  # config.allow_unconfirmed_access_for = 2.days

  # A period that the user is allowed to confirm their account before their
  # token becomes invalid. For example, if set to 3.days, the user can confirm
  # their account within 3 days after the mail was sent, but on the fourth day
  # their account can't be confirmed with the token any more.
  # Default is nil, meaning there is no restriction on how long a user can take
  # before confirming their account.
  # config.confirm_within = 3.days

  # If true, requires any email changes to be confirmed (exactly the same way as
  # initial account confirmation) to be applied. Requires additional unconfirmed_email
  # db field (see migrations). Until confirmed, new email is stored in
  # unconfirmed_email column, and copied to email column on successful confirmation.
  config.reconfirmable = true

  # Defines which key will be used when confirming an account
  # config.confirmation_keys = [:email]
```

## Recoverable

這是預設啟用的功能模組之一，讓使用者可以重設密碼。

我們可以使用提供的方法寄送重設密碼信：

```ruby
User.find(1).send_reset_password_instructions
```

我們可以在 `config/initializers/devise.rb` 調整相關設定：

```ruby
  # ==> Configuration for :recoverable
  #
  # Defines which key will be used when recovering the password for an account
  # config.reset_password_keys = [:email]

  # Time interval you can reset your password with a reset password key.
  # Don't put a too small interval or your users won't have the time to
  # change their passwords.
  config.reset_password_within = 6.hours

  # When set to false, does not sign a user in automatically after their password is
  # reset. Defaults to true, so a user is signed in automatically after a reset.
  # config.sign_in_after_reset_password = true
```

## Registerable

這是預設啟用的功能模組之一，讓使用者可以自行註冊帳號。

我們可以在 `config/initializers/devise.rb` 調整相關設定：

```ruby
  # ==> Configuration for :registerable

  # When set to false, does not sign a user in automatically after their password is
  # changed. Defaults to true, so a user is signed in automatically after changing a password.
  # config.sign_in_after_change_password = true
```

## Rememberable

這是預設啟用的功能模組之一，用來設定有期限的 Cookie 來記住登入資訊。

預設是儲存兩個星期，但我們可以修改 `config/initializers/devise.rb` 中的設定：

```ruby
  # ==> Configuration for :rememberable
  # The time the user will be remembered without asking for credentials again.
  # config.remember_for = 2.weeks

  # Invalidates all the remember me tokens when the user signs out.
  config.expire_all_remember_me_on_sign_out = true

  # If true, extends the user's remember period when remembered via cookie.
  # config.extend_remember_period = false

  # Options to be passed to the created cookie. For instance, you can set
  # secure: true in order to force SSL only cookies.
  # config.rememberable_options = {}
```

預設情況下，Rails 會用 Cookie 儲存 Session 資訊，而該 Cookie 的 Expires 欄位為 `Session`。

{{< figure src="/images/2024-01-04/20240104002.jpg" alt="Default cookies" >}}

而當我們於登入同時傳入 `user[remember_me]=1`，登入表單上可以加上 Checkbox：

```ruby
<div class="flex items-center justify-end gap-2">
  <p><%= f.check_box :remember_me %></p>
  <p><%= f.label :remember_me %></p>
</div>
```

{{< figure src="/images/2024-01-04/20240104003.jpg" alt="Remember me cookies" >}}

可以發現多了一個叫做 `remember_user_token` 的 cookies，而且 `Expires` 欄位有填入特定的過期時間。

## Trackable

這個功能模組可以紀錄使用者登入次數、時間、與登入 IP。

需要在驗證 Model 的資料表上新增欄位，可以在 `rails generate devise User` 產生的 Migration 檔案中移除註解：

```ruby
def change
  create_table :users do |t|
    ...

    ## Trackable
    t.integer  :sign_in_count, default: 0, null: false
    t.datetime :current_sign_in_at
    t.datetime :last_sign_in_at
    t.string   :current_sign_in_ip
    t.string   :last_sign_in_ip

    ...
  end
end
```

或是自行新增 Migration 檔案，為現有的 Model 的資料表加上欄位：

```ruby
def change
  ## Trackable
  add_column :users, :sign_in_count, :integer, default: 0, null: false
  add_column :users, :current_sign_in_at, :datetime
  add_column :users, :last_sign_in_at, :datetime
  add_column :users, :current_sign_in_ip, :string
  add_column :users, :last_sign_in_ip, :string
end
```

另外如果資料庫使用的是 `postgresql`，儲存 IP 的欄位型態會用 `inet` 而不是 `string`。

在驗證 Model 中啟用，編輯 `app/models/user.rb`

```ruby
class User < ApplicationRecord
  # Include default devise modules. Others available are:
  # :confirmable, :lockable, :timeoutable, :trackable and :omniauthable
  devise :database_authenticatable, :registerable,
         :recoverable, :rememberable, :validatable,
         :trackable
end
```

從資料欄位我們就可以察覺到，該功能只會紀錄累積的登入次數、上次與最近的登入資訊。如果要能夠紀錄每一次登入歷史，還是需要自行開發，或是使用 [AuthTrail Gem][AuthTrail Gem]。

## Timeoutable

這個功能模組是用來自動登出超時閒置的使用者。啟用此模組不需要新增資料表欄位。直接在驗證 Model 中啟用，編輯 `app/models/user.rb`

```ruby
class User < ApplicationRecord
  # Include default devise modules. Others available are:
  # :confirmable, :lockable, :timeoutable, :trackable and :omniauthable
  devise :database_authenticatable, :registerable,
         :recoverable, :rememberable, :validatable,
         :timeoutable
end
```

並編輯 `config/initializers/devise.rb`，設定超時時間：

```ruby
  # ==> Configuration for :timeoutable
  # The time you want to timeout the user session without activity. After this
  # time the user will be asked for credentials again. Default is 30 minutes.
  config.timeout_in = 30.minutes
```

我們可以在[原始碼][Devise after_set_user hook]中確認行為：

```ruby
Warden::Manager.after_set_user do |record, warden, options|
  ...

  if record && record.respond_to?(:timedout?) && warden.authenticated?(scope) &&
     options[:store] != false && !env['devise.skip_timeoutable']
    last_request_at = warden.session(scope)['last_request_at']

    ...

    if !env['devise.skip_timeout'] &&
        record.timedout?(last_request_at) &&
        !proxy.remember_me_is_active?(record)
      Devise.sign_out_all_scopes ? proxy.sign_out : proxy.sign_out(scope)
      throw :warden, scope: scope, message: :timeout
    end

    unless env['devise.skip_trackable']
      warden.session(scope)['last_request_at'] = Time.now.utc.to_i
    end
  end
end
```

每次已登入的請求都會更新 Session 中 `last_request_at` 的值，當使用者閒置超過 `timeout_in` 設定的時間，再次送出請求時就會被登出。

## Validatable

這是預設啟用的功能模組之一，提供電子信箱與密碼欄位的格式驗證設定。

我們可以在 `config/initializers/devise.rb` 調整欄位驗證的條件：

```ruby
  # ==> Configuration for :validatable
  # Range for password length.
  config.password_length = 6..128

  # Email regex used to validate email formats. It simply asserts that
  # one (and only one) @ exists in the given string. This is mainly
  # to give user feedback and not to assert the e-mail validity.
  config.email_regexp = /\A[^@\s]+@[^@\s]+\z/
```

## Lockable

功能模組提供登入失敗即鎖定帳號的功能。

需要在驗證 Model 的資料表上新增欄位，可以在 `rails generate devise User` 產生的 Migration 檔案中移除註解：

```ruby
def change
  create_table :users do |t|
    ...

    ## Lockable
    t.integer  :failed_attempts, default: 0, null: false # Only if lock strategy is :failed_attempts
    t.string   :unlock_token # Only if unlock strategy is :email or :both
    t.datetime :locked_at

    ...
  end
end
```

或是自行新增 Migration 檔案，為現有的 Model 的資料表加上欄位：

```ruby
def change
  ## Lockable
  add_column :users, :failed_attempts, :integer, default: 0, null: false
  add_column :users, :unlock_token, :string
  add_column :users, :locked_at, :datetime
end
```

在驗證 Model 中啟用，編輯 `app/models/user.rb`

```ruby
class User < ApplicationRecord
  # Include default devise modules. Others available are:
  # :confirmable, :lockable, :timeoutable, :trackable and :omniauthable
  devise :database_authenticatable, :registerable,
         :recoverable, :rememberable, :validatable,
         :lockable
end
```

我們可以在 `config/initializers/devise.rb` 調整相關設定：

```ruby
  # ==> Configuration for :lockable
  # Defines which strategy will be used to lock an account.
  # :failed_attempts = Locks an account after a number of failed attempts to sign in.
  # :none            = No lock strategy. You should handle locking by yourself.
  # config.lock_strategy = :failed_attempts

  # Defines which key will be used when locking and unlocking an account
  # config.unlock_keys = [:email]

  # Defines which strategy will be used to unlock an account.
  # :email = Sends an unlock link to the user email
  # :time  = Re-enables login after a certain amount of time (see :unlock_in below)
  # :both  = Enables both strategies
  # :none  = No unlock strategy. You should handle unlocking by yourself.
  # config.unlock_strategy = :both

  # Number of authentication tries before locking an account if lock_strategy
  # is failed attempts.
  # config.maximum_attempts = 20

  # Time interval to unlock the account if :time is enabled as unlock_strategy.
  # config.unlock_in = 1.hour

  # Warn on the last attempt before the account is locked.
  # config.last_attempt_warning = true
```

範例設定：

```ruby
config.lock_strategy = :failed_attempts
config.unlock_keys = [:email]
config.unlock_strategy = :both
config.maximum_attempts = 3
config.unlock_in = 1.minutes
config.last_attempt_warning = true
```

嘗試登入 3 次失敗，會收到：

{{< figure src="/images/2024-01-04/20240104004.jpg" alt="Unlock link" >}}

點擊信件中的連結就可以解鎖，或是等到設定的 1 分鐘後才重試。

我們也可以自行鎖定或解鎖使用者：

```ruby
# 鎖定，只有當 config.unlock_strategy 包含 :email 才會寄送 unlock link
User.find(1).lock_access!

# 鎖定但不寄送 unlock link
User.find(1).lock_access!({ send_instructions: false })

# 解鎖
User.find(1).unlock_access!
```

## 總結

Devise 提供 10 大模組涵蓋了實作使用者身份驗證功能大多數的需求，除了直接使用 Devise 的預設行為，我們也能夠透過調整設定或是使用模組提供的方法來依據需求客製化，減少開發的時間。

## 參考資料
- [Devise Gem][Devise Gem]
- [OmniAuth Gem][OmniAuth Gem]
- [AuthTrail Gem][AuthTrail Gem]
- [Devise after_set_user hook][Devise after_set_user hook]
- [OmniAuth Facebook example][OmniAuth Facebook example]

<!-- Links -->
[Devise Gem]: https://github.com/heartcombo/devise
[OmniAuth Gem]: https://github.com/omniauth/omniauth
[AuthTrail Gem]: https://github.com/ankane/authtrail
[Devise after_set_user hook]: https://github.com/heartcombo/devise/blob/e2242a95f3bb2e68ec0e9a064238ff7af6429545/lib/devise/hooks/timeoutable.rb#L24
[OmniAuth Facebook example]: https://github.com/heartcombo/devise/wiki/OmniAuth:-Overview#facebook-example
[localhost]: http://localhost:3000/
