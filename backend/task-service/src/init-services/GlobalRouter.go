package initServices

import (
	restServerInterface "github.com/revazcus/task-tracker/backend/infrastructure/restServer/interface"
	routerInterface "github.com/revazcus/task-tracker/backend/task-service/init-services/routers/interface"
)

type GlobalRouter struct {
	server  restServerInterface.Server
	routers []routerInterface.Router
}

func NewGlobalRouter(server restServerInterface.Server, routers ...routerInterface.Router) *GlobalRouter {
	return NewGlobalRouterFromSlice(server, routers)
}

func NewGlobalRouterFromSlice(server restServerInterface.Server, routers []routerInterface.Router) *GlobalRouter {
	return &GlobalRouter{
		server:  server,
		routers: routers,
	}
}

func (gr *GlobalRouter) RegisterAllRoutes() {
	for _, router := range gr.routers {
		router.RegisterRoutes(gr.server)
	}
}
