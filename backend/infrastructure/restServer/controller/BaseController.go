package restServerController

import (
	"context"
	"infrastructure/errors"
	loggerInterface "infrastructure/logger/interface"
	restServerInterface "infrastructure/restServer/interface"
	"infrastructure/restServer/response"
	"infrastructure/security/jwtService"
	"io"
	"net/http"
)

type BaseController struct {
	responseService *response.ResponseService
	logger          loggerInterface.Logger
}

// NewBaseController TODO переписать на билдер
func NewBaseController(responseService *response.ResponseService, logger loggerInterface.Logger) (*BaseController, error) {
	if responseService == nil {
		return nil, errors.NewError("SYS", "ResponseService is required")
	}
	if logger == nil {
		return nil, errors.NewError("SYS", "Logger is required")
	}
	return &BaseController{responseService: responseService, logger: logger}, nil
}

func (bc *BaseController) FillReqModel(r *http.Request, reqModel restServerInterface.RequestModel) error {
	requestBody, err := bc.GetRequestBody(r)
	if err != nil {
		return err
	}

	err = reqModel.FillFromBytes(requestBody)
	if err != nil {
		return response.ErrUnmarshalRequest(err.Error())
	}
	return err
}

func (bc *BaseController) GetRequestBody(r *http.Request) ([]byte, error) {
	if r.Body == nil {
		return nil, response.ErrUnmarshalRequest("Request body is nil")
	}
	return io.ReadAll(r.Body)
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

func (bc *BaseController) Response(w http.ResponseWriter, r *http.Request, result []byte, responseCode int) {
	bc.responseService.Response(w, r, result, responseCode)
}

func (bc *BaseController) JsonResponse(w http.ResponseWriter, r *http.Request, result interface{}, responseCode int) {
	bc.responseService.JsonResponse(w, r, result, responseCode)
}

func (bc *BaseController) ErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	bc.responseService.ErrorResponse(w, r, err)
}
