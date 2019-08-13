package main

import (
	"fmt"
	"go_schedule/data"
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

type Switchable interface {
	On()
	Off()
	State()
}
type TimerService struct {
	timers                   []data.TimerEventDescription
	runningTimerStopChannels []chan bool
	db                       data.DataAccess
	mode                     TimerMode
	switchable 				 Switchable
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

func(t *TimerService) Init(db data.DataAccess, switchable Switchable) {
	t.db, t.mode, t.switchable = db, ModeOff, switchable

	timers, err := db.GetAllTimers()
	if err != nil {
		log("error retrieving timer list from db, initializing empty: %s", err)
		t.timers = make([]data.TimerEventDescription, 0 )
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