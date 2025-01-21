package agreementPrimitive

import (
	"infrastructure/errors"
	commonTime "infrastructure/tools/time"
)

type Builder struct {
	agreement *Agreement
	errors    *errors.Errors
}

func NewBuilder() *Builder {
	return &Builder{
		agreement: &Agreement{},
		errors:    errors.NewErrors(),
	}
}

func (b *Builder) Accepted(accepted bool) *Builder {
	b.agreement.accepted = accepted
	return b
}

func (b *Builder) AcceptedDate(acceptedDate *commonTime.Time) *Builder {
	b.agreement.acceptedDate = acceptedDate
	return b
}

func (b *Builder) Build() (*Agreement, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	return b.agreement, nil
}

func (b *Builder) checkRequiredFields() {
	if b.agreement.IsAccepted() && b.agreement.AcceptedDate() == nil {
		b.errors.AddError(ErrAcceptedDateIsRequired)
	}
	if !b.agreement.IsAccepted() && b.agreement.AcceptedDate() != nil {
		b.errors.AddError(ErrNotAcceptedAgreement)
	}
}
