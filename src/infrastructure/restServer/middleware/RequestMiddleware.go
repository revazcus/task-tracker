package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
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

		// Читаем и сохраняем тело запроса для логирования
		bodyBytes, _ := context.GetRawData()

		r.logger.LogInfo(context.Request.Method, context.Request.URL.Path, string(bodyBytes))

		// Восстанавливаем тело запроса для последующих обработчиков (иначе вернёт EOF (конец файла))
		context.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		context.Next()
	}
}
