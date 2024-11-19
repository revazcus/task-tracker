package middleware

import (
	"github.com/gin-gonic/gin"
	loggerInterface "task-tracker/infrastructure/logger/interface"
)

type RequestMiddleware struct {
	logger loggerInterface.Logger
}

func NewRequestMiddleware(logger loggerInterface.Logger) *RequestMiddleware {
	return &RequestMiddleware{
		logger,
	}
}

func (r *RequestMiddleware) Handler() gin.HandlerFunc {
	return func(context *gin.Context) {
		r.logger.LogInfo(context.Request.Method, context.Request.URL.Path)
		context.Next()
	}
}
