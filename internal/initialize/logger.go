package initialize

import (
	"github.com/hoangxt/go-ecommerce-ba-api/global"
	"github.com/hoangxt/go-ecommerce-ba-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
