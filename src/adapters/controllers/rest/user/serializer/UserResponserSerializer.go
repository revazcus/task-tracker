package serializer

import (
	userDto "task-tracker/boundary/dto/user"
	jsonApiModel "task-tracker/infrastructure/jsonapi/model"
)

const (
	ResponseUser = "user"
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
			"username": user.Username().String(),
			"email":    user.Email().String(),
			"token":    responseDto.Token,
		},
		Relationships: jsonApiModel.JsonApiObjectRelationships{},
	}
	return response
}
