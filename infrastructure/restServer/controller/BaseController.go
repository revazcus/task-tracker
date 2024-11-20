package restServerController

import (
	"encoding/json"
	"net/http"
	userRestRequest "task-tracker/adapters/controllers/rest/requests"
)

type BaseController struct {
}

func NewBaseController() *BaseController {
	return &BaseController{}
}

// FillReqModel преобразует данные из входящего запроса в json и далее в модель UserRequest
func (bc BaseController) FillReqModel(r *http.Request, requestData *userRestRequest.UserRequest) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(requestData)
}
