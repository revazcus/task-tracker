package notificationRest

import (
	"net/http"
	"task-tracker/boundary/domain/usecase"
	"task-tracker/infrastructure/restServer/controller"
)

type NotificationController struct {
	*restServerController.BaseController
	notificationUseCase usecase.NotificationUseCaseInterface
}

func NewNotificationController(controller *restServerController.BaseController, notificationUseCase usecase.NotificationUseCaseInterface) *NotificationController {
	return &NotificationController{BaseController: controller, notificationUseCase: notificationUseCase}
}

func (c *NotificationController) GetNotificationById(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *NotificationController) CreateNotification(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *NotificationController) UpdateNotification(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *NotificationController) DeleteNotificationById(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
