package userRest

import (
	"net/http"
	"task-tracker/adapters/controllers/rest/user/request"
	"task-tracker/adapters/controllers/rest/user/serializer"
	"task-tracker/boundary/domain/usecase"
	"task-tracker/infrastructure/errors"
	loggerInterface "task-tracker/infrastructure/logger/interface"
	"task-tracker/infrastructure/restServer/controller"
	"task-tracker/infrastructure/security/jwtService"
)

type UserController struct {
	*restServerController.BaseController
	userUseCase usecase.UserUseCaseInterface
	logger      loggerInterface.Logger
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	requestData := &request.CreateUserRequest{}

	if err := c.FillReqModel(r, requestData); err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	createdUser, err := c.userUseCase.CreateUser(r.Context(), requestData.CreateUserDto())
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	response, err := serializer.SerializeUserResponse(createdUser)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.JsonResponse(w, r, response, http.StatusCreated)
}

func (c *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	// Вытаскиваем id из строки запроса вида v1/foundUser?id=1
	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		c.ErrorResponse(w, r, errors.NewError("SYS", "UserId отсутствует"))
		return
	}

	foundUser, err := c.userUseCase.GetUserById(r.Context(), idStr)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	response, err := serializer.SerializeUser(foundUser)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.JsonResponse(w, r, response, http.StatusOK)
}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	requestData := &request.CreateUserRequest{}

	if err := c.FillReqModel(r, requestData); err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	createdUser, err := c.userUseCase.UpdateUser(r.Context(), requestData.CreateUserDto())
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	response, err := serializer.SerializeUser(createdUser)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.JsonResponse(w, r, response, http.StatusOK)
}

func (c *UserController) UpdateUserEmail(w http.ResponseWriter, r *http.Request) {
	requestData := &request.CreateUserRequest{}

	if err := c.FillReqModel(r, requestData); err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	createdUser, err := c.userUseCase.UpdateUserEmail(r.Context(), requestData.CreateUserDto())
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	response, err := serializer.SerializeUser(createdUser)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.JsonResponse(w, r, response, http.StatusOK)
}

func (c *UserController) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	requestData := &request.CreateUserRequest{}

	if err := c.FillReqModel(r, requestData); err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	createdUser, err := c.userUseCase.UpdateUserPassword(r.Context(), requestData.CreateUserDto())
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	response, err := serializer.SerializeUser(createdUser)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.JsonResponse(w, r, response, http.StatusOK)
}

func (c *UserController) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	// Вытаскиваем id из строки запроса вида v1/foundUser?id=1
	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		c.ErrorResponse(w, r, errors.NewError("SYS", "UserId отсутствует"))
		return
	}

	if err := c.userUseCase.DeleteUser(r.Context(), idStr); err != nil {
		c.ErrorResponse(w, r, err)
		return
	}
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	requestData := &request.CreateUserRequest{}

	if err := c.FillReqModel(r, requestData); err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	foundUser, err := c.userUseCase.LoginUser(r.Context(), requestData.CreateUserDto())
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	response, err := serializer.SerializeUserResponse(foundUser)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.JsonResponse(w, r, response, http.StatusOK)
}

func (c *UserController) Me(w http.ResponseWriter, r *http.Request) {
	userId, err := c.GetStrParamFromCtx(r.Context(), jwtService.UserIdKey)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	foundUser, err := c.userUseCase.GetUserById(r.Context(), userId)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	response, err := serializer.SerializeUser(foundUser)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.JsonResponse(w, r, response, http.StatusOK)
}
