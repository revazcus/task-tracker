package serializer

import (
	userEntity "task-tracker/domain/entity/user"
	jsonApiModel "task-tracker/infrastructure/jsonapi/model"
)

func SerializeUser(user *userEntity.User) (jsonApiModel.JsonApiPayload, error) {
	responseBuilder := jsonApiModel.NewJsonApiPayloadBuilder()
	responseBuilder.AddData(CreateUserObject(user))
	return responseBuilder.Build()
}

func CreateUserObject(user *userEntity.User) *jsonApiModel.JsonApiObject {
	response := &jsonApiModel.JsonApiObject{
		Id:   user.ID().String(),
		Type: ResponseUser,
		Attributes: map[string]interface{}{
			"username": user.Username(),
			"email":    user.Email(),
		},
		Relationships: jsonApiModel.JsonApiObjectRelationships{},
	}
	return response
}
