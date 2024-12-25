package taskSerializer

import (
	taskEntity "task-tracker/domain/entity/task"
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
			"tag":         task.Tag(),
			"creatorId":   task.CreatorId(),
			"performerId": task.PerformerId(),
			"createAt":    task.CreateAt(),
			"updateAt":    task.UpdateAt(),
			"deadline":    task.Deadline(),
			"comments":    task.Comments(),
			"estimation":  task.Estimation(),
			"spentTime":   task.SpentTime(),
		},
		Relationships: jsonApiModel.JsonApiObjectRelationships{},
	}
	return response
}
