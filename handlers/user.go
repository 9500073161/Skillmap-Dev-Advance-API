package handlers

import (
	"net/http"

	"github.com/9500073161/skill-map-app/managers"
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
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "api version2",
	})

}
