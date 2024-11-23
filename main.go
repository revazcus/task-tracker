package main

import (
	"task-tracker/adapters/controllers/rest"
	"task-tracker/domain/usecase"
	"task-tracker/infrastructure/logger"
	"task-tracker/infrastructure/restServer"
	"task-tracker/infrastructure/restServer/controller"
	"task-tracker/init-services"
	"task-tracker/init-services/routers"
)

func main() {

	simpleLogger := logger.NewSimpleLogger()

	server := restServer.NewGinServer(simpleLogger)

	baseController := restServerController.NewBaseController()

	// Lifecycle
	lifecycleUseCase := &usecase.LifecycleUseCase{}
	lifecycleController := rest.NewLifeCycleController(baseController, lifecycleUseCase)
	lifecycleRouter := router.NewLifecycleRouter(lifecycleController)

	// Notification
	notificationUseCase := &usecase.NotificationUseCase{}
	notificationController := rest.NewNotificationController(baseController, notificationUseCase)
	notificationRouter := router.NewNotificationRouter(notificationController)

	// Permission
	permissionUseCase := &usecase.PermissionUseCase{}
	permissionController := rest.NewPermissionController(baseController, permissionUseCase)
	permissionRouter := router.NewPermissionRouter(permissionController)

	// Project
	projectUseCase := &usecase.ProjectUseCase{}
	projectController := rest.NewProjectController(baseController, projectUseCase)
	projectRouter := router.NewProjectRouter(projectController)

	// Report
	reportUseCase := &usecase.ReportUseCase{}
	reportController := rest.NewReportController(baseController, reportUseCase)
	reportRouter := router.NewReportRouter(reportController)

	// Role
	roleUseCase := &usecase.RoleUseCase{}
	roleController := rest.NewRoleController(baseController, roleUseCase)
	roleRouter := router.NewRoleRouter(roleController)

	// Rule
	ruleUseCase := &usecase.RuleUseCase{}
	ruleController := rest.NewRuleController(baseController, ruleUseCase)
	ruleRouter := router.NewRuleRouter(ruleController)

	// Task
	taskUseCase := &usecase.TaskUseCase{}
	taskController := rest.NewTaskController(baseController, taskUseCase)
	taskRouter := router.NewTaskRouter(taskController)

	// Team
	teamUseCase := &usecase.TeamUseCase{}
	teamController := rest.NewTeamController(baseController, teamUseCase)
	teamRouter := router.NewTeamRouter(teamController)

	// User
	userUseCase := &usecase.UserUseCase{}
	userController := rest.NewUserController(baseController, userUseCase)
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
