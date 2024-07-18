package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hoangxt/go-ecommerce-ba-api/internal/service"
	"github.com/hoangxt/go-ecommerce-ba-api/pkg/response"
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
	if err != nil {
		response.ErrorResponse(c, 40001, "error")
	}
	return response.SuccessResponse(c, 20001, []string{"Hoang", "24", "Vietnam"})
	// c.JSON(http.StatusOK, response.ResponseData{
	// 	Code:    20001,
	// 	Message: "success",
	// 	Data:    []string{"Hoang", "24", "Vietnam"},
	// })
	// response.SuccessResponse(c, 20001, []string{"Hoang", "24", "Vietnam"})
	// response.ErrorResponse(c, 40001, "error")
}
