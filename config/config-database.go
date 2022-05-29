package config

import (
	"fmt"
	"os"

	// "log"

	// "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// setup config database
func SetupDatabaseConnection() *gorm.DB {
	// err := godotenv.Load()
	// if err != nil {
	// 	panic("failed to load file")
	// }

	// db_user := os.Getenv("MYSQL_USER")
	// db_pass := os.Getenv("MYSQL_PASSWORD")
	// db_host := os.Getenv("MYSQL_HOST")
	// db_name := os.Getenv("MYSQL_DATABASE")

	// db_user := "root"
	// db_pass := ""
	// db_host := "localhost"
	// db_name := "db_golang"

	db_url := os.Getenv("DATABASE_URL")

	// dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_pass, db_host, db_name)
	dsn := fmt.Sprintf("%ssslmode=disable", db_url)
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
