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
			Agreement bool   `json:"agreement"`
		} `json:"attributes"`
	} `json:"data"`
}

func (r *CreateUserRequest) FillFromBytes(jsonBytes []byte) error {
	return json.NewDecoder(bytes.NewReader(jsonBytes)).Decode(r)
}

func (r *CreateUserRequest) CreateUserDto() *userDto.UserDto {
	return &userDto.UserDto{
		Id:        r.Data.Id,
		FirstName: r.Data.Attributes.FirstName,
		LastName:  r.Data.Attributes.LastName,
		Email:     r.Data.Attributes.Email,
		Username:  r.Data.Attributes.Username,
		Password:  r.Data.Attributes.Password,
		Agreement: r.Data.Attributes.Agreement,
	}
}
