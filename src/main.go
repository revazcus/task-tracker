package main

import (
	"context"
	"fmt"
	lifecycleRest "task-tracker/adapters/controllers/rest/lifecycle"
	notificationRest "task-tracker/adapters/controllers/rest/notification"
	permissionRest "task-tracker/adapters/controllers/rest/permission"
	projectRest "task-tracker/adapters/controllers/rest/project"
	reportRest "task-tracker/adapters/controllers/rest/report"
	roleRest "task-tracker/adapters/controllers/rest/role"
	ruleRest "task-tracker/adapters/controllers/rest/rule"
	taskRest "task-tracker/adapters/controllers/rest/task"
	teamRest "task-tracker/adapters/controllers/rest/team"
	userRest "task-tracker/adapters/controllers/rest/user"
	userRepo "task-tracker/adapters/repository/user"
	lifecycleUseCase "task-tracker/domain/usecase/lifecycle"
	notificationUseCase "task-tracker/domain/usecase/notification"
	permissionUseCase "task-tracker/domain/usecase/permission"
	projectUseCase "task-tracker/domain/usecase/project"
	reportUseCase "task-tracker/domain/usecase/report"
	roleUseCase "task-tracker/domain/usecase/role"
	ruleUseCase "task-tracker/domain/usecase/rule"
	taskUseCase "task-tracker/domain/usecase/task"
	teamUseCase "task-tracker/domain/usecase/team"
	userUseCase "task-tracker/domain/usecase/user"
	"task-tracker/infrastructure/logger"
	mongoRepo "task-tracker/infrastructure/mongo"
	"task-tracker/infrastructure/restServer"
	restServerController "task-tracker/infrastructure/restServer/controller"
	"task-tracker/infrastructure/security/jwtService"
	initServices "task-tracker/init-services"
	router "task-tracker/init-services/routers"
)

func main() {

	mongoDB, err := mongoRepo.InitMongoDatabase("mongodb://root:root@localhost:27017", "task-tracker")
	if err != nil {
		fmt.Printf("Ошибка подключения к MongoDB: %v\n", err)
		return
	}
	defer mongoDB.Client().Disconnect(context.Background())

	mongoRepository := mongoRepo.NewMongoRepo(mongoDB)

	simpleLogger := logger.NewSimpleLogger()

	jwtService, _ := jwtService.NewBuilder().Secret("1").Build()

	server := restServer.NewGinServer(simpleLogger, jwtService)

	baseController := restServerController.NewBaseController()

	// Lifecycle
	lifecycleUseCase := &lifecycleUseCase.LifecycleUseCase{}
	lifecycleController := lifecycleRest.NewLifeCycleController(baseController, lifecycleUseCase)
	lifecycleRouter := router.NewLifecycleRouter(lifecycleController)

	// Notification
	notificationUseCase := &notificationUseCase.NotificationUseCase{}
	notificationController := notificationRest.NewNotificationController(baseController, notificationUseCase)
	notificationRouter := router.NewNotificationRouter(notificationController)

	// Permission
	permissionUseCase := &permissionUseCase.PermissionUseCase{}
	permissionController := permissionRest.NewPermissionController(baseController, permissionUseCase)
	permissionRouter := router.NewPermissionRouter(permissionController)

	// Project
	projectUseCase := &projectUseCase.ProjectUseCase{}
	projectController := projectRest.NewProjectController(baseController, projectUseCase)
	projectRouter := router.NewProjectRouter(projectController)

	// Report
	reportUseCase := &reportUseCase.ReportUseCase{}
	reportController := reportRest.NewReportController(baseController, reportUseCase)
	reportRouter := router.NewReportRouter(reportController)

	// Role
	roleUseCase := &roleUseCase.RoleUseCase{}
	roleController := roleRest.NewRoleController(baseController, roleUseCase)
	roleRouter := router.NewRoleRouter(roleController)

	// Rule
	ruleUseCase := &ruleUseCase.RuleUseCase{}
	ruleController := ruleRest.NewRuleController(baseController, ruleUseCase)
	ruleRouter := router.NewRuleRouter(ruleController)

	// Task
	taskUseCase := &taskUseCase.TaskUseCase{}
	taskController := taskRest.NewTaskController(baseController, taskUseCase)
	taskRouter := router.NewTaskRouter(taskController)

	// Team
	teamUseCase := &teamUseCase.TeamUseCase{}
	teamController := teamRest.NewTeamController(baseController, teamUseCase)
	teamRouter := router.NewTeamRouter(teamController)

	// User
	userRepo, _ := userRepo.NewBuilder().
		Table("User").
		MongoRepo(mongoRepository).
		Logger(simpleLogger).
		Build()
	userUseCase, _ := userUseCase.NewBuilder().
		UserRepo(userRepo).
		JwtService(jwtService).
		Build()
	userController, _ := userRest.NewBuilder().
		BaseController(baseController).
		UserUseCase(userUseCase).
		Logger(simpleLogger).
		Build()
	userRouter := router.NewUserRouter(userController)

	globalRouter := initServices.NewGlobalRouter(server,
		userRouter,
		lifecycleRouter,
		notificationRouter,
		permissionRouter,
		projectRouter,
		reportRouter,
		roleRouter,
		ruleRouter,
		taskRouter,
		teamRouter)

	globalRouter.RegisterAllRoutes()

	server.Start(":8080")
}
