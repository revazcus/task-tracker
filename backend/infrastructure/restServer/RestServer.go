package restServer

import (
	"github.com/gin-gonic/gin"
	loggerInterface "infrastructure/logger/interface"
	restServerInterface "infrastructure/restServer/interface"
	"infrastructure/restServer/middleware"
	jwtServiceInterface "infrastructure/security/jwtService/interface"
	"net/http"
)

type GinServer struct {
	server        *gin.Engine
	logger        loggerInterface.Logger
	jwtMiddleware gin.HandlerFunc
}

func NewGinServer(logger loggerInterface.Logger, jwtService jwtServiceInterface.JWTService) restServerInterface.Server {
	server := gin.New()

	server.Use(
		gin.Recovery(),
		middleware.NewRequestMiddleware(logger).Handler(),
		middleware.NewErrorMiddleware(logger).Handler())

	// TODO переписать
	// Работа с фронтом
	server.Static("/assets", "C:\\Users\\Rezo\\IdeaProjects\\Monetization\\Task-Tracker\\frontend\\assets") // Шрифты и изображения
	server.Static("/css", "C:\\Users\\Rezo\\IdeaProjects\\Monetization\\Task-Tracker\\frontend\\css")       // Статические CSS файлы
	server.Static("/js", "C:\\Users\\Rezo\\IdeaProjects\\Monetization\\Task-Tracker\\frontend\\js")         // Статические JS файлы

	// Регистрируем маршрут для главной страницы
	server.GET("/", func(c *gin.Context) {
		// Указываем путь к главному файлу start.html
		c.File("C:\\Users\\Rezo\\IdeaProjects\\Monetization\\Task-Tracker\\frontend\\start.html")
	})

	// Кастомный обработчик при отсутствии ресурса по переданному маршруту
	server.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Упс, ничего не найдено",
			"message": "Когда-нибудь тут обязательно что-то появится :)",
		})
	})

	return &GinServer{
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

func (s *GinServer) Start(address string) error {
	return s.server.Run(address)
}
