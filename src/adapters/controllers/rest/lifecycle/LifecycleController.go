package lifecycleRest

import (
	"net/http"
	"task-tracker/boundary/domain/usecase"
	"task-tracker/infrastructure/restServer/controller"
)

type LifecycleController struct {
	*restServerController.BaseController
	lifecycleUseCase usecase.LifecycleUseCaseInterface
}

func NewLifeCycleController(controller *restServerController.BaseController, lifecycleUseCase usecase.LifecycleUseCaseInterface) *LifecycleController {
	return &LifecycleController{BaseController: controller, lifecycleUseCase: lifecycleUseCase}
}

func (c *LifecycleController) GetLifecycleById(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *LifecycleController) CreateLifecycle(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *LifecycleController) UpdateLifecycle(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *LifecycleController) DeleteLifecycleById(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
