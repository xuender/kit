package logs_test

import (
	"github.com/samber/lo"
	"github.com/xuender/kit/logs"
)

func ExampleFile() {
	writer := lo.Must1(logs.File("/tmp/log", "link.log"))
	defer logs.Close()

	lo.Must1(writer.Write([]byte("xxx")))
	// Output:
}
