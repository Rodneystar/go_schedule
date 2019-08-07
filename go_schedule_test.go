package go_schedule

import (
	"testing"
	"time"
)

func Test_ScheduleAt(t *testing.T) {
	t.Logf("ScheduleAt(now ): %s", ScheduleAt(time.Now()))
}
