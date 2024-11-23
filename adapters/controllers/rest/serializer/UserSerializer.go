package serializer

import (
	restResponse "task-tracker/adapters/controllers/rest/response"
	"task-tracker/boundary/dto"
)

func SerializeUser(user *dto.UserDto) (*restResponse.UserResponse, error) {

	return &restResponse.UserResponse{
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
