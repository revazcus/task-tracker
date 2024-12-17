package restServerController

import (
	"context"
	"encoding/json"
	"net/http"
	"task-tracker/infrastructure/errors"
	"task-tracker/infrastructure/security/jwtService"
)

type BaseController struct {
}

func NewBaseController() *BaseController {
	return &BaseController{}
}

// FillReqModel преобразует данные из входящего запроса в json и далее в модель UserRequest
// TODO спорный момент с requestData в виде interface{}, возможно лучше использовать T any или иной подход
func (bc *BaseController) FillReqModel(r *http.Request, requestData interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(requestData)
}

func (bc *BaseController) GetStrParamFromCtx(ctx context.Context, key jwtService.ContextKey) (string, error) {
	value := ctx.Value(key)
	if value == nil {
		return "", errors.NewError("SYS", "Parameter not found by context key")
	}
	valueStr, ok := value.(string)
	if !ok {
		return "", errors.NewError("SYS", "Couldn't parse ctx params")
	}
	return valueStr, nil
}
