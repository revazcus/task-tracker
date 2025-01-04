package taskDto

import (
	commentDto "github.com/revazcus/task-tracker/backend/task-service/boundary/dto/task/comment"
	timeCostsDto "github.com/revazcus/task-tracker/backend/task-service/boundary/dto/task/timeCosts"
)

type TaskDto struct {
	Id          string
	Title       string
	Description string
	Status      string
	Priority    string
	Tags        []string
	CreatorId   string
	PerformerId string
	DeadLine    string
	Assessment  int
	TimeCosts   timeCostsDto.TimeCostsDto
	Comments    commentDto.CommentDto
}
