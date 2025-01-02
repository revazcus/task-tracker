package profilePrimitive

import (
	"fmt"
	"strings"
	errors "task-tracker/infrastructure/errors"
)

type Profile struct {
	firstName string
	lastName  string
}

func NewProfile(firstName, lastName string) (*Profile, error) {
	errs := errors.NewErrors()
	if firstName == "" {
		errs.AddError(ErrFirstNameIsRequired)
	}
	if lastName == "" {
		errs.AddError(ErrLastNameIsRequired)
	}
	if errs.IsPresent() {
		return nil, errs
	}
	return &Profile{firstName: firstName, lastName: lastName}, nil
}

func (p Profile) FullName() string {
	fullName := fmt.Sprintf("%s %s", p.firstName, p.lastName)
	return strings.TrimSpace(fullName)
}

func (p Profile) FirstName() string {
	return p.firstName
}

func (p Profile) LastName() string {
	return p.lastName
}
