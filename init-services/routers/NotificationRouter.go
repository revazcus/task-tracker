package router

import (
	"net/http"
	notificationRest "task-tracker/adapters/controllers/rest/notification"
	restServerInterface "task-tracker/infrastructure/restServer/interface"
)

type NotificationRouter struct {
	controller *notificationRest.NotificationController
}

func NewNotificationRouter(controller *notificationRest.NotificationController) *NotificationRouter {
	return &NotificationRouter{
		controller: controller,
	}
}

func (r *NotificationRouter) RegisterRoutes(server restServerInterface.Server) {
	server.RegisterPublicRoute(http.MethodGet, "v1/notification", r.controller.GetNotificationById)
	server.RegisterPublicRoute(http.MethodPost, "v1/notification/create", r.controller.CreateNotification)
	server.RegisterPublicRoute(http.MethodPut, "v1/notification/update", r.controller.UpdateNotification)
	server.RegisterPublicRoute(http.MethodDelete, "v1/notification", r.controller.DeleteNotificationById)
}
