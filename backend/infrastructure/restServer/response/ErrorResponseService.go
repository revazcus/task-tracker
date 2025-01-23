package restResponse

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"infrastructure/errors"
	loggerInterface "infrastructure/logger/interface"
	restServerInterface "infrastructure/restServer/interface"
	"net/http"
)

type ErrorResponseService struct {
	errorResolver restServerInterface.ErrorResolver
	logger        loggerInterface.Logger
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
		s.logger.Error(ctx, fmt.Errorf("%s: %v", "JSON marshal failed", err))
		return
	}

	w.Header().Set(HeaderContentType, AppJsonContentType)
	w.Header().Set(HeaderXContentTypeOptions, NoSniffXContentTypeOptions)
	w.WriteHeader(code)

	prettyJson, err := s.prettyJson(body)
	if err != nil {
		s.logger.Error(ctx, fmt.Errorf("%s: %v", "Prettify response failed", err))
	}
	_, err = w.Write(prettyJson)
	if err != nil {
		s.logger.Error(ctx, fmt.Errorf("%s: %v", "ResponseWriter Error", err))
		return
	}
}

func (s *ErrorResponseService) prettyJson(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}
