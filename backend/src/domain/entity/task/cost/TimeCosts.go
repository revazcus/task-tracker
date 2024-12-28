package taskTimeCosts

import (
	"fmt"
	idPrimitive "task-tracker/common/domainPrimitive/id"
	commonTime "task-tracker/infrastructure/tools/time"
)

type TimeEntry struct {
	minutes int
	date    *commonTime.Time
	userId  *idPrimitive.EntityId
}

func (e *TimeEntry) Minutes() int {
	return e.minutes
}

func (e *TimeEntry) Date() *commonTime.Time {
	return e.date
}

func (e *TimeEntry) UserId() *idPrimitive.EntityId {
	return e.userId
}

type TimeCosts struct {
	totalMinutes int
	timeEntries  []*TimeEntry
}

func NewTimeCosts() *TimeCosts {
	return &TimeCosts{timeEntries: make([]*TimeEntry, 0)}
}

func (c *TimeCosts) AddEntry(minutes int, userIdStr string) error {
	if minutes <= 0 {
		return ErrInvalidMinutes
	}

	userId, err := idPrimitive.EntityIdFrom(userIdStr)
	if err != nil {
		return err
	}

	timeEntry := TimeEntry{
		minutes: minutes,
		date:    commonTime.Now(),
		userId:  &userId,
	}

	c.totalMinutes += minutes
	c.timeEntries = append(c.timeEntries, &timeEntry)
	return nil
}

func (c *TimeCosts) TotalTime() string {
	hours := c.totalMinutes / 60
	minutes := c.totalMinutes % 60
	return fmt.Sprintf("Oбщее время работы %d ч %d мин", hours, minutes)
}

func (c *TimeCosts) TotalMinutes() int {
	return c.totalMinutes
}

func (c *TimeCosts) TimeEntries() []*TimeEntry {
	return c.timeEntries
}
