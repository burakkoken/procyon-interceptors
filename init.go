package main

import (
	"github.com/burakkoken/procyon-interceptors/controller"
	"github.com/burakkoken/procyon-interceptors/interceptor"
	core "github.com/procyon-projects/procyon-core"
)

func init() {
	core.Register(controller.NewUserController)
	core.Register(interceptor.NewCustomInterceptor)
}
