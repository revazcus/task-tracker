package restResponse

import (
	"encoding/json"
	"fmt"
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
