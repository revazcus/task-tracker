package profilePrimitive

import (
	"fmt"
	"strings"
)

type Profile struct {
	firstName string
	lastName  string
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
