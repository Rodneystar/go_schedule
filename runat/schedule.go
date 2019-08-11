package runat

import (
	"time"
)

func ScheduleAt(atTime time.Time, action func()) chan<- bool {
	until := time.Until(atTime)
	tick := make(chan time.Time)
	stop := make(chan bool)
	go func() {
		time.Sleep(until)
		tick <- time.Now()
	}()
	go func() {
		for {
			select {
			case <-tick:
				action()
				close(tick)
				return
			case <-stop:
				close(stop)
				close(tick)
				return
			}
		}
	}()
	return stop
}

func ScheduleEvery(duration time.Duration, action func()) chan<- bool {
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

func ScheduleEveryAfterDelay(initDelay time.Duration, period time.Duration, action func()) chan<- bool {
	var ticker *time.Ticker
	stop := make(chan bool)

	go func() {
		time.Sleep(initDelay)
		ticker = time.NewTicker(period)
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
	}()
	return stop
}
