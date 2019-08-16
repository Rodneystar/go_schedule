package runat

import (
	"schedule/clock"
	"time"
)

func ScheduleAt(atTime time.Time, action func()) chan<- bool {
	timer, stop := getInitDelayChan(time.Until(atTime))
	go func() {
		for {
			select {
			case <-timer.C:
				action()
				return
			case <-stop:
				close(stop)
				timer.Stop()
				return
			}
		}
	}()
	return stop
}

func ScheduleEvery(duration time.Duration, action func()) chan<- bool {
	return scheduleEveryAfterDelay(0, duration, action)
}

func ScheduleEveryAfterDelay(initDelay time.Duration, period time.Duration, action func()) chan<- bool {
	return scheduleEveryAfterDelay(initDelay, period, action)
}

func EachDayAt(t clock.Clock, action func()) chan<- bool {
	now := clock.NewClock(time.Now())
	return scheduleEveryAfterDelay(now.Until(t), time.Hour*24, action)
}

func scheduleEveryAfterDelay(initDelay time.Duration, period time.Duration, action func()) chan<- bool {
	timer, stop := getInitDelayChan(initDelay)
	go func() {
		for {
			select {
			case <-stop:
				timer.Stop()
				close(stop)
				return
			case <-timer.C:
				go func() {
					ticker := time.NewTicker(period)
					for {
						select {
						case <-ticker.C:
							action()
						case <-stop:
							ticker.Stop()
							return
						}
					}
				}()

			}

		}
	}()
	return stop
}

func getInitDelayChan(delay time.Duration) (*time.Timer, chan bool) {
	alarmRinging := time.NewTimer(delay)
	stop := make(chan bool)

	return alarmRinging, stop
}
