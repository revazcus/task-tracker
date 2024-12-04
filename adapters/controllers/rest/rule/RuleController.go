package ruleRest

import (
	"net/http"
	"task-tracker/boundary/domain/usecase"
	"task-tracker/infrastructure/restServer/controller"
)

type RuleController struct {
	*restServerController.BaseController
	ruleUseCase usecase.RuleUseCaseInterface
}

func NewRuleController(controller *restServerController.BaseController, ruleUseCase usecase.RuleUseCaseInterface) *RuleController {
	return &RuleController{BaseController: controller, ruleUseCase: ruleUseCase}
}

func (c *RuleController) GetRuleById(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *RuleController) CreateRule(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *RuleController) UpdateRule(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *RuleController) DeleteRuleById(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
