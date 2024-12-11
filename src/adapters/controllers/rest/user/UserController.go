package userRest

import (
	"encoding/json"
	"net/http"
	"task-tracker/adapters/controllers/rest/user/request"
	"task-tracker/adapters/controllers/rest/user/serializer"
	"task-tracker/boundary/domain/usecase"
	loggerInterface "task-tracker/infrastructure/logger/interface"
	"task-tracker/infrastructure/restServer/controller"
)

type UserController struct {
	*restServerController.BaseController
	userUseCase usecase.UserUseCaseInterface
	logger      loggerInterface.Logger
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	requestData := &request.UserRequest{}

	if err := c.FillReqModel(r, requestData); err != nil {
		http.Error(w, `{"error":"Invalid request", "message": "`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	createdUser, err := c.userUseCase.CreateUser(r.Context(), requestData.CreateUserDto())
	if err != nil {
		http.Error(w, `{"error":"Couldn't create user", "message": "`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	response, err := serializer.SerializeUser(createdUser)
	if err != nil {
		http.Error(w, `{"error":"Couldn't serialize response", "message": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	requestData := &request.UserRequest{}

	if err := c.FillReqModel(r, requestData); err != nil {
		http.Error(w, `{"error":"Invalid request", "message": "`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	createdUser, err := c.userUseCase.UpdateUser(r.Context(), requestData.CreateUserDto())
	if err != nil {
		http.Error(w, `{"error":"Couldn't update user", "message": "`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	response, err := serializer.SerializeUser(createdUser)
	if err != nil {
		http.Error(w, `{"error":"Couldn't serialize response", "message": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (c *UserController) UpdateUserEmail(w http.ResponseWriter, r *http.Request) {
	requestData := &request.UserRequest{}

	if err := c.FillReqModel(r, requestData); err != nil {
		http.Error(w, `{"error":"Invalid request", "message": "`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	createdUser, err := c.userUseCase.UpdateUserEmail(r.Context(), requestData.CreateUserDto())
	if err != nil {
		http.Error(w, `{"error":"Couldn't update user", "message": "`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	response, err := serializer.SerializeUser(createdUser)
	if err != nil {
		http.Error(w, `{"error":"Couldn't serialize response", "message": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (c *UserController) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	requestData := &request.UserRequest{}

	if err := c.FillReqModel(r, requestData); err != nil {
		http.Error(w, `{"error":"Invalid request", "message": "`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	createdUser, err := c.userUseCase.UpdateUserPassword(r.Context(), requestData.CreateUserDto())
	if err != nil {
		http.Error(w, `{"error":"Couldn't update user", "message": "`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	response, err := serializer.SerializeUser(createdUser)
	if err != nil {
		http.Error(w, `{"error":"Couldn't serialize response", "message": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (c *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	// Вытаскиваем id из строки запроса вида v1/foundUser?id=1
	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		http.Error(w, `{"error":"Missing id parameter"}`, http.StatusBadRequest)
		return
	}

	foundUser, err := c.userUseCase.GetUserById(r.Context(), idStr)
	if err != nil {
		http.Error(w, `{"error":"User not found"}`, http.StatusNotFound)
		return
	}

	response, err := serializer.SerializeUser(foundUser)
	if err != nil {
		http.Error(w, `{"error":"Couldn't serialize response", "message": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (c *UserController) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	// Вытаскиваем id из строки запроса вида v1/foundUser?id=1
	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		http.Error(w, `{"error":"Missing id parameter"}`, http.StatusBadRequest)
		return
	}

	if err := c.userUseCase.DeleteUser(r.Context(), idStr); err != nil {
		http.Error(w, `{"error":"User not found"}`, http.StatusNotFound)
		return
	}
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	requestData := &request.UserRequest{}

	if err := c.FillReqModel(r, requestData); err != nil {
		http.Error(w, `{"error":"Invalid request", "message": "`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	if requestData.Data.Id == "" {
		http.Error(w, `{"error":"Invalid id parameter, "message": "id отсутствует или невалиден"}`, http.StatusBadRequest)
		return
	}

	foundUser, err := c.userUseCase.LoginUser(r.Context(), requestData.CreateUserDto())
	if err != nil {
		http.Error(w, `{"error":"Couldn't login", "message": "`+err.Error()+`"}`, http.StatusUnauthorized)
		return
	}

	response, err := serializer.SerializeUser(foundUser)
	if err != nil {
		http.Error(w, `{"error":"Couldn't serialize response", "message": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
