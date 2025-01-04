package commonTime

import "time"

type Time struct {
	time.Time
}

func (t *Time) Local() *Time {
	return &Time{time.Unix(0, t.UnixNano()).Local()}
}

func (t *Time) Add(duration time.Duration) *Time {
	newTime := t.Time.Add(duration)
	return &Time{newTime}
}

func (t *Time) Sub(duration time.Duration) *Time {
	newTime := t.Time.Add(-duration)
	return &Time{newTime}
}

func (t *Time) Equal(time *Time) bool {
	return t.Time.Equal(time.Time)
}

func (t *Time) Before(time *Time) bool {
	return t.Time.Before(time.Time)
}

func (t *Time) After(time *Time) bool {
	return t.Time.After(time.Time)
}

func (t *Time) Unix() int64 {
	var unix int64
	if t != nil && !t.IsZero() {
		unix = t.Time.Unix()
	}
	return unix
}

func (t *Time) UnixMilli() int64 {
	var unixMilli int64
	if t != nil && !t.IsZero() {
		unixMilli = t.Time.UnixMilli()
	}
	return unixMilli
}

func (t *Time) UnixNano() int64 {
	var unixNano int64
	if t != nil && !t.IsZero() {
		unixNano = t.Time.UnixNano()
	}
	return unixNano
}

func FromUnixNano(nanoseconds int64) *Time {
	if nanoseconds == 0 {
		return Empty()
	}
	return &Time{time.Unix(0, nanoseconds).UTC()}
}

func FromUnixMillis(milliseconds int64) *Time {
	if milliseconds == 0 {
		return Empty()
	}
	return &Time{time.UnixMilli(milliseconds).UTC()}
}

func Now() *Time {
	return &Time{time.Now().UTC()}
}

func Empty() *Time {
	return &Time{}
}

func Parse(layout, value string) (*Time, error) {
	parsedValue, err := time.Parse(layout, value)
	if err != nil {
		return nil, err
	}
	return FromTime(parsedValue), nil
}

func FromTime(t time.Time) *Time {
	return &Time{time.Unix(0, t.UnixNano()).UTC()}
}
