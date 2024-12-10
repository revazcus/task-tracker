package serializer

import (
	"task-tracker/adapters/controllers/rest/user/response"
	"task-tracker/boundary/dto"
)

func SerializeUser(user *dto.UserDto) (*response.UserResponse, error) {
	return &response.UserResponse{
		Data: struct {
			Id         string `json:"id"`
			Attributes struct {
				Username string `json:"username"`
				Email    string `json:"email"`
				Token    string `json:"token"`
			} `json:"attributes"`
		}{
			Id: user.Id,
			Attributes: struct {
				Username string `json:"username"`
				Email    string `json:"email"`
				Token    string `json:"token"`
			}{
				Username: user.Username,
				Email:    user.Email,
				Token:    user.Token,
			},
		},
	}, nil
}
