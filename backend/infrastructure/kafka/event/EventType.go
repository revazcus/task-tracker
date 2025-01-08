package kafkaEvent

type EventType string

func (t EventType) String() string {
	return string(t)
}
