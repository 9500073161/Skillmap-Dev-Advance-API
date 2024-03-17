package database

import (
	"github.com/9500073161/Skillmap-Dev-Advance-API/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Initialize(){

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")

	}

	// Migrate the schema
	db.AutoMigrate(&models.User{})

	

}
