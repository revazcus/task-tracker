package main

import (
	"context"
	"fmt"
	saramaClient "infrastructure/kafka"
	commonLogger "infrastructure/logger"
	"infrastructure/logger/zapLogger"
	mongoRepo "infrastructure/mongo"
	"infrastructure/restServer"
	restServerController "infrastructure/restServer/controller"
	"infrastructure/restServer/response"
	"infrastructure/security/jwtService"
	"user-service/src/adapters/controllers/grpc"
	userRest "user-service/src/adapters/controllers/rest"
	"user-service/src/adapters/controllers/rest/resolver"
	userRepo "user-service/src/adapters/repository/user"
	userUseCase "user-service/src/domain/usecase"
	initServices "user-service/src/init-services"
	router "user-service/src/init-services/routers"
)

const AppId = "user-service"
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

	// Kafka
	kafkaClient, _ := saramaClient.NewSaramaClient([]string{"localhost:9093"}, "user-service-group", logger)
	if err := kafkaClient.CreateTopic(context.Background(), "user-info", 3, 1); err != nil {
		logger.Error(context.Background(), err)
	}

	// GRPC
	grpcController := grpc.NewUserController(userUseCase)
	grpcServer := grpc.NewUserServer(":50051", grpcController, logger)
	grpcServer.Start()

	// Register all routes
	globalRouter := initServices.NewGlobalRouter(server,
		userRouter)

	globalRouter.RegisterAllRoutes()

	//// Открывает стартовую страницу в браузере (работает только с Windows)
	//if err := exec.Command("explorer", "http://localhost:8080").Run(); err != nil {
	//	logger.Error(context.Background(), err)
	//}

	server.Start(":8081")

}
