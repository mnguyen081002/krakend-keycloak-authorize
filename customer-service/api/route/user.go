package route

import (
	controller "erp/api/controllers"
	"erp/lib"
)

type UserRoutes struct {
	handler *lib.Handler
}

func NewUserRoutes(handler *lib.Handler, userController *controller.UserController) *UserRoutes {
	r := handler.Group("/customers")

	r.GET("/profile", userController.Profile)
	return &UserRoutes{
		handler: handler,
	}
}
