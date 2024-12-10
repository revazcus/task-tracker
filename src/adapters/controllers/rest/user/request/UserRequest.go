package request

import "task-tracker/boundary/dto"

// UserRequest структура входящего запроса по стандарту https://jsonapi.org/
type UserRequest struct {
	Data struct {
		Id         string `json:"id"`
		Attributes struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
		} `json:"attributes"`
	} `json:"data"`
}

// CreateUserDto маппит данные из UserRequest
func (r *UserRequest) CreateUserDto() *dto.UserDto {
	return &dto.UserDto{
		Id:       r.Data.Id,
		Username: r.Data.Attributes.Username,
		Email:    r.Data.Attributes.Email,
		Password: r.Data.Attributes.Password,
	}
}
