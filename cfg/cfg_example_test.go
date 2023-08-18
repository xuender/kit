package cfg_test

import (
	"fmt"
	"os"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/xuender/kit/cfg"
)

func ExampleCfg_Bytes() {
	patchRead := gomonkey.ApplyFuncReturn(os.ReadFile, []byte(_data), nil)
	defer patchRead.Reset()

	patchWrite := gomonkey.ApplyFuncReturn(os.WriteFile, nil)
	defer patchWrite.Reset()

	_, err := cfg.New("key").Bytes("test.toml")
	fmt.Println(err)

	// Output:
	// <nil>
}

func ExampleCfg_Reader() {
	patchRead := gomonkey.ApplyFuncReturn(os.ReadFile, []byte(_data), nil)
	defer patchRead.Reset()

	patchWrite := gomonkey.ApplyFuncReturn(os.WriteFile, nil)
	defer patchWrite.Reset()

	_, err := cfg.New("key").Reader("test.toml")
	fmt.Println(err)

	// Output:
	// <nil>
}

func ExampleCfg_String() {
	patchRead := gomonkey.ApplyFuncReturn(os.ReadFile, []byte(_data), nil)
	defer patchRead.Reset()

	patchWrite := gomonkey.ApplyFuncReturn(os.WriteFile, nil)
	defer patchWrite.Reset()

	_, err := cfg.New("key").String("test.toml")
	fmt.Println(err)

	// Output:
	// <nil>
}

func ExampleCfg_Read() {
	data, err := cfg.New("key").Read([]byte(`a=AES(A/43wTj2AVQboZZ0lNMqbw==)
b=DES(LABOK5l6Q64=)
c=DES[abc]`))

	fmt.Println(string(data))
	fmt.Println(err)

	// Output:
	// a=aaa
	// b=test2
	// c=abc
	// <nil>
}
