package taskDuration

import "fmt"

type Duration struct {
	days    int
	hours   int
	minutes int
	str     string // TODO убрать
}

func (d *Duration) AddDays(days int) error {
	if days <= 0 {
		return ErrInvalidDayValue
	}
	d.days += days
	return nil
}

func (d *Duration) AddHours(hours int) error {
	if hours <= 0 {
		return ErrInvalidHourValue
	}
	d.hours += hours
	return nil
}

func (d *Duration) AddMinutes(minutes int) error {
	if minutes <= 0 {
		return ErrInvalidMinuteValue
	}
	d.minutes += minutes
	return nil
}

// DurationFrom TODO переписать
func DurationFrom(duration string) *Duration {
	return &Duration{str: duration}
}

func (d *Duration) String() string {
	return fmt.Sprintf("%dD %dH %dM", d.days, d.hours, d.minutes)
}
