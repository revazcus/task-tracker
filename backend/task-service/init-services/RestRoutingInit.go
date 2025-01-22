package initService

import "net/http"

type RestRoutingInit struct {
	dc *DependencyContainer
}

func NewRestRoutingInit(dc *DependencyContainer) *RestRoutingInit {
	return &RestRoutingInit{
		dc: dc,
	}
}

func (i *RestRoutingInit) InitServices() error {
	i.InitRouting(i.dc)
	return nil
}

func (i *RestRoutingInit) StartServices() error {
	if err := i.dc.RestServer.Start(); err != nil {
		return err
	}
	return nil
}

func (i *RestRoutingInit) InitRouting(dc *DependencyContainer) {
	dc.RestServer.RegisterPrivateRoute(http.MethodPost, "v1/task/create", dc.RestControllers.TaskController.CreateTask)
	dc.RestServer.RegisterPrivateRoute(http.MethodGet, "v1/tasks", dc.RestControllers.TaskController.GetAllTasks)
	dc.RestServer.RegisterPrivateRoute(http.MethodGet, "v1/task", dc.RestControllers.TaskController.GetTaskById)
	dc.RestServer.RegisterPrivateRoute(http.MethodPut, "v1/task/update", dc.RestControllers.TaskController.UpdateTask)
	dc.RestServer.RegisterPrivateRoute(http.MethodPut, "v1/task/takeOn", dc.RestControllers.TaskController.TakeOn)
	dc.RestServer.RegisterPrivateRoute(http.MethodPut, "v1/task/addPerformer", dc.RestControllers.TaskController.AddPerformer)
	dc.RestServer.RegisterPrivateRoute(http.MethodPut, "v1/task/addTimeCosts", dc.RestControllers.TaskController.AddTimeCosts)
	dc.RestServer.RegisterPrivateRoute(http.MethodPut, "v1/task/addComment", dc.RestControllers.TaskController.AddComment)
	dc.RestServer.RegisterPrivateRoute(http.MethodDelete, "v1/task", dc.RestControllers.TaskController.DeleteTaskById)
}
