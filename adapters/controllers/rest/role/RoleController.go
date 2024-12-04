package roleRest

import (
	"net/http"
	"task-tracker/boundary/domain/usecase"
	"task-tracker/infrastructure/restServer/controller"
)

type RoleController struct {
	*restServerController.BaseController
	roleUseCase usecase.RoleUseCaseInterface
}

func NewRoleController(controller *restServerController.BaseController, roleUseCase usecase.RoleUseCaseInterface) *RoleController {
	return &RoleController{BaseController: controller, roleUseCase: roleUseCase}
}

func (c *RoleController) GetRoleById(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *RoleController) CreateRole(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *RoleController) UpdateRole(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *RoleController) DeleteRoleById(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
