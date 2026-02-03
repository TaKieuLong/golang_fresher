//go:build wireinject

package wire

import (
	"github.com/TaKieuLong/golang_fresher/internal/controller"
	"github.com/TaKieuLong/golang_fresher/internal/repo"
	"github.com/TaKieuLong/golang_fresher/internal/service"
	"github.com/google/wire"
)

func InitUserRouterHanlder() (*controller.UserController, error){
	wire.Build(
		repo.NewUserRepository,
		repo.NewUserAuthRepository,
		service.NewUserService,
		controller.NewUserController,
		
	)
	return new(controller.UserController), nil
}	