package request

import (
	"bytes"
	"encoding/json"
	taskDto "task-tracker/boundary/dto/task"
)

type CreateTaskRequest struct {
	Data struct {
		Id         string `json:"id"`
		Attributes struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			Status      string `json:"status"`
			Priority    string `json:"priority"`
			Tag         string `json:"Tag"`
			CreatorId   string `json:"creatorId"`
			PerformerId string `json:"performerId"`
			Deadline    string `json:"deadline"`
			Comments    string `json:"comments"` // TODO нужно поддержать массив
			Estimation  string `json:"estimation"`
			SpentTime   string `json:"spentTime"`
		} `json:"attributes"`
	} `json:"data"`
}

func (r *CreateTaskRequest) FillFromBytes(jsonBytes []byte) error {
	return json.NewDecoder(bytes.NewReader(jsonBytes)).Decode(r)
}

func (r *CreateTaskRequest) CreateTaskDto() *taskDto.TaskDto {
	return &taskDto.TaskDto{
		Id:          r.Data.Id,
		Title:       r.Data.Attributes.Title,
		Description: r.Data.Attributes.Description,
		Status:      r.Data.Attributes.Status,
		Priority:    r.Data.Attributes.Priority,
		Tag:         r.Data.Attributes.Tag,
		CreatorId:   r.Data.Attributes.CreatorId,
		PerformerId: r.Data.Attributes.PerformerId,
		DeadLine:    r.Data.Attributes.Deadline,
		Comments:    r.Data.Attributes.Comments,
		Estimation:  r.Data.Attributes.Estimation,
		SpentTime:   r.Data.Attributes.SpentTime,
	}
}
