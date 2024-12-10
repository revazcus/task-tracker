package response

// UserResponse структура ответа по стандарту https://jsonapi.org/
type UserResponse struct {
	Data struct {
		Id         string `json:"id"`
		Attributes struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Token    string `json:"token"`
		} `json:"attributes"`
	} `json:"data"`
}
