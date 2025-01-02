package taskSerializer

import (
	userObject "task-tracker/common/domainObject/shortUser"
	taskEntity "task-tracker/domain/entity/task"
	taskComment "task-tracker/domain/entity/task/comment"
	taskTimeCosts "task-tracker/domain/entity/task/cost"
	jsonApiModel "task-tracker/infrastructure/jsonapi/model"
)

const (
	ResponseTask = "task"
)

func SerializeTask(task *taskEntity.Task) (jsonApiModel.JsonApiPayload, error) {
	responseBuilder := jsonApiModel.NewJsonApiPayloadBuilder()
	responseBuilder.AddData(CreateTaskObject(task))
	return responseBuilder.Build()
}

func SerializeTasks(tasks []*taskEntity.Task) (jsonApiModel.JsonApiPayload, error) {
	responseBuilder := jsonApiModel.NewJsonApiPayloadBuilder()
	for _, task := range tasks {
		responseBuilder.AddData(CreateTaskObject(task))
	}
	return responseBuilder.Build()
}

func CreateTaskObject(task *taskEntity.Task) *jsonApiModel.JsonApiObject {
	response := &jsonApiModel.JsonApiObject{
		Id:   task.ID().String(),
		Type: ResponseTask,
		Attributes: map[string]interface{}{
			"title":       task.Title(),
			"description": task.Description(),
			"status":      task.Status(),
			"priority":    task.Priority(),
			"tags":        task.Tags(),
			"creator":     serializeShortUser(task.Creator()),
			"performer":   serializeShortUser(task.Performer()),
			"createAt":    task.CreateAt(),
			"updateAt":    task.UpdateAt(),
			"deadline":    task.Deadline(),
			"assessment":  task.Assessment(),
			"timeCosts":   serializeTimeCosts(task.TimeCosts()),
			"totalTime":   task.TimeCosts().TotalTime(),
			"comments":    serializeComments(task.Comments()),
		},
		Relationships: jsonApiModel.JsonApiObjectRelationships{},
	}
	return response
}

func serializeTimeCosts(timeCosts *taskTimeCosts.TimeCosts) interface{} {
	if timeCosts == nil {
		return nil
	}
	return map[string]interface{}{
		"totalMinutes":    timeCosts.TotalMinutes(),
		"timeInvestments": serializeTimeCost(timeCosts.TimeInvestments()),
	}
}

func serializeTimeCost(timeCosts []*taskTimeCosts.TimeInvestment) []interface{} {
	serializedTimeCosts := make([]interface{}, len(timeCosts))
	for i, timeCost := range timeCosts {
		serializedTimeCosts[i] = map[string]interface{}{
			"worker":  serializeShortUser(timeCost.Worker()),
			"date":    timeCost.Date(),
			"minutes": timeCost.Minutes(),
		}
	}
	return serializedTimeCosts
}

func serializeComments(comments *taskComment.Comments) interface{} {
	if comments == nil {
		return nil
	}
	serializedComments := make([]interface{}, len(comments.Comments()))
	for i, comment := range comments.Comments() {
		serializedComments[i] = map[string]interface{}{
			"author": serializeShortUser(comment.Author()),
			"date":   comment.Date(),
			"text":   comment.Text(),
		}
	}
	return serializedComments
}

func serializeShortUser(shortUser *userObject.ShortUser) interface{} {
	if shortUser == nil {
		return nil
	}
	return map[string]interface{}{
		"userId":   shortUser.ID().String(),
		"email":    shortUser.Email().String(),
		"fullName": shortUser.Profile().FullName(),
	}
}
