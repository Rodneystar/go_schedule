package runat

import (
	"fmt"
	"testing"
	"time"
)

func Test_ScheduleAt(t *testing.T) {
	future := ScheduleAt(time.Now().Add(time.Second * 5))
	t.Logf("ScheduleAt(now ): %t", <-future)
}

func Test_ScheduleEvery(t *testing.T) {
	count := 0
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
