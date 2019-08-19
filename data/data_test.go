package data

import (
	"io/ioutil"
	"schedule/clock"
	"testing"
	"time"
)

var now = clock.NewClock(time.Now())

var events = []TimerActiveSpan{
	TimerActiveSpan{
		AtTime:   *now,
		Duration: time.Second * 2,
	}, TimerActiveSpan{
		AtTime:   *now.Add(-time.Hour * 2),
		Duration: time.Second * 2,
	},
}

func Test_Add(t *testing.T) {
	data := NewDataAccess()
	data.AddTimer(events[0])
	persistedEvents, _ := data.GetAllTimers()
	if noOfEvents := len(persistedEvents); noOfEvents != 1 {
		t.Errorf("expected length 1, was: %d", noOfEvents)
	}
}

func Test_delall(t *testing.T) {
	data := NewDataAccess()
	t.Error()
	data.AddTimer(events[0])
	data.AddTimer(events[1])
	data.DelAll()
	persistedEvents, _ := data.GetAllTimers()
	if noOfEvents := len(persistedEvents); noOfEvents != 0 {
		t.Errorf("expected length 1, was: %d", noOfEvents)
	}
}
func Test_getAllTimers(t *testing.T) {

	data := NewDataAccess()
	data.DelAll()
	data.AddTimer(events[0])

	allTimers, _ := data.GetAllTimers()
	result := allTimers[0].AtTime
	if !result.Equal(now) {
		t.Errorf("expected %s, was: %s", now, result)
	}

	data.AddTimer(events[1])
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
	event := TimerActiveSpan{
		AtTime:   now,
		Duration: time.Second * 2,
	}
	arr := []TimerActiveSpan{event}

	jsonEvent, _ := toJSONArr(arr)
	ioutil.WriteFile("timers.dat", jsonEvent, 0644)

}

func Test_readFiles(t *testing.T) {
	jsonIn, err := ioutil.ReadFile("timers.dat")
	if err != nil {
		t.Error("read error, ", err)
	}

	timerEvents := fromJSONArr(jsonIn)
	t.Logf("time 1: %s, action1: %s", timerEvents[0].AtTime, timerEvents[0].Duration.String())

}

func Test_Remove(t *testing.T) {
	data := NewDataAccess()
	data.DelAll()
	data.AddTimer(events[0])
	data.AddTimer(events[1])
	data.RemoveByAtTime(events[0].AtTime)

	timers, _ := data.GetAllTimers()
	length := len(timers)
	if length != 1 {
		t.Errorf("expected length 1, was %d", length)
	}

}
