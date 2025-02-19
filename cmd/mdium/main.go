package main

import (
	"flag"
	"log"

	"github.com/0xydev/mdium"
)

func main() {
	username := flag.String("user", "", "Medium username")
	domain := flag.String("domain", "", "Medium publication domain (e.g., blog.medium.com)")
	outputDir := flag.String("output", "articles", "Output directory for markdown files")
	limit := flag.Int("limit", 0, "Limit the number of latest articles to process (0 means all articles)")
	flag.Parse()

	if *username == "" && *domain == "" {
		log.Fatal("Either username (-user) or domain (-domain) is required")
	}

	if *username != "" && *domain != "" {
		log.Fatal("Please specify either username (-user) or domain (-domain), not both")
	}

	source := *username
	if *domain != "" {
		source = *domain
	}

	client, err := mdium.NewClient(source, *outputDir)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	if err := client.DownloadArticles(*limit); err != nil {
		log.Fatalf("Error: %v", err)
	}
} 