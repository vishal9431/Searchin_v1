package utils

import (
	"gorm.io/gorm"

	"crawler/database"
	"crawler/models"
)

// IsContentExists checks if the given URL already exists in the database.
func IsContentExists(url string) (bool, error) {
	var existingURL string

	err := database.DB.
		Model(&models.Content{}).
		Where("url = ?", url).
		Pluck("url", &existingURL).
		Error

	// Return error unless it's the "record not found" case
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	return existingURL != "", nil
}

// CreateContent inserts a new content record into the database.
func CreateContent(url, title, content string) (bool, error) {
	newContent := models.Content{
		URL:     url,
		Title:   title,
		Content: content,
	}

	if err := database.DB.Create(&newContent).Error; err != nil {
		return false, err
	}

	return true, nil
}

// GetAllContents retrieves all content records from the database.
func GetAllContents() ([]models.Content, error) {
	var contents []models.Content

	if err := database.DB.Find(&contents).Error; err != nil {
		return nil, err
	}

	return contents, nil
}

// AddContentIfNotExists adds new content only if the URL does not already exist.
func AddContentIfNotExists(url, title, content string) (bool, error) {
	exists, err := IsContentExists(url)
	if err != nil {
		return false, err
	}

	if exists {
		return false, nil // Already exists, no need to insert
	}

	return CreateContent(url, title, content)
}
