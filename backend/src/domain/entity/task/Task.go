package taskEntity

import (
	commentPrimitive "task-tracker/common/domainPrimitive/comment"
	descriptionPrimitive "task-tracker/common/domainPrimitive/description"
	idPrimitive "task-tracker/common/domainPrimitive/id"
	titlePrimitive "task-tracker/common/domainPrimitive/title"
	assessmentPrimitive "task-tracker/domain/entity/task/assessment"
	taskTimeCosts "task-tracker/domain/entity/task/cost"
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
	tags        []*taskTag.Tag
	creatorId   string // userId TODO подумать над lite user ver
	performerId string // userId TODO подумать над lite user ver
	createAt    *commonTime.Time
	updateAt    *commonTime.Time
	deadline    *commonTime.Time
	//attachments  string // скрины / видео TODO подумать над реализацией
	assessment *assessmentPrimitive.Assessment
	timeCosts  *taskTimeCosts.TimeCosts
	comments   []*commentPrimitive.Comment
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

func (t *Task) Tags() []*taskTag.Tag {
	return t.tags
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

func (t *Task) Assessment() *assessmentPrimitive.Assessment {
	return t.assessment
}

func (t *Task) TimeCosts() *taskTimeCosts.TimeCosts {
	return t.timeCosts
}
