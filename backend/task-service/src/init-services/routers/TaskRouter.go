package router

import (
	restServerInterface "infrastructure/restServer/interface"
	"net/http"
	taskRest "task-service/src/adapters/controllers/rest/task"
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
	server.RegisterPrivateRoute(http.MethodPost, "v1/task/create", r.controller.CreateTask)
	server.RegisterPrivateRoute(http.MethodGet, "v1/tasks", r.controller.GetAllTasks)
	server.RegisterPrivateRoute(http.MethodGet, "v1/task", r.controller.GetTaskById)
	server.RegisterPrivateRoute(http.MethodPut, "v1/task/update", r.controller.UpdateTask)
	server.RegisterPrivateRoute(http.MethodPut, "v1/task/takeOn", r.controller.TakeOn)
	server.RegisterPrivateRoute(http.MethodPut, "v1/task/addPerformer", r.controller.AddPerformer)
	server.RegisterPrivateRoute(http.MethodPut, "v1/task/addTimeCosts", r.controller.AddTimeCosts)
	server.RegisterPrivateRoute(http.MethodPut, "v1/task/addComment", r.controller.AddComment)
	server.RegisterPrivateRoute(http.MethodDelete, "v1/task", r.controller.DeleteTaskById)
}
