package agreementPrimitive

import commonTime "github.com/revazcus/task-tracker/backend/infrastructure/tools/time"

type Agreement struct {
	accepted     bool
	acceptedDate *commonTime.Time
}

func (a *Agreement) Accept() {
	a.accepted = true
	a.acceptedDate = commonTime.Now()
}

func (a *Agreement) IsAccepted() bool {
	return a.accepted
}

func (a *Agreement) AcceptedDate() *commonTime.Time {
	return a.acceptedDate
}
