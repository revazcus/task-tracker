package main

import (
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
	lifecycleController := lifecycleRest.NewLifeCycleController(baseController, lifecycleUseCase)
	lifecycleRouter := router.NewLifecycleRouter(lifecycleController)

	// Notification
	notificationUseCase := &usecase.NotificationUseCase{}
	notificationController := notificationRest.NewNotificationController(baseController, notificationUseCase)
	notificationRouter := router.NewNotificationRouter(notificationController)

	// Permission
	permissionUseCase := &usecase.PermissionUseCase{}
	permissionController := permissionRest.NewPermissionController(baseController, permissionUseCase)
	permissionRouter := router.NewPermissionRouter(permissionController)

	// Project
	projectUseCase := &usecase.ProjectUseCase{}
	projectController := projectRest.NewProjectController(baseController, projectUseCase)
	projectRouter := router.NewProjectRouter(projectController)

	// Report
	reportUseCase := &usecase.ReportUseCase{}
	reportController := reportRest.NewReportController(baseController, reportUseCase)
	reportRouter := router.NewReportRouter(reportController)

	// Role
	roleUseCase := &usecase.RoleUseCase{}
	roleController := roleRest.NewRoleController(baseController, roleUseCase)
	roleRouter := router.NewRoleRouter(roleController)

	// Rule
	ruleUseCase := &usecase.RuleUseCase{}
	ruleController := ruleRest.NewRuleController(baseController, ruleUseCase)
	ruleRouter := router.NewRuleRouter(ruleController)

	// Task
	taskUseCase := &usecase.TaskUseCase{}
	taskController := taskRest.NewTaskController(baseController, taskUseCase)
	taskRouter := router.NewTaskRouter(taskController)

	// Team
	teamUseCase := &usecase.TeamUseCase{}
	teamController := teamRest.NewTeamController(baseController, teamUseCase)
	teamRouter := router.NewTeamRouter(teamController)

	// User
	userUseCase := &usecase.UserUseCase{}
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
