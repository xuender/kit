package times_test

import (
	"time"

	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/times"
)

func ExampleHour() {
	logs.I.Println("start")

	cancel := times.Hour(func() {
		logs.I.Println("run")
	})

	go func() {
		time.Sleep(time.Second)
		cancel()
		logs.I.Println("stop")
	}()

	logs.I.Println("sleep")
	time.Sleep(2 * time.Second)
	logs.I.Println("end")

	// Output:
}
