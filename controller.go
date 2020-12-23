package main

import (
	context "github.com/procyon-projects/procyon-context"
	web "github.com/procyon-projects/procyon-web"
)

type User struct {
	Name  string
	Email string
}

type UserController interface {
	GetUser(context *web.WebRequestContext)
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
		web.Get(controller.GetUser, web.Path("/api/user")),
	)
}

func (controller ImpUserController) GetUser(context *web.WebRequestContext) {
	controller.logger.Info(context, "GetUser is invoked")
	context.SetModel(User{
		Name:  "Burak",
		Email: "burakkokenn@gmail.com",
	}).SetResponseContentType(web.MediaTypeApplicationJson)
}
