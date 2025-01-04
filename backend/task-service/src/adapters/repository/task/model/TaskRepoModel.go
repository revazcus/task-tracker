package taskRepoModel

import (
	descriptionPrimitive "github.com/revazcus/task-tracker/backend/common/domainPrimitive/description"
	idPrimitive "github.com/revazcus/task-tracker/backend/common/domainPrimitive/id"
	titlePrimitive "github.com/revazcus/task-tracker/backend/common/domainPrimitive/title"
	shortUserRepoModel "github.com/revazcus/task-tracker/backend/common/repoModel/shortUser"
	commonTime "github.com/revazcus/task-tracker/backend/infrastructure/tools/time"
	taskEntity "github.com/revazcus/task-tracker/backend/task-service/domain/entity/task"
	assessmentPrimitive "github.com/revazcus/task-tracker/backend/task-service/domain/entity/task/assessment"
	taskPriority "github.com/revazcus/task-tracker/backend/task-service/domain/entity/task/spec/priority"
	taskStatus "github.com/revazcus/task-tracker/backend/task-service/domain/entity/task/spec/status"
)

type TaskRepoModel struct {
	Id          string                                 `bson:"task_id"`
	Title       string                                 `bson:"title"`
	Description string                                 `bson:"description"`
	Status      string                                 `bson:"status"`
	Priority    string                                 `bson:"priority"`
	Tags        []string                               `bson:"tags"`
	Creator     *shortUserRepoModel.ShortUserRepoModel `bson:"creator"`
	Performer   *shortUserRepoModel.ShortUserRepoModel `bson:"performer"`
	CreateAt    int64                                  `bson:"create_at"`
	UpdateAt    int64                                  `bson:"update_at"`
	Deadline    int64                                  `bson:"deadline"`
	Assessment  int                                    `bson:"assessment"`
	TimeCosts   *TimeCostsRepoModel                    `bson:"time_costs"`
	Comments    *CommentsRepoModel                     `bson:"comments"`
}

func TaskToRepoModel(task *taskEntity.Task) *TaskRepoModel {
	return &TaskRepoModel{
		Id:          task.ID().String(),
		Title:       task.Title().String(),
		Description: task.Description().String(),
		Status:      task.Status().String(),
		Priority:    task.Priority().String(),
		Tags:        taskTag.TagsToStrings(task.Tags()),
		Creator:     shortUserRepoModel.ShortUserToRepoModel(task.Creator()),
		Performer:   shortUserRepoModel.ShortUserToRepoModel(task.Performer()),
		CreateAt:    task.CreateAt().UnixNano(),
		UpdateAt:    task.UpdateAt().UnixNano(),
		Deadline:    task.Deadline().UnixNano(),
		Assessment:  task.Assessment().Int(),
		TimeCosts:   TimeInvestmentsToRepoModels(task.TimeCosts()),
		Comments:    CommentsToRepoModel(task.Comments()),
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

	creator, err := m.Creator.GetObject()
	if err != nil {
		return nil, err
	}

	performer, err := m.Performer.GetObject()
	if err != nil {
		return nil, err
	}

	createAt := commonTime.FromUnixNano(m.CreateAt)

	var updateAt *commonTime.Time
	if m.UpdateAt == 0 {
		updateAt = nil
	} else {
		updateAt = commonTime.FromUnixNano(m.UpdateAt)
	}

	deadline := commonTime.FromUnixNano(m.Deadline)

	assessment, err := assessmentPrimitive.AssessmentFrom(m.Assessment)
	if err != nil {
		return nil, err
	}

	timeCosts, err := m.TimeCosts.GetObject()
	if err != nil {
		return nil, err
	}

	comments, err := m.Comments.GetObject()
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
		Creator(creator).
		Performer(performer).
		CreatedAt(createAt).
		UpdateAt(updateAt).
		Deadline(deadline).
		Assessment(assessment).
		TimeCosts(timeCosts).
		Comments(comments).
		Build()
}
