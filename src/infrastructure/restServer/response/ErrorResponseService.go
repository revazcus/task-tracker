package response

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"task-tracker/infrastructure/errors"
	loggerInterface "task-tracker/infrastructure/logger/interface"
	restServerInterface "task-tracker/infrastructure/restServer/interface"
)

type ErrorResponseService struct {
	errorResolver restServerInterface.ErrorResolver
	logger        loggerInterface.Logger
}

// NewErrorResponseService TODO переписать на билдер
func NewErrorResponseService(errorResolver restServerInterface.ErrorResolver, logger loggerInterface.Logger) (*ErrorResponseService, error) {
	if errorResolver == nil {
		return nil, errors.NewError("SYS", "ErrorResolver is required")
	}
	if logger == nil {
		return nil, errors.NewError("SYS", "Logger is required")
	}
	return &ErrorResponseService{logger: logger, errorResolver: errorResolver}, nil
}

func (s *ErrorResponseService) ErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	errorResponse := s.CreateErrorResponse(err)
	s.writeErrorResponse(r.Context(), w, errorResponse, errorResponse.FirstHttpCode())
}

func (s *ErrorResponseService) ErrorsResponse(w http.ResponseWriter, r *http.Request, errors []error) {
	errorResponse := s.CreateErrorResponse(errors...)
	s.writeErrorResponse(r.Context(), w, errorResponse, errorResponse.FirstHttpCode())
}

func (s *ErrorResponseService) CreateErrorResponse(errs ...error) ErrorResponse {
	errorsData := make([]ErrorResponseData, 0, len(errs))
	for _, err := range errs {
		switch err := err.(type) {
		case *errors.Errors:
			for _, errItem := range err.ToArray() {
				errData := s.createErrorResponseData(errItem)
				errorsData = append(errorsData, errData)
			}
		default:
			errData := s.createErrorResponseData(err)
			errorsData = append(errorsData, errData)
		}
	}
	return NewErrorResponse(errorsData)
}

func (s *ErrorResponseService) createErrorResponseData(err error) ErrorResponseData {
	responseCode := s.errorResolver.GetHttpCode(err)
	errorCode := s.errorResolver.GetErrorCode(err)
	errText := s.errorResolver.GetErrorText(err)
	return ErrorResponseData{HttpCode: responseCode, ErrorCode: errorCode, Text: errText}
}

func (s *ErrorResponseService) writeErrorResponse(ctx context.Context, w http.ResponseWriter, response ErrorResponse, code int) {
	body, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.logger.LogError(err, "json.Marshal", "src/infrastructure/restServer/response/ErrorResponseService.go:57", "JSON marshal failed")
		return
	}

	w.Header().Set(HeaderContentType, AppJsonContentType)
	w.Header().Set(HeaderXContentTypeOptions, NoSniffXContentTypeOptions)
	w.WriteHeader(code)

	prettyJson, err := s.prettyJson(body)
	if err != nil {
		s.logger.LogError(err, "prettyJson", "src/infrastructure/restServer/response/ErrorResponseService.go:69", "Prettify response failed")
	}
	_, err = w.Write(prettyJson)
	if err != nil {
		s.logger.LogError(err, "Write", "src/infrastructure/restServer/response/ErrorResponseService.go:73", "ResponseWriter Error")
		return
	}
}

func (s *ErrorResponseService) prettyJson(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}
