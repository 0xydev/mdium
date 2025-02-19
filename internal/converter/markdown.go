package converter

import (
	"fmt"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/0xydev/mdium/internal/models"
	"gopkg.in/yaml.v3"
)

// MarkdownConverter handles converting HTML content to Markdown
type MarkdownConverter struct {
	converter *md.Converter
}

// NewMarkdownConverter creates a new MarkdownConverter instance
func NewMarkdownConverter() *MarkdownConverter {
	return &MarkdownConverter{
		converter: md.NewConverter("", true, nil),
	}
}

// ConvertArticle converts an article's HTML content to Markdown with frontmatter
func (m *MarkdownConverter) ConvertArticle(article models.Article) (string, error) {
	// Convert HTML content to Markdown
	markdown, err := m.converter.ConvertString(article.Content)
	if err != nil {
		return "", fmt.Errorf("failed to convert HTML to markdown: %w", err)
	}

	// Create frontmatter
	frontMatter := models.FrontMatter{
		Title:  article.Title,
		Date:   article.PublishDate,
		Link:   article.Link,
		Author: article.Author,
	}

	// Marshal frontmatter to YAML
	frontMatterBytes, err := yaml.Marshal(frontMatter)
	if err != nil {
		return "", fmt.Errorf("failed to create frontmatter: %w", err)
	}

	// Combine frontmatter and content
	result := fmt.Sprintf("---\n%s---\n\n%s", string(frontMatterBytes), markdown)
	return result, nil
} 