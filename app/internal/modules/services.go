package modules

import (
	"projects/LDmitryLD/repository/app/internal/modules/user/service"
	"projects/LDmitryLD/repository/app/internal/storages"
)

type Services struct {
	User service.Userer
}

func NewServices(storages *storages.Storages) *Services {
	userService := service.NewUserService(storages.User)
	return &Services{
		User: userService,
	}
}
