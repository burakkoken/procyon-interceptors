package main

import (
	core "github.com/procyon-projects/procyon-core"
)

func init() {
	core.Register(NewUserController)
	core.Register(NewCustomInterceptor)
}
