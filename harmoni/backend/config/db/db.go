package db

import (
	"fmt"
	"harmoni/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() error {
	dsn := "host=localhost user=sena password=sena123 dbname=harmoni port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	err = db.AutoMigrate(&model.LoginRequest{}, &model.User{})
	if err != nil {
		log.Fatal("Failed to migrate schema:", err)
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}
	if err := sqlDB.Close(); err != nil {
		log.Fatal("Failed to close database connection:", err)
	}
}
