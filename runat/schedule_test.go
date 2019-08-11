package runat

import (
	"fmt"
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

	time.Sleep(2 * time.Second)
	if count != 1 {
		t.Error("should be 1, is: ", count)
	}
	time.Sleep(2 * time.Second)
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
	fmt.Println("not sending stop false")
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
