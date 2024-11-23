package routerInterface

import restServerInterface "task-tracker/infrastructure/restServer/interface"

type Router interface {
	RegisterRoutes(server restServerInterface.Server)
}
