package event

import "errors"

type EventType string

func (t EventType) String() string {
	return string(t)
}

type EventTypeEnum map[string]EventType

func EventTypesFrom(eventTypes []string) ([]*EventType, error) {
	var result []*EventType
	for _, eventTypeStr := range eventTypes {
		eventType := EventType(eventTypeStr)
		result = append(result, &eventType)
	}
	return result, nil
}

func (e EventTypeEnum) Of(eventTypeStr string) (EventType, error) {
	eventType, ok := e[eventTypeStr]
	if !ok {
		return "", errors.New("unsupported event type")
	}
	return eventType, nil

}
