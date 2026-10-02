package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"internity/Backend/models"
)

func ConnectDB() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("load .env: %w", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	if host == "" || user == "" || database == "" {
		return nil, fmt.Errorf("DB_HOST, DB_USER, and DB_NAME must be set")
	}
	if port == "" {
		port = "3306"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("connect to MySQL: %w", err)
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.Attendance{},
		&models.News{},
		&models.EmailOTP{},
	); err != nil {
		return nil, fmt.Errorf("auto-migrate database: %w", err)
	}

	return db, nil
}
