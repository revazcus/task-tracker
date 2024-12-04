package router

import (
	"net/http"
	userRest "task-tracker/adapters/controllers/rest/user"
	restServerInterface "task-tracker/infrastructure/restServer/interface"
)

type UserRouter struct {
	controller *userRest.UserController
}

func NewUserRouter(controller *userRest.UserController) *UserRouter {
	return &UserRouter{
		controller: controller,
	}
}

func (r *UserRouter) RegisterRoutes(server restServerInterface.Server) {
	server.RegisterPublicRoute(http.MethodGet, "v1/user", r.controller.GetUserById)
	server.RegisterPublicRoute(http.MethodPost, "v1/user/create", r.controller.CreateUser)
	server.RegisterPublicRoute(http.MethodPut, "v1/user/update", r.controller.UpdateUser)
	server.RegisterPublicRoute(http.MethodDelete, "v1/user", r.controller.DeleteUserById)
}
