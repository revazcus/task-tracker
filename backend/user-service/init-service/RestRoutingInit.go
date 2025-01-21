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
	dc.RestServer.RegisterPublicRoute(http.MethodPost, "v1/user/register", dc.RestControllers.UserController.CreateUser)
	dc.RestServer.RegisterPublicRoute(http.MethodPost, "v1/user/login", dc.RestControllers.UserController.Login)
	dc.RestServer.RegisterPrivateRoute(http.MethodGet, "v1/users", dc.RestControllers.UserController.GetAllUsers)
	dc.RestServer.RegisterPrivateRoute(http.MethodGet, "v1/user", dc.RestControllers.UserController.GetUserById)
	dc.RestServer.RegisterPrivateRoute(http.MethodGet, "v1/user/me", dc.RestControllers.UserController.Me)
	dc.RestServer.RegisterPrivateRoute(http.MethodPut, "v1/user/update", dc.RestControllers.UserController.UpdateUser)
	dc.RestServer.RegisterPrivateRoute(http.MethodPut, "v1/user/updateEmail", dc.RestControllers.UserController.UpdateEmail)
	dc.RestServer.RegisterPrivateRoute(http.MethodPut, "v1/user/updateUsername", dc.RestControllers.UserController.UpdateUsername)
	dc.RestServer.RegisterPrivateRoute(http.MethodPut, "v1/user/updatePassword", dc.RestControllers.UserController.UpdatePassword)
	dc.RestServer.RegisterPrivateRoute(http.MethodDelete, "v1/user", dc.RestControllers.UserController.DeleteUserById)
}
