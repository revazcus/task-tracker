package userRepo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"strings"
	userRepoModel "task-tracker/adapters/repository/user/repoModel/user"
	emailPrimitive "task-tracker/domain/domainPrimitive/email"
	idPrimitive "task-tracker/domain/domainPrimitive/id"
	passwordPrimitive "task-tracker/domain/domainPrimitive/password"
	usernamePrimitive "task-tracker/domain/domainPrimitive/username"
	userEntity "task-tracker/domain/entity/user"
	"task-tracker/infrastructure/errors"
	loggerInterface "task-tracker/infrastructure/logger/interface"
	mongoInterface "task-tracker/infrastructure/mongo/interface"
	mongoModel "task-tracker/infrastructure/mongo/model"
)

const (
	indexUserId  = "uniqUserId"
	indexUserKey = "user_id"

	indexEmail    = "uniqEmail"
	indexEmailKey = "email"

	indexUsername    = "uniqUsername"
	indexUsernameKey = "username"
)

type UserRepo struct {
	collection string
	mongoRepo  mongoInterface.MongoRepository
	logger     loggerInterface.Logger
}

func (r *UserRepo) Init(ctx context.Context) error {
	userIdIndex := &mongoModel.DBIndex{
		Collection: r.collection,
		Name:       indexUserId,
		Keys:       []string{indexUserKey},
		Type:       mongoModel.DBIndexAsc,
		Uniq:       true,
	}

	if err := r.mongoRepo.TryCreateIndex(ctx, userIdIndex); err != nil {
		return err
	}

	emailIndex := &mongoModel.DBIndex{
		Collection: r.collection,
		Name:       indexEmail,
		Keys:       []string{indexEmailKey},
		Type:       mongoModel.DBIndexAsc,
		Uniq:       true,
	}

	if err := r.mongoRepo.TryCreateIndex(ctx, emailIndex); err != nil {
		return err
	}

	usernameIndex := &mongoModel.DBIndex{
		Collection: r.collection,
		Name:       indexUsername,
		Keys:       []string{indexUsernameKey},
		Type:       mongoModel.DBIndexAsc,
		Uniq:       true,
	}

	if err := r.mongoRepo.TryCreateIndex(ctx, usernameIndex); err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) Create(ctx context.Context, user *userEntity.User) error {
	userModel := userRepoModel.UserToRepoModel(user)
	if err := r.mongoRepo.InsertOne(ctx, r.collection, userModel); err != nil {
		duplicateErr, isDuplicateErr := err.(*errors.Error)
		msg := duplicateErr.Message()
		if isDuplicateErr {
			if strings.Contains(msg, indexEmail) {
				return ErrEmailAlreadyExist
			}
			if strings.Contains(msg, indexUsername) {
				return ErrUsernameAlreadyExist
			}
		}
		return err
	}
	return nil
}

func (r *UserRepo) GetById(ctx context.Context, userId *idPrimitive.EntityId) (*userEntity.User, error) {
	find := bson.D{{"user_id", userId.String()}}
	var userModel *userRepoModel.UserRepoModel

	if err := r.mongoRepo.FindOne(ctx, r.collection, find, &userModel); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	user, err := userModel.GetEntity()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepo) GetByUsername(ctx context.Context, username *usernamePrimitive.Username) (*userEntity.User, error) {
	find := bson.D{{"username", username.String()}}
	var userModel *userRepoModel.UserRepoModel

	if err := r.mongoRepo.FindOne(ctx, r.collection, find, &userModel); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrInvalidUsernameOrPassword
		}
		return nil, err
	}

	user, err := userModel.GetEntity()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepo) Update(ctx context.Context, user *userEntity.User) (*userEntity.User, error) {
	find := bson.D{{"user_id", user.ID().String()}}
	updatedUserModel := userRepoModel.UserToRepoModel(user)

	if err := r.mongoRepo.UpdateOne(ctx, r.collection, find, updatedUserModel); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (r *UserRepo) UpdateEmail(ctx context.Context, userId *idPrimitive.EntityId, email *emailPrimitive.Email) (*userEntity.User, error) {
	change := bson.D{
		{"$set", // указываем, что обновляем по ключам
			bson.M{"email": email.String()},
		},
	}
	return r.updateUser(ctx, userId, change)
}

func (r *UserRepo) UpdatePassword(ctx context.Context, userId *idPrimitive.EntityId, password *passwordPrimitive.Password) (*userEntity.User, error) {
	change := bson.D{
		{"$set",
			bson.M{"password": password.String()},
		},
	}
	return r.updateUser(ctx, userId, change)
}

func (r *UserRepo) DeleteById(ctx context.Context, userId *idPrimitive.EntityId) error {
	find := bson.D{{"user_id", userId.String()}}

	if err := r.mongoRepo.DeleteOne(ctx, r.collection, find); err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) updateUser(ctx context.Context, userId *idPrimitive.EntityId, change bson.D) (*userEntity.User, error) {
	find := bson.D{{"user_id", userId.String()}}

	var userModel *userRepoModel.UserRepoModel

	if err := r.mongoRepo.FindOneAndUpdate(ctx, r.collection, &userModel, find, change, options.FindOneAndUpdate().SetReturnDocument(options.After)); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	updatedUser, err := userModel.GetEntity()
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
