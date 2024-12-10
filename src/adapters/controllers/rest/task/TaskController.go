package taskRest

import (
	"net/http"
	"task-tracker/boundary/domain/usecase"
	"task-tracker/infrastructure/restServer/controller"
)

type TaskController struct {
	*restServerController.BaseController
	taskUseCase usecase.TaskUseCaseInterface
}

func NewTaskController(controller *restServerController.BaseController, taskUseCase usecase.TaskUseCaseInterface) *TaskController {
	return &TaskController{BaseController: controller, taskUseCase: taskUseCase}
}

func (c *TaskController) GetTaskById(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *TaskController) DeleteTaskById(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
