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
	"task-tracker/infrastructure/errors"
	commonTime "task-tracker/infrastructure/tools/time"
)

type Builder struct {
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
	assessment  *assessmentPrimitive.Assessment
	timeCosts   *taskTimeCosts.TimeCosts
	comments    []*commentPrimitive.Comment

	errors *errors.Errors
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

func (b *Builder) Tags(tags []*taskTag.Tag) *Builder {
	b.tags = tags
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

func (b *Builder) Assessment(assessment *assessmentPrimitive.Assessment) *Builder {
	b.assessment = assessment
	return b
}

func (b *Builder) TimeCosts(timeCosts *taskTimeCosts.TimeCosts) *Builder {
	b.timeCosts = timeCosts
	return b
}

func (b *Builder) Comments(comments []*commentPrimitive.Comment) *Builder {
	b.comments = comments
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
}

func (b *Builder) fillDefaultFields() {
	if b.id == nil {
		entityId := idPrimitive.NewEntityId()
		b.id = &entityId
	}
	if b.status == "" {
		b.status = taskStatus.Statuses.New()
	}
	if b.priority == "" {
		b.priority = taskPriority.Priorities.Low()
	}
	if b.createAt == nil {
		b.createAt = commonTime.Now()
	}
}

func (b *Builder) createFromBuilder() *Task {
	return &Task{
		id:          b.id,
		title:       b.title,
		description: b.description,
		status:      b.status,
		priority:    b.priority,
		tags:        b.tags,
		creatorId:   b.creatorId,
		performerId: b.performerId,
		createAt:    b.createAt,
		updateAt:    b.updateAt,
		deadline:    b.deadline,
		assessment:  b.assessment,
		timeCosts:   b.timeCosts,
		comments:    b.comments,
	}
}
