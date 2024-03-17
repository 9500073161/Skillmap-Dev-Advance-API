package main

import (
	"fmt"

	"github.com/9500073161/Skillmap-Dev-Advance-API/handlers"
	"github.com/9500073161/Skillmap-Dev-Advance-API/managers"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Developing Skill-MAP")

	router := gin.Default()

	userManager := managers.NewUserManager()
	userHandler := handlers.NewUserHandlerFrom(userManager)
	userHandler.RegisterUserApis(router)

	router.Run()

}
