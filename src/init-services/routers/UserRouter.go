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
	server.RegisterPublicRoute(http.MethodPost, "v1/user/register", r.controller.CreateUser)
	server.RegisterPublicRoute(http.MethodPost, "v1/user/login", r.controller.Login)
	server.RegisterPrivateRoute(http.MethodGet, "v1/user", r.controller.GetUserById)
	server.RegisterPrivateRoute(http.MethodGet, "v1/user/me", r.controller.Me)
	server.RegisterPrivateRoute(http.MethodPut, "v1/user/update", r.controller.UpdateUser)
	server.RegisterPrivateRoute(http.MethodPut, "v1/user/updateEmail", r.controller.UpdateUserEmail)
	server.RegisterPrivateRoute(http.MethodPut, "v1/user/updatePassword", r.controller.UpdateUserPassword)
	server.RegisterPrivateRoute(http.MethodDelete, "v1/user", r.controller.DeleteUserById)
}
