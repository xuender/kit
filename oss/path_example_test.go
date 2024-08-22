package oss_test

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xuender/kit/los"
	"github.com/xuender/kit/oss"
)

func ExampleAbs() {
	home, _ := os.UserHomeDir()
	fmt.Println(los.Must(oss.Abs("~/file")) == filepath.Join(home, "file"))

	// Output:
	// true
}
