---
categories:
  - gcp
keywords:
  - note
  - gcp
  - google cloud
  - gdg
  - google developer
comments: true
date: 2023-12-17T12:00:00+08:00
title: "[筆記] GDG DevFest Taipei 2023"
url: /2023/12/17/note-gdg-devfest-taipei-2023/
images:
  - /images/2023-12-17/note-gdg-devfest-taipei-2023.png
---

不意外的，這次 [GDG DevFest Taipei 2023](https://gdg.community.dev/events/details/google-gdg-taipei-presents-devfest-taipei-2023/) 有許多討論關於 AI 的議題，但我還是選擇跟 Cloud 或是 Application Development 比較相關且有興趣的議題去聽，憑記憶紀錄一下想法。

## Opening/Keynote - 展望 Generative AI 的現在與未來

講者：EvA ChU、Eric ShangKuan、薛良斌、KJ Wu (吳貴融)、林雅芳 (Tina Lin)

印象最深的是布丁的觀點：
> LLM 的本質就是幻覺。It's Not a Bug, It's a Feature!

運用 LLM 的時候就是讓它根據素材進行幻想，而我們將這些產生的幻覺當作素材去進行後續的創作。要求 LLM 不要幻想，就相當於否定它的本質了，LLM 就是個不斷做夢的機器裝置。

> 你不會因為 LLM，就比測試工程師還要懂測試，但是你獲得一個測試實習生。

我們可能不會開車，但是我們知道目的地方向。LLM 與我們的關係可以是 LLM 為駕駛不停進行創作的發想；而我們則是副駕駛，用 prompt 控制 LLM 發想的方向，透過這樣的配合降低創作的門檻。

簡報連結：[Generating New Possibilities in the Generative AI Era](https://docs.google.com/presentation/d/1_bYMSgdwKaT8IekoxstEvpKIBRBGdvQt8HWxZNU4jnk/)

## OpenTelemetry on GCP實戰

講者：Shawn Ho

講者是 Google Cloud 很厲害的一位工程師，經常分享 Kubernetes 相關知識與技巧。

這一場 Shawn 提供了線上 Lab，透過 Lab 操作帶大家認識如何在 Google Kubernetes Engine 上透過 OpenTelemetry 將 traces、metrics、logs 資訊整合到 Google Cloud 的監控體系中。

OpenTelemetry 是 CNCF 底下其中一個專案，定位是可觀測性框架與工具集(Observability framework and toolkit)。

OpenTelemetry 標準化了 traces、metrics、logs 的處理流程，透過 OTel Collector 進行處理與發送到儲存目的地。尤其近幾年越來越多廠商支援 OpenTelemetry 的 Protocol: OpenTelemetry Protocol (OTLP)。Google Cloud 的 Cloud Trace 也建議使用 OpenTelemetry 將 traces 資料送到 Google Cloud 的監控體系中。

<figure>
  <img src="/images/2023-12-17/20231217001.jpg" alt="otel diagram">
  <figcaption>Source: <a href="https://opentelemetry.io/docs/">https://opentelemetry.io/docs/</a></figcaption>
</figure>

OpenTelemetry 主要使用 YAML 來進行設定，我們也可以使用 BindPlane 提供的網頁介面，透過 UI 操作來進行一般設定。

<figure>
  <img src="/images/2023-12-17/20231217002.png" alt="bindplane diagram">
  <figcaption>Source: <a href="https://cloud.google.com/stackdriver/bindplane">https://cloud.google.com/stackdriver/bindplane</a></figcaption>
</figure>

<figure>
  <img src="/images/2023-12-17/20231217003.jpg" alt="bindplane diagram">
  <figcaption>Source: <a href="https://observiq.com/blog/integrating-opentelemetry-into-a-fluentbit-environment-using-bindplane-op">https://observiq.com/blog/integrating-opentelemetry-into-a-fluentbit-environment-using-bindplane-op</a></figcaption>
</figure>


Blog: [輕鬆小品-k8s的點滴](https://medium.com/@shawn.ho)

References:
- [OpenTelemetry](https://opentelemetry.io/)
- [BindPlane](https://cloud.google.com/stackdriver/bindplane)

## GCP 監控、自動化與合規

講者：Johnny Yeng

Google Cloud 內建安全性與風險管理的解決方案為 Cloud Security Command Center(CSCC)，這個服務主要是針對 Organization 層級即企業組織帳號，如果是一般個人用 gmail 註冊使用，則無法使用。

CSCC 能夠偵測到 Google Cloud 資源的使用情況，便能夠發現雲端資源是否有不當的地方，尤其當公司的專案一多，透過 CSCC 就能用單一介面進行監控。部分資安工具例如 Cloudflare 或是 Fortinet，也能夠將防護資訊也送到 CSCC 中。

<figure>
  <img src="/images/2023-12-17/20231217004.png" alt="cscc finding">
  <figcaption>Source: <a href="https://cloud.google.com/blog/products/identity-security/cloud-security-command-center-is-now-in-beta">https://cloud.google.com/blog/products/identity-security/cloud-security-command-center-is-now-in-beta</a></figcaption>
</figure>

企業重視的合規，CSCC 也有對應的儀表板可以檢視並下載相關報告。
<figure>
  <img src="/images/2023-12-17/20231217005.png" alt="CSCC Compliance Dashboard">
  <figcaption>Source: <a href="https://cloud.google.com/security-command-center/docs/concepts-managing-monitoring-for-compliance">https://cloud.google.com/security-command-center/docs/concepts-managing-monitoring-for-compliance</a></figcaption>
</figure>

演講中還提到可以使用 osquery 在 VM 收集資訊，並送到 Google Cloud 的新服務 Chronicle SIEM。
<figure>
  <img src="/images/2023-12-17/20231217006.png" alt="Chronicle SIEM">
  <figcaption>Source: <a href="https://cloud.google.com/chronicle/docs/overview">https://cloud.google.com/chronicle/docs/overview</a></figcaption>
</figure>

References:
- [Cloud Security Command Center](https://cloud.google.com/security/products/security-command-center)
- [Manage and monitor for compliance](https://cloud.google.com/security-command-center/docs/concepts-managing-monitoring-for-compliance)
- [osquery](https://www.osquery.io/)
- [Collect osquery logs](https://cloud.google.com/chronicle/docs/ingestion/default-parsers/collect-osquery)

## 前端模組解放運動

講者：高見龍

前端的複雜度，主要依賴瀏覽器對標準的支援程度。當瀏覽器支援的越多，開發者或工具要處理的事情就越少。使用 ESM 我們可以更方便將程式碼拆開與引用。

在 ESM 的部分：
- 我們可以用 `base` tag 來指定目前頁面 base URL 的值，所有相關路徑的網址都會基於此值。
- 我們可以在 `script` tag 指定 type 為 module，來使用 ESM 的語法。

```html
<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
        <title>Document</title>

        <base href="https://unpkg.com">

        <script type="module">
            import sum from "/lodash-es/sum.js"
            console.log(sum([1, 4, 5, 0]))
        </script>
    </head>

    <body>
    </body>
</html>
```

用瀏覽器開始後，可以在開發者工具的 Console 中看到結果是 `10`。

如果使用 `importmap`，我們先定義 module 的 mapping 後，再引入使用。

```html
<html>
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <title>Document</title>

    <script type="importmap">
      {
        "imports": {
          "sum": "https://unpkg.com/lodash-es/sum.js"
        }
      }
    </script>

    <script type="module">
      import sum from "sum"
      console.log(sum([1, 4, 5, 0]))
    </script>
  </head>

  <body>
  </body>
</html>
```

透過這樣的做法：
1. 可以在網頁上直接使用 CDN 上的第三方模組，而不用進行繁瑣的打包工具設定。
2. 專案可以不依賴 Nodejs 環境。

但如果需要使用體積較大或是較多的第三方模組，網頁的載入時間就會受到影響，例如我們改一下程式，測試的時候可以很明顯地感受到顯示結果的速度變慢了許多。

```html
<html>
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <title>Document</title>

    <script type="importmap">
      {
        "imports": {
          "lodash": "https://unpkg.com/lodash-es/lodash.js"
        }
      }
    </script>

    <script type="module">
      import { sum } from "lodash"
      console.log(sum([1, 4, 5, 0]))
    </script>
  </head>

  <body>
  </body>
</html>
```

另外提一下，現在 Rails7 預設前端也是使用 importmap。

簡報連結：[前端模組解放運動 - importmap](https://speakerdeck.com/eddie/qian-duan-mo-zu-jie-fang-yun-dong-importmap)

References:
- [base HTML tag](https://developer.mozilla.org/docs/Web/HTML/Element/base)

## Empowering Community-Driven Learning through Serverless Practice

講者：NiJia LIn

這個時間大多數人都跑去聽保哥的演講了：[使用 Google AI Studio 與 Gemini API 快速打造 Generative AI 原型設計](https://drive.google.com/file/d/1T-rBossVpr1M5mf5L8jWuIZlX5Qz_3UN/view)

這一場鼓勵大家能夠在業餘做一些 Side Project，除了練手、在社群中建立 Credit，也有機會在求職過程展現一下能力。

使用 Google Cloud 生態系的原因是因為它是一個大房東，並提供了各種服務。

自己處理基礎建設很麻煩，小型的專案可以先使用 Serverless 的服務，例如 Google Cloud 的服務：
- [Cloud Run](https://cloud.google.com/run)
- [Cloud Functions](https://cloud.google.com/functions)

Serverless 的好處是只需要專心開發業務邏輯，基礎建設都由服務商提供，並且通常有 Free Tier 足夠小流量的服務運作。

以我的經驗來說，如果要走完成的開發流程，Cloud Functions 最好配合框架開發，不然測試有點麻煩。Cloud Run 就是依照容器的方式進行開發，只是部署的目的地不同，我通常會用 Cloud Run，易用程式的可攜性比較好，我之後可以很容易的選擇要在 Serverless 服務、VM、或是 Kubernetes Cluster 上部署。

如果單純寫一小段程式碼不在乎版控的話，在 Cloud Console 上可以直接用線上編輯器開發跑在 Cloud Functions 上的程式(就跟在 Production 機器上面改 code 一樣嗨 lol)。

Firebase 提供多項工具與服務，資料庫相關的服務有：
- [Firebase Realtime Database](https://firebase.google.com/docs/database)
- [Cloud Firestore](https://firebase.google.com/docs/firestore)

儲存檔案的部分可以使用：
- [Cloud Storage](https://cloud.google.com/storage)

講者也提了一些 Side Project 範例，主要是 Line bot 上的應用，例如個人助理、行事曆小工具、場地觀測等。從自身需求出發是一個很好的方向。

Blog: [忍者工坊](https://nijialin.com/)
