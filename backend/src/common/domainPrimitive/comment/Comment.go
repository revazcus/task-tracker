package commentPrimitive

import (
	"strings"
	"unicode/utf8"
)

const CommentMaxLength = 500

type Comment string

func CommentFrom(comment string) (Comment, error) {
	if comment == "" {
		return "", ErrCommentIsEmpty
	}

	comment = strings.TrimSpace(comment)

	if utf8.RuneCountInString(comment) > CommentMaxLength {
		return "", ErrCommentIsToLong
	}

	return Comment(comment), nil
}

func (c Comment) String() string {
	return string(c)
}
