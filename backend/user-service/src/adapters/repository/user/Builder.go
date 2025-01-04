package userRepo

import (
	"github.com/revazcus/task-tracker/backend/infrastructure/errors"
	loggerInterface "github.com/revazcus/task-tracker/backend/infrastructure/logger/interface"
	mongoInterface "github.com/revazcus/task-tracker/backend/infrastructure/mongo/interface"
)

type Builder struct {
	userRepo *UserRepo
	errors   *errors.Errors
}

func NewBuilder() *Builder {
	return &Builder{
		userRepo: &UserRepo{},
		errors:   errors.NewErrors(),
	}
}

func (b *Builder) Collection(table string) *Builder {
	b.userRepo.collection = table
	return b
}

func (b *Builder) MongoRepo(mongoRepo mongoInterface.MongoRepository) *Builder {
	b.userRepo.mongoRepo = mongoRepo
	return b
}

func (b *Builder) Logger(logger loggerInterface.Logger) *Builder {
	b.userRepo.logger = logger
	return b
}

func (b *Builder) Build() (*UserRepo, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}
	return b.userRepo, nil
}

func (b *Builder) checkRequiredFields() {
	if b.userRepo.collection == "" {
		b.errors.AddError(errors.NewError("SYS", "UserRepoBuilder: Collection is required"))
	}
	if b.userRepo.mongoRepo == nil {
		b.errors.AddError(errors.NewError("SYS", "UserRepoBuilder: MongoRepository is required"))
	}
	if b.userRepo.logger == nil {
		b.errors.AddError(errors.NewError("SYS", "UserRepoBuilder: Logger is required"))
	}
}
