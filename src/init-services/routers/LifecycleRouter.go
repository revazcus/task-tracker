package router

import (
	"net/http"
	lifecycleRest "task-tracker/adapters/controllers/rest/lifecycle"
	restServerInterface "task-tracker/infrastructure/restServer/interface"
)

type LifecycleRouter struct {
	controller *lifecycleRest.LifecycleController
}

func NewLifecycleRouter(controller *lifecycleRest.LifecycleController) *LifecycleRouter {
	return &LifecycleRouter{
		controller: controller,
	}
}

func (r *LifecycleRouter) RegisterRoutes(server restServerInterface.Server) {
	server.RegisterPublicRoute(http.MethodGet, "v1/lifecycle", r.controller.GetLifecycleById)
	server.RegisterPublicRoute(http.MethodPost, "v1/lifecycle/create", r.controller.CreateLifecycle)
	server.RegisterPublicRoute(http.MethodPut, "v1/lifecycle/update", r.controller.UpdateLifecycle)
	server.RegisterPublicRoute(http.MethodDelete, "v1/lifecycle", r.controller.DeleteLifecycleById)
}
