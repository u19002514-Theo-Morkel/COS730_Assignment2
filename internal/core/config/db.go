package config

import (
	"Assignment2/internal/core/Models"
	"fmt"
	"log/slog"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s database=%s port=%s sslmode=disable TimeZone=Africa/Johannesburg", host, user, password, dbname, port)
	slog.Info("Connecting to database at " + host + ":" + port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error(err.Error())
		return nil
	}

	slog.Info("Connected to database and migrating tables")
	err = db.AutoMigrate(&Models.Page{})
	if err != nil {
		slog.Error(err.Error())
	}
	err = db.AutoMigrate(&Models.Comment{})
	if err != nil {
		slog.Error(err.Error())
	}
	err = db.AutoMigrate(&Models.Moderation{})
	if err != nil {
		slog.Error(err.Error())
	}

	return db
}
