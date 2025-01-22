package initService

import (
	restServerController "infrastructure/restServer/controller"
	"infrastructure/restServer/response"
	taskRest "task-service/adapters/controllers/rest/task"
	"task-service/adapters/controllers/rest/task/resolver"
)

type RestControllers struct {
	BaseController *restServerController.BaseController
	TaskController *taskRest.TaskController
}

type RestControllersInit struct {
	dc *DependencyContainer

	restControllers *RestControllers
	initFunctions   []func(*DependencyContainer) error
}

func NewRestControllersInit(dc *DependencyContainer) *RestControllersInit {
	restControllers := &RestControllers{}
	return &RestControllersInit{
		dc:              dc,
		restControllers: restControllers,
		initFunctions: []func(*DependencyContainer) error{
			restControllers.initBaseController,
			restControllers.initTaskController,
		},
	}
}

func (i *RestControllersInit) InitServices() error {
	for _, initFunc := range i.initFunctions {
		if err := initFunc(i.dc); err != nil {
			return err
		}
	}

	i.dc.RestControllers = i.restControllers
	return nil
}

func (i *RestControllersInit) StartServices() error {
	return nil
}

func (i *RestControllers) initBaseController(dc *DependencyContainer) error {
	errResponseService, err := response.NewErrorResponseService(resolver.NewErrorResolver(), dc.Logger)
	if err != nil {
		dc.LogError(err)
		return err
	}

	responseService, err := response.NewResponseService(errResponseService, dc.Logger)
	if err != nil {
		dc.LogError(err)
		return err
	}

	baseController, err := restServerController.NewBaseController(responseService, dc.Logger)
	if err != nil {
		dc.LogError(err)
		return err
	}

	i.BaseController = baseController
	return nil
}

func (i *RestControllers) initTaskController(dc *DependencyContainer) error {
	taskController, err := taskRest.NewBuilder().
		BaseController(i.BaseController).
		TaskUseCase(dc.UseCases.TaskUseCase).
		Logger(dc.Logger).
		Build()
	if err != nil {
		dc.LogError(err)
		return err
	}

	i.TaskController = taskController
	return err
}
