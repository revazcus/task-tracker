package taskRepo

import (
	userObject "common/domainObject/shortUser"
	idPrimitive "common/domainPrimitive/id"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	loggerInterface "infrastructure/logger/interface"
	logModel "infrastructure/logger/model"
	mongoInterface "infrastructure/mongo/interface"
	mongoModel "infrastructure/mongo/model"
	commonTime "infrastructure/tools/time"
	taskRepoModel "task-service/adapters/repository/task/model"
	taskEntity "task-service/domain/entity/task"
	taskComment "task-service/domain/entity/task/comment"
	taskTimeCosts "task-service/domain/entity/task/cost"
	taskStatus "task-service/domain/entity/task/spec/status"
	taskTag "task-service/domain/entity/task/spec/tag"
)

const (
	indexTaskId  = "uniqTaskId"
	indexTaskKey = "task_id"
)

type TaskRepo struct {
	collection string
	mongoRepo  mongoInterface.MongoRepository
	logger     loggerInterface.Logger
}

func (r *TaskRepo) Init(ctx context.Context) error {
	taskIdIndex := &mongoModel.DBIndex{
		Collection: r.collection,
		Name:       indexTaskId,
		Keys:       []string{indexTaskKey},
		Type:       mongoModel.DBIndexAsc,
		Uniq:       true,
	}

	if err := r.mongoRepo.TryCreateIndex(ctx, taskIdIndex); err != nil {
		return err
	}

	return nil
}

func (r *TaskRepo) Create(ctx context.Context, task *taskEntity.Task) error {
	taskModel := taskRepoModel.TaskToRepoModel(task)
	if err := r.mongoRepo.InsertOne(ctx, r.collection, taskModel); err != nil {
		return err
	}
	return nil
}

func (r *TaskRepo) GetAll(ctx context.Context) ([]*taskEntity.Task, error) {
	var taskModels []*taskRepoModel.TaskRepoModel
	if err := r.mongoRepo.Find(ctx, r.collection, &taskModels, bson.D{}, options.Find().SetComment("Get all tasks")); err != nil {
		return nil, err
	}

	var tasks []*taskEntity.Task
	for _, taskModel := range taskModels {
		task, err := taskModel.GetEntity()
		if err != nil {
			r.logger.Error(ctx, err, logModel.WithComponent("Mongo"))
			continue
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *TaskRepo) GetById(ctx context.Context, taskId *idPrimitive.EntityId) (*taskEntity.Task, error) {
	find := bson.D{{"task_id", taskId.String()}}
	var taskModel *taskRepoModel.TaskRepoModel

	if err := r.mongoRepo.FindOne(ctx, r.collection, find, &taskModel); err != nil {
		// TODO добавить кастомную ошибку
		return nil, err
	}

	task, err := taskModel.GetEntity()
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *TaskRepo) Update(ctx context.Context, task *taskEntity.Task) (*taskEntity.Task, error) {
	change := bson.D{
		{"$set", bson.M{
			"title":       task.Title().String(),
			"description": task.Description().String(),
			"priority":    task.Priority().String(),
			"status":      task.Status().String(),
			"tags":        taskTag.TagsToStrings(task.Tags()),
			"deadline":    task.Deadline().UnixNano(),
			"assessment":  task.Assessment().Int(),
			"update_at":   commonTime.Now().UnixNano(),
		}},
	}
	return r.updateTask(ctx, task.ID(), change)
}

func (r *TaskRepo) UpdatePerformer(ctx context.Context, taskId *idPrimitive.EntityId, performer *userObject.ShortUser) (*taskEntity.Task, error) {
	change := bson.D{
		{"$set", bson.M{
			"performer": bson.M{
				"user_id": performer.ID().String(),
				"email":   performer.Email().String(),
				"profile": bson.M{
					"first_name": performer.Profile().FirstName(),
					"last_name":  performer.Profile().LastName(),
				},
			},
			"update_at": commonTime.Now().UnixNano(),
		}},
	}
	return r.updateTask(ctx, taskId, change)
}

func (r *TaskRepo) UpdatePerformerAndStatus(ctx context.Context, taskId *idPrimitive.EntityId, performer *userObject.ShortUser, status taskStatus.Status) (*taskEntity.Task, error) {
	change := bson.D{
		{"$set", bson.M{
			"performer": bson.M{
				"user_id": performer.ID().String(),
				"email":   performer.Email().String(),
				"profile": bson.M{
					"first_name": performer.Profile().FirstName(),
					"last_name":  performer.Profile().LastName(),
				},
			},
			"status":    status.String(),
			"update_at": commonTime.Now().UnixNano(),
		}},
	}
	return r.updateTask(ctx, taskId, change)
}

func (r *TaskRepo) UpdateTimeCosts(ctx context.Context, taskId *idPrimitive.EntityId, timeCost *taskTimeCosts.TimeInvestment) (*taskEntity.Task, error) {
	timeCostRepoModel := taskRepoModel.TimeInvestmentToRepoModel(timeCost)
	change := bson.D{
		{"$push", bson.M{"time_costs.time_investments": timeCostRepoModel}},
		{"$set", bson.M{"update_at": commonTime.Now().UnixNano()}},
	}
	return r.updateTask(ctx, taskId, change)
}

func (r *TaskRepo) UpdateComments(ctx context.Context, taskId *idPrimitive.EntityId, comment *taskComment.Comment) (*taskEntity.Task, error) {
	commentRepoModel := taskRepoModel.CommentToRepoModel(comment)
	change := bson.D{
		{"$push", bson.M{"comments.comments": commentRepoModel}},
	}
	return r.updateTask(ctx, taskId, change)
}

func (r *TaskRepo) DeleteById(ctx context.Context, taskId *idPrimitive.EntityId) error {
	find := bson.D{{"task_id", taskId.String()}}
	if err := r.mongoRepo.DeleteOne(ctx, r.collection, find); err != nil {
		return err
	}
	return nil
}

func (r *TaskRepo) updateTask(ctx context.Context, taskId *idPrimitive.EntityId, change bson.D) (*taskEntity.Task, error) {
	find := bson.D{{"task_id", taskId.String()}}

	var taskModel *taskRepoModel.TaskRepoModel

	if err := r.mongoRepo.FindOneAndUpdate(ctx, r.collection, &taskModel, find, change, options.FindOneAndUpdate().SetReturnDocument(options.After)); err != nil {
		return nil, err
	}

	updatedTask, err := taskModel.GetEntity()
	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}
