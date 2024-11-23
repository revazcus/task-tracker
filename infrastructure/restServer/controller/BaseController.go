package restServerController

import (
	"encoding/json"
	"net/http"
)

type BaseController struct {
}

func NewBaseController() *BaseController {
	return &BaseController{}
}

// FillReqModel преобразует данные из входящего запроса в json и далее в модель UserRequest
// TODO спорный момент с requestData в виде interface{}, возможно лучше использовать T any или иной подход
func (bc BaseController) FillReqModel(r *http.Request, requestData interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(requestData)
}
