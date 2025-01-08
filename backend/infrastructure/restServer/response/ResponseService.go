package response

import (
	"encoding/json"
	"fmt"
	"infrastructure/errors"
	loggerInterface "infrastructure/logger/interface"
	"net/http"
)

const (
	HeaderContentType         = "Content-Type"
	HeaderXContentTypeOptions = "X-Content-Type-Options"

	AppJsonContentType         = "application/json; charset=utf-8"
	NoSniffXContentTypeOptions = "nosniff"
)

type ResponseService struct {
	errorResponseService *ErrorResponseService
	logger               loggerInterface.Logger
}

// NewResponseService TODO переписать на билдер
func NewResponseService(errorResponseService *ErrorResponseService, logger loggerInterface.Logger) (*ResponseService, error) {
	if logger == nil {
		return nil, errors.NewError("SYS", "Logger is required")
	}
	if errorResponseService == nil {
		return nil, errors.NewError("SYS", "ErrorResponseService is required")
	}
	return &ResponseService{errorResponseService: errorResponseService, logger: logger}, nil
}

func (s *ResponseService) JsonResponse(w http.ResponseWriter, r *http.Request, result interface{}, responseCode int) {
	body, err := s.marshalBody(result)
	if err != nil {
		s.logger.Error(r.Context(), fmt.Errorf("%s: %v", "Marshal json error", err))
		s.errorResponseService.ErrorResponse(w, r, err)
		return
	}
	s.Response(w, r, body, responseCode)
}

func (s *ResponseService) Response(w http.ResponseWriter, r *http.Request, result []byte, responseCode int) {
	w.Header().Set(HeaderContentType, AppJsonContentType)
	w.Header().Set(HeaderXContentTypeOptions, NoSniffXContentTypeOptions)
	w.WriteHeader(responseCode)
	if responseCode != http.StatusNoContent {
		if _, err := w.Write(result); err != nil {
			s.logger.Error(r.Context(), fmt.Errorf("%s: %v", "Response Writer Error", err))
			return
		}
	}
}

func (s *ResponseService) ErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	s.errorResponseService.ErrorResponse(w, r, err)
}

func (s *ResponseService) marshalBody(result interface{}) ([]byte, error) {
	if result == nil || result == "" {
		return []byte{}, nil
	}

	body, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return body, nil
}
