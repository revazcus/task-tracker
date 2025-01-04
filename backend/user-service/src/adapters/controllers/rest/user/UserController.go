package userRest

import (
	"github.com/revazcus/task-tracker/backend/infrastructure/errors"
	loggerInterface "github.com/revazcus/task-tracker/backend/infrastructure/logger/interface"
	restServerController "github.com/revazcus/task-tracker/backend/infrastructure/restServer/controller"
	"github.com/revazcus/task-tracker/backend/infrastructure/security/jwtService"
	"github.com/revazcus/task-tracker/backend/user-service/adapters/controllers/rest/user/request"
	userSerializer "github.com/revazcus/task-tracker/backend/user-service/adapters/controllers/rest/user/serializer"
	"github.com/revazcus/task-tracker/backend/user-service/boundary/domain/usecase"
	"net/http"
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

	response, err := userSerializer.SerializeUserResponse(createdUser)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.JsonResponse(w, r, response, http.StatusCreated)
}

func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	foundUsers, err := c.userUseCase.GetAllUsers(r.Context())
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	response, err := userSerializer.SerializeUsers(foundUsers)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.JsonResponse(w, r, response, http.StatusOK)
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

	response, err := userSerializer.SerializeUser(foundUser)
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

	response, err := userSerializer.SerializeUser(createdUser)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.JsonResponse(w, r, response, http.StatusOK)
}

func (c *UserController) UpdateEmail(w http.ResponseWriter, r *http.Request) {
	requestData := &request.CreateUserRequest{}

	if err := c.FillReqModel(r, requestData); err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	updatedUser, err := c.userUseCase.UpdateEmail(r.Context(), requestData.CreateUserDto())
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	response, err := userSerializer.SerializeUser(updatedUser)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.JsonResponse(w, r, response, http.StatusOK)
}

func (c *UserController) UpdateUsername(w http.ResponseWriter, r *http.Request) {
	requestData := &request.CreateUserRequest{}
	if err := c.FillReqModel(r, requestData); err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	updatedUser, err := c.userUseCase.UpdateUsername(r.Context(), requestData.CreateUserDto())
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	response, err := userSerializer.SerializeUser(updatedUser)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.JsonResponse(w, r, response, http.StatusOK)
}

func (c *UserController) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	requestData := &request.CreateUserRequest{}

	if err := c.FillReqModel(r, requestData); err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	updatedUser, err := c.userUseCase.UpdatePassword(r.Context(), requestData.CreateUserDto())
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	response, err := userSerializer.SerializeUser(updatedUser)
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

	response, err := userSerializer.SerializeUserResponse(foundUser)
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

	response, err := userSerializer.SerializeUser(foundUser)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.JsonResponse(w, r, response, http.StatusOK)
}
