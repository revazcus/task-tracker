package taskRest

import (
	"fmt"
	"infrastructure/errors"
	loggerInterface "infrastructure/logger/interface"
	restServerController "infrastructure/restServer/controller"
	"infrastructure/security/jwtService"
	"net/http"
	"task-service/adapters/controllers/rest/task/request"
	taskSerializer "task-service/adapters/controllers/rest/task/serializer"
	"task-service/boundary/domain/usecase"
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
	userId, err := c.GetStrParamFromCtx(r.Context(), jwtService.UserIdKey)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	requestData := &request.CreateTaskRequest{}
	if err := c.FillReqModel(r, requestData); err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	dto := requestData.CreateTaskDto()
	c.logger.Info(r.Context(), fmt.Sprintf("User %s try to update task %s", userId, dto.Id))

	updatedTask, err := c.taskUseCase.UpdateTask(r.Context(), dto)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	response, err := taskSerializer.SerializeTask(updatedTask)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.JsonResponse(w, r, response, http.StatusOK)
}

func (c *TaskController) TakeOn(w http.ResponseWriter, r *http.Request) {
	performerId, err := c.GetStrParamFromCtx(r.Context(), jwtService.UserIdKey)
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
	dto := requestData.CreateTaskDto()
	dto.PerformerId = performerId

	updatedTask, err := c.taskUseCase.TakeOnTask(r.Context(), dto)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	response, err := taskSerializer.SerializeTask(updatedTask)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.JsonResponse(w, r, response, http.StatusOK)
}

func (c *TaskController) AddPerformer(w http.ResponseWriter, r *http.Request) {
	requestData := &request.CreateTaskRequest{}
	if err := c.FillReqModel(r, requestData); err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	updatedTask, err := c.taskUseCase.AddPerformer(r.Context(), requestData.CreateTaskDto())
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	response, err := taskSerializer.SerializeTask(updatedTask)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.JsonResponse(w, r, response, http.StatusOK)
}

func (c *TaskController) AddTimeCosts(w http.ResponseWriter, r *http.Request) {
	userId, err := c.GetStrParamFromCtx(r.Context(), jwtService.UserIdKey)
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
	dto := requestData.CreateTaskDto()
	dto.TimeCosts.UserId = userId

	updatedTask, err := c.taskUseCase.AddTimeCosts(r.Context(), dto)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	response, err := taskSerializer.SerializeTask(updatedTask)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.JsonResponse(w, r, response, http.StatusOK)
}

func (c *TaskController) AddComment(w http.ResponseWriter, r *http.Request) {
	authorId, err := c.GetStrParamFromCtx(r.Context(), jwtService.UserIdKey)
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
	dto := requestData.CreateTaskDto()
	dto.Comments.UserId = authorId

	updatedTask, err := c.taskUseCase.AddComment(r.Context(), dto)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	response, err := taskSerializer.SerializeTask(updatedTask)
	if err != nil {
		c.ErrorResponse(w, r, err)
		return
	}

	c.JsonResponse(w, r, response, http.StatusOK)
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
