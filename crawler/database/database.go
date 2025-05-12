package database

import (
	// "fmt"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"


	"crawler/models"
)

var DB *gorm.DB


func InitDB() {
	dsn := getEnv("PG_URL", "")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	// Auto migrate the model (create table if not exists)
	err = DB.AutoMigrate(&models.Content{})
	if err != nil {
		log.Fatalf("❌ AutoMigrate failed: %v", err)
	}

	log.Println("✅ Connected and migrated PostgreSQL database.")
}

func getEnv(key, fallback string) string {
	godotenv.Load(".env")
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
