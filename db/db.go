package db

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

func GetDB() *gorm.DB {
	dsn := viper.GetString("DATABASE_URL")
	var db *gorm.DB
	var err error
	for i := 0; i < 10; i++ { // Пробуем 10 раз
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			return db
		}
		log.Printf("Failed to connect to database: %v, retrying in 5s...", err)
		time.Sleep(5 * time.Second)
	}
	log.Fatalf("Could not connect to database after retries: %v", err)
	return nil
}
