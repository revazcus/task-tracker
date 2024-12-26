package taskStatus

type Status string

type StatusEnum map[string]Status

func (s Status) String() string {
	return string(s)
}

const (
	created    = "New"
	inProgress = "InProgress"
	done       = "Done"
)

var Statuses = StatusEnum{
	created:    created,
	inProgress: inProgress,
	done:       done,
}

func (e StatusEnum) New() Status {
	return e[created]
}

func (e StatusEnum) InProgress() Status {
	return e[inProgress]
}

func (e StatusEnum) Done() Status {
	return e[done]
}

func (e StatusEnum) Of(code string) (Status, error) {
	status, ok := e[code]
	if !ok {
		return "", ErrUnsupportedStatus(code)
	}

	return status, nil
}
