package models


type Content struct {
	ID      uint   `gorm:"primaryKey"`
	URL     string `gorm:"type:text;not null"`
	Title   string `gorm:"type:text"`
	Content string `gorm:"type:text"`
}