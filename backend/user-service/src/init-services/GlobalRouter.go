package initServices

import (
	restServerInterface "infrastructure/restServer/interface"
	routerInterface "user-service/src/init-services/routers/interface"
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
