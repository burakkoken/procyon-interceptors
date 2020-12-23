package main

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
	customInterceptor.logger.Info(requestContext, "HandleBefore is invoked : "+requestContext.GetPath())
}

func (customInterceptor CustomInterceptor) HandleAfter(requestContext *web.WebRequestContext) {
	customInterceptor.logger.Info(requestContext, "HandleAfter is invoked")
}

func (customInterceptor CustomInterceptor) AfterCompletion(requestContext *web.WebRequestContext) {
	customInterceptor.logger.Info(requestContext, "AfterCompletion is invoked")
	customInterceptor.logger.Info(requestContext, "Response Body :"+string(requestContext.GetResponseBody()))

	if requestContext.IsSuccess() {

		if requestContext.IsCanceled() {
			customInterceptor.logger.Warning(requestContext, "Request is cancelled")
		} else if requestContext.IsCompleted() {
			customInterceptor.logger.Info(requestContext, "Request is completed successfully")
		}

	} else {
		customInterceptor.logger.Warning(requestContext, "Request is not completed successfully")

		if requestContext.GetInternalError() != nil {
			customInterceptor.logger.Error(requestContext, "Internal Error : "+requestContext.GetInternalError().Error())
		} else if requestContext.GetHTTPError() != nil {
			customInterceptor.logger.Error(requestContext, "Http Error : "+requestContext.GetHTTPError().Message.(string))
		}
	}
}
