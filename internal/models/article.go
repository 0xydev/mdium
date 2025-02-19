package models

import "time"

// Article represents a Medium article with its metadata and content
type Article struct {
	Title       string
	Content     string
	Link        string
	PublishDate time.Time
	Author      string
}

// FrontMatter represents the metadata section of a markdown file
type FrontMatter struct {
	Title  string    `yaml:"title"`
	Date   time.Time `yaml:"date"`
	Link   string    `yaml:"link"`
	Author string    `yaml:"author"`
}
