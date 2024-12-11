package userRepo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	userRepoModel "task-tracker/adapters/repository/user/repoModel/user"
	emailPrimitive "task-tracker/domain/domainPrimitive/email"
	idPrimitive "task-tracker/domain/domainPrimitive/id"
	passwordPrimitive "task-tracker/domain/domainPrimitive/password"
	userEntity "task-tracker/domain/entity/user"
	loggerInterface "task-tracker/infrastructure/logger/interface"
	mongoInterface "task-tracker/infrastructure/mongo/interface"
)

type UserRepo struct {
	collection string
	mongoRepo  mongoInterface.MongoRepository
	logger     loggerInterface.Logger
}

func (r *UserRepo) Create(ctx context.Context, user *userEntity.User) error {
	userModel := userRepoModel.UserToRepoModel(user)
	if err := r.mongoRepo.InsertOne(ctx, r.collection, userModel); err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) Update(ctx context.Context, user *userEntity.User) error {
	find := bson.D{{"user_id", user.ID().String()}}
	updatedUserModel := userRepoModel.UserToRepoModel(user)

	if err := r.mongoRepo.UpdateOne(ctx, r.collection, find, updatedUserModel); err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) UpdateEmail(ctx context.Context, userId *idPrimitive.EntityId, email *emailPrimitive.Email) error {
	change := bson.D{
		{"$set", // указываем, что обновляем по ключам
			bson.M{"email": email.String()},
		},
	}
	return r.updateUser(ctx, userId, change)
}

func (r *UserRepo) UpdatePassword(ctx context.Context, userId *idPrimitive.EntityId, password *passwordPrimitive.Password) error {
	change := bson.D{
		{"$set",
			bson.M{"password": password.String()},
		},
	}
	return r.updateUser(ctx, userId, change)
}

func (r *UserRepo) GetById(ctx context.Context, userId *idPrimitive.EntityId) (*userEntity.User, error) {
	find := bson.D{{"user_id", userId.String()}}
	var userModel *userRepoModel.UserRepoModel

	if err := r.mongoRepo.FindOne(ctx, r.collection, find, &userModel); err != nil {
		return nil, err
	}

	user, err := userModel.GetEntity()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepo) DeleteById(ctx context.Context, userId *idPrimitive.EntityId) error {
	find := bson.D{{"user_id", userId.String()}}

	if err := r.mongoRepo.DeleteOne(ctx, r.collection, find); err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) updateUser(ctx context.Context, userId *idPrimitive.EntityId, change bson.D) error {
	find := bson.D{{"user_id", userId.String()}}

	if err := r.mongoRepo.UpdateOne(ctx, r.collection, find, change); err != nil {
		return err
	}

	return nil
}
