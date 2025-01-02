package taskUseCase

import (
	"context"
	taskDto "task-tracker/boundary/dto/task"
	repositoryInterface "task-tracker/boundary/repository"
	userObject "task-tracker/common/domainObject/shortUser"
	descriptionPrimitive "task-tracker/common/domainPrimitive/description"
	idPrimitive "task-tracker/common/domainPrimitive/id"
	titlePrimitive "task-tracker/common/domainPrimitive/title"
	taskEntity "task-tracker/domain/entity/task"
	assessmentPrimitive "task-tracker/domain/entity/task/assessment"
	taskComment "task-tracker/domain/entity/task/comment"
	taskTimeCosts "task-tracker/domain/entity/task/cost"
	taskPriority "task-tracker/domain/entity/task/spec/priority"
	taskStatus "task-tracker/domain/entity/task/spec/status"
	taskTag "task-tracker/domain/entity/task/spec/tag"
	userUseCase "task-tracker/domain/usecase/user"
	"task-tracker/infrastructure/errors"
	commonTime "task-tracker/infrastructure/tools/time"
	"time"
)

type TaskUseCase struct {
	taskRepo repositoryInterface.TaskRepository

	// TODO уедет в отдельый микросервис с общением через кафку
	userUseCase *userUseCase.UserUseCase
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

	// TODO переписать
	user, err := u.userUseCase.GetUserById(ctx, taskCreateDto.CreatorId)
	if err != nil {
		return nil, err
	}
	creator, err := userObject.NewShortUser(user.ID(), user.Email(), user.Profile())
	if err != nil {
		return nil, err
	}

	// TODO вынести в доменный примитив
	deadline, err := commonTime.Parse(time.RFC3339Nano, taskCreateDto.DeadLine)
	if err != nil {
		return nil, errors.NewError("SYS", "Deadline должен быть указан в формате RFC3339Nano")
	}
	if deadline.Before(commonTime.Now()) {
		return nil, errors.NewError("SYS", "Deadline не может быть меньше текущего времени")
	}

	assessment, err := assessmentPrimitive.AssessmentFrom(taskCreateDto.Assessment)
	if err != nil {
		return nil, err
	}

	timeCosts := taskTimeCosts.NewTimeCosts()

	comments := taskComment.NewComments()

	task, err := taskEntity.NewBuilder().
		Title(&title).
		Description(&description).
		Priority(priority).
		Tags(tags).
		Creator(creator).
		Deadline(deadline).
		Assessment(assessment).
		TimeCosts(timeCosts).
		Comments(comments).
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
	taskId, err := idPrimitive.EntityIdFrom(dto.Id)
	if err != nil {
		return nil, err
	}

	title, err := titlePrimitive.TitleFrom(dto.Title)
	if err != nil {
		return nil, err
	}

	description, err := descriptionPrimitive.DescriptionFrom(dto.Description)
	if err != nil {
		return nil, err
	}

	priority, err := taskPriority.Priorities.Of(dto.Priority)
	if err != nil {
		return nil, err
	}

	status, err := taskStatus.Statuses.Of(dto.Status)
	if err != nil {
		return nil, err
	}

	tags, err := taskTag.TagsFrom(dto.Tags)
	if err != nil {
		return nil, err
	}

	// TODO вынести в доменный примитив
	deadline, err := commonTime.Parse(time.RFC3339Nano, dto.DeadLine)
	if err != nil {
		return nil, errors.NewError("SYS", "Deadline должен быть указан в формате RFC3339Nano")
	}
	if deadline.Before(commonTime.Now()) {
		return nil, errors.NewError("SYS", "Deadline не может быть меньше текущего времени")
	}

	// TODO переписать
	user, err := u.userUseCase.GetUserById(ctx, dto.CreatorId)
	if err != nil {
		return nil, err
	}
	creator, err := userObject.NewShortUser(user.ID(), user.Email(), user.Profile())
	if err != nil {
		return nil, err
	}

	assessment, err := assessmentPrimitive.AssessmentFrom(dto.Assessment)
	if err != nil {
		return nil, err
	}

	task, err := taskEntity.NewBuilder().
		Id(&taskId).
		Title(&title).
		Description(&description).
		Priority(priority).
		Status(status).
		Tags(tags).
		Creator(creator).
		Deadline(deadline).
		Assessment(assessment).
		Build()
	if err != nil {
		return nil, err
	}

	updatedTask, err := u.taskRepo.Update(ctx, task)
	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func (u TaskUseCase) TakeOnTask(ctx context.Context, dto *taskDto.TaskDto) (*taskEntity.Task, error) {
	taskId, err := idPrimitive.EntityIdFrom(dto.Id)
	if err != nil {
		return nil, err
	}

	// TODO переписать
	user, err := u.userUseCase.GetUserById(ctx, dto.PerformerId)
	if err != nil {
		return nil, err
	}
	performer, err := userObject.NewShortUser(user.ID(), user.Email(), user.Profile())
	if err != nil {
		return nil, err
	}

	updatedTask, err := u.taskRepo.UpdatePerformerAndStatus(ctx, &taskId, performer, taskStatus.Statuses.InProgress())
	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func (u TaskUseCase) AddPerformer(ctx context.Context, dto *taskDto.TaskDto) (*taskEntity.Task, error) {
	taskId, err := idPrimitive.EntityIdFrom(dto.Id)
	if err != nil {
		return nil, err
	}

	// TODO переписать
	user, err := u.userUseCase.GetUserById(ctx, dto.PerformerId)
	if err != nil {
		return nil, err
	}
	performer, err := userObject.NewShortUser(user.ID(), user.Email(), user.Profile())
	if err != nil {
		return nil, err
	}

	updatedTask, err := u.taskRepo.UpdatePerformer(ctx, &taskId, performer)
	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func (u TaskUseCase) AddTimeCosts(ctx context.Context, dto *taskDto.TaskDto) (*taskEntity.Task, error) {
	taskId, err := idPrimitive.EntityIdFrom(dto.Id)
	if err != nil {
		return nil, err
	}

	// TODO переписать
	user, err := u.userUseCase.GetUserById(ctx, dto.TimeCosts.UserId)
	if err != nil {
		return nil, err
	}
	worker, err := userObject.NewShortUser(user.ID(), user.Email(), user.Profile())
	if err != nil {
		return nil, err
	}

	timeCost, err := taskTimeCosts.AddTimeCost(worker, commonTime.Now(), dto.TimeCosts.Minutes)
	if err != nil {
		return nil, err
	}

	updatedTask, err := u.taskRepo.UpdateTimeCosts(ctx, &taskId, timeCost)
	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func (u TaskUseCase) AddComment(ctx context.Context, dto *taskDto.TaskDto) (*taskEntity.Task, error) {
	taskId, err := idPrimitive.EntityIdFrom(dto.Id)
	if err != nil {
		return nil, err
	}

	// TODO переписать
	user, err := u.userUseCase.GetUserById(ctx, dto.Comments.UserId)
	if err != nil {
		return nil, err
	}
	author, err := userObject.NewShortUser(user.ID(), user.Email(), user.Profile())
	if err != nil {
		return nil, err
	}

	comment, err := taskComment.AddComment(author, commonTime.Now(), dto.Comments.Text)
	if err != nil {
		return nil, err
	}

	updatedTask, err := u.taskRepo.UpdateComments(ctx, &taskId, comment)
	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func (u TaskUseCase) DeleteTask(ctx context.Context, id string) error {
	taskId := idPrimitive.EntityId(id)
	if err := u.taskRepo.DeleteById(ctx, &taskId); err != nil {
		return err
	}
	return nil
}
