package taskTimeCosts

import (
	userObject "common/domainObject/shortUser"
	"fmt"
	commonTime "infrastructure/tools/time"
)

type TimeInvestment struct {
	minutes int
	date    *commonTime.Time
	worker  *userObject.ShortUser
}

func (e *TimeInvestment) Minutes() int {
	return e.minutes
}

func (e *TimeInvestment) Date() *commonTime.Time {
	return e.date
}

func (e *TimeInvestment) Worker() *userObject.ShortUser {
	return e.worker
}

type TimeCosts struct {
	totalMinutes    int
	timeInvestments []*TimeInvestment
}

func NewTimeCosts() *TimeCosts {
	return &TimeCosts{timeInvestments: make([]*TimeInvestment, 0)}
}

func (c *TimeCosts) AddTimeCost(worker *userObject.ShortUser, date *commonTime.Time, minutes int) error {
	timeCost, err := AddTimeCost(worker, date, minutes)
	if err != nil {
		return err
	}
	c.totalMinutes += minutes
	c.timeInvestments = append(c.timeInvestments, timeCost)
	return nil
}

func AddTimeCost(worker *userObject.ShortUser, date *commonTime.Time, minutes int) (*TimeInvestment, error) {
	if minutes <= 0 {
		return nil, ErrInvalidMinutes
	}

	timeCost := TimeInvestment{
		minutes: minutes,
		date:    date,
		worker:  worker,
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

func (c *TimeCosts) TimeInvestments() []*TimeInvestment {
	return c.timeInvestments
}
