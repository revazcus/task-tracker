package rest

import (
	"encoding/json"
	"net/http"
	"strconv"
	"task-tracker/adapters/controllers/rest/serializer"
	"task-tracker/boundary/domain/usecase"
	"task-tracker/infrastructure/restServer/controller"
)

type LifecycleController struct {
	*restServerController.BaseController
	lifecycleUseCase usecase.LifecycleUseCaseInterface
}

func NewLifeCycleController(controller *restServerController.BaseController, lifecycleUseCase usecase.LifecycleUseCaseInterface) *LifecycleController {
	return &LifecycleController{BaseController: controller, lifecycleUseCase: lifecycleUseCase}
}

func (c *LifecycleController) GetLifecycleById(w http.ResponseWriter, r *http.Request) {

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

	foundLifecycle, err := c.lifecycleUseCase.GetById(id)
	if err != nil {
		http.Error(w, `{"error":"User not found"}`, http.StatusNotFound)
		return
	}

	response, err := serializer.SerializeLifecycle(foundLifecycle)
	if err != nil {
		http.Error(w, `{"error":"Couldn't serialize response", "message": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func (c *LifecycleController) CreateLifecycle(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (c *LifecycleController) UpdateLifecycle(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (c *LifecycleController) DeleteLifecycleById(w http.ResponseWriter, r *http.Request) {

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

	if err := c.lifecycleUseCase.Delete(id); err != nil {
		http.Error(w, `{"error":"User not found"}`, http.StatusNotFound)
		return
	}
}
