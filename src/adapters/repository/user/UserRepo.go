package userRepo

import (
	"context"
	"task-tracker/adapters/repository/user/repoModel"
	userEntity "task-tracker/domain/entity/user"
	loggerInterface "task-tracker/infrastructure/logger/interface"
	mongoInterface "task-tracker/infrastructure/mongo/interface"
)

type UserRepo struct {
	table     string
	mongoRepo mongoInterface.MongoRepository
	logger    loggerInterface.Logger
}

func (r *UserRepo) Init() error {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepo) Create(ctx context.Context, user *userEntity.User) (string, error) {
	userModel := repoModel.NewUserRepoModel(user)
	createdUserId, err := r.mongoRepo.Create(ctx, r.table, userModel)
	if err != nil {
		return "", err
	}
	return createdUserId, nil
}

func (r *UserRepo) Update(ctx context.Context, user *userEntity.User) error {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepo) GetById(ctx context.Context, userId string) (*userEntity.User, error) {
	var userModel *repoModel.UserRepoModel
	err := r.mongoRepo.GetByID(ctx, r.table, userId, &userModel)
	if err != nil {
		return nil, err
	}

	user, err := userModel.GetEntity()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepo) GetAll(ctx context.Context) ([]*userEntity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepo) Delete(ctx context.Context, userId string) error {
	//TODO implement me
	panic("implement me")
}
