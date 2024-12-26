package loggerInterface

import logModel "task-tracker/infrastructure/logger/model"

type LogPublisher interface {
	SendMessage(data *logModel.LogData)
}
