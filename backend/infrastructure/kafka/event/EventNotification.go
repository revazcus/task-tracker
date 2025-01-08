package kafkaEvent

import (
	"encoding/json"
	"github.com/google/uuid"
	commonTime "infrastructure/tools/time"
)

type EventNotification struct {
	EventId   string                 `json:"event_id"`
	EventType *EventType             `json:"event_type"`
	Date      *commonTime.Time       `json:"date"`
	Source    string                 `json:"source"`
	Payload   map[string]interface{} `json:"payload"`
}

func NewEventNotification(eventType *EventType, source string, payload map[string]interface{}) *EventNotification {
	return &EventNotification{
		EventId:   uuid.NewString(),
		EventType: eventType,
		Date:      commonTime.Now(),
		Source:    source,
		Payload:   payload,
	}
}

func (n *EventNotification) ToBytes() ([]byte, error) {
	return json.Marshal(n)
}

func (n *EventNotification) FromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, n)
}
