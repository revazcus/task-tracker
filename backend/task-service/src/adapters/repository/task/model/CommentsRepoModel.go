package taskRepoModel

import (
	shortUserRepoModel "github.com/revazcus/task-tracker/backend/common/repoModel/shortUser"
	commonTime "github.com/revazcus/task-tracker/backend/infrastructure/tools/time"
	taskComment "github.com/revazcus/task-tracker/backend/task-service/domain/entity/task/comment"
)

type CommentsRepoModel struct {
	Comments []CommentRepoModel `bson:"comments"`
}

type CommentRepoModel struct {
	Date   int64                                  `bson:"date"`
	Text   string                                 `bson:"text"`
	Author *shortUserRepoModel.ShortUserRepoModel `bson:"author"`
}

func CommentsToRepoModel(comments *taskComment.Comments) *CommentsRepoModel {
	commentsRepoModel := CommentsRepoModel{Comments: make([]CommentRepoModel, 0)}
	for _, comment := range comments.Comments() {
		commentsRepoModel.Comments = append(commentsRepoModel.Comments, *CommentToRepoModel(comment))
	}
	return &commentsRepoModel
}

func CommentToRepoModel(comment *taskComment.Comment) *CommentRepoModel {
	return &CommentRepoModel{
		Date:   comment.Date().UnixNano(),
		Text:   comment.Text(),
		Author: shortUserRepoModel.ShortUserToRepoModel(comment.Author()),
	}
}

func (m *CommentsRepoModel) GetObject() (*taskComment.Comments, error) {
	comments := taskComment.NewComments()
	for _, commentRepoModel := range m.Comments {
		author, err := commentRepoModel.Author.GetObject()
		if err != nil {
			return nil, err
		}
		if err := comments.AddComment(author, commonTime.FromUnixNano(commentRepoModel.Date), commentRepoModel.Text); err != nil {
			return nil, err
		}
	}
	return comments, nil
}
