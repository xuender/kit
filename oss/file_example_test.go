package oss_test

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xuender/kit/oss"
)

// ExampleAppendFile is an example function.
func ExampleAppendFile() {
	file, err := oss.CreateFile(filepath.Join(os.TempDir(), "go-cli", "create.txt"))
	fmt.Println(err)
	fmt.Println(file.Write([]byte("123")))
	file.Close()

	file, err = oss.AppendFile(filepath.Join(os.TempDir(), "go-cli", "create.txt"))
	fmt.Println(err)

	_, _ = file.Write([]byte("aaaa"))
	file.Close()

	data, _ := os.ReadFile(file.Name())
	fmt.Println(string(data))
	os.Remove(file.Name())

	// Output:
	// <nil>
	// 3 <nil>
	// <nil>
	// 123aaaa
}

// ExampleCreateFile is an example function.
func ExampleCreateFile() {
	file, err := oss.CreateFile(filepath.Join(os.TempDir(), "go-cli", "create.txt"))

	fmt.Println(err)
	file.Close()
	os.Remove(file.Name())

	// Output:
	// <nil>
}
