package data

import (
	"io/ioutil"

	"testing"
	"time"
)

var now = time.Now()
var events = []TimerEventDescription{
	TimerEventDescription{
		AtTime: now,
		State:  true,
	}, TimerEventDescription{
		AtTime: now.Add(-time.Hour * 2),
		State:  false,
	},
}

func Test_Add(t *testing.T) {
	data
}
func Test_delall(t *testing.T) {

}
func Test_getAllTimers(t *testing.T) {

	now := time.Now()

	_ = events
	data := NewDataAccess()
	data.DelAll()

	data.AddTimer(event)

	allTimers, _ := data.GetAllTimers()
	result := allTimers[0].AtTime
	if !result.Equal(now) {
		t.Errorf("expected %s, was: %s", now, result)
	}

	data.AddTimer(event2)
	allTimers, _ = data.GetAllTimers()
	length := len(allTimers)
	if length != 2 {
		t.Errorf("expected length 2, was %d", length)
	}

	allTimers, _ = data.GetAllTimers()
	result = allTimers[1].AtTime
	if !result.Equal(now.Add(-time.Hour * 2)) {
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
