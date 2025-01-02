package taskComments

import (
	"strings"
	userObject "task-tracker/common/domainObject/shortUser"
	commonTime "task-tracker/infrastructure/tools/time"
	"unicode/utf8"
)

const CommentMaxLength = 500

type Comment struct {
	author *userObject.ShortUser
	date   *commonTime.Time
	text   string
}

func (c *Comment) Author() *userObject.ShortUser {
	return c.author
}

func (c *Comment) Date() *commonTime.Time {
	return c.date
}

func (c *Comment) Text() string {
	return c.text
}

type Comments struct {
	comments []*Comment
}

func NewComments() *Comments {
	return &Comments{comments: make([]*Comment, 0)}
}

func (c *Comments) AddComment(author *userObject.ShortUser, date *commonTime.Time, text string) error {
	comment, err := AddComment(author, date, text)
	if err != nil {
		return err
	}
	c.comments = append(c.comments, comment)
	return nil
}

func AddComment(author *userObject.ShortUser, date *commonTime.Time, text string) (*Comment, error) {
	if text == "" {
		return nil, ErrCommentIsEmpty
	}

	text = strings.TrimSpace(text)

	if utf8.RuneCountInString(text) > CommentMaxLength {
		return nil, ErrCommentIsToLong
	}

	comment := Comment{
		author: author,
		date:   date,
		text:   text,
	}

	return &comment, nil
}

func (c *Comments) Comments() []*Comment {
	return c.comments
}
