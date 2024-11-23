package rest

import (
	"encoding/json"
	"net/http"
	"strconv"
	"task-tracker/adapters/controllers/rest/request"
	"task-tracker/adapters/controllers/rest/serializer"
	"task-tracker/boundary/domain/usecase"
	"task-tracker/infrastructure/restServer/controller"
)

type UserController struct {
	*restServerController.BaseController
	userUseCase usecase.UserUseCaseInterface
}

func NewUserController(controller *restServerController.BaseController, userUseCase usecase.UserUseCaseInterface) *UserController {
	return &UserController{BaseController: controller, userUseCase: userUseCase}
}

func (c *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	// Вытаскиваем id из строки запроса вида v1/foundUser?id=1
	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		http.Error(w, `{"error":"Missing id parameter"}`, http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error":"Invalid id format"}`, http.StatusBadRequest)
		return
	}

	foundUser, err := c.userUseCase.GetById(id)
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

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	requestData := &restRequest.UserRequest{}

	if err := c.FillReqModel(r, requestData); err != nil {
		http.Error(w, `{"error":"Invalid request", "message": "`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	createdUser, err := c.userUseCase.Create(requestData.CreateUserDto())
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
	requestData := &restRequest.UserRequest{}

	if err := c.FillReqModel(r, requestData); err != nil {
		http.Error(w, `{"error":"Invalid request", "message": "`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	createdUser, err := c.userUseCase.Update(requestData.CreateUserDto())
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

func (c *UserController) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	// Вытаскиваем id из строки запроса вида v1/foundUser?id=1
	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		http.Error(w, `{"error":"Missing id parameter"}`, http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error":"Invalid id format"}`, http.StatusBadRequest)
		return
	}

	if err := c.userUseCase.Delete(id); err != nil {
		http.Error(w, `{"error":"User not found"}`, http.StatusNotFound)
		return
	}
}
