package titlePrimitive

import (
	"strings"
	"unicode/utf8"
)

const TitleMaxLength = 255

type Title string

func TitleFrom(title string) (Title, error) {
	if title == "" {
		return "", ErrTitleIsEmpty
	}

	title = strings.TrimSpace(title)

	if utf8.RuneCountInString(title) > TitleMaxLength {
		return "", ErrTitleIsToLong
	}

	return Title(title), nil
}

func (t Title) String() string {
	return string(t)
}
