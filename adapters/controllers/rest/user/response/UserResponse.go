package response

// UserResponse структура ответа по стандарту https://jsonapi.org/
type UserResponse struct {
	Data struct {
		Id         int `json:"id"`
		Attributes struct {
			Username string `json:"username"`
			Email    string `json:"email"`
		} `json:"attributes"`
	} `json:"data"`
}
