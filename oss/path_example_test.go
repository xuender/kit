package oss_test

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xuender/kit/v2/ops"
	"github.com/xuender/kit/v2/oss"
)

func ExampleAbs() {
	home, _ := os.UserHomeDir()
	fmt.Println(ops.Must(oss.Abs("~/file")) == filepath.Join(home, "file"))

	// Output:
	// true
}
