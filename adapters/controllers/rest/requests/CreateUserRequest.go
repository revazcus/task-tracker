package userRestRequest

import "task-tracker/boundary/dto"

// CreateUserRequest структура входящего запроса по стандарту https://jsonapi.org/
type CreateUserRequest struct {
	Data struct {
		Attributes struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
		} `json:"attributes"`
	} `json:"data"`
}

// CreateUserDto маппинг данных входящего запроса в dto
func (r *CreateUserRequest) CreateUserDto() *dto.CreateUserDto {
	return &dto.CreateUserDto{
		Username: r.Data.Attributes.Username,
		Email:    r.Data.Attributes.Email,
		Password: r.Data.Attributes.Password,
	}
}
