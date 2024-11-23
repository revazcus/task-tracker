package rest

import (
	"net/http"
	"task-tracker/boundary/domain/usecase"
	"task-tracker/infrastructure/restServer/controller"
)

type PermissionController struct {
	*restServerController.BaseController
	permissionUseCase usecase.PermissionUseCaseInterface
}

func NewPermissionController(controller *restServerController.BaseController, permissionUseCase usecase.PermissionUseCaseInterface) *PermissionController {
	return &PermissionController{BaseController: controller, permissionUseCase: permissionUseCase}
}

func (c *PermissionController) GetPermissionById(w http.ResponseWriter, r *http.Request) {

}

func (c *PermissionController) CreatePermission(w http.ResponseWriter, r *http.Request) {

}

func (c *PermissionController) UpdatePermission(w http.ResponseWriter, r *http.Request) {

}

func (c *PermissionController) DeletePermissionById(w http.ResponseWriter, r *http.Request) {

}
