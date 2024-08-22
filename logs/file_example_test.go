package logs_test

import (
	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/los"
)

func ExampleFile() {
	writer := los.Must1(logs.File("/tmp/log", "link.log"))
	defer logs.Close()

	los.Must(writer.Write([]byte("xxx")))
	// Output:
}
