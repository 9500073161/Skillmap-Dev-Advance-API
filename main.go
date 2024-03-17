package main

import (
	"fmt"

	"github.com/9500073161/Skillmap-Dev-Advance-API/handlers"
	"github.com/9500073161/Skillmap-Dev-Advance-API/managers"
	"github.com/9500073161/Skillmap-Dev-Advance-API/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	
)

func main() {
	fmt.Println("Developing Skill-MAP")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")

	}

	// Migrate the schema
	db.AutoMigrate(&models.User{})

	router := gin.Default()

	userManager := managers.NewUserManager()
	userHandler := handlers.NewUserHandlerFrom(userManager)
	userHandler.RegisterUserApis(router)

	router.Run()

}
