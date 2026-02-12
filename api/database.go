// Based on https://neon.com/guides/golang-gorm-postgres
package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/joho/godotenv"
)

var DB *gorm.DB

func InitDB() {
	//Load Environment
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Connection string for Neon Postgres
	dsn := os.Getenv("NEON_DB_URL")

	// Connect to the database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Get the underlying SQL DB object
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get DB object: %v", err)
	}

	// Verify connection
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	fmt.Println("Successfully connected to Neon Postgres database!")
}
