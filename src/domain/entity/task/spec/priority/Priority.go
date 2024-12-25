package taskPriority

type Priority string

type PriorityEnum map[string]Priority

func (p Priority) String() string {
	return string(p)
}

const (
	low      = "Low"
	medium   = "Medium"
	high     = "High"
	critical = "Critical"
)

var Priorities = PriorityEnum{
	low:      low,
	medium:   medium,
	high:     high,
	critical: critical,
}

func (e PriorityEnum) Low() Priority {
	return e[low]
}

func (e PriorityEnum) Medium() Priority {
	return e[medium]
}

func (e PriorityEnum) High() Priority {
	return e[high]
}

func (e PriorityEnum) Critical() Priority {
	return e[critical]
}

func (e PriorityEnum) Of(code string) (Priority, error) {
	priority, ok := e[code]
	if !ok {
		return "", ErrUnsupportedPriority(code)
	}

	return priority, nil
}
