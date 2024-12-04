package serializer

import (
	"task-tracker/adapters/controllers/rest/user/response"
	"task-tracker/boundary/dto"
)

func SerializeUser(user *dto.UserDto) (*response.UserResponse, error) {
	return &response.UserResponse{
		Data: struct {
			Id         int `json:"id"`
			Attributes struct {
				Username string `json:"username"`
				Email    string `json:"email"`
			} `json:"attributes"`
		}{
			Id: user.Id,
			Attributes: struct {
				Username string `json:"username"`
				Email    string `json:"email"`
			}{
				Username: user.Username,
				Email:    user.Email,
			},
		},
	}, nil
}
