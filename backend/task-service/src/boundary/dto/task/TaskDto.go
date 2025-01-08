package taskDto

import (
	commentDto "task-service/src/boundary/dto/task/comment"
	timeCostsDto "task-service/src/boundary/dto/task/timeCosts"
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
