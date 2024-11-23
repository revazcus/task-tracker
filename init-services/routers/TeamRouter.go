package router

import (
	"net/http"
	teamRest "task-tracker/adapters/controllers/rest"
	restServerInterface "task-tracker/infrastructure/restServer/interface"
)

type TeamRouter struct {
	controller *teamRest.TeamController
}

func NewTeamRouter(controller *teamRest.TeamController) *TeamRouter {
	return &TeamRouter{
		controller: controller,
	}
}

func (r *TeamRouter) RegisterRoutes(server restServerInterface.Server) {
	server.RegisterPublicRoute(http.MethodGet, "v1/team", r.controller.GetTeamById)
	server.RegisterPublicRoute(http.MethodPost, "v1/team/create", r.controller.CreateTeam)
	server.RegisterPublicRoute(http.MethodPut, "v1/team/update", r.controller.UpdateTeam)
	server.RegisterPublicRoute(http.MethodDelete, "v1/team", r.controller.DeleteTeamById)
}
