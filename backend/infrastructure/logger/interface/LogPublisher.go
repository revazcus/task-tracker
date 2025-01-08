package loggerInterface

import logModel "infrastructure/logger/model"

type LogPublisher interface {
	SendMessage(data *logModel.LogData)
}
