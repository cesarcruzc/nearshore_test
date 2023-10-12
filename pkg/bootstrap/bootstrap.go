package bootstrap

import (
	"fmt"
	"github.com/cesarcruzc/nearshore_test/internal/core/domain"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func InitLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}

func DBConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if os.Getenv("DB_DEBUG") == "true" {
		db = db.Debug()
	}

	if os.Getenv("DB_AUTO_MIGRATE") == "true" {
		if err := db.AutoMigrate(&domain.Device{}); err != nil {
			return nil, err
		}

		if err := db.AutoMigrate(&domain.Firmware{}); err != nil {
			return nil, err
		}
	}

	return db, nil
}

func InitLoadEnv() error {
	return godotenv.Load(".env")
}
