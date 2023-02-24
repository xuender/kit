package logs_test

import (
	"github.com/xuender/kit/base"
	"github.com/xuender/kit/logs"
)

func ExampleFile() {
	writer := base.Must1(logs.File("/tmp/log", "link.log"))
	defer logs.Close()

	base.Must1(writer.Write([]byte("xxx")))
	// Output:
}
