# mdium

mdium is a Go package for downloading Medium articles as markdown files.

[![EN](https://img.shields.io/badge/lang-EN-blue.svg)](README.md)
[![TR](https://img.shields.io/badge/lang-TR-red.svg)](README_TR.md)

## Tech Stack

[![Go](https://img.shields.io/badge/go-1.23-00ADD8.svg?style=flat&logo=go)](https://go.dev/)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/0xydev/mdium)](https://pkg.go.dev/github.com/0xydev/mdium)
[![Go Report Card](https://goreportcard.com/badge/github.com/0xydev/mdium)](https://goreportcard.com/report/github.com/0xydev/mdium)
[![RSS](https://img.shields.io/badge/RSS-Feed-FFA500.svg?style=flat&logo=rss)](https://en.wikipedia.org/wiki/RSS)
[![Markdown](https://img.shields.io/badge/Markdown-000000.svg?style=flat&logo=markdown)](https://daringfireball.net/projects/markdown/)
[![YAML](https://img.shields.io/badge/YAML-CB171E.svg?style=flat&logo=yaml)](https://yaml.org/)

## Requirements
- Go 1.23 or higher

## Installation

### As a Package
```bash
go get github.com/0xydev/mdium@v0.1.1
```

### As a Command Line Tool
```bash
go install github.com/0xydev/mdium/cmd/mdium@v0.1.1
```

## Features
- Download articles from:
  - User profiles (e.g., @Medium)
  - Medium blogs (e.g., blog.medium.com)
  - Publications (e.g., medium.com/netflix-techblog)
- Convert articles to markdown with YAML frontmatter
- Limit the number of articles to download
- Customizable output directory

## Usage

### Command Line
```bash
# Download from a user profile
go run cmd/mdium/main.go -user Medium -output user_articles -limit 5

# Download from Medium's blog
go run cmd/mdium/main.go -domain blog.medium.com -output blog_articles -limit 3

# Download from a publication
go run cmd/mdium/main.go -domain medium.com/netflix-techblog -output pub_articles
```

### As a Package
```go
package main

import (
    "log"
    "github.com/0xydev/mdium"
)

func main() {
    // Download from a user profile
    client, err := mdium.NewClient("Medium", "articles")
    if err != nil {
        log.Fatal(err)
    }
    
    err = client.DownloadArticles(5)
    if err != nil {
        log.Fatal(err)
    }

    // Download from a publication
    pubClient, err := mdium.NewClient("medium.com/netflix-techblog", "netflix_articles")
    if err != nil {
        log.Fatal(err)
    }
    
    err = pubClient.DownloadArticles(0) // Download all available articles
    if err != nil {
        log.Fatal(err)
    }
}
```

## Output Format
Articles are saved as markdown files with YAML frontmatter:
```yaml
---
title: "Article Title"
date: 2024-02-19T10:00:00Z
link: "https://medium.com/..."
author: "Author Name"
---

Article content in markdown...
```

## Limitations
- Medium's RSS feed is limited to the latest 10 articles
- Some Medium articles might have formatting differences when converted to markdown

## License
MIT License - see [LICENSE](LICENSE) file for details
