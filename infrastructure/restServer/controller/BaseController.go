package restServerController

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BaseController struct {
}

func NewBaseController() *BaseController {
	return &BaseController{}
}

func (bc BaseController) FillReqModel(r *http.Request) {
	requestBody := json.NewDecoder(r.Body)
	fmt.Println(requestBody)
}
