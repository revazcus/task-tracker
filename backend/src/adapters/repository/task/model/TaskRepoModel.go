package taskRepoModel

import (
	commentPrimitive "task-tracker/common/domainPrimitive/comment"
	descriptionPrimitive "task-tracker/common/domainPrimitive/description"
	idPrimitive "task-tracker/common/domainPrimitive/id"
	titlePrimitive "task-tracker/common/domainPrimitive/title"
	taskEntity "task-tracker/domain/entity/task"
	assessmentPrimitive "task-tracker/domain/entity/task/assessment"
	taskPriority "task-tracker/domain/entity/task/spec/priority"
	taskStatus "task-tracker/domain/entity/task/spec/status"
	taskTag "task-tracker/domain/entity/task/spec/tag"
	commonTime "task-tracker/infrastructure/tools/time"
)

type TaskRepoModel struct {
	Id          string              `bson:"task_id"`
	Title       string              `bson:"title"`
	Description string              `bson:"description"`
	Status      string              `bson:"status"`
	Priority    string              `bson:"priority"`
	Tags        []string            `bson:"tags"`
	CreatorId   string              `bson:"creatorId"`
	PerformerId string              `bson:"performerId"`
	CreateAt    int64               `bson:"create_at"`
	UpdateAt    int64               `bson:"update_at"`
	Deadline    int64               `bson:"deadline"`
	Comments    []string            `bson:"comments"`
	Assessment  int                 `bson:"assessment"`
	TimeCosts   *TimeCostsRepoModel `bson:"time_costs"`
}

func TaskToRepoModel(task *taskEntity.Task) *TaskRepoModel {
	return &TaskRepoModel{
		Id:          task.ID().String(),
		Title:       task.Title().String(),
		Description: task.Description().String(),
		Status:      task.Status().String(),
		Priority:    task.Priority().String(),
		Tags:        taskTag.TagsToStrings(task.Tags()),
		CreatorId:   task.CreatorId(),
		PerformerId: task.PerformerId(),
		CreateAt:    task.CreateAt().UnixNano(),
		UpdateAt:    task.UpdateAt().UnixNano(),
		Deadline:    task.Deadline().UnixNano(),
		Comments:    commentPrimitive.CommentsToStrings(task.Comments()),
		Assessment:  task.Assessment().Int(),
		TimeCosts:   TimeCostsToRepoModel(task.TimeCosts()),
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

	tags, err := taskTag.TagsFrom(m.Tags)
	if err != nil {
		return nil, err
	}

	createAt := commonTime.FromUnixNano(m.CreateAt)

	updateAt := commonTime.FromUnixNano(m.UpdateAt)

	deadline := commonTime.FromUnixNano(m.Deadline)

	assessment, err := assessmentPrimitive.AssessmentFrom(m.Assessment)
	if err != nil {
		return nil, err
	}

	timeCosts, err := m.TimeCosts.GetObject()
	if err != nil {
		return nil, err
	}

	comments, err := commentPrimitive.CommentsFrom(m.Comments)
	if err != nil {
		return nil, err
	}

	return taskEntity.NewBuilder().
		Id(&id).
		Title(&title).
		Description(&description).
		Status(status).
		Priority(priority).
		Tags(tags).
		CreatorId(m.CreatorId).
		PerformerId(m.PerformerId).
		CreatedAt(createAt).
		UpdateAt(updateAt).
		Deadline(deadline).
		Assessment(assessment).
		TimeCosts(timeCosts).
		Comments(comments).
		Build()
}
