package routers

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	c "github.com/hoangxt/go-ecommerce-ba-api/internal/controller"
// 	"github.com/hoangxt/go-ecommerce-ba-api/internal/middlewares"
// )

// func AA() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		fmt.Println("Before AA")
// 		c.Next()
// 		fmt.Println("After AA")

// 	}
// }

// func BB() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		fmt.Println("Before BB")
// 		c.Next()
// 		fmt.Println("After BB")

// 	}
// }

// func CC(c *gin.Context) {
// 	fmt.Println("Before CC")
// 	c.Next()
// 	fmt.Println("After CC")
// }

// func NewRouter() *gin.Engine {
// 	r := gin.Default()
// 	// use middleware
// 	r.Use(middlewares.AuthMiddleware(), BB(), CC)

// 	v1 := r.Group("/v1/2024")
// 	{
// 		v1.GET("/ping", c.NewPongController().Pong)          // /v1/2024/ping curl http://localhost:8002/user/123
// 		v1.GET("/user/1", c.NewUserController().GetUserByID) // curl http://localhost:8002/v1/2024/user/1
// 		// v1.PUT("/ping", Pong)
// 		// v1.POST("/ping", Pong)
// 		// v1.DELETE("/ping", Pong)
// 		// v1.PATCH("/ping", Pong)
// 		// v1.OPTIONS("/ping", Pong)
// 		// v1.HEAD("/ping", Pong)
// 	}

// 	// v2 := r.Group("/v2/2024")
// 	// {
// 	// 	v2.GET("/ping/", Pong) // /v2/2024/ping curl http://localhost:8002/v2/2024/ping
// 	// 	v2.PUT("/ping", Pong)
// 	// 	v2.POST("/ping", Pong)
// 	// 	v2.DELETE("/ping", Pong)
// 	// 	v2.PATCH("/ping", Pong)
// 	// 	v2.OPTIONS("/ping", Pong)
// 	// 	v2.HEAD("/ping", Pong)
// 	// }
// 	return r
// }

// api or controller
// func Pong(c *gin.Context) {
// 	name := c.DefaultQuery("name", "hoang")
// 	uid := c.Query("uid")
// 	c.JSON(http.StatusOK, gin.H{ // map string
// 		"message": "pong ... ping" + name,
// 		"uid":     uid,
// 		"users":   []string{"hoang", "thanh", "tuan"},
// 	})
// }
