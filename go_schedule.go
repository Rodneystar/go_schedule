package go_schedule

import "time"

func ScheduleAt(atTime time.Time) string {

	outString := atTime.Format("15:04")
	return outString
}
