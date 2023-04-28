package logs_test

import (
	"os"

	"github.com/xuender/kit/logs"
)

func Example() {
	// logs.SetLogFile("/var/tmp", "test.log")
	logs.SetLog(os.Stdout)
	logs.SetLevel(logs.Info)
	logs.I.SetFlags(0)

	logs.T.Println("trace")
	logs.D.Printf("debug: %d", 1)
	logs.I.Println("info")

	// Output:
	// [I] info
}

func ExampleLog() {
	logs.SetLog(os.Stdout)
	logs.I.SetFlags(0)
	logs.E.SetFlags(0)

	logs.Log(nil)
	logs.Log(1)
	logs.Log(2, os.ErrClosed)

	// Output:
	// [I] 1
	// [E] 2 file already closed
}
