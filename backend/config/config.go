package config;

import (
	"os"
	"log"
	"github.com/joho/godotenv"

)

var PORT string = Getenv("PORT")


func Getenv(val string) string {
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatalf("Error loading .env file: %s", err)
	}
	greeting := os.Getenv(val)
	return greeting
}