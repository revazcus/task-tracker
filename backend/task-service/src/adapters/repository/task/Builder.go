package taskRepo

import (
	"github.com/revazcus/task-tracker/backend/infrastructure/errors"
	loggerInterface "github.com/revazcus/task-tracker/backend/infrastructure/logger/interface"
	mongoInterface "github.com/revazcus/task-tracker/backend/infrastructure/mongo/interface"
)

type Builder struct {
	taskRepo *TaskRepo
	errors   *errors.Errors
}

func NewBuilder() *Builder {
	return &Builder{
		taskRepo: &TaskRepo{},
		errors:   errors.NewErrors(),
	}
}

func (b *Builder) Collection(table string) *Builder {
	b.taskRepo.collection = table
	return b
}

func (b *Builder) MongoRepo(mongoRepo mongoInterface.MongoRepository) *Builder {
	b.taskRepo.mongoRepo = mongoRepo
	return b
}

func (b *Builder) Logger(logger loggerInterface.Logger) *Builder {
	b.taskRepo.logger = logger
	return b
}

func (b *Builder) Build() (*TaskRepo, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}
	return b.taskRepo, nil
}

func (b *Builder) checkRequiredFields() {
	if b.taskRepo.collection == "" {
		b.errors.AddError(errors.NewError("SYS", "TaskRepoBuilder: Collection is required"))
	}
	if b.taskRepo.mongoRepo == nil {
		b.errors.AddError(errors.NewError("SYS", "TaskRepoBuilder: MongoRepository is required"))
	}
	if b.taskRepo.logger == nil {
		b.errors.AddError(errors.NewError("SYS", "TaskRepoBuilder: Logger is required"))
	}
}
