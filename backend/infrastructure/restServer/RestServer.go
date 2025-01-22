package restServer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	loggerInterface "infrastructure/logger/interface"
	restServerInterface "infrastructure/restServer/interface"
	"infrastructure/restServer/middleware"
	restModel "infrastructure/restServer/model"
	jwtServiceInterface "infrastructure/security/jwtService/interface"
	"net/http"
	"strconv"
)

type GinServer struct {
	server        *gin.Engine
	serverConfig  *restModel.RestServerConfig
	logger        loggerInterface.Logger
	jwtMiddleware gin.HandlerFunc
}

// NewGinServer TODO переписать на Builder
func NewGinServer(logger loggerInterface.Logger, jwtService jwtServiceInterface.JWTService, serverConfig *restModel.RestServerConfig) restServerInterface.Server {
	server := gin.New()

	server.Use(
		gin.Recovery(),
		middleware.NewRequestMiddleware(logger).Handler(),
		middleware.NewErrorMiddleware(logger).Handler())

	// TODO переписать
	// Работа с фронтом
	server.Static("/assets", "D:\\Development\\Monetization\\task-tracker\\frontend\\assets") // Шрифты и изображения
	server.Static("/css", "D:\\Development\\Monetization\\task-tracker\\frontend\\css")       // Статические CSS файлы
	server.Static("/js", "D:\\Development\\Monetization\\task-tracker\\frontend\\js")         // Статические JS файлы

	// TODO переписать
	// Регистрируем маршрут для главной страницы
	server.GET("/", func(c *gin.Context) {
		// Указываем путь к главному файлу start.html
		c.File("D:\\Development\\Monetization\\task-tracker\\frontend\\start.html")
	})

	// TODO переписать
	// Кастомный обработчик при отсутствии ресурса по переданному маршруту
	server.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Упс, ничего не найдено",
			"message": "Когда-нибудь тут обязательно что-то появится :)",
		})
	})

	return &GinServer{
		serverConfig:  serverConfig,
		server:        server,
		logger:        logger,
		jwtMiddleware: middleware.NewJWTMiddleware(logger, jwtService).Handler(),
	}
}

func (s *GinServer) RegisterPublicRoute(method, path string, handler http.HandlerFunc) {

	// Обёртка http.HandlerFunc в gin.HandlerFunc
	ginHandlerFunc := func(c *gin.Context) {
		handler(c.Writer, c.Request)
	}

	s.registerGinRouts(method, path, ginHandlerFunc)
}

func (s *GinServer) RegisterPrivateRoute(method, path string, handler http.HandlerFunc) {

	// Обёртка http.HandlerFunc в gin.HandlerFunc
	ginHandlerFunc := func(c *gin.Context) {
		handler(c.Writer, c.Request)
	}

	s.registerGinRouts(method, path, s.jwtMiddleware, ginHandlerFunc)
}

// Регистрирует маршрут на основе переданного метода и конвертирует в методы Gin
func (s *GinServer) registerGinRouts(method, path string, handlers ...gin.HandlerFunc) {
	switch method {
	case http.MethodGet:
		s.server.GET(path, handlers...)
	case http.MethodPost:
		s.server.POST(path, handlers...)
	case http.MethodPut:
		s.server.PUT(path, handlers...)
	case http.MethodDelete:
		s.server.DELETE(path, handlers...)
	default:
		panic("Unsupported method")
	}
}

func (s *GinServer) Start() error {
	return s.startWorker()
}

func (s *GinServer) startWorker() error {
	address := fmt.Sprintf(":%s", strconv.Itoa(s.serverConfig.Port()))
	return s.server.Run(address)
}
