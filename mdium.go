package mdium

import (
	"strings"

	"github.com/0xydev/mdium/internal/converter"
	"github.com/0xydev/mdium/internal/fetcher"
	"github.com/0xydev/mdium/internal/storage"
)

// SourceType represents the type of Medium source (user or domain)
type SourceType int

const (
	// UserSource represents a Medium user profile
	UserSource SourceType = iota
	// DomainSource represents a Medium blog or publication
	DomainSource
)

// Client represents a Medium article downloader
type Client struct {
	source     string
	sourceType SourceType
	fetcher    *fetcher.MediumFetcher
	converter  *converter.MarkdownConverter
	storage    *storage.FileStorage
}

// NewClient creates a new Medium client
// source can be either a username or a domain (blog.medium.com or medium.com/publication)
// outputDir is the directory where markdown files will be saved
func NewClient(source string, outputDir string) (*Client, error) {
	sourceType := detectSourceType(source)

	return &Client{
		source:     source,
		sourceType: sourceType,
		fetcher:    fetcher.NewMediumFetcher(source, sourceType == DomainSource),
		converter:  converter.NewMarkdownConverter(),
		storage:    storage.NewFileStorage(outputDir),
	}, nil
}

// detectSourceType determines if the source is a username or domain
func detectSourceType(source string) SourceType {
	if strings.Contains(source, ".") {
		return DomainSource
	}
	return UserSource
}

// DownloadArticles downloads Medium articles
// limit: number of latest articles to download (0 means all articles)
// Returns an error if the download process fails
func (c *Client) DownloadArticles(limit int) error {
	articles, err := c.fetcher.FetchArticles()
	if err != nil {
		return err
	}

	if limit > 0 && limit < len(articles) {
		articles = articles[:limit]
	}

	for _, article := range articles {
		markdown, err := c.converter.ConvertArticle(article)
		if err != nil {
			return err
		}

		if err := c.storage.SaveArticle(article, markdown); err != nil {
			return err
		}
	}

	return nil
}
