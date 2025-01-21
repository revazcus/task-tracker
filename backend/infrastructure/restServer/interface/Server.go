package restServerInterface

import (
	"net/http"
)

type Server interface {

	// RegisterPublicRoute регистрирует роуты через gin HandlerFunc
	RegisterPublicRoute(method, path string, handler http.HandlerFunc)

	// RegisterPrivateRoute регистрирует приватные роуты с хендлером для проверки JWT токена
	RegisterPrivateRoute(method, path string, handler http.HandlerFunc)

	// Start производит запуск сервера на определённом адресе
	Start() error
}
