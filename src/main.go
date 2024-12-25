package main

import (
	"context"
	"fmt"
	taskRest "task-tracker/adapters/controllers/rest/task"
	userRest "task-tracker/adapters/controllers/rest/user"
	"task-tracker/adapters/controllers/rest/user/resolver"
	taskRepo "task-tracker/adapters/repository/task"
	userRepo "task-tracker/adapters/repository/user"
	taskUseCase "task-tracker/domain/usecase/task"
	userUseCase "task-tracker/domain/usecase/user"
	commonLogger "task-tracker/infrastructure/logger"
	"task-tracker/infrastructure/logger/zapLogger"
	mongoRepo "task-tracker/infrastructure/mongo"
	"task-tracker/infrastructure/restServer"
	restServerController "task-tracker/infrastructure/restServer/controller"
	"task-tracker/infrastructure/restServer/response"
	"task-tracker/infrastructure/security/jwtService"
	initServices "task-tracker/init-services"
	router "task-tracker/init-services/routers"
)

const AppId = "task_tracker"
const Environment = "develop"

func main() {

	mongoDB, err := mongoRepo.InitMongoDatabase("mongodb://root:root@localhost:27017", "task-tracker")
	if err != nil {
		fmt.Printf("Ошибка подключения к MongoDB: %v\n", err)
		return
	}
	defer mongoDB.Client().Disconnect(context.Background())

	mongoRepository := mongoRepo.NewMongoRepo(mongoDB)

	stopChan := make(chan struct{})
	logService := commonLogger.NewLoggerService(stopChan)
	zapLogger := zapLogger.NewZapLogger(AppId, Environment)
	logService.AddLogger("zap", zapLogger)
	logService.Start()

	logger := commonLogger.NewLogger(logService.GetInputChan())

	jwtService, _ := jwtService.NewBuilder().Secret("1").Build()

	server := restServer.NewGinServer(logger, jwtService)

	errResponseService, _ := response.NewErrorResponseService(resolver.NewErrorResolver(), logger)

	responseService, _ := response.NewResponseService(errResponseService, logger)

	baseController, _ := restServerController.NewBaseController(responseService, logger)

	// Task
	taskRepo, _ := taskRepo.NewBuilder().
		Collection("Task").
		MongoRepo(mongoRepository).
		Logger(logger).
		Build()

	taskRepo.Init(context.Background())

	taskUseCase, _ := taskUseCase.NewBuilder().
		TaskRepo(taskRepo).
		Build()

	taskController, _ := taskRest.NewBuilder().
		BaseController(baseController).
		TaskUseCase(taskUseCase).
		Logger(logger).
		Build()

	taskRouter := router.NewTaskRouter(taskController)

	// User
	userRepo, _ := userRepo.NewBuilder().
		Collection("User").
		MongoRepo(mongoRepository).
		Logger(logger).
		Build()

	userRepo.Init(context.Background())

	userUseCase, _ := userUseCase.NewBuilder().
		UserRepo(userRepo).
		JwtService(jwtService).
		Build()

	userController, _ := userRest.NewBuilder().
		BaseController(baseController).
		UserUseCase(userUseCase).
		Logger(logger).
		Build()

	userRouter := router.NewUserRouter(userController)

	globalRouter := initServices.NewGlobalRouter(server,
		userRouter,
		taskRouter)

	globalRouter.RegisterAllRoutes()

	server.Start(":8080")
}
