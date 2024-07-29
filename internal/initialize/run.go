package initialize

import (
	"fmt"

	"github.com/hoangxt/go-ecommerce-ba-api/global"
	"go.uber.org/zap"
)

func Run() {
	// load config
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading config mysql", m.Username, m.Password)
	// Initialize logger
	InitLogger()
	global.Logger.Info("Logger init success", zap.String("ok", "success"))
	// Initialize mysql
	// InitMysql()
	// Initialize redis
	InitRedis()
	// Initialize router
	r := InitRouter()

	r.Run(":8002")
}
