package descriptionPrimitive

import (
	"strings"
	"unicode/utf8"
)

const DescriptionMaxLength = 500

type Description string

func DescriptionFrom(description string) (Description, error) {
	if description == "" {
		return "", ErrDescriptionIsEmpty
	}

	description = strings.TrimSpace(description)

	if utf8.RuneCountInString(description) > DescriptionMaxLength {
		return "", ErrDescriptionToLong
	}

	return Description(description), nil
}

func (d Description) String() string {
	return string(d)
}
