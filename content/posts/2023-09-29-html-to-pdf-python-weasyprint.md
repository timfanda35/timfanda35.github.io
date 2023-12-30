---
categories:
  - rails
keywords:
  - Rails
  - Ruby
  - PDFkit
  - wkhtmltopdf
  - grover
  - python
  - weasyprint
  - fastapi
  - gem
  - PDF
  - HTML
  - print
comments: true
date: 2023-09-29T12:00:00+08:00
title: "HTML to PDF PDFkit 的替代方案 - WeasyPrint 介紹"
url: /2023/09/29/html-to-pdf-python-weasyprint/
images:
  - /images/2023-09-29/html-to-pdf-python-weasyprint.png
---

本文接續之前文章：
- 方案一：[HTML to PDF PDFkit 的替代方案 - Grover 介紹](/2023/09/09/html-to-pdf-grover-gem/)
- 方案二：[HTML to PDF PDFkit 的替代方案 - puppeteer-ruby 介紹](/2023/09/23/html-to-pdf-puppeteer-ruby-gem/)

## 方案三：WeasyPrint

現在問題來到 Chrome/Chromium Container 對需求來說似乎使用了太多資源。

嘗試尋找其他 Ruby gem 還沒有看到滿意的，如果您有推薦的 Ruby gem 還請讓我知道。

我想了想：
1. 我現在另外起了一個 Chrome/Chromium Container 來產生 PDF，它就是個外部服務。
2. 主要需求是輸入 HTML 輸出 PDF，而且日後應該也不用再支援其他格式，不太會有頻繁更新的需要

於是我決定去尋找其他語言的 HTML to PDF 方案，打算包成 API 取代方案二。最後我使用了 [WeasyPrint](https://github.com/Kozea/WeasyPrint)。

WeasyPrint 使用 Python 開發，對此方案有信心是因為我發現：
- 近期仍有活躍開發活動
- GitHub Start 超過 6k
- 有獨立的[官方網站](https://weasyprint.org/) 與說明文件

WeasyPrint 提供 Command line Tool 與 Python Package 的用法，本文主要著重在於 Python Package 的用法。

## Demo

我們可以先從 [Quickstart](https://doc.courtbouillon.org/weasyprint/stable/first_steps.html#quickstart) 了解如何使用 WeasyPrint 用 HTML 產生 PDF。

本文使用的環境：
- macOS
- Ruby 3.2.2
- Rails 7.0.7.2
- Docker Desktop 4.21.1

建立測試專案目錄

```bash
mkdir demo
cd demo
```

我們使用 Python Container 來進行測試。Binding 本地的 `8000` port 是為了之後的步驟可以方便本地瀏覽器存取。

```bash
docker run -it --rm -p 8000:8000 -v $(pwd):/app python:3 bash
```

[Container] 安裝 CJK 字型

```bash
apt update && apt install -y fonts-noto-cjk
```

[Container] 移動工作目錄

```bash
cd /app
```

[Container] 安裝 WeasyPrint package

```bash
pip install weasyprint
```

[Container] 啟動 Python Console

```bash
python
```

[Python Console] 我們預期會輸入 HTML 字串，然後產生 PDF 檔案：

```python
from weasyprint import HTML

html='''
<h1>Hello 世界！</h1>
'''

HTML(string=html).write_pdf('./demo.pdf')
```

我們在本機的專案目錄打開 `demo.pdf`，可以確認 PDF 有正確顯示。

![demo.pdf](/images/2023-09-29/20230929001.jpg)

## Fast API

因為平常都是在跟第三方服務的 API 打交道，所以比起用 Ruby 做 WeasyPrint Command Line Tool 的 Wrapper，我更偏好將 WeasyPrint 包成 RESTful API 來使用。

之前聽過說 [FastAPI](https://github.com/tiangolo/fastapi)，掃了一下官方文件，覺得內容豐富就決定這次使用它來包裝 WeasyPrint。

### Demo

我們繼續使用上一章節的 Python Container。

[Python Console] 我們先從 Python Console 退出。

```python
exit()
```

[Container] 安裝 FastAPI package

```bash
pip install fastapi uvicorn[standard]
```

由於我們已將專案目錄掛載到 Python Container 的 `/app`，也就是我們當前的工作目錄，所以我們可以在本地開啟文字編輯器於專案目錄下建立檔案 `main.py`，新增以下內容：

```python
import io
from typing import Union

from fastapi import FastAPI, Response
from fastapi.responses import StreamingResponse
from pydantic import BaseModel

from weasyprint import HTML

class PrintPdfRequest(BaseModel):
    html: str

# https://fastapi.tiangolo.com/
app = FastAPI()

@app.post("/pdfs")
async def print_pdf(response: Response, body: PrintPdfRequest):
    filename = 'demo'
    byte_string = HTML(string=body.html).write_pdf()

    headers = {
        'Content-Type': 'application/pdf',
        'Content-Disposition': '%s; name="%s"; filename="%s.%s"' % (
            'attachment',
            filename,
            filename,
            'pdf'
        )
    }
    return StreamingResponse(io.BytesIO(byte_string), headers=headers)
```

我們寫了一個 POST API，接受名為 `html` 的參數作為輸入，回傳結果為下載 PDF 檔案。

[Python Container] 啟動 API

```bash
uvicorn main:app --reload --host 0.0.0.0
```

我們用瀏覽器開啟 http://localhost:8000/docs

![swagger ui](/images/2023-09-29/20230929002.jpg)

這是我覺得 FastAPI 其中一個很酷的地方，寫完程式碼後不用做特別的設定，就產生了 Swagger UI 可以用瀏覽器直接進行測試。

點擊 Try it out。

![click Try it out](/images/2023-09-29/20230929003.jpg)

於 Request Body 輸入以下內容後，點擊 Execute。

```json
{
  "html": "<h1>Hello 世界！</h1>"
}
```

![click Execute](/images/2023-09-29/20230929004.jpg)

往下可以看到執行結果為成功，並且出現 Download file 的連結，我們可以點擊連結下載檔案。

![download file](/images/2023-09-29/20230929005.jpg)

確認下載的 PDF 有正確顯示。

![check result](/images/2023-09-29/20230929006.jpg)

## Rails

透過 FastAPI 將 WeasyPrint 包裝成 RESTful API，在 Rails 中您可以用您喜歡的 Ruby HTTP Client 去呼叫使用。

我在 Rails 中將其包成 `WeasyPrintService::Printer` Class，先前的 Demo 程式碼會變成像這樣：

```ruby
class WeasyPrintPdfsController < ApplicationController
  def index
    filename = 'weasyprint.pdf'
    pack_slip = render_to_string layout: false

    response.headers['content-disposition'] = "attachment; filename=#{filename}"
    render_pdf pack_slip, filename: filename
  rescue => e
    response.headers['content-disposition'] = ''
    Rails.logger.error "#{e.class}: #{e.message}"
    render plain: "#{e.class}: #{e.message}"
  end

  def render_pdf(html, filename:)
    printer = WeasyPrintService::Printer.new
    pdf = printer.to_pdf(html)

    unless printer.ok?
      response.headers['content-disposition'] = ''
      render(
        plain: "Error: Can not print PDF, because system can not connect to WeasyPrint service.",
        status: :internal_server_error
      )
      return
    end

    send_data pdf, filename: filename, type: "application/pdf"
  end
end
```

如果有多個 Controller 需要產生 PDF，那麼就可以將 `render_pdf` 方法抽出來變成 Concern 共用。

## 成果

目前本文的方案三已經在正式環境上線數週，運作正常 🚀

### WeasyPrint PDF API

打包成 Container 使用的部分可以參考我開源出來的專案：[WeasyPrint PDF API](https://github.com/timfanda35/weasyprint-pdf-api)

可以從 GitHub 上拉取預先建置好的 Container Image 玩玩看。

### Container Image Size

以類似方式去打包各方案所需的 Container Image，總大小從低往高排序：

| 方案 | Rails APP Image Size | External Image | External Image Size | Total Size |
|---  |---:             |---            |---:                  |---:        |
|pdfkit        | 953MB|                                     |   0MB| 953MB|
|weasyprint    | 662MB|ghcr.io/timfanda35/weasyprint-pdf-api| 552MB|1214MB|
|puppeteer-ruby| 665MB|browserless/chrome                   |3160MB|3825MB|
|grover        |1340MB|browserless/chrome                   |3160MB|4500MB|

![container image size](/images/2023-09-29/20230929007.jpg)

總大小是原始方案 pdfkit 最小。

grover 與 puppeteer-ruby 因為需要 Chrome/Chromium，所以直接加上了 3GB。(可能有體積更小的 Chrome/Chromium Container Image，這裡是使用本文測試過的 Image。)

我們最後的方案是 weasyprint，比起原始方案總大小多了約 250MB，weasyprint-pdf-api 的 Container Image Size 應該還有再瘦身的空間。主要會常更新的只有 Rails APP Image，因為 Size 變小了，跑 CI/CD 都變快了許多。

### 結論

建立外部服務來達成需求：
- 找到替代方案避免 pdfkit 使用的 wkhtmltopdf 停止維護所造成的安全性風險或升級困難

缺點：
- 多了一個服務需要維護

優點：
- 減少了主要 Rails APP Container Image 的 Size，加速 CI/CD

## 參考資料
- [WeasyPrint PDF API](https://github.com/timfanda35/weasyprint-pdf-api)
- [PDFkit](https://github.com/pdfkit/pdfkit)
- [wkhtmltopdf](https://github.com/wkhtmltopdf/wkhtmltopdf)
- [Grover](https://github.com/Studiosity/grover)
- [Google Puppeteer](https://github.com/puppeteer/puppeteer)
- [puppeteer-ruby](https://github.com/YusukeIwaki/puppeteer-ruby)
- [WeasyPrint](https://github.com/Kozea/WeasyPrint)
- [FastAPI](https://github.com/tiangolo/fastapi)
- [browserless](https://www.browserless.io/docs/docker-quickstart)
- [browserless/chrome](https://github.com/browserless/chrome)
- [Converting HTML to PDF using Rails](https://dev.to/ayushn21/converting-html-to-pdf-using-rails-54e7)
