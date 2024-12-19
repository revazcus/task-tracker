package logModel

import "strings"

type LogMessage struct {
	Message   string
	Component string
	Lvl       LogLvl
	Fields    []*LogField
}

func NewLogMessage(lvl LogLvl, msg string) *LogMessage {
	return &LogMessage{
		Message: msg,
		Lvl:     lvl,
	}
}

func (m *LogMessage) SetMessage(messages ...string) *LogMessage {
	m.Message = strings.Join(messages, " / ")
	return m
}

func (m *LogMessage) SetComponent(component string) *LogMessage {
	m.Component = component
	return m
}

func (m *LogMessage) SetFields(fields ...*LogField) *LogMessage {
	m.Fields = fields
	return m
}
