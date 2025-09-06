---
categories:
  - note
keywords:
  - note
  - disk
  - partition
  - ubuntu
  - testdisk
comments: true
date: 2025-08-31T08:00:00+08:00
title: "[筆記] 拯救硬碟資料小記"
url: /2025/08/31/note-save-the-disk/
images:
  - /images/2025-08-31/save-the-disk.png
---

在一次線上增加 VM 硬碟空間時，明明發現介面不是以往的 SCSI 而是 NVME，分割區也跟印象中的不同，卻還是按照以往的步驟進行調整，結果導致分割區損毀。

## 問題發生

在 GCP 的 VM 需要增加硬碟空間，在 GCP Console 上面將硬碟從 20G 增加到 50G 後還需要在 VM 裡面進行調整。

於是我參考以前的筆記，用 `parted` 對 `nvme0n1p1` 做 `resizepart`：

```bash
nvme0n1      259:0    0    20G  0 disk
├─nvme0n1p1  259:1    0  19.9G  0 part /
├─nvme0n1p14 259:2    0     4M  0 part
└─nvme0n1p15 259:3    0   106M  0 part /boot/efi
```

想從 20G 變到 50G

```bash
nvme0n1      259:0    0    50G  0 disk
├─nvme0n1p1  259:1    0  19.9G  0 part /
├─nvme0n1p14 259:2    0     4M  0 part
└─nvme0n1p15 259:3    0   106M  0 part /boot/efi
```

以前的筆記上面這樣紀錄：

```bash
sudo sgdisk --move-second-header /dev/sda
sudo partprobe /dev/sda
sudo resize2fs /dev/sda1
```

然而當我執行以下指令之後，檔案系統就壞掉了

```bash
sudo sgdisk --move-second-header /dev/nvme0n1p1
```

壞掉的情況是 SSH Session 還連著，可以輸入指令，但執行任何指令卻出現錯誤

```bash
-bash: /usr/bin/ls: Input/output error
```

我才突然意識到，這次的硬碟是 NVME 而不是 SCSI，不知道為什麼只有這一台 VM 是使用 NVME。

## 解決過程

### 能否重開機自動修復？

思考了許久，無法進行任何實際操作，便立即製作硬碟快照。

我嘗試使用快照作為開機硬碟建立一台新的 VM，結果是無法開機，這表示分割區損壞了。

重開機無法治療這一次的手殘。

### 使用工具修復

我重新建立一台新的 VM，將快照作為附加硬碟，使用 SSH 登入後開始進行嘗試修復硬碟資料。

先執行指令進行掛載：

```bash
sudo mkdir /mnt/recovered
sudo mount -t ext4 /dev/sdb1 /mnt/recovered
```

但卻出現了錯誤：

```bash
mount: /mnt/recovered: wrong fs type, bad option, bad superblock on /dev/sdb1, missing codepage or helper program, or other error.
dmesg(1) may have more information after failed mount system call.
```

我對這方面的知識是一竅不通，用 google 搜尋也不知道該用什麼關鍵字比較好，於是我就在 [Google AI Studio](https://aistudio.google.com) 問問了 Gemini，它推薦了一個我從沒聽過的的工具 [testdisk](https://blog.gtwang.org/linux/testdisk-linux-recover-deleted-files/)。

執行指令安裝：

```bash
sudo apt update
sudo apt install testdisk
```

執行 `testdisk` 修復硬碟，因為是 ubuntu 系統，所以附加硬碟是 `/dev/sdb`，可以使用 `lsblk` 指令查看。

```bash
sudo testdisk /dev/sdb
```

按照直覺操作，進行分析後，測試查看能否列出檔案，發現可以，則確認寫入變更。

重新執行掛載指令就成功，沒有出現錯誤了：

```bash
sudo mount -t ext4 /dev/sdb1 /mnt/recovered
```

雖然修復了硬碟可以正常存取資料，但卻已經不能再作為 GCE 的開機硬碟，我試著作為開機硬碟建立 VM，但卻無法成功開機。

是什麼原因我就沒有再深入研究了。

### 搬移資料

因為不想花太多時機修復成開機硬碟，所以我決定將資料搬出來。

原本的 VM 上是用 Docker Compose 在管理應用程式，所以在硬碟上面有幾個重要的資料：

1. 應用程式設定檔
2. 容器映像檔
3. 資料庫

應用程式設定檔是文字檔，很簡單的從搬移檔案就好。

容器映像檔與資料庫都是跟 Docker 有關係，相比資料就比較麻煩了。

Docker 通常會將資料存放在 `/var/lib/docker` 目錄下，我們可以在 `/etc/docker/daemon.json` 中設定 `data-root` 的值來改成其他目錄位置。

我們先停止 Docker daemon：

```bash
sudo systemctl stop docker.service
sudo systemctl stop docker.socket
```

修改檔案 `/etc/docker/daemon.json`，其中 `/mnt/recovered` 是剛剛掛載硬碟的目錄。

```json
{
    "data-root": "/mnt/recovered/var/lib/docker"
}
```

啟動 Docker

```bash
sudo systemctl start docker.socket
sudo systemctl start docker.service
```

這時候 Docker 就會使用新的 `data-root` 位置，我們可以使用 `docker images` 指令查看容器映像檔是否有成功讀取。

我們可以使用指定將掛載硬碟中的容器映像檔匯出：

```bash
docker save app:prod > app.tar
```

因為也有用到資料庫容器，所以我們要先啟動資料庫容器，然後再使用 `docker exec` 來進行資料庫備份，將資料庫備份檔複製到本地檔案系統。

當容器映像檔與資料庫資料都備份完後，我們停止 Docker：

```bash
sudo systemctl stop docker.service
sudo systemctl stop docker.socket
```

修改檔案 `/etc/docker/daemon.json` 移除 `data-root`，這樣就會指向預設的位置 `/var/lib/docker`。

啟動 Docker：

```bash
sudo systemctl start docker.socket
sudo systemctl start docker.service
```

我們匯入容器映像檔：

```bash
docker load < app.tar
```

後續重新建立容器，並且將資料庫資料匯入，這顆硬碟就可以作為開機硬碟使用了。

## 參考
- [更改 Docker 預設路徑](https://ithelp.ithome.com.tw/articles/10235112)
