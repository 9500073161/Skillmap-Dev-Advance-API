package handlers

import (
	"net/http"

	"github.com/9500073161/Skillmap-Dev-Advance-API/database"
	"github.com/9500073161/Skillmap-Dev-Advance-API/managers"
	"github.com/9500073161/Skillmap-Dev-Advance-API/models"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	groupName   string
	userManager *managers.UserManager
}

func NewUserHandlerFrom(userManager *managers.UserManager) *UserHandler {
	return &UserHandler{
		"api/users",
		userManager,
	}

}

func (UserHandler *UserHandler) RegisterUserApis(r *gin.Engine) {
	userGroup := r.Group(UserHandler.groupName)
	userGroup.GET("", UserHandler.Create)

}
func (UserHandler *UserHandler) Create(ctx *gin.Context) {

	database.DB.Create(&models.User{FullName: "Prashob", Email: "admin@tcs.com"})
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "api version2",
	})

}
