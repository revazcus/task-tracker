package router

import (
	"net/http"
	reportRest "task-tracker/adapters/controllers/rest"
	restServerInterface "task-tracker/infrastructure/restServer/interface"
)

type ReportRouter struct {
	controller *reportRest.ReportController
}

func NewReportRouter(controller *reportRest.ReportController) *ReportRouter {
	return &ReportRouter{
		controller: controller,
	}
}

func (r *ReportRouter) RegisterRoutes(server restServerInterface.Server) {
	server.RegisterPublicRoute(http.MethodGet, "v1/report", r.controller.GetReportById)
	server.RegisterPublicRoute(http.MethodPost, "v1/report/create", r.controller.CreateReport)
	server.RegisterPublicRoute(http.MethodPut, "v1/report/update", r.controller.UpdateReport)
	server.RegisterPublicRoute(http.MethodDelete, "v1/report", r.controller.DeleteReportById)
}
