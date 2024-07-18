package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangxt/go-ecommerce-ba-api/internal/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// uc user controller
// us user service

// controller -> service -> repo -> models -> dbs
func (uc *UserController) GetUserByID(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{ // map string
		"message": uc.userService.GetInfoUser(),
		"users":   []string{"hoang", "thanh", "tuan"},
	})
}
