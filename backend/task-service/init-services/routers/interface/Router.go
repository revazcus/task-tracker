package routerInterface

import restServerInterface "infrastructure/restServer/interface"

type Router interface {
	RegisterRoutes(server restServerInterface.Server)
}
