package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (p *PongController) Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "hoang")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{ // map string
		"message": "pong ... ping" + name,
		"uid":     uid,
		"users":   []string{"hoang", "thanh", "tuan"},
	})
}
