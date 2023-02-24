package logs_test

import (
	"os"

	"github.com/xuender/kit/logs"
)

func Example() {
	// logs.SetLogFile("/var/tmp", "test")
	logs.SetLog(os.Stdout)
	logs.SetLevel(logs.Info)

	logs.T.Println("trace")
	logs.D.Printf("debug: %d", 1)

	// Output:
}
