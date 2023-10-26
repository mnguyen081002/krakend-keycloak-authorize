package controller

import (
	"erp/domain"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserController struct {
	userService domain.UserService
	logger      *zap.Logger
}

func NewUserController(userService domain.UserService, logger *zap.Logger) *UserController {
	controller := &UserController{
		userService: userService,
		logger:      logger,
	}
	return controller
}

func (controller *UserController) Profile(context *gin.Context) {
	Response(context, 200, "success", "Nguyen Van A")
}
