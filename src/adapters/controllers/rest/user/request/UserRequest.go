package request

import "task-tracker/boundary/dto"

// UserRequest структура входящего запроса по стандарту https://jsonapi.org/
type UserRequest struct {
	Data struct {
		Id         string `json:"id"`
		Attributes struct {
			FirstName string `json:"firstName"`
			LastName  string `json:"lastName"`
			Email     string `json:"email"`
			Username  string `json:"username"`
			Password  string `json:"password"`
		} `json:"attributes"`
	} `json:"data"`
}

// CreateUserDto маппит данные из UserRequest
func (r *UserRequest) CreateUserDto() *dto.UserDto {
	return &dto.UserDto{
		Id:        r.Data.Id,
		FirstName: r.Data.Attributes.FirstName,
		LastName:  r.Data.Attributes.LastName,
		Email:     r.Data.Attributes.Email,
		Username:  r.Data.Attributes.Username,
		Password:  r.Data.Attributes.Password,
	}
}
