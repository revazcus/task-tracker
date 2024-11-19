package restServer

import (
	"github.com/gin-gonic/gin"
	"net/http"
	loggerInterface "task-tracker/infrastructure/logger/interface"
	restServerInterface "task-tracker/infrastructure/restServer/interface"
	"task-tracker/infrastructure/restServer/middleware"
)

type GinServer struct {
	server *gin.Engine
	logger loggerInterface.Logger
}

func NewGinServer(logger loggerInterface.Logger) restServerInterface.Server {
	server := gin.New()

	server.Use(
		gin.Recovery(),
		middleware.NewRequestMiddleware(logger).Handler(),
		middleware.NewErrorMiddleware(logger).Handler())

	return &GinServer{
		server: server,
		logger: logger,
	}
}

func (s *GinServer) RegisterPublicRoute(method, path string, handler http.HandlerFunc) {

	// Обёртка http.HandlerFunc в gin.HandlerFunc
	ginHandlerFunc := func(c *gin.Context) {
		handler(c.Writer, c.Request)
	}

	s.registerGinRouts(method, path, ginHandlerFunc)
}

// Регистрирует маршрут на основе переданного метода и конвертирует в методы Gin
func (s *GinServer) registerGinRouts(method, path string, handler gin.HandlerFunc) {
	switch method {
	case http.MethodGet:
		s.server.GET(path, handler)
	case http.MethodPost:
		s.server.POST(path, handler)
	case http.MethodPut:
		s.server.PUT(path, handler)
	case http.MethodDelete:
		s.server.DELETE(path, handler)
	default:
		panic("Unsupported method")
	}
}

func (s *GinServer) Start(address string) error {
	return s.server.Run(address)
}
