package projectRest

import (
	"net/http"
	"task-tracker/boundary/domain/usecase"
	"task-tracker/infrastructure/restServer/controller"
)

type ProjectController struct {
	*restServerController.BaseController
	projectUseCase usecase.ProjectUseCaseInterface
}

func NewProjectController(controller *restServerController.BaseController, projectUseCase usecase.ProjectUseCaseInterface) *ProjectController {
	return &ProjectController{BaseController: controller, projectUseCase: projectUseCase}
}

func (c *ProjectController) GetProjectById(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *ProjectController) CreateProject(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *ProjectController) UpdateProject(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *ProjectController) DeleteProjectById(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
