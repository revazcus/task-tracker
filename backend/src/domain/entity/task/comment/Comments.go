package taskComments

import (
	"strings"
	idPrimitive "task-tracker/common/domainPrimitive/id"
	commonTime "task-tracker/infrastructure/tools/time"
	"unicode/utf8"
)

const CommentMaxLength = 500

type Comment struct {
	userId *idPrimitive.EntityId
	date   *commonTime.Time
	text   string
}

func (c *Comment) UserId() *idPrimitive.EntityId {
	return c.userId
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

func (c *Comments) AddComment(authorIdSts string, date *commonTime.Time, text string) error {
	comment, err := AddComment(authorIdSts, date, text)
	if err != nil {
		return err
	}
	c.comments = append(c.comments, comment)
	return nil
}

func AddComment(authorIdSts string, date *commonTime.Time, text string) (*Comment, error) {
	authorId, err := idPrimitive.EntityIdFrom(authorIdSts)
	if err != nil {
		return nil, err
	}

	if text == "" {
		return nil, ErrCommentIsEmpty
	}

	text = strings.TrimSpace(text)

	if utf8.RuneCountInString(text) > CommentMaxLength {
		return nil, ErrCommentIsToLong
	}

	comment := Comment{
		userId: &authorId,
		date:   date,
		text:   text,
	}

	return &comment, nil
}

func (c *Comments) Comments() []*Comment {
	return c.comments
}
