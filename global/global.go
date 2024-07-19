package global

import (
	"github.com/hoangxt/go-ecommerce-ba-api/pkg/logger"
	"github.com/hoangxt/go-ecommerce-ba-api/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)

/*
	Config
	Redis
	MySQL
	...
*/
