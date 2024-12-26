package taskRest

import (
	"net/http"
	"task-tracker/adapters/controllers/rest/task/request"
	"task-tracker/adapters/controllers/rest/task/serializer"
	"task-tracker/boundary/domain/usecase"
	"task-tracker/infrastructure/errors"
	loggerInterface "task-tracker/infrastructure/logger/interface"
	"task-tracker/infrastructure/restServer/controller"
	"task-tracker/infrastructure/security/jwtService"
)

type TaskController struct {
	*restServerController.BaseController
	taskUseCase usecase.TaskUseCaseInterface
	logger      loggerInterface.Logger
}

func (c *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	creatorId, err := c.GetStrParamFromCtx(r.Context(), jwtService.UserIdKey)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	requestData := &request.CreateTaskRequest{}
	if err := c.FillReqModel(r, requestData); err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	// TODO подумать как лучше сделать
	requestData.Data.Attributes.CreatorId = creatorId

	createdTask, err := c.taskUseCase.CreateTask(r.Context(), requestData.CreateTaskDto())
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	response, err := taskSerializer.SerializeTask(createdTask)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.JsonResponse(w, r, response, http.StatusCreated)
}

func (c *TaskController) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	foundTasks, err := c.taskUseCase.GetAllTasks(r.Context())
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	response, err := taskSerializer.SerializeTasks(foundTasks)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.JsonResponse(w, r, response, http.StatusOK)
}

func (c *TaskController) GetTaskById(w http.ResponseWriter, r *http.Request) {
	// Вытаскиваем id из строки запроса вида v1/foundUser?id=1
	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		c.ErrorResponse(w, r, errors.NewError("SYS", "TaskId отсутствует"))
		return
	}

	foundTask, err := c.taskUseCase.GetTaskById(r.Context(), idStr)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	response, err := taskSerializer.SerializeTask(foundTask)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.JsonResponse(w, r, response, http.StatusOK)
}

func (c *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *TaskController) DeleteTaskById(w http.ResponseWriter, r *http.Request) {
	// Вытаскиваем id из строки запроса вида v1/foundUser?id=1
	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		c.ErrorResponse(w, r, errors.NewError("SYS", "TaskId отсутствует"))
		return
	}

	if err := c.taskUseCase.DeleteTask(r.Context(), idStr); err != nil {
		c.ErrorResponse(w, r, err)
		return
	}
}
