package main

import (
	"context"
	"fmt"
	"github.com/revazcus/task-tracker/backend/common/gateways/user/gateways"
	commonUserGateways "github.com/revazcus/task-tracker/backend/common/gateways/user/gateways/user"
	saramaClient "github.com/revazcus/task-tracker/backend/infrastructure/kafka"
	commonLogger "github.com/revazcus/task-tracker/backend/infrastructure/logger"
	"github.com/revazcus/task-tracker/backend/infrastructure/logger/zapLogger"
	mongoRepo "github.com/revazcus/task-tracker/backend/infrastructure/mongo"
	"github.com/revazcus/task-tracker/backend/infrastructure/restServer"
	restServerController "github.com/revazcus/task-tracker/backend/infrastructure/restServer/controller"
	"github.com/revazcus/task-tracker/backend/infrastructure/restServer/response"
	"github.com/revazcus/task-tracker/backend/infrastructure/security/jwtService"
	taskRest "github.com/revazcus/task-tracker/backend/task-service/adapters/controllers/rest/task"
	"github.com/revazcus/task-tracker/backend/task-service/adapters/controllers/rest/task/resolver"
	taskRepo "github.com/revazcus/task-tracker/backend/task-service/adapters/repository/task"
	taskUseCase "github.com/revazcus/task-tracker/backend/task-service/domain/usecase/task"
	initServices "github.com/revazcus/task-tracker/backend/task-service/init-services"
	router "github.com/revazcus/task-tracker/backend/task-service/init-services/routers"
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

	jwtService, _ := jwtService.NewBuilder().Secret("1").Build()

	server := restServer.NewGinServer(logger, jwtService)

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

	server.Start(":8082")
}
