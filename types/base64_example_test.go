package types_test

import (
	"fmt"

	"github.com/xuender/kit/types"
)

func ExampleNumToB64() {
	fmt.Println(types.NumToB64(3))
	fmt.Println(types.NumToB64(0))
	fmt.Println(types.NumToB64(-40000000000))
	fmt.Println(types.NumToB64(6284325805349880))

	// Output:
	// D
	// A
	// -lQL5AA
	// WU4!fR9v4
}

func ExampleB64ToNum() {
	fmt.Println(types.B64ToNum[int]("D"))
	fmt.Println(types.B64ToNum[int]("A"))
	fmt.Println(types.B64ToNum[int]("-lQL5AA"))
	fmt.Println(types.B64ToNum[int]("WU4!fR9v4"))

	// Output:
	// 3
	// 0
	// -40000000000
	// 6284325805349880
}
