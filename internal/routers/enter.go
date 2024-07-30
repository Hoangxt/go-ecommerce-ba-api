package routers

import (
	"github.com/hoangxt/go-ecommerce-ba-api/internal/routers/manager"
	"github.com/hoangxt/go-ecommerce-ba-api/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)
