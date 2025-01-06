package repositoryInterface

import (
	"context"
	userObject "github.com/revazcus/task-tracker/backend/common/domainObject/shortUser"
	idPrimitive "github.com/revazcus/task-tracker/backend/common/domainPrimitive/id"
	taskEntity "github.com/revazcus/task-tracker/backend/task-service/domain/entity/task"
	taskComment "github.com/revazcus/task-tracker/backend/task-service/domain/entity/task/comment"
	taskTimeCosts "github.com/revazcus/task-tracker/backend/task-service/domain/entity/task/cost"
	taskStatus "github.com/revazcus/task-tracker/backend/task-service/domain/entity/task/spec/status"
)

type TaskRepository interface {
	Init(ctx context.Context) error

	Create(ctx context.Context, task *taskEntity.Task) error

	GetAll(ctx context.Context) ([]*taskEntity.Task, error)
	GetById(ctx context.Context, taskId *idPrimitive.EntityId) (*taskEntity.Task, error)

	Update(ctx context.Context, task *taskEntity.Task) (*taskEntity.Task, error)
	UpdatePerformer(ctx context.Context, taskId *idPrimitive.EntityId, performer *userObject.ShortUser) (*taskEntity.Task, error)
	UpdatePerformerAndStatus(ctx context.Context, taskId *idPrimitive.EntityId, performer *userObject.ShortUser, status taskStatus.Status) (*taskEntity.Task, error)
	UpdateTimeCosts(ctx context.Context, taskId *idPrimitive.EntityId, timeCost *taskTimeCosts.TimeInvestment) (*taskEntity.Task, error)
	UpdateComments(ctx context.Context, taskId *idPrimitive.EntityId, comment *taskComment.Comment) (*taskEntity.Task, error)

	DeleteById(ctx context.Context, taskId *idPrimitive.EntityId) error
}
