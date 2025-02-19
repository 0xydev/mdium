package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/0xydev/mdium/internal/models"
)

// FileStorage handles saving articles to the filesystem
type FileStorage struct {
	outputDir string
}

// NewFileStorage creates a new FileStorage instance
func NewFileStorage(outputDir string) *FileStorage {
	return &FileStorage{
		outputDir: outputDir,
	}
}

// SaveArticle saves the markdown content to a file
func (f *FileStorage) SaveArticle(article models.Article, content string) error {
	if err := os.MkdirAll(f.outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	fileName := f.sanitizeFileName(article.Title) + ".md"
	filePath := filepath.Join(f.outputDir, fileName)

	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", fileName, err)
	}

	return nil
}

// sanitizeFileName removes invalid characters from the filename
func (f *FileStorage) sanitizeFileName(title string) string {
	invalid := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|"}
	result := title

	for _, char := range invalid {
		result = strings.ReplaceAll(result, char, "-")
	}

	return result
}
