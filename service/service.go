package service

import (
	"fmt"
	"schedule/clock"
	"schedule/data"
	"schedule/physical"
	"schedule/runat"
	"time"
)

func log(s string, args ...interface{}) {
	fmt.Printf("%s "+s, time.Now(), args)
}

type TimerMode int8

const (
	ModeOff TimerMode = 0 + iota
	ModeOn
	ModeTimed
)

type TimerEvent struct {
	data.TimerActiveSpan
}
type TimerService struct {
	// timers                   []TimerEvent
	runningTimerStopChannels []chan<- bool
	db                       data.DataAccess
	mode                     TimerMode
	switchable               physical.Switchable
}

func (t *TimerService) cycleTimers() {
	t.stopTimers()
	t.startTimers()
}

func (t *TimerService) startTimers() {
	timers, err := t.db.GetAllTimers()
	if err != nil {
		log("nothing")
	}

	i := 0
	for _, e := range timers {
		stopTime := *e.AtTime.Add(e.Duration)
		t.runningTimerStopChannels[i] = runat.EachDayAt(e.AtTime, t.switchable.On)
		t.runningTimerStopChannels[i+1] = runat.EachDayAt(stopTime, t.switchable.Off)
		i += 2
	}

}

func (t *TimerService) stopTimers() {
	for _, c := range t.runningTimerStopChannels {
		c <- true
	}
}

func (t *TimerService) SetMode(mode TimerMode) {
	switch mode {
	case ModeOff:
		t.off()
	case ModeOn:
		t.on()
	case ModeTimed:
		t.timed()
	}
}

func (t *TimerService) AddTimer(startTime clock.Clock, duration time.Duration) error {
	event := data.TimerActiveSpan{
		AtTime:   startTime,
		Duration: duration,
	}
	err := t.db.AddTimer(event)
	if err != nil {
		return err
	}

	if t.mode == ModeTimed {
		t.cycleTimers()
	}
	return nil
}

func (t *TimerService) RemoveByAtTime(ti clock.Clock) {
	t.db.RemoveByAtTime(ti)
}
func (t *TimerService) GetAllTimers() []data.TimerActiveSpan {
	timers, err := t.db.GetAllTimers()
	if err != nil {
		log("error getting all timers: %s", err)
	}
	return timers
}

func (t *TimerService) Init(db data.DataAccess, switchable physical.Switchable) {
	t.db, t.mode, t.switchable = db, ModeOff, switchable
}

func (t *TimerService) off() {
	//turn the relay off
	t.switchable.Off()
}

func (t *TimerService) on() {
	//on
}

func (t *TimerService) timed() {
	t.startTimers()
	//timed
}
