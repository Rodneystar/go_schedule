package clock

import (
	"fmt"
	"time"
)

const (
	minsInDay uint16 = 60 * 24
)

// Clock represents the time of day in minutes
type Clock struct {
	time uint16 // time of day in minutes since midnight
}

// Until returns minutes until laterClock, answer is always positive
// eg 23:50 Until 00:10 = 20
func (c *Clock) Until(laterClock Clock) time.Duration {
	start, until := int(c.time), int(laterClock.time)
	duration := until - start
	if duration < 0 {
		duration += int(minsInDay)
	}

	return time.Duration(duration) * time.Minute
}

// Equal returns true if the clock times are equal
func (c *Clock) Equal(c2 Clock) bool {
	return c.time == c2.time

}

// Add adds d duration to the clock, rolling back to the start if the result is past 23:59.
// modifies the supplied reference value.
func (c *Clock) Add(d time.Duration) *Clock {
	// mins := d / time.Minute
	sTime := c.time + uint16(d/time.Minute)
	if sTime > 1440 {
		sTime -= 1440
	}
	return &Clock{
		time: sTime,
	}
}

// Now returns a new clock with current time
func Now() *Clock {
	return NewClock(time.Now())
}

// Set sets the hours and minutes of this clock
func (c *Clock) Set(h uint8, m uint8) error {
	if (h > 23 || h < 0) ||
		(m > 59 || m < 0) {
		return fmt.Errorf("invalid clock values, %d:%d", h, m)
	}
	c.time = uint16(h*60 + m)
	return nil
}

// GetHoursMins returns 2 int values, the hours and minutes
func (c *Clock) GetHoursMins() (hours uint8, minutes uint8) {
	hours = uint8(c.time / 60)
	minutes = uint8(c.time % 60)
	return
}

// GetMins returns the time of clock in minutes
func (c *Clock) GetMins() uint16 {
	return c.time
}

// NewClock clock from a time.Time, seconds always rounded down
func NewClock(t time.Time) *Clock {
	h, m, _ := t.Clock()
	return &Clock{
		time: uint16(h*60 + m),
	}
}
