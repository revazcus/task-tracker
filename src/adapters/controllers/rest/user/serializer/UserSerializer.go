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
			"firstName": user.Profile().FirstName(),
			"lastName":  user.Profile().LastName(),
			"email":     user.Email(),
			"username":  user.Username(),
			"role":      user.Role(),
			"createdAt": user.CreatedAt(),
		},
		Relationships: jsonApiModel.JsonApiObjectRelationships{},
	}
	return response
}
