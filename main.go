package main

import (
	userRest "task-tracker/adapters/controllers/rest"
	"task-tracker/domain/usecase"
	logger "task-tracker/infrastructure/logger"
	"task-tracker/infrastructure/restServer"
	restServerController "task-tracker/infrastructure/restServer/controller"
	initservices "task-tracker/init-services"
)

func main() {

	simpleLogger := logger.NewSimpleLogger()

	server := restServer.NewGinServer(simpleLogger)

	baseController := restServerController.NewBaseController()

	userUseCase := &usecase.UserUseCase{}

	controller := userRest.NewUserController(baseController, userUseCase)

	router := initservices.NewUserRouter(server, controller)

	router.RegisterRoutes()

	server.Start(":8080")
}
