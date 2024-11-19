package userRest

import (
	"net/http"
	userRestRequest "task-tracker/adapters/controllers/rest/requests"
	usecaseInterface "task-tracker/boundary/domain/usecase"
	restServerController "task-tracker/infrastructure/restServer/controller"
)

type UserController struct {
	*restServerController.BaseController
	userUseCase usecaseInterface.UserUseCaseInterface
}

func NewUserController(controller *restServerController.BaseController, userUseCase usecaseInterface.UserUseCaseInterface) *UserController {
	return &UserController{BaseController: controller, userUseCase: userUseCase}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	requestData := &userRestRequest.CreateUserRequest{}
	requestData.CreateUserDto()
}

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	requestData := &userRestRequest.CreateUserRequest{}
	requestData.CreateUserDto()
}
