package rest

import (
	"net/http"
	"task-tracker/boundary/domain/usecase"
	"task-tracker/infrastructure/restServer/controller"
)

type ReportController struct {
	*restServerController.BaseController
	reportUseCase usecase.ReportUseCaseInterface
}

func NewReportController(controller *restServerController.BaseController, reportUseCase usecase.ReportUseCaseInterface) *ReportController {
	return &ReportController{BaseController: controller, reportUseCase: reportUseCase}
}

func (c *ReportController) GetReportById(w http.ResponseWriter, r *http.Request) {

}

func (c *ReportController) CreateReport(w http.ResponseWriter, r *http.Request) {

}

func (c *ReportController) UpdateReport(w http.ResponseWriter, r *http.Request) {

}

func (c *ReportController) DeleteReportById(w http.ResponseWriter, r *http.Request) {

}
