package runat

import (
	"fmt"
	"schedule/clock"
	"testing"
	"time"
)

func Test_ScheduleAt(t *testing.T) {
	count := 0
	_ = ScheduleAt(time.Now().Add(time.Second*3), func() {
		fmt.Println("action happening")
		count++
	})

	time.Sleep(2 * time.Second)
	if count != 0 {
		t.Error("should be 0, is: ", count)
	}

	//stop <- true

	time.Sleep(2 * time.Second)
	if count != 1 {
		t.Error("should be 1, is: ", count)
	}
	time.Sleep(2 * time.Second)
}

func Test_ScheduleEveryAfterDelay(t *testing.T) {
	fmt.Println("starting --")
	stop := ScheduleEveryAfterDelay(time.Second*5, time.Second, func() {
		fmt.Println("we are running ")
	})
	fmt.Println("sent the stop --")
	stop <- true
	time.Sleep(8 * time.Second)
	time.Sleep(10 * time.Second)
}

func Test_ScheduleEveryAfterDelayMultiple(t *testing.T) {
	fmt.Println("starting --")
	stop := ScheduleEveryAfterDelay(time.Millisecond*500, time.Second, func() {
		fmt.Println("we are running 1st")
	})

	stop2 := ScheduleEveryAfterDelay(time.Millisecond*1000, time.Second, func() {
		fmt.Println("we are running 2nd")

	})
	time.Sleep(4 * time.Second)
	stop <- true

	time.Sleep(3950 * time.Millisecond)
	fmt.Println("sent the stop --")

	stop2 <- true

	time.Sleep(5 * time.Second)
}

func Test_ScheduleEvery(t *testing.T) {
	count := 0
	fmt.Println("starting")
	stop := ScheduleEvery(1*time.Second, func() {
		count++
		fmt.Println("actiosn: ", count)

	})
	time.Sleep(time.Millisecond * 800)
	if count != 0 {
		t.Error("0.8 seconds, count: ", count)
	}

	time.Sleep(time.Second)
	if count != 1 {
		t.Error("1.8 seconds, count: ", count)
	}

	time.Sleep(time.Second * 3)
	if count != 4 {
		t.Error("4.8 seconds, count: ", count)
	}

	time.Sleep(time.Second * 4)
	fmt.Println("sending stop false")
	stop <- false
	time.Sleep(6 * time.Second)
}

func Test_ticker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	ticker.Stop()
	count := 0
	go func() {
		for {
			count++
			tick := <-ticker.C
			if count >= 10 {
				ticker.Stop()
			}
			fmt.Printf("ticker val : %s", tick)
		}
	}()

	time.Sleep(10 * time.Second)
}

func Test_EachDayAt(t *testing.T) {
	now := clock.Now()
	hours, mins := now.GetHoursMins()
	now.Set(hours, mins+1)
	cancel := EachDayAt(*now, func() { fmt.Println("ticked") })

	time.Sleep(65 * time.Second)
	cancel <- true

}

func Test_wrapping(t *testing.T) {

}
