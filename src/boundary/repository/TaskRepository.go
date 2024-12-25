package repositoryInterface

import (
	"context"
	idPrimitive "task-tracker/common/domainPrimitive/id"
	taskEntity "task-tracker/domain/entity/task"
)

type TaskRepository interface {
	Init(ctx context.Context) error

	Create(ctx context.Context, task *taskEntity.Task) error

	GetAll(ctx context.Context) ([]*taskEntity.Task, error)
	GetById(ctx context.Context, taskId *idPrimitive.EntityId) (*taskEntity.Task, error)

	Update(ctx context.Context, task *taskEntity.Task) (*taskEntity.Task, error)

	DeleteById(ctx context.Context, taskId *idPrimitive.EntityId) error
}
