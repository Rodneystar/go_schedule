package main

import (
	"fmt"
	"go_schedule/data"
	"go_schedule/physical"
	"time"
)

func log(s string, args ...interface{}) {
	fmt.Printf("%s " + s, time.Now(), args, )
}

type TimerMode int8

const (
	ModeOff		TimerMode = 0 + iota
	ModeOn
	ModeTimed
)

type TimerService struct {
	timers                   []data.TimerActiveSpan
	runningTimerStopChannels []chan bool
	db                       data.DataAccess
	mode                     TimerMode
	switchable 				 physical.Switchable
}


func(t *TimerService) cycleTimers() {
	t.stopTimers()
	t.startTimers()
}

func (t *TimerService) startTimers() {
	timers, err := t.db.GetAllTimers()
	if err != nil {
		log("nothing")
	}
	t.timers = timers
}

func(t *TimerService) stopTimers() {
	for _, c := range t.runningTimerStopChannels {
		c <- true
	}
}

func(t *TimerService) SetMode(mode TimerMode) {
	switch mode{
		case ModeOff:
			t.off()
		case ModeOn:
			t.on()
		case ModeTimed:
			t.timed()
	}
}

func(t *TimerService) AddTimer(startTime time.Time, duration time.Duration) error {
	event := data.TimerActiveSpan{
		AtTime: startTime,
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

func(t *TimerService) RemoveByAtTime( ti time.Time) {
	t.db.RemoveByAtTime(ti)
}
func(t *TimerService) GetAllTimers() []data.TimerActiveSpan {
	timers, err := t.db.GetAllTimers()
	if err != nil {
		log("error getting all timers")
	}
	return timers
}

func(t *TimerService) Init(db data.DataAccess, switchable physical.Switchable) {
	t.db, t.mode, t.switchable = db, ModeOff, switchable

	timers, err := db.GetAllTimers()
	if err != nil {
		log("error retrieving timer list from db, initializing empty: %s", err)
		t.timers = make([]data.TimerActiveSpan, 0 )
	}
	t.timers = timers
}

func (t *TimerService) off() {
	//turn the relay off
}

func (t *TimerService) on() {
	//on
}

func (t *TimerService) timed() {

	//timed
}



func main() {

}