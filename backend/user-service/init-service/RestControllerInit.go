package initService

import (
	restServerController "infrastructure/restServer/controller"
	"infrastructure/restServer/response"
	userRest "user-service/adapters/controllers/rest"
	"user-service/adapters/controllers/rest/resolver"
)

type RestControllers struct {
	BaseController *restServerController.BaseController
	UserController *userRest.UserController
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
			restControllers.initUserController,
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
	errResponseService, err := restResponse.NewErrorResponseServiceBuilder().
		ErrorResolver(resolver.NewErrorResolver()).
		Logger(dc.Logger).
		Build()
	if err != nil {
		dc.LogError(err)
		return err
	}

	responseService, err := restResponse.NewResponseServiceBuilder().
		ErrorResponseService(errResponseService).
		Logger(dc.Logger).
		Build()
	if err != nil {
		dc.LogError(err)
		return err
	}

	baseController, err := restServerController.NewBaseControllerBuilder().
		ResponseService(responseService).
		Logger(dc.Logger).
		Build()
	if err != nil {
		dc.LogError(err)
		return err
	}

	i.BaseController = baseController
	return nil
}

func (i *RestControllers) initUserController(dc *DependencyContainer) error {
	userController, err := userRest.NewBuilder().
		BaseController(i.BaseController).
		UserUseCase(dc.UseCases.UserUseCase).
		Logger(dc.Logger).
		Build()
	if err != nil {
		dc.LogError(err)
		return err
	}

	i.UserController = userController
	return err
}
