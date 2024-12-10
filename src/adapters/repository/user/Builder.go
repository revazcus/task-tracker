package userRepo

import (
	"task-tracker/infrastructure/errors"
	loggerInterface "task-tracker/infrastructure/logger/interface"
	mongoInterface "task-tracker/infrastructure/mongo/interface"
)

type Builder struct {
	userRepo *UserRepo
}

func NewBuilder() *Builder {
	return &Builder{
		userRepo: &UserRepo{},
	}
}

func (b *Builder) Table(table string) *Builder {
	b.userRepo.table = table
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
	err := b.checkRequiredFields()
	if err != nil {
		return nil, err
	}
	return b.userRepo, nil
}

func (b *Builder) checkRequiredFields() error {
	if b.userRepo.table == "" {
		return errors.ErrTableIsRequired
	}
	if b.userRepo.mongoRepo == nil {
		return errors.ErrRepositoryIsRequired
	}
	if b.userRepo.logger == nil {
		return errors.ErrLoggerIsRequired
	}
	return nil
}
