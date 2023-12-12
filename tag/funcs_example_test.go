// nolint: dupword
package tag_test

import (
	"fmt"

	"github.com/xuender/kit/tag"
)

// ExampleGet 获取例子.
func ExampleGet() {
	num := tag.Tag[int](1, 2)
	nums := tag.Get(num)

	fmt.Println(nums)

	// Output:
	// [1 2]
}

// ExampleAdd 增加例子.
func ExampleAdd() {
	num := tag.Add(0, 3, 4)

	fmt.Println(num)

	// Output:
	// 24
}

// ExampleDel 删除例子.
func ExampleDel() {
	num := tag.Del(tag.Tag[int](3, 4), 3, 5)

	fmt.Println(num)

	// Output:
	// 16
}

// ExampleHas 包含例子.
func ExampleHas() {
	num := tag.Tag[int](3, 4)

	fmt.Println(tag.Has(num, 3))
	fmt.Println(tag.Has(num, 3, 4))
	fmt.Println(tag.Has(num, 2))
	fmt.Println(tag.Has(num, 3, 2))

	// Output:
	// true
	// true
	// false
	// true
}

// ExampleHit 命中例子.
func ExampleHit() {
	num := tag.Tag[int](3, 4)

	fmt.Println(tag.Hit(num, 3))
	fmt.Println(tag.Hit(num, 3, 4))
	fmt.Println(tag.Hit(num, 2))
	fmt.Println(tag.Hit(num, 2, 3, 4))

	// Output:
	// true
	// true
	// false
	// false
}
