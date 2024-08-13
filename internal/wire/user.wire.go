//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/hoangxt/go-ecommerce-ba-api/internal/controller"
	"github.com/hoangxt/go-ecommerce-ba-api/internal/repo"
	"github.com/hoangxt/go-ecommerce-ba-api/internal/service"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
