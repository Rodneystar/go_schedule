package runat

import (
	"fmt"
	"time"
)

func ScheduleAt(atTime time.Time) <-chan bool {
	until := time.Until(atTime)
	back := make(chan bool)
	go func() {
		time.Sleep(until)
		fmt.Printf("we slept for a while")
		back <- true
	}()
	return back
}

func ScheduleEvery(duration time.Duration, action func()) (chan<- bool) {
	ticker := time.NewTicker(duration)
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				ticker.Stop()
				close(stop)
				return
			case <-ticker.C:
				action()
			}
		}
	}()
	return stop
}
