# mdium

Medium makalelerini markdown formatına dönüştüren Go paketi.

[![EN](https://img.shields.io/badge/lang-EN-blue.svg)](README.md)
[![TR](https://img.shields.io/badge/lang-TR-red.svg)](README_TR.md)

## Teknolojiler

[![Go](https://img.shields.io/badge/go-1.23-00ADD8.svg?style=flat&logo=go)](https://go.dev/)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/0xydev/mdium)](https://pkg.go.dev/github.com/0xydev/mdium)
[![Go Report Card](https://goreportcard.com/badge/github.com/0xydev/mdium)](https://goreportcard.com/report/github.com/0xydev/mdium)
[![RSS](https://img.shields.io/badge/RSS-Feed-FFA500.svg?style=flat&logo=rss)](https://en.wikipedia.org/wiki/RSS)
[![Markdown](https://img.shields.io/badge/Markdown-000000.svg?style=flat&logo=markdown)](https://daringfireball.net/projects/markdown/)
[![YAML](https://img.shields.io/badge/YAML-CB171E.svg?style=flat&logo=yaml)](https://yaml.org/)

## Gereksinimler
- Go 1.23 veya üstü

## Kurulum

### Paket Olarak
```bash
go get github.com/0xydev/mdium@v0.1.1
```

### Komut Satırı Aracı Olarak
```bash
go install github.com/0xydev/mdium/cmd/mdium@v0.1.1
```

## Özellikler
- Farklı kaynaklardan makale indirme:
  - Kullanıcı profilleri (örn: @Medium)
  - Medium blogları (örn: blog.medium.com)
  - Yayınlar (örn: medium.com/netflix-techblog)
- Makaleleri YAML frontmatter ile markdown formatına dönüştürme
- İndirilecek makale sayısını sınırlama
- Özelleştirilebilir çıktı dizini

## Kullanım

### Komut Satırı
```bash
# Kullanıcı profilinden indirme
go run cmd/mdium/main.go -user Medium -output user_articles -limit 5

# Medium blogdan indirme
go run cmd/mdium/main.go -domain blog.medium.com -output blog_articles -limit 3

# Yayından indirme
go run cmd/mdium/main.go -domain medium.com/netflix-techblog -output pub_articles
```

### Paket Olarak
```go
package main

import (
    "log"
    "github.com/0xydev/mdium"
)

func main() {
    // Kullanıcı profilinden indirme
    client, err := mdium.NewClient("Medium", "articles")
    if err != nil {
        log.Fatal(err)
    }
    
    err = client.DownloadArticles(5)
    if err != nil {
        log.Fatal(err)
    }

    // Yayından indirme
    pubClient, err := mdium.NewClient("medium.com/netflix-techblog", "netflix_articles")
    if err != nil {
        log.Fatal(err)
    }
    
    err = pubClient.DownloadArticles(0) // Tüm mevcut makaleleri indir
    if err != nil {
        log.Fatal(err)
    }
}
```

## Çıktı Formatı
Makaleler YAML frontmatter ile markdown dosyaları olarak kaydedilir:
```yaml
---
title: "Makale Başlığı"
date: 2024-02-19T10:00:00Z
link: "https://medium.com/..."
author: "Yazar Adı"
---

Makalenin markdown içeriği...
```

## Kısıtlamalar
- Medium'un RSS beslemesi en son 10 makale ile sınırlıdır
- Bazı Medium makaleleri markdown'a dönüştürülürken format farklılıkları gösterebilir

## Lisans
MIT Lisansı - detaylar için [LICENSE](LICENSE) dosyasına bakın 