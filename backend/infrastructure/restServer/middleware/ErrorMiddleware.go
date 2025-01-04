package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	loggerInterface "github.com/revazcus/task-tracker/backend/infrastructure/logger/interface"
	"net/http"
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

			r.logger.Error(context, fmt.Errorf("error: %v, Path: %s, Method: %s", err, context.Request.URL.Path, context.Request.Method))

			context.JSON(code, gin.H{
				"error": err.Error(),
			})

		}
	}
}
