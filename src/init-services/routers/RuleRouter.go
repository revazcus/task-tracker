package router

import (
	"net/http"
	ruleRest "task-tracker/adapters/controllers/rest/rule"
	restServerInterface "task-tracker/infrastructure/restServer/interface"
)

type RuleRouter struct {
	controller *ruleRest.RuleController
}

func NewRuleRouter(controller *ruleRest.RuleController) *RuleRouter {
	return &RuleRouter{
		controller: controller,
	}
}

func (r *RuleRouter) RegisterRoutes(server restServerInterface.Server) {
	server.RegisterPublicRoute(http.MethodGet, "v1/rule", r.controller.GetRuleById)
	server.RegisterPublicRoute(http.MethodPost, "v1/rule/create", r.controller.CreateRule)
	server.RegisterPublicRoute(http.MethodPut, "v1/rule/update", r.controller.UpdateRule)
	server.RegisterPublicRoute(http.MethodDelete, "v1/rule", r.controller.DeleteRuleById)
}
