package global

import (
	"github.com/hoangxt/go-ecommerce-ba-api/pkg/logger"
	"github.com/hoangxt/go-ecommerce-ba-api/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
)

/*
	Config
	Redis
	MySQL
	...
*/
