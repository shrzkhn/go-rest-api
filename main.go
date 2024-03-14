package main

import (
	"fmt"
	"log"
	"os"
	"shrzkhn/go-rest-api/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()

	connectToDb()

	router.Run()
}

func connectToDb() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}

	// Define a PostgreSQL connection
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Connect to PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	// Create table based on our model
	err = db.AutoMigrate(&models.City{})
	if err != nil {
		panic("failed to perform migrations: " + err.Error())
	}
}
