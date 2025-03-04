// database/database.go
package database

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"todo-app/models"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// Auto migrate model
	DB.AutoMigrate(&models.Task{})
}