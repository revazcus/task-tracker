package main

import (
	"common/gateways"
	commonUserGateways "common/gateways/user"
	"context"
	"fmt"
	saramaClient "infrastructure/kafka"
	commonLogger "infrastructure/logger"
	"infrastructure/logger/zapLogger"
	mongoRepo "infrastructure/mongo"
	"infrastructure/restServer"
	restServerController "infrastructure/restServer/controller"
	restModel "infrastructure/restServer/model"
	"infrastructure/restServer/response"
	"infrastructure/security/jwtService"
	taskRest "task-service/adapters/controllers/rest/task"
	"task-service/adapters/controllers/rest/task/resolver"
	taskRepo "task-service/adapters/repository/task"
	taskUseCase "task-service/domain/usecase/task"
	initServices "task-service/init-services"
	router "task-service/init-services/routers"
)

const AppId = "task-service"
const Environment = "develop"

func main() {

	mongoDB, err := mongoRepo.InitMongoDatabase("mongodb://root:root@localhost:27017", "task-tracker")
	if err != nil {
		fmt.Printf("Ошибка подключения к MongoDB: %v\n", err)
		return
	}
	defer mongoDB.Client().Disconnect(context.Background())

	stopChan := make(chan struct{})
	logService := commonLogger.NewLoggerService(stopChan)
	zapLogger := zapLogger.NewZapLogger(AppId, Environment)
	logService.AddLogger("zap", zapLogger)
	logService.Start()

	logger := commonLogger.NewLogger(logService.GetInputChan())

	mongoRepository, _ := mongoRepo.NewBuilder().
		MongoDB(mongoDB).
		Logger(logger).
		Build()

	jwtService, _ := jwtService.NewBuilder().Secret("Pro100User_Service322").Build()

	restConfig := restModel.NewDefaultRestConfig(8082)

	server := restServer.NewGinServer(logger, jwtService, restConfig)

	errResponseService, _ := response.NewErrorResponseService(resolver.NewErrorResolver(), logger)

	responseService, _ := response.NewResponseService(errResponseService, logger)

	baseController, _ := restServerController.NewBaseController(responseService, logger)

	// Kafka
	kafkaClient, _ := saramaClient.NewSaramaClient([]string{"localhost:9093"}, "task-service-group", logger)
	if err := kafkaClient.CreateTopic(context.Background(), "task-info", 3, 1); err != nil {
		logger.Error(context.Background(), err)
	}

	// GRPS
	baseGRPCGateway := gateways.NewBaseGRPCGateway("localhost:50051", logger)
	baseGRPCGateway.Start()
	userGateway := commonUserGateways.NewUserGateway(baseGRPCGateway, logger)

	// Task
	taskRepo, _ := taskRepo.NewBuilder().
		Collection("Task").
		MongoRepo(mongoRepository).
		Logger(logger).
		Build()

	taskRepo.Init(context.Background())

	taskUseCase, _ := taskUseCase.NewBuilder().
		TaskRepo(taskRepo).
		KafkaClient(kafkaClient).
		UserGateway(userGateway).
		Build()

	taskController, _ := taskRest.NewBuilder().
		BaseController(baseController).
		TaskUseCase(taskUseCase).
		Logger(logger).
		Build()

	taskRouter := router.NewTaskRouter(taskController)

	globalRouter := initServices.NewGlobalRouter(server,
		taskRouter)

	globalRouter.RegisterAllRoutes()

	//// Открывает стартовую страницу в браузере (работает только с Windows)
	//if err := exec.Command("explorer", "http://localhost:8080").Run(); err != nil {
	//	logger.Error(context.Background(), err)
	//}

	server.Start()
}
