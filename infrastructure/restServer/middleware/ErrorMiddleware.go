package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	loggerInterface "task-tracker/infrastructure/logger/interface"
)

type ErrorMiddleware struct {
	logger loggerInterface.Logger
}

func NewErrorMiddleware(logger loggerInterface.Logger) *ErrorMiddleware {
	return &ErrorMiddleware{
		logger,
	}
}

func (r *ErrorMiddleware) Handler() gin.HandlerFunc {
	return func(context *gin.Context) {

		context.Next()

		if len(context.Errors) > 0 {
			err := context.Errors.Last()
			code := http.StatusInternalServerError

			if err.Type == gin.ErrorTypePublic {
				code = http.StatusBadRequest
			}

			r.logger.LogError(err, context.Request.URL.Path, context.Request.Method)

			context.JSON(code, gin.H{
				"error": err.Error(),
			})

		}
	}
}
