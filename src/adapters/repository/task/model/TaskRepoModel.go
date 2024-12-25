package taskRepoModel

import (
	descriptionPrimitive "task-tracker/common/domainPrimitive/description"
	idPrimitive "task-tracker/common/domainPrimitive/id"
	titlePrimitive "task-tracker/common/domainPrimitive/title"
	taskEntity "task-tracker/domain/entity/task"
	taskDuration "task-tracker/domain/entity/task/duration"
	taskPriority "task-tracker/domain/entity/task/spec/priority"
	taskStatus "task-tracker/domain/entity/task/spec/status"
	taskTag "task-tracker/domain/entity/task/spec/tag"
	commonTime "task-tracker/infrastructure/tools/time"
)

type TaskRepoModel struct {
	Id          string   `bson:"task_id"`
	Title       string   `bson:"title"`
	Description string   `bson:"description"`
	Status      string   `bson:"status"`
	Priority    string   `bson:"priority"`
	Tag         string   `bson:"tag"`
	CreatorId   string   `bson:"creatorId"`
	PerformerId string   `bson:"performerId"`
	CreateAt    int64    `bson:"create_at"`
	UpdateAt    int64    `bson:"update_at"`
	Deadline    int64    `bson:"deadline"`
	Comments    []string `bson:"comments"` // TODO подумать как сделать
	Estimation  string   `bson:"estimation"`
	SpentTime   string   `bson:"spentTime"`
}

func TaskToRepoModel(task *taskEntity.Task) *TaskRepoModel {
	return &TaskRepoModel{
		Id:          string(*task.ID()),
		Title:       string(*task.Title()),
		Description: string(*task.Description()),
		Status:      task.Status().String(),
		Priority:    task.Priority().String(),
		Tag:         task.Tag().String(),
		CreatorId:   task.CreatorId(),
		PerformerId: task.PerformerId(),
		CreateAt:    task.CreateAt().UnixNano(),
		UpdateAt:    task.UpdateAt().UnixNano(),
		Deadline:    task.Deadline().UnixNano(),
		// TODO подумать, как хранить в БД
		//Estimation:  task.Estimation().String(),
		//SpentTime:   task.SpentTime().String(),
	}
}

func (m *TaskRepoModel) GetEntity() (*taskEntity.Task, error) {
	id, err := idPrimitive.EntityIdFrom(m.Id)
	if err != nil {
		return nil, err
	}

	title, err := titlePrimitive.TitleFrom(m.Title)
	if err != nil {
		return nil, err
	}

	description, err := descriptionPrimitive.DescriptionFrom(m.Description)
	if err != nil {
		return nil, err
	}

	status, err := taskStatus.Statuses.Of(m.Status)
	if err != nil {
		return nil, err
	}

	priority, err := taskPriority.Priorities.Of(m.Priority)
	if err != nil {
		return nil, err
	}

	tag, err := taskTag.Tags.Of(m.Tag)
	if err != nil {
		return nil, err
	}

	createAt := commonTime.FromUnixNano(m.CreateAt)

	updateAt := commonTime.FromUnixNano(m.UpdateAt)

	deadline := commonTime.FromUnixNano(m.Deadline)

	estimation := taskDuration.DurationFrom(m.Estimation)

	spentTime := taskDuration.DurationFrom(m.SpentTime)

	return taskEntity.NewBuilder().
		Id(&id).
		Title(&title).
		Description(&description).
		Status(status).
		Priority(priority).
		Tag(tag).
		CreatorId(m.CreatorId).
		PerformerId(m.PerformerId).
		CreatedAt(createAt).
		UpdateAt(updateAt).
		Deadline(deadline).
		Estimation(estimation).
		SpentTime(spentTime).
		Build()
}
