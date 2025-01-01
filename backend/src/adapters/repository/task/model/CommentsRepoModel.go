package taskRepoModel

import (
	taskComment "task-tracker/domain/entity/task/comment"
	commonTime "task-tracker/infrastructure/tools/time"
)

type CommentsRepoModel struct {
	Comments []CommentRepoModel `bson:"comments"`
}

type CommentRepoModel struct {
	AuthorId string `bson:"author_id"`
	Date     int64  `bson:"date"`
	Text     string `bson:"text"`
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
		AuthorId: comment.UserId().String(),
		Date:     comment.Date().UnixNano(),
		Text:     comment.Text(),
	}
}

func (m *CommentsRepoModel) GetObject() (*taskComment.Comments, error) {
	comments := taskComment.NewComments()
	for _, commentRepoModel := range m.Comments {
		if err := comments.AddComment(commentRepoModel.AuthorId, commonTime.FromUnixNano(commentRepoModel.Date), commentRepoModel.Text); err != nil {
			return nil, err
		}
	}
	return comments, nil
}
