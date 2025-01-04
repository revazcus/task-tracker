package userSerializer

import (
	jsonApiModel "github.com/revazcus/task-tracker/backend/infrastructure/jsonapi/model"
	userDto "github.com/revazcus/task-tracker/backend/user-service/boundary/dto/user"
)

func SerializeUserResponse(responseDto *userDto.UserResponseDto) (jsonApiModel.JsonApiPayload, error) {
	responseBuilder := jsonApiModel.NewJsonApiPayloadBuilder()
	responseBuilder.AddData(CreateUserResponseObject(responseDto))
	return responseBuilder.Build()
}

func CreateUserResponseObject(responseDto *userDto.UserResponseDto) *jsonApiModel.JsonApiObject {
	user := responseDto.User
	response := &jsonApiModel.JsonApiObject{
		Id:   user.ID().String(),
		Type: ResponseUser,
		Attributes: map[string]interface{}{
			"firstName": user.Profile().FirstName(),
			"lastName":  user.Profile().LastName(),
			"username":  user.Username().String(),
			"email":     user.Email().String(),
			"token":     responseDto.Token,
			"role":      user.Role(),
			"createdAt": user.CreatedAt(),
		},
		Relationships: jsonApiModel.JsonApiObjectRelationships{},
	}
	return response
}
