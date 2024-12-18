package request

import (
	"bytes"
	"encoding/json"
	userDto "task-tracker/boundary/dto/user"
)

type CreateUserRequest struct {
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

// FillFromBytes создаёт JSON-декодер, который читает данные из преобразованного массива байт в поток (Reader) и распаковывает JSON-данные в структуру CreateUserRequest
func (r *CreateUserRequest) FillFromBytes(jsonBytes []byte) error {
	return json.NewDecoder(bytes.NewReader(jsonBytes)).Decode(r)
}

func (r *CreateUserRequest) CreateUserDto() *userDto.UserDto {
	return &userDto.UserDto{
		Id:       r.Data.Id,
		Email:    r.Data.Attributes.Email,
		Username: r.Data.Attributes.Username,
		Password: r.Data.Attributes.Password,
	}
}
