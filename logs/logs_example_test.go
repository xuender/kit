package logs_test

import (
	"os"

	"github.com/xuender/kit/logs"
)

func Example() {
	// logs.SetLogFile("/var/tmp", "test")
	logs.SetLog(os.Stdout)
	logs.SetLevel(logs.Warn)

	logs.T.Print("trace")
	logs.D.Printf("debug: %d", 1)
	logs.I.Println("info")

	// Output:
}
