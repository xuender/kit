package main

import (
	"github.com/xuender/kit/base"
	"github.com/xuender/kit/logs"
)

func main() {
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
	base.Must(logs.SetLogFile("/tmp/log", "test"))
	base.Must(logs.SetErrorFile("/tmp/log", "test_error"))

	defer logs.Close()

	printLog("log")
}

func printLog(msg string) {
	logs.T.Println(msg)
	logs.D.Println(msg)
	logs.I.Println(msg)
	logs.W.Println(msg)
	logs.E.Println(msg)
}
