package init_services

import (
	"net/http"
	userRest "task-tracker/adapters/controllers/rest"
	restServerInterface "task-tracker/infrastructure/restServer/interface"
)

type Router interface {
	RegisterRoutes()
}

type UserRouter struct {
	server     restServerInterface.Server
	controller *userRest.UserController
}

func NewUserRouter(server restServerInterface.Server, controller *userRest.UserController) *UserRouter {
	return &UserRouter{
		server:     server,
		controller: controller,
	}
}

func (r *UserRouter) RegisterRoutes() {
	r.server.RegisterPublicRoute(http.MethodPost, "v1/user/create", r.controller.CreateUser)
	r.server.RegisterPublicRoute(http.MethodGet, "/", r.controller.GetUser)
}
