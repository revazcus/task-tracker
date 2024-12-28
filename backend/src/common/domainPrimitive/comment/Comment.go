package commentPrimitive

import (
	"strings"
	"unicode/utf8"
)

const CommentMaxLength = 500

type Comment string

func (c Comment) String() string {
	return string(c)
}

func CommentsFrom(comments []string) ([]*Comment, error) {
	var result []*Comment
	for _, str := range comments {
		comment, err := commentFrom(str)
		if err != nil {
			return nil, err
		}
		result = append(result, &comment)
	}
	return result, nil
}

func CommentsToStrings(comments []*Comment) []string {
	result := make([]string, len(comments))
	for i, tag := range comments {
		result[i] = string(*tag)
	}
	return result
}

func commentFrom(comment string) (Comment, error) {
	if comment == "" {
		return "", ErrCommentIsEmpty
	}

	comment = strings.TrimSpace(comment)

	if utf8.RuneCountInString(comment) > CommentMaxLength {
		return "", ErrCommentIsToLong
	}

	return Comment(comment), nil
}
