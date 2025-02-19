package main

import (
	"log"

	"github.com/0xydev/mdium"
)

func main() {
	// Example 1: Download from a user profile
	userClient, err := mdium.NewClient("Medium", "user_articles")
	if err != nil {
		log.Fatalf("Error creating client for user: %v", err)
	}

	err = userClient.DownloadArticles(5)
	if err != nil {
		log.Printf("Error downloading user articles: %v", err)
	}

	// Example 2: Download from Medium's official blog
	blogClient, err := mdium.NewClient("blog.medium.com", "blog_articles")
	if err != nil {
		log.Fatalf("Error creating client for blog: %v", err)
	}

	err = blogClient.DownloadArticles(3)
	if err != nil {
		log.Printf("Error downloading blog articles: %v", err)
	}

	// Example 3: Download from a publication
	pubClient, err := mdium.NewClient("medium.com/netflix-techblog", "netflix_articles")
	if err != nil {
		log.Fatalf("Error creating client for publication: %v", err)
	}

	// Download all articles from Medium (limited to 10 articles by Medium RSS)
	err = pubClient.DownloadArticles(0)
	if err != nil {
		log.Printf("Error downloading publication articles: %v", err)
	}
}
