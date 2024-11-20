package serializer

import (
	"task-tracker/adapters/controllers/rest/responses"
	"task-tracker/boundary/dto"
)

func SerializeUser(user *dto.UserDto) (*responses.UserResponse, error) {

	return &responses.UserResponse{
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
