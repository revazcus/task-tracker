package taskUseCase

import (
	"context"
	taskDto "task-tracker/boundary/dto/task"
	repositoryInterface "task-tracker/boundary/repository"
	commentPrimitive "task-tracker/common/domainPrimitive/comment"
	descriptionPrimitive "task-tracker/common/domainPrimitive/description"
	idPrimitive "task-tracker/common/domainPrimitive/id"
	titlePrimitive "task-tracker/common/domainPrimitive/title"
	taskEntity "task-tracker/domain/entity/task"
	assessmentPrimitive "task-tracker/domain/entity/task/assessment"
	taskTimeCosts "task-tracker/domain/entity/task/cost"
	taskPriority "task-tracker/domain/entity/task/spec/priority"
	taskTag "task-tracker/domain/entity/task/spec/tag"
	commonTime "task-tracker/infrastructure/tools/time"
	"time"
)

type TaskUseCase struct {
	taskRepo repositoryInterface.TaskRepository
}

func (u TaskUseCase) CreateTask(ctx context.Context, taskCreateDto *taskDto.TaskDto) (*taskEntity.Task, error) {
	title, err := titlePrimitive.TitleFrom(taskCreateDto.Title)
	if err != nil {
		return nil, err
	}

	description, err := descriptionPrimitive.DescriptionFrom(taskCreateDto.Description)
	if err != nil {
		return nil, err
	}

	priority, err := taskPriority.Priorities.Of(taskCreateDto.Priority)
	if err != nil {
		return nil, err
	}

	tags, err := taskTag.TagsFrom(taskCreateDto.Tags)
	if err != nil {
		return nil, err
	}

	deadline, err := commonTime.Parse(time.RFC3339Nano, taskCreateDto.DeadLine)
	if err != nil {
		return nil, err
	}

	assessment, err := assessmentPrimitive.AssessmentFrom(taskCreateDto.Assessment)
	if err != nil {
		return nil, err
	}

	timeCosts := taskTimeCosts.NewTimeCosts()
	if err := timeCosts.AddEntry(taskCreateDto.TimeCosts, taskCreateDto.CreatorId); err != nil {
		return nil, err
	}

	comments, err := commentPrimitive.CommentsFrom(taskCreateDto.Comments)
	if err != nil {
		return nil, err
	}

	task, err := taskEntity.NewBuilder().
		Title(&title).
		Description(&description).
		Priority(priority).
		Tags(tags).
		CreatorId(taskCreateDto.CreatorId).
		PerformerId(taskCreateDto.PerformerId).
		Deadline(deadline).
		Comments(comments).
		Assessment(assessment).
		Build()
	if err != nil {
		return nil, err
	}

	if err := u.taskRepo.Create(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}

func (u TaskUseCase) GetAllTasks(ctx context.Context) ([]*taskEntity.Task, error) {
	foundTasks, err := u.taskRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return foundTasks, nil
}

func (u TaskUseCase) GetTaskById(ctx context.Context, id string) (*taskEntity.Task, error) {
	taskId := idPrimitive.EntityId(id)
	foundTask, err := u.taskRepo.GetById(ctx, &taskId)
	if err != nil {
		return nil, err
	}
	return foundTask, nil
}

func (u TaskUseCase) UpdateTask(ctx context.Context, dto *taskDto.TaskDto) (*taskEntity.Task, error) {
	panic("implement me")
}

func (u TaskUseCase) DeleteTask(ctx context.Context, id string) error {
	taskId := idPrimitive.EntityId(id)
	if err := u.taskRepo.DeleteById(ctx, &taskId); err != nil {
		return err
	}
	return nil
}
