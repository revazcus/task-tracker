package taskSerializer

import (
	taskEntity "task-tracker/domain/entity/task"
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
			"creatorId":   task.CreatorId(),
			"performerId": task.PerformerId(),
			"createAt":    task.CreateAt(),
			"updateAt":    task.UpdateAt(),
			"deadline":    task.Deadline(),
			"assessment":  task.Assessment(),
			"timeCosts":   serializeTimeCosts(task.TimeCosts()),
			"totalTime":   task.TimeCosts().TotalTime(),
			"comments":    task.Comments(),
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
		"totalMinutes": timeCosts.TotalMinutes(),
		"timeEntries":  serializeTimeEntries(timeCosts.TimeEntries()),
	}
}

func serializeTimeEntries(timeEntries []*taskTimeCosts.TimeEntry) []interface{} {
	serializedEntries := make([]interface{}, len(timeEntries))
	for i, timeEntry := range timeEntries {
		serializedEntries[i] = map[string]interface{}{
			"minutes": timeEntry.Minutes(),
			"date":    timeEntry.Date(),
			"userId":  timeEntry.UserId(),
		}
	}
	return serializedEntries
}
