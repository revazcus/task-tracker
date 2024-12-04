package teamRest

import (
	"net/http"
	"task-tracker/boundary/domain/usecase"
	"task-tracker/infrastructure/restServer/controller"
)

type TeamController struct {
	*restServerController.BaseController
	teamUseCase usecase.TeamUseCaseInterface
}

func NewTeamController(controller *restServerController.BaseController, teamUseCase usecase.TeamUseCaseInterface) *TeamController {
	return &TeamController{BaseController: controller, teamUseCase: teamUseCase}
}

func (c *TeamController) GetTeamById(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *TeamController) CreateTeam(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *TeamController) UpdateTeam(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *TeamController) DeleteTeamById(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
