package taskEntity

import (
	commentPrimitive "task-tracker/common/domainPrimitive/comment"
	descriptionPrimitive "task-tracker/common/domainPrimitive/description"
	idPrimitive "task-tracker/common/domainPrimitive/id"
	titlePrimitive "task-tracker/common/domainPrimitive/title"
	taskDuration "task-tracker/domain/entity/task/duration"
	taskPriority "task-tracker/domain/entity/task/spec/priority"
	taskStatus "task-tracker/domain/entity/task/spec/status"
	taskTag "task-tracker/domain/entity/task/spec/tag"
	commonTime "task-tracker/infrastructure/tools/time"
)

type Task struct {
	id          *idPrimitive.EntityId
	title       *titlePrimitive.Title
	description *descriptionPrimitive.Description
	status      taskStatus.Status
	priority    taskPriority.Priority
	tag         taskTag.Tag
	//severity     string // сложность задачи TODO подумать над целесообразностью
	creatorId   string // userId TODO подумать над lite user ver
	performerId string // userId TODO подумать над lite user ver
	createAt    *commonTime.Time
	updateAt    *commonTime.Time
	deadline    *commonTime.Time
	comments    []*commentPrimitive.Comment
	//attachments  string // скрины / видео TODO подумать над реализацией
	estimation *taskDuration.Duration
	spentTime  *taskDuration.Duration
}

func (t *Task) ID() *idPrimitive.EntityId {
	return t.id
}

func (t *Task) Title() *titlePrimitive.Title {
	return t.title
}

func (t *Task) Description() *descriptionPrimitive.Description {
	return t.description
}

func (t *Task) Status() taskStatus.Status {
	return t.status
}

func (t *Task) Priority() taskPriority.Priority {
	return t.priority
}

func (t *Task) Tag() taskTag.Tag {
	return t.tag
}

func (t *Task) CreatorId() string {
	return t.creatorId
}

func (t *Task) PerformerId() string {
	return t.performerId
}

func (t *Task) CreateAt() *commonTime.Time {
	return t.createAt
}

func (t *Task) UpdateAt() *commonTime.Time {
	return t.updateAt
}

func (t *Task) Deadline() *commonTime.Time {
	return t.deadline
}

func (t *Task) Comments() []*commentPrimitive.Comment {
	return t.comments
}

func (t *Task) Estimation() *taskDuration.Duration {
	return t.estimation
}

func (t *Task) SpentTime() *taskDuration.Duration {
	return t.spentTime
}
