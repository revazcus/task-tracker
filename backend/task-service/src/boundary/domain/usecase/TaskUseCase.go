package usecase

import (
	"context"
	taskDto "github.com/revazcus/task-tracker/backend/task-service/boundary/dto/task"
	taskEntity "github.com/revazcus/task-tracker/backend/task-service/domain/entity/task"
)

type TaskUseCaseInterface interface {
	CreateTask(ctx context.Context, taskCreateDto *taskDto.TaskDto) (*taskEntity.Task, error)

	GetAllTasks(ctx context.Context) ([]*taskEntity.Task, error)
	GetTaskById(ctx context.Context, id string) (*taskEntity.Task, error)

	UpdateTask(ctx context.Context, dto *taskDto.TaskDto) (*taskEntity.Task, error)
	TakeOnTask(ctx context.Context, dto *taskDto.TaskDto) (*taskEntity.Task, error)
	AddPerformer(ctx context.Context, dto *taskDto.TaskDto) (*taskEntity.Task, error)
	AddTimeCosts(ctx context.Context, dto *taskDto.TaskDto) (*taskEntity.Task, error)
	AddComment(ctx context.Context, dto *taskDto.TaskDto) (*taskEntity.Task, error)

	DeleteTask(ctx context.Context, id string) error
}
