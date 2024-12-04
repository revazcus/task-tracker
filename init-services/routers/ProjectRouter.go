package router

import (
	"net/http"
	projectRest "task-tracker/adapters/controllers/rest/project"
	restServerInterface "task-tracker/infrastructure/restServer/interface"
)

type ProjectRouter struct {
	controller *projectRest.ProjectController
}

func NewProjectRouter(controller *projectRest.ProjectController) *ProjectRouter {
	return &ProjectRouter{
		controller: controller,
	}
}

func (r *ProjectRouter) RegisterRoutes(server restServerInterface.Server) {
	server.RegisterPublicRoute(http.MethodGet, "v1/project", r.controller.GetProjectById)
	server.RegisterPublicRoute(http.MethodPost, "v1/project/create", r.controller.CreateProject)
	server.RegisterPublicRoute(http.MethodPut, "v1/project/update", r.controller.UpdateProject)
	server.RegisterPublicRoute(http.MethodDelete, "v1/project", r.controller.DeleteProjectById)
}
