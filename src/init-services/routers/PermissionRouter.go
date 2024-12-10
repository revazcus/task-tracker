package router

import (
	"net/http"
	permissionRest "task-tracker/adapters/controllers/rest/permission"
	restServerInterface "task-tracker/infrastructure/restServer/interface"
)

type PermissionRouter struct {
	controller *permissionRest.PermissionController
}

func NewPermissionRouter(controller *permissionRest.PermissionController) *PermissionRouter {
	return &PermissionRouter{
		controller: controller,
	}
}

func (r *PermissionRouter) RegisterRoutes(server restServerInterface.Server) {
	server.RegisterPublicRoute(http.MethodGet, "v1/permission", r.controller.GetPermissionById)
	server.RegisterPublicRoute(http.MethodPost, "v1/permission/create", r.controller.CreatePermission)
	server.RegisterPublicRoute(http.MethodPut, "v1/permission/update", r.controller.UpdatePermission)
	server.RegisterPublicRoute(http.MethodDelete, "v1/permission", r.controller.DeletePermissionById)
}
