package database

import (
	"github.com/9500073161/Skillmap-Dev-Advance-API/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize(){

	var err error
	DB, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")

	}

	// Migrate the schema
	DB.AutoMigrate(&models.User{})

	

}
