package user

import (
	"task-tracker/infrastructure/errors"
	loggerInterface "task-tracker/infrastructure/logger/interface"
)

type Builder struct {
	userRepo *UserRepo
}

func NewBuilder() *Builder {
	return &Builder{
		userRepo: &UserRepo{},
	}
}

// Logger инициализирует поле logger в структуре UserRepo
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
	if b.userRepo.logger == nil {
		return errors.ErrLoggerIsRequired
	}
	return nil
}
