package main

import (
	"go_schedule/data"
)

type TimerService struct {
	timers 						[]data.TimerEventDescription
	runningTimerStopChannels 	[]chan bool
}



func main() {

}