package taskTimeCosts

import (
	"fmt"
	idPrimitive "task-tracker/common/domainPrimitive/id"
	commonTime "task-tracker/infrastructure/tools/time"
)

type TimeCost struct {
	minutes int
	date    *commonTime.Time
	userId  *idPrimitive.EntityId
}

func (e *TimeCost) Minutes() int {
	return e.minutes
}

func (e *TimeCost) Date() *commonTime.Time {
	return e.date
}

func (e *TimeCost) UserId() *idPrimitive.EntityId {
	return e.userId
}

type TimeCosts struct {
	totalMinutes int
	timeCosts    []*TimeCost
}

func NewTimeCosts() *TimeCosts {
	return &TimeCosts{timeCosts: make([]*TimeCost, 0)}
}

func (c *TimeCosts) AddTimeCost(userIdStr string, date *commonTime.Time, minutes int) error {
	timeCost, err := AddTimeCost(userIdStr, date, minutes)
	if err != nil {
		return err
	}
	c.totalMinutes += minutes
	c.timeCosts = append(c.timeCosts, timeCost)
	return nil
}

func AddTimeCost(userIdStr string, date *commonTime.Time, minutes int) (*TimeCost, error) {
	userId, err := idPrimitive.EntityIdFrom(userIdStr)
	if err != nil {
		return nil, err
	}

	if minutes <= 0 {
		return nil, ErrInvalidMinutes
	}

	timeCost := TimeCost{
		minutes: minutes,
		date:    date,
		userId:  &userId,
	}

	return &timeCost, nil
}

func (c *TimeCosts) TotalTime() string {
	hours := c.totalMinutes / 60
	minutes := c.totalMinutes % 60
	return fmt.Sprintf("Oбщее время работы %d ч %d мин", hours, minutes)
}

func (c *TimeCosts) TotalMinutes() int {
	return c.totalMinutes
}

func (c *TimeCosts) TimeCosts() []*TimeCost {
	return c.timeCosts
}
