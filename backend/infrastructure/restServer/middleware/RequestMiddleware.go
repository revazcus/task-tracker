package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	loggerInterface "infrastructure/logger/interface"
	"io"
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

		r.logger.Info(context, fmt.Sprintf("Request: Path=%s, Method=%s", context.Request.URL.Path, context.Request.Method))

		// Восстанавливаем тело запроса для последующих обработчиков (иначе вернёт EOF (конец файла))
		context.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		context.Next()
	}
}
