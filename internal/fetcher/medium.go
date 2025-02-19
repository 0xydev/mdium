// Package fetcher provides functionality to fetch articles from Medium
package fetcher

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/0xydev/mdium/internal/models"
)

// RSSFeed represents the structure of Medium's RSS feed
type RSSFeed struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

// Channel represents the channel element in RSS feed
type Channel struct {
	Items []Item `xml:"item"`
}

// Item represents a single article in the RSS feed
type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	Description string `xml:"description"`
	Content     string `xml:"encoded"`
}

// MediumFetcher handles fetching articles from Medium
type MediumFetcher struct {
	source   string
	isDomain bool
}

// NewMediumFetcher creates a new MediumFetcher instance
func NewMediumFetcher(source string, isDomain bool) *MediumFetcher {
	return &MediumFetcher{
		source:   source,
		isDomain: isDomain,
	}
}

// FetchArticles retrieves all articles for the configured source
func (m *MediumFetcher) FetchArticles() ([]models.Article, error) {
	feed, err := m.fetchRSSFeed()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch RSS feed: %w", err)
	}

	return m.convertToArticles(feed)
}

// fetchRSSFeed retrieves the RSS feed from Medium
func (m *MediumFetcher) fetchRSSFeed() (*RSSFeed, error) {
	var url string
	source := strings.TrimPrefix(m.source, "https://")
	source = strings.TrimPrefix(source, "http://")
	
	if m.isDomain {
		if strings.Contains(source, "medium.com/") {
			// Publication feed (e.g., medium.com/netflix-techblog)
			url = fmt.Sprintf("https://medium.com/feed/%s", strings.TrimPrefix(source, "medium.com/"))
		} else {
			// Blog feed (e.g., blog.medium.com)
			url = fmt.Sprintf("https://%s/feed", source)
		}
	} else {
		// User feed (e.g., @username)
		url = fmt.Sprintf("https://medium.com/feed/@%s", m.source)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var feed RSSFeed
	if err := xml.Unmarshal(body, &feed); err != nil {
		return nil, err
	}

	return &feed, nil
}

// convertToArticles converts RSS items to Article models
// Articles are sorted by publish date (newest first)
func (m *MediumFetcher) convertToArticles(feed *RSSFeed) ([]models.Article, error) {
	articles := make([]models.Article, 0, len(feed.Channel.Items))

	for _, item := range feed.Channel.Items {
		pubDate, err := time.Parse("Mon, 02 Jan 2006 15:04:05 GMT", item.PubDate)
		if err != nil {
			pubDate, err = time.Parse("Mon, 2 Jan 2006 15:04:05 GMT", item.PubDate)
			if err != nil {
				return nil, fmt.Errorf("failed to parse date %s: %w", item.PubDate, err)
			}
		}

		articles = append(articles, models.Article{
			Title:       item.Title,
			Content:     item.Content,
			Link:        item.Link,
			PublishDate: pubDate,
			Author:      m.source,
		})
	}

	// Sort articles by publish date (newest first)
	sort.Slice(articles, func(i, j int) bool {
		return articles[i].PublishDate.After(articles[j].PublishDate)
	})

	return articles, nil
} 