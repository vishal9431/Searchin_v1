package search

import (
	"backend/config"
	"fmt"
)

type Content struct {
	ID      int
	Url string
	Title   string
	Content string
}

func Search(query string) ([]Content, error) {
	var results []Content

	// Full-text search misoli (ILIKE - case-insensitive LIKE)
	sql := `
		SELECT id, title, content, url
		FROM contents
		WHERE url ILIKE ? OR title ILIKE ? OR content ILIKE ?
	`

	searchTerm := fmt.Sprintf("%%%s%%", query)

	if err := config.DB.Raw(sql, searchTerm, searchTerm, searchTerm).Scan(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}
