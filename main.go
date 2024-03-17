package main

import (
	"fmt"

	"github.com/9500073161/skill-map-app/apis"
	"github.com/9500073161/skill-map-app/handlers"
	"github.com/9500073161/skill-map-app/managers"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Developing Skill-MAP")

	router := gin.Default()

	userManager := managers.NewUserManager()
	userHandler := handlers.NewUserHandlerFrom(userManager)
	userHandler.RegisterUserApis()

	router.Run()

}
