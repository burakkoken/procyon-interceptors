package controller

import (
	context "github.com/procyon-projects/procyon-context"
	web "github.com/procyon-projects/procyon-web"
)

type UserController interface {
	GetUsers(context *web.WebRequestContext)
}

type ImpUserController struct {
	logger context.Logger
}

func NewUserController(logger context.Logger) ImpUserController {
	return ImpUserController{
		logger,
	}
}

func (controller ImpUserController) RegisterHandlers(registry web.HandlerRegistry) {
	registry.Register(
		web.Get(controller.GetUsers, web.Path("/api/users")),
	)
}

func (controller ImpUserController) GetUsers(context *web.WebRequestContext) {
	controller.logger.Debug(context, "GetUsers is invoked")
}
