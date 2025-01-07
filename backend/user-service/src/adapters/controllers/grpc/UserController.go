package grpc

import (
	"context"
	"github.com/revazcus/task-tracker/backend/common/gateways/user/gateways/user/protoModel"
	"github.com/revazcus/task-tracker/backend/user-service/boundary/domain/usecase"
	userEntity "github.com/revazcus/task-tracker/backend/user-service/domain/entity"
)

type UserController struct {
	userUseCase                               usecase.UserUseCaseInterface
	protoModel.UnimplementedUserServiceServer // Позволяет реализовать не весь интерфейс UserServiceServer
}

// NewUserController TODO переписать на билдер
func NewUserController(userUseCase usecase.UserUseCaseInterface) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

func (u UserController) GetUser(ctx context.Context, request *protoModel.UserRequest) (*protoModel.User, error) {
	user, err := u.userUseCase.GetUserById(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return createProtoUserFromEntity(user), nil
}

func createProtoUserFromEntity(user *userEntity.User) *protoModel.User {
	return &protoModel.User{
		Id:        user.ID().String(),
		Email:     user.Email().String(),
		FirstName: user.Profile().FirstName(),
		LastName:  user.Profile().LastName(),
	}
}
