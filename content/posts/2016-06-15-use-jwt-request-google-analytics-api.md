---
categories:
  - ruby
  - jwt
comments: true
date: 2016-06-15T14:21:00+08:00
title: 透過 jwt 使用 Google Analytics API
url: /2016/06/15/use-jwt-request-google-analytics-api/
---

最近需要希望能夠不用網頁登入就能取得 Gooogle Analytics 的資料

但因為跟 [google-api-ruby-client](ref1) 不熟，只會用網頁登入

在搜尋後看到了這一篇[《在伺服器上使用 Google Analytics API》][ref2]

I got it!!!

<!--more-->

## 設定

設定就照著 [Analytics Reporting API - Authorization][ref3] 做

在 [Google Developers Console][ref4] 建立一個新專案，並啟用 **Analytics API**

然後在憑證頁面建立 **服務帳戶金鑰**，並將下載的 `JSON` 檔案改名成 `client-secret.json` 稍後使用

![1.png](http://user-image.logdown.io/user/6114/blog/6123/post/737458/a7I54QGYStublsdKRH0m_1.png)

建立好服務帳戶後，點擊管理服務帳戶，

![2.png](http://user-image.logdown.io/user/6114/blog/6123/post/737458/ACst1ndT0GLmbn3AIxEA_2.png)

複製剛剛新增的服務帳戶的**服務帳戶 ID**

![3.png](http://user-image.logdown.io/user/6114/blog/6123/post/737458/pC6xYwZQTIKckGKNmUtv_3.png)

到 [Google Analytics][ref5] 的管理頁面，點擊**資源**下的**使用者管理**

![4.png](http://user-image.logdown.io/user/6114/blog/6123/post/737458/4VRh5c2pQVOAp2fdPTla_4.png)

加入剛剛複製的**服務帳戶 ID**

![5.png](http://user-image.logdown.io/user/6114/blog/6123/post/737458/zerjIhaQaCk5872MgOow_5.png)

再到 [Google Analytics][ref5] 的管理頁面，點擊**資料檢視**下的**檢視設定**

複製 **資料檢視編號** 稍後使用

![6.png](http://user-image.logdown.io/user/6114/blog/6123/post/737458/VPzNGyZ5T76E4LjPmJ8v_6.png)

## Coding

立即來段程式碼 `main.rb`，記得修改 `ga_id` 的值

在同一目錄下放置設定步驟所下載的 `client-secret.json`

執行指令為： `ruby main.rb`

在這之前，你可能會需要執行 `gem install oj jwt faraday awesome_print` 安裝依賴的 gem

```ruby
# main.rb

require 'oj'
require 'jwt'
require 'faraday'
require 'awesome_print'

# 將設定步驟所下載的 JSON 改名成 client-secret.json
# 並放置在此 .rb 檔在相同的目錄下
# 以下將 client-secret.json 的內容讀取至 config hash 中
secert_json_path = File.expand_path('../client-secret.json', __FILE__)

unless File.exist?(secert_json_path)
  puts "client-secret.json not found!!!"
  exit
end

secert_json = File.read(secert_json_path)
config = Oj.load(secert_json)

# 產生 jwt sign ，這裡使用 jwt gem
# 依照 Google 說明文件建立 body 的部分
# 再透過 jwt gem 用 private key 對 jwt_body 簽名產生所需要的字串
jwt_body = {
  "iss": config['client_email'],
  "scope": "https://www.googleapis.com/auth/analytics.readonly",
  "aud": "https://www.googleapis.com/oauth2/v4/token",
  "iat": Time.now.to_i,
  "exp": Time.now.to_i + 3600
}

rsa_private = OpenSSL::PKey::RSA.new(config['private_key'])
jwt_sign = JWT.encode jwt_body, rsa_private, 'RS256'

# 這裡利用 faraday gem 送出 post 請求
# 參數依照 Google 說明文件設定
conn = Faraday.new(:url => 'https://www.googleapis.com') do |faraday|
  faraday.request  :url_encoded
  faraday.adapter  Faraday.default_adapter
end

resp = conn.post do |req|
  req.url '/oauth2/v4/token'
  req.body = {
    grant_type: "urn:ietf:params:oauth:grant-type:jwt-bearer",
    assertion: jwt_sign
  }
end

jwt_token =  Oj.load(resp.body)
puts "jwt_token"
ap jwt_token

#回傳成功，取得時效為 3,600 秒的 jwt token
#{
#    "access_token" => "bla bla token",
#      "token_type" => "Bearer",
#      "expires_in" => 3600
#}

# 這裡使用 V3 API 做示範
# ga_id 填入設定步驟取得的"資料檢視編號"
ga_id = "replace to your info"
ga_resp = conn.get do |req|
  req.url '/analytics/v3/data/ga'
  req.headers['Authorization'] = "Bearer #{jwt_token['access_token']}"
  req.params = {
    "ids": "ga:#{ga_id}",
    "start-date": Time.now.strftime("%Y-%m-01"),
    "end-date": Time.now.strftime("%Y-%m-%d"),
    "metrics": "ga:sessions,ga:bounces"
  }
end

puts "ga response"
ap Oj.load(ga_resp.body)

# 請求成功就可以取得 JSON 資料了
# {
#                    "kind" => "analytics#gaData",
#                      "id" => "https://www.googleapis.com/analytics/v3/data/ga?ids=ga:---------&metrics=ga:sessions,ga:bounces&start-date=2016-06-01&end-date=2016-06-15",
#                   "query" => {
#          "start-date" => "2016-06-01",
#            "end-date" => "2016-06-15",
#                 "ids" => "ga:---------",
#             "metrics" => [
#             [0] "ga:sessions",
#             [1] "ga:bounces"
#         ],
#         "start-index" => 1,
#         "max-results" => 1000
#     },
#            "itemsPerPage" => 1000,
#            "totalResults" => 1,
#                "selfLink" => "https://www.googleapis.com/analytics/v3/data/ga?ids=ga:---------&metrics=ga:sessions,ga:bounces&start-date=2016-06-01&end-date=2016-06-15",
#             "profileInfo" => {
#                     "profileId" => "---------",
#                     "accountId" => "---------",
#                 "webPropertyId" => "UA-----------",
#         "internalWebPropertyId" => "---------",
#                   "profileName" => "所有網站資料",
#                       "tableId" => "ga:---------"
#     },
#     "containsSampledData" => false,
#           "columnHeaders" => [
#         [0] {
#                   "name" => "ga:sessions",
#             "columnType" => "METRIC",
#               "dataType" => "INTEGER"
#         },
#         [1] {
#                   "name" => "ga:bounces",
#             "columnType" => "METRIC",
#               "dataType" => "INTEGER"
#         }
#     ],
#     "totalsForAllResults" => {
#         "ga:sessions" => "106",
#          "ga:bounces" => "89"
#     },
#                    "rows" => [
#         [0] [
#             [0] "106",
#             [1] "89"
#         ]
#     ]
# }
```

接下來就是嗑文件開發符合需求的功能了！

## 參考

[《在伺服器上使用 Google Analytics API》][ref2]

[Analytics Reporting API - Authorization][ref3]

[Migrate Analytics API V3 to V4][ref6]

[Dimensions & Metrics Explorer][ref7]

[API - Reference Guide][ref8]

[ref1]: https://github.com/google/google-api-ruby-client
[ref2]: https://zespia.tw/blog/2014/07/28/use-google-analytics-api-on-server/
[ref3]: https://developers.google.com/analytics/devguides/reporting/core/v4/authorization
[ref4]: https://console.developers.google.com/
[ref5]: https://analytics.google.com
[ref6]: https://developers.google.com/analytics/devguides/reporting/core/v4/migration
[ref7]: https://developers.google.com/analytics/devguides/reporting/core/dimsmets
[ref8]: https://developers.google.com/analytics/devguides/reporting/core/v3/reference