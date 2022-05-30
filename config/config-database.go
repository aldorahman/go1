package config

import (
	"fmt"
	"os"

	// "log"

	// "github.com/joho/godotenv"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// setup config database
func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load file")
	}

	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_DATABASE")
	db_port := os.Getenv("DB_PORT")
	db_sslmode := os.Getenv("DB_SSLMODE")

	// db_user := "root"
	// db_pass := ""
	// db_host := "localhost"
	// db_name := "db_golang"

	// db_url := os.Getenv("DATABASE_URL")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", db_host, db_user, db_pass, db_name, db_port, db_sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("db error connection")
	}

	return db
}

// close connection between yur app and database
func CloseConnectionDatabase(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		fmt.Println("Failed to close connection from database")
	}

	dbSQL.Close()
}
