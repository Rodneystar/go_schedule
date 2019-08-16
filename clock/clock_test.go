package clock

import (
	"fmt"
	"testing"
	"time"
)

func Test_add(t *testing.T) {
	now := time.Now()
	c1 := NewClock(now)
	c2 := NewClock(now.Add(15 * time.Hour))
	duration := 15 * time.Hour

	c1.Add(duration)
	expected := c2
	fmt.Printf("expected: %v, result: %v\n", expected.GetMins(), c1.GetMins())
	if !expected.Equals(*c1) {
		fmt.Printf("expected: %v, result: %v\n", expected.GetMins(), c1.GetMins())
	}

}

func Test_SetGet(t *testing.T) {
	timeNow := time.Now().Add(16 * time.Hour)
	clock := NewClock(timeNow)

	h, m, _ := timeNow.Clock()
	clockH, clockM := clock.GetHoursMins()
	if h != int(clockH) {
		t.Errorf("expected: %d, result: %d\n", h, clockH)
	}
	if m != int(clockM) {
		t.Errorf("expected: %d, result: %d\n", m, clockM)
	}
}

func Test_Until(t *testing.T) {
	timeNow := time.Now()
	clockNow := NewClock(timeNow)

	duration := 83
	laterTime := time.Now().Add(time.Duration(duration) * time.Minute)
	clockLater := NewClock(laterTime)

	expected := time.Duration(duration)
	result := clockNow.Until(*clockLater)
	t.Log(result)

	if expected != result {
		t.Errorf("expected: %d, was: %d", expected, result)
	}

}
func Test_UntilLarger(t *testing.T) {
	timeNow := time.Now()
	clockNow := NewClock(timeNow)

	duration := minsInDay - 5
	laterTime := time.Now().Add(time.Duration(duration) * time.Minute)
	clockLater := NewClock(laterTime)

	expected := time.Duration(duration)
	result := clockNow.Until(*clockLater)
	t.Log(result)
	if expected != result {
		t.Errorf("expected: %d, was: %d", expected, result)
	}

}
