package loggerInterface

import logModel "github.com/revazcus/task-tracker/backend/infrastructure/logger/model"

type LogPublisher interface {
	SendMessage(data *logModel.LogData)
}
