package oss_test

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xuender/kit/base"
	"github.com/xuender/kit/oss"
)

func ExampleAbs() {
	home, _ := os.UserHomeDir()
	fmt.Println(base.Must1(oss.Abs("~/file")) == filepath.Join(home, "file"))

	// Output:
	// true
}
