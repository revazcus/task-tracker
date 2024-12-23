package profilePrimitive

import "task-tracker/infrastructure/errors"

type Builder struct {
	firstName string
	lastName  string
	errors    *errors.Errors
}

func NewBuilder() *Builder {
	return &Builder{
		errors: errors.NewErrors(),
	}
}

func (b *Builder) FirstName(firstName string) *Builder {
	b.firstName = firstName
	return b
}

func (b *Builder) LastName(lastName string) *Builder {
	b.lastName = lastName
	return b
}

func (b *Builder) Build() (*Profile, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	b.fillDefaultFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	return b.createFromBuilder(), nil
}

func (b *Builder) checkRequiredFields() {
	if b.firstName == "" {
		b.errors.AddError(ErrFirstNameIsRequired)
	}
	if b.lastName == "" {
		b.errors.AddError(ErrLastNameIsRequired)
	}
}

func (b *Builder) fillDefaultFields() {

}

func (b *Builder) createFromBuilder() *Profile {
	return &Profile{
		firstName: b.firstName,
		lastName:  b.lastName,
	}
}
