package routerInterface

import restServerInterface "github.com/revazcus/task-tracker/backend/infrastructure/restServer/interface"

type Router interface {
	RegisterRoutes(server restServerInterface.Server)
}
