package main

import (
	"time"

	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/times"
)

func SetLogFile() {
	rotate(SetLogFile)
}

func rotate(yield func()) {
	times.Hour(Rotating(yield))
}

func Rotating(yield func()) func() {
	logs.Log("Rotating")

	return func() {
		logs.Log("bbb Rotating")
		yield()
	}
}

func main() {
	logs.I.Println("Start")

	SetLogFile()

	time.Sleep(time.Minute)
	logs.I.Println("end")
}

func main1() {
	logs.I.Println("start")

	cancel := times.Hour(func() {
		logs.I.Println("run")
	})

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
		logs.I.Println("stop")
	}()

	logs.I.Println("sleep")
	time.Sleep(5 * time.Second)
	logs.I.Println("end")
}
