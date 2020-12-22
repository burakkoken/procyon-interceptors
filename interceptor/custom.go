package interceptor

import (
	context "github.com/procyon-projects/procyon-context"
	web "github.com/procyon-projects/procyon-web"
)

type CustomInterceptor struct {
	logger context.Logger
}

func NewCustomInterceptor(logger context.Logger) CustomInterceptor {
	return CustomInterceptor{
		logger,
	}
}

func (customInterceptor CustomInterceptor) HandleBefore(requestContext *web.WebRequestContext) {
	customInterceptor.logger.Debug(requestContext, "HandleBefore is invoked")
}

func (customInterceptor CustomInterceptor) HandleAfter(requestContext *web.WebRequestContext) {
	customInterceptor.logger.Debug(requestContext, "HandleAfter is invoked")
}

func (customInterceptor CustomInterceptor) AfterCompletion(requestContext *web.WebRequestContext) {
	customInterceptor.logger.Debug(requestContext, "AfterCompletion is invoked")
}
