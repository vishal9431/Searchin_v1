package utils

import (
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

// GetData crawls the given URL and returns its title, description, and discovered links.
func GetData(targetURL string) (map[string]any, error) {
	collector := colly.NewCollector(
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 Chrome/119.0.0.0 Safari/537.36"),
	)

	// Rate limiting configuration
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 1,
		Delay:       5 * time.Second,
		RandomDelay: 2 * time.Second,
	})

	// Set custom headers
	collector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Referer", "https://www.google.com")
	})

	var (
		title       string
		description string
		links       []string
	)

	// Extract <title> tag
	collector.OnHTML("title", func(e *colly.HTMLElement) {
		title = strings.TrimSpace(e.Text)
	})

	// Extract meta description
	collector.OnHTML("meta[name='description']", func(e *colly.HTMLElement) {
		description = strings.TrimSpace(e.Attr("content"))
	})

	// Extract all useful <a> href links
	collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		href := e.Attr("href")

		// Skip irrelevant or unsafe links
		if strings.HasPrefix(href, "#") ||
			strings.HasPrefix(href, "mailto:") ||
			strings.HasPrefix(href, "javascript:") ||
			strings.Contains(href, "twitter.com/intent") {
			return
		}

		absoluteURL := e.Request.AbsoluteURL(href)
		links = append(links, absoluteURL)
	})

	if err := collector.Visit(targetURL); err != nil {
		return nil, err
	}

	collector.Wait() // Wait for all asynchronous tasks

	return map[string]any{
		"url":         targetURL,
		"title":       title,
		"description": description,
		"links":       links,
	}, nil
}


// LoadUrls loads all stored URLs from the database into the provided slice pointer.
func LoadUrls(storage *[]string) error {
	records, err := GetAllContents()
	if err != nil {
		return err
	}

	for _, record := range records {
		*storage = append(*storage, record.URL)
	}
	return nil
}
