package modules

import "projects/LDmitryLD/repository/app/internal/modules/user/controller"

type Controllers struct {
	User controller.Userer
}

func NewControllers(services *Services) *Controllers {
	userController := controller.NewUser(services.User)

	return &Controllers{
		User: userController,
	}
}
