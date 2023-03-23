package main

import (
	"time"

	"github.com/samber/lo"
	"github.com/xuender/kit/logs"
)

func main() {
	lo.Must0(logs.SetLogFile("/tmp/log", "log.log"))
	defer logs.Close()

	logs.I.Println("start")
	go run()
	time.Sleep(time.Hour)
}

func main1() {
	printLog("default")

	logs.SetLevel(logs.Error)
	printLog("error")

	logs.SetLevel(logs.Warn)
	printLog("warn")

	logs.SetLevel(logs.Info)
	printLog("info")

	logs.SetLevel(logs.Debug)
	printLog("debug")

	logs.SetLevel(logs.Trace)
	printLog("trace")

	lo.Must0(logs.SetLogFile("/tmp/log", "test.log"))
	lo.Must0(logs.SetErrorFile("/tmp/log", "test_error.log"))

	defer logs.Close()

	printLog("log")

	go run()
	time.Sleep(3 * time.Hour)
}

func run() {
	for range time.NewTicker(time.Second).C {
		logs.E.Printf("time: %v", time.Now())
	}
}

func printLog(msg string) {
	logs.T.Println(msg)
	logs.D.Println(msg)
	logs.I.Println(msg)
	logs.W.Println(msg)
	logs.E.Println(msg)
}
