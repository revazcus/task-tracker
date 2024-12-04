package router

import (
	"net/http"
	taskRest "task-tracker/adapters/controllers/rest/task"
	restServerInterface "task-tracker/infrastructure/restServer/interface"
)

type TaskRouter struct {
	controller *taskRest.TaskController
}

func NewTaskRouter(controller *taskRest.TaskController) *TaskRouter {
	return &TaskRouter{
		controller: controller,
	}
}

func (r *TaskRouter) RegisterRoutes(server restServerInterface.Server) {
	server.RegisterPublicRoute(http.MethodGet, "v1/task", r.controller.GetTaskById)
	server.RegisterPublicRoute(http.MethodPost, "v1/task/create", r.controller.CreateTask)
	server.RegisterPublicRoute(http.MethodPut, "v1/task/update", r.controller.UpdateTask)
	server.RegisterPublicRoute(http.MethodDelete, "v1/task", r.controller.DeleteTaskById)
}
