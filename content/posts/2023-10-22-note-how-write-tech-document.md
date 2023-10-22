---
categories:
  - note
keywords:
  - note
  - writing
comments: true
date: 2023-10-22T12:00:00+08:00
title: "[筆記] 如何寫一份技術文件"
url: /2023/10/22/note-how-write-tech-document/
---

原文：[Diátaxis](https://diataxis.fr/)

## 筆記

原文解說了技術文件的類型、原則，並提供 "The language of tutorials" 幫助寫作。

技術文件依據象限可以分為四大類：
- Tutorials
- How-To Guides
- Explanation
- Reference

![graph](/images/2023-10-22/20231022001.png)
image source: [https://diataxis.fr/](https://diataxis.fr/)

## Tutorials

主要核心目標是學習實踐，讓初學者能夠對產品有基本的了解與操作能力。

通常會是一個 Getting Start，讓初學者能夠輕鬆上手使用。

在這類型的文件並不會解釋太多理論知識，因為目的是讓使用者能夠實踐產品的入門操作，從而在腦中建立認知地圖。當初學者完成 Tutorials，在腦中建立認知地圖後，應該就要能夠有能力理解其他文件從而更深入理解產品。

一份好的 Tutorials 需要明確地指示每一個步驟，並且：

1. 有意義：完成 Tutorials 讀者會有成就感
2. 明確成功：每一個 Tutorials 的步驟都卻要明確的結果反饋
3. 有邏輯：每一個 Tutorials 的步驟順序能夠連貫相通
4. 實際有用：Tutorials 的步驟要讓讀者能夠熟悉你希望讀者熟悉的工作、概念和工具

Tutorials 是為讓了初學者能夠入門並轉變為使用者，而不是變成專家，試圖在步驟間加入太多艱深的內容會嚇跑初學者，這些額外知識我們可以在另外的篇幅開展說明。在 Tutorials 中我們需要讓初學者能夠專注在我們的指示上去完成每一個步驟。

在 Tutorials 開頭說明本文目的與提供路徑圖對於初學者會很有幫助。

## How-to Guides

主要核心目標是協助讀者解決實際問題。

How-to Guides 可能也會被稱為 Recipes，說明實現特定目標的步驟。這也可能是在技術文件中最多人閱讀的類型。

How-to Guides 與 Tutorials 不同，雖然兩者都會有一系列的實踐步驟，但兩者的目的是不同的。How-to Guides 是幫助已經有能力的使用者，完成特定任務，或是解決特定問題。Tutorials 是幫助初學者入門，擁有基本操作能力。

How-to Guides 面對的是已擁有基礎能力的讀者，在原文中將 How-to Guides 比做食譜，Tutorials 比做烹飪課，即使遵循食譜也至少需要基本能力。以前從未烹飪的人不應該指望他成功地遵循食譜，所以食譜不能代替烹飪課。

由於是面對已擁有基礎能力的讀者，如果在文件中不斷提及基礎知識或是背景，對讀者來說可能會是不必要的內容。一份好的 How-To Guides 最好能夠遵循固定的格式，不包括教學與討論，並且：

1. 必須是可靠的
2. 可以解決問題
3. 不要解釋概念，如果需要，請使用超連結關聯
4. 準確命名標題，對讀者與 SEO 都有所幫助

## Explanation

主要核心目標是協助讀者理解相關知識。

Explanation 幫助讀者深入與拓寬對技術的理解，以 Google Cloud 的產品文件為例，通常會有一個 Concept Section 來解釋產品的架構與概念。

Tutorials 與 How-To Guides 中所提供的知識都在特定範圍內，透過 Explanation 類型的文件主題我們可以幫相關知識都串連起來。

在原文中以《論食物和烹飪》（On Food and Cook）為例，這本書沒有教如何烹飪，但是它介紹了烹飪的歷史背景，解釋了我們在廚房做的事情，以及這些情況是如何隨時代變化的。

一份好的 Explanation 應該能夠：

1. 建立連結，將相關的知識組織起來
2. 提供脈絡(Context)
3. 主題性，避免太過發散
4. 討論性，可以在文中提供替代方案、反例或其他意見幫助讀者思考

## Reference

主要核心目標是提供讀者所需的資料。

Reference 應要是簡潔有序地進行描述，不同於 Tutorials 與 How-To Guides 是使用者導向的文件，Reference 是產品導向的文件。

Reference 圍繞著產品的技術規格提供資訊，例如 API 介面、參數，如何使用與限制等等。

一份好的 Reference 應該能夠：

1. 反應產品結構，這有助於讀者查詢
2. 保持一致，不要造成讀者困惑
3. 清晰描述，不要解釋、討論
4. 提供範例，幫助讀者理解如何使用
5. 最重要的是準確，不要誤導讀者
