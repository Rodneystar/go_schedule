package data

import (
	"io/ioutil"

	"testing"
	"time"
)

func Test_getAllTimers(t *testing.T) {
	now := time.Now()
	event := TimerEventDescription{
		AtTime: now,
		State:  true,
	}

	event2 := TimerEventDescription{
		AtTime: now.Add(-time.Hour * 2),
		State:  false,
	}
	events := []TimerEventDescription{event, event2}
	_ = events
	data := NewDataAccess()

	data.AddTimer(event)
	allTimers, _ := data.GetAllTimers()
	result := allTimers[0].AtTime
	if result != now {
		t.Errorf("expected %s, was: %s", now, result)
	}

}

func Test_files(t *testing.T) {
	now := time.Now()
	event := TimerEventDescription{
		AtTime: now,
		State:  true,
	}
	arr := []TimerEventDescription{event}

	jsonEvent, _ := toJSONArr(arr)
	ioutil.WriteFile("timers.dat", jsonEvent, 0644)

}

func Test_readFiles(t *testing.T) {
	jsonIn, err := ioutil.ReadFile("timers.dat")
	if err != nil {
		t.Error("read error, ", err)
	}

	timerEvents := fromJSONArr(jsonIn)
	t.Logf("time 1: %s, action1: %t", timerEvents[0].AtTime, timerEvents[0].State)

}
