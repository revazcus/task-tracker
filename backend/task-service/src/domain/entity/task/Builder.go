package taskEntity

import (
	userObject "common/domainObject/shortUser"
	descriptionPrimitive "common/domainPrimitive/description"
	idPrimitive "common/domainPrimitive/id"
	titlePrimitive "common/domainPrimitive/title"
	"infrastructure/errors"
	commonTime "infrastructure/tools/time"
	assessmentPrimitive "task-service/src/domain/entity/task/assessment"
	taskComment "task-service/src/domain/entity/task/comment"
	taskTimeCosts "task-service/src/domain/entity/task/cost"
	taskPriority "task-service/src/domain/entity/task/spec/priority"
	taskStatus "task-service/src/domain/entity/task/spec/status"
	taskTag "task-service/src/domain/entity/task/spec/tag"
)

type Builder struct {
	id          *idPrimitive.EntityId
	title       *titlePrimitive.Title
	description *descriptionPrimitive.Description
	status      taskStatus.Status
	priority    taskPriority.Priority
	tags        []*taskTag.Tag
	creator     *userObject.ShortUser
	performer   *userObject.ShortUser
	createAt    *commonTime.Time
	updateAt    *commonTime.Time
	deadline    *commonTime.Time
	assessment  *assessmentPrimitive.Assessment
	timeCosts   *taskTimeCosts.TimeCosts
	comments    *taskComment.Comments

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

func (b *Builder) Creator(creator *userObject.ShortUser) *Builder {
	b.creator = creator
	return b
}

func (b *Builder) Performer(performer *userObject.ShortUser) *Builder {
	b.performer = performer
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

func (b *Builder) Comments(comments *taskComment.Comments) *Builder {
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
	if b.creator == nil {
		b.errors.AddError(ErrCreatorIsRequired)
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
		creator:     b.creator,
		performer:   b.performer,
		createAt:    b.createAt,
		updateAt:    b.updateAt,
		deadline:    b.deadline,
		assessment:  b.assessment,
		timeCosts:   b.timeCosts,
		comments:    b.comments,
	}
}
