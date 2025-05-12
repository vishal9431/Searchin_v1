package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"crawler/database"
	"crawler/models"
)

// check is a simple error handler
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Initialize the database
	database.InitDB()

	// Prompt user for input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter query: ")

	searchTerm, err := reader.ReadString('\n')
	check(err)

	// Clean up input
	searchTerm = strings.TrimSpace(searchTerm)

	// Search contents from DB
	results, err := searchContents(searchTerm)
	check(err)

	// Display results
	if len(results) == 0 {
		fmt.Println("No results found.")
		return
	}

	for _, content := range results {
		fmt.Printf("URL: %s\nTitle: %s\nContent: %s\n\n", content.URL, content.Title, content.Content)
	}
}

// searchContents looks for the search term in url, title, or content
func searchContents(term string) ([]models.Content, error) {
	var contents []models.Content

	err := database.DB.Where(
		"url ILIKE ? OR title ILIKE ? OR content ILIKE ?",
		"%"+term+"%", "%"+term+"%", "%"+term+"%",
	).Find(&contents).Error

	return contents, err
}
