package restServer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	loggerInterface "infrastructure/logger/interface"
	"infrastructure/restServer/middleware"
	restModel "infrastructure/restServer/model"
	jwtServiceInterface "infrastructure/security/jwtService/interface"
	"net/http"
	"strconv"
)

type RestServer struct {
	server       *gin.Engine
	serverConfig *restModel.RestServerConfig
	logger       loggerInterface.Logger
	jwtService   jwtServiceInterface.JWTService
}

func (s *RestServer) RegisterPublicRoute(method, path string, handler http.HandlerFunc) {

	// Обёртка http.HandlerFunc в gin.HandlerFunc
	ginHandlerFunc := func(c *gin.Context) {
		handler(c.Writer, c.Request)
	}

	s.registerGinRouts(method, path, ginHandlerFunc)
}

func (s *RestServer) RegisterPrivateRoute(method, path string, handler http.HandlerFunc) {

	// Обёртка http.HandlerFunc в gin.HandlerFunc
	ginHandlerFunc := func(c *gin.Context) {
		handler(c.Writer, c.Request)
	}

	// TODO вынести
	jwtMiddleware := middleware.NewJWTMiddleware(s.logger, s.jwtService).Handler()

	s.registerGinRouts(method, path, jwtMiddleware, ginHandlerFunc)
}

func (s *RestServer) Start() error {
	return s.startWorker()
}

func (s *RestServer) createHttpServer() {
	server := gin.New()

	// TODO переписать
	// Кастомный обработчик при отсутствии ресурса по переданному маршруту
	server.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Упс, ничего не найдено",
			"message": "Когда-нибудь тут обязательно что-то появится :)",
		})
	})

	s.server = server

	s.initMiddlewares()
	s.initFront()
}

func (s *RestServer) initMiddlewares() {
	s.server.Use(
		gin.Recovery(),
		middleware.NewRequestMiddleware(s.logger).Handler(),
		middleware.NewErrorMiddleware(s.logger).Handler())
}

// TODO переписать
func (s *RestServer) initFront() {
	// Шрифты и изображения
	s.server.Static("/assets", "D:\\Development\\Monetization\\task-tracker\\frontend\\assets")
	// Статические CSS файлы
	s.server.Static("/css", "D:\\Development\\Monetization\\task-tracker\\frontend\\css")
	// Статические JS файлы
	s.server.Static("/js", "D:\\Development\\Monetization\\task-tracker\\frontend\\js")

	// Регистрируем маршрут для главной страницы
	s.server.GET("/", func(c *gin.Context) {
		// Указываем путь к главному файлу start.html
		c.File("D:\\Development\\Monetization\\task-tracker\\frontend\\start.html")
	})
}

// Регистрирует маршрут на основе переданного метода и конвертирует в методы Gin
func (s *RestServer) registerGinRouts(method, path string, handlers ...gin.HandlerFunc) {
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

func (s *RestServer) startWorker() error {
	address := fmt.Sprintf(":%s", strconv.Itoa(s.serverConfig.Port()))
	return s.server.Run(address)
}
