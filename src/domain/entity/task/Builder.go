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
	"task-tracker/infrastructure/errors"
	commonTime "task-tracker/infrastructure/tools/time"
	"time"
)

type Builder struct {
	id          *idPrimitive.EntityId
	title       *titlePrimitive.Title
	description *descriptionPrimitive.Description
	status      taskStatus.Status
	priority    taskPriority.Priority
	tag         taskTag.Tag
	creatorId   string // userId TODO подумать над lite user ver
	performerId string // userId TODO подумать над lite user ver
	createAt    *commonTime.Time
	updateAt    *commonTime.Time
	deadline    *commonTime.Time
	comments    []*commentPrimitive.Comment
	estimation  *taskDuration.Duration
	spentTime   *taskDuration.Duration
	errors      *errors.Errors
}

func NewBuilder() *Builder {
	return &Builder{
		errors: errors.NewErrors(),
	}
}

func (b *Builder) Id(id *idPrimitive.EntityId) *Builder {
	b.id = id
	return b
}

func (b *Builder) Title(title *titlePrimitive.Title) *Builder {
	b.title = title
	return b
}

func (b *Builder) Description(description *descriptionPrimitive.Description) *Builder {
	b.description = description
	return b
}

func (b *Builder) Status(status taskStatus.Status) *Builder {
	b.status = status
	return b
}

func (b *Builder) Priority(priority taskPriority.Priority) *Builder {
	b.priority = priority
	return b
}

func (b *Builder) Tag(tag taskTag.Tag) *Builder {
	b.tag = tag
	return b
}

func (b *Builder) CreatorId(creatorId string) *Builder {
	b.creatorId = creatorId
	return b
}

func (b *Builder) PerformerId(performerId string) *Builder {
	b.performerId = performerId
	return b
}

func (b *Builder) CreatedAt(createdAt *commonTime.Time) *Builder {
	b.createAt = createdAt
	return b
}

func (b *Builder) UpdateAt(updateAt *commonTime.Time) *Builder {
	b.updateAt = updateAt
	return b
}

func (b *Builder) Deadline(deadline *commonTime.Time) *Builder {
	b.deadline = deadline
	return b
}

func (b *Builder) Comments(comments []*commentPrimitive.Comment) *Builder {
	b.comments = comments
	return b
}

func (b *Builder) Estimation(estimation *taskDuration.Duration) *Builder {
	b.estimation = estimation
	return b
}

func (b *Builder) SpentTime(spentTime *taskDuration.Duration) *Builder {
	b.spentTime = spentTime
	return b
}

func (b *Builder) Build() (*Task, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	b.fillDefaultFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	return b.createFromBuilder(), nil
}

func (b *Builder) checkRequiredFields() {
	if b.title == nil {
		b.errors.AddError(ErrTitleIsRequired)
	}
	if b.description == nil {
		b.errors.AddError(ErrDescriptionIsRequired)
	}
	if b.creatorId == "" {
		b.errors.AddError(ErrCreatorIdIsRequired)
	}
	//if b.deadline == nil {
	//	b.errors.AddError(ErrDeadlineIsRequired)
	//}
}

func (b *Builder) fillDefaultFields() {
	if b.id == nil {
		entityId := idPrimitive.NewEntityId()
		b.id = &entityId
	}
	// TODO подумать как фронт будет получать доступный список статусов
	if b.status == "" {
		b.status = taskStatus.Statuses.New()
	}
	// TODO подумать как фронт будет получать доступный список приоритетов
	if b.priority == "" {
		b.priority = taskPriority.Priorities.Low()
	}
	// TODO подумать как фронт будет получать доступный список тегов
	if b.tag == "" {
		b.tag = taskTag.Tags.Quest()
	}
	if b.createAt == nil {
		b.createAt = commonTime.Now()
	}
	// TODO подумать что делать с updateAt (сейчас заглушка)
	if b.updateAt == nil {
		b.updateAt = commonTime.Now()
	}
	// TODO подумать как фронт будет передавать дедлайн (сейчас заглушка)
	if b.deadline == nil {
		b.deadline = commonTime.FromTime(time.Time.Add(time.Now(), time.Duration(2)))
	}
}

func (b *Builder) createFromBuilder() *Task {
	return &Task{
		id:          b.id,
		title:       b.title,
		description: b.description,
		status:      b.status,
		priority:    b.priority,
		tag:         b.tag,
		creatorId:   b.creatorId,
		performerId: b.performerId,
		createAt:    b.createAt,
		updateAt:    b.updateAt,
		deadline:    b.deadline,
		comments:    b.comments,
		estimation:  b.estimation,
		spentTime:   b.spentTime,
	}
}
