// nolint: dupword
package tag_test

import (
	"fmt"

	"github.com/xuender/kit/tag"
)

// ExampleAdd 增加例子.
func ExampleAdd() {
	num := tag.Add(1<<3, 1<<4)

	fmt.Println(num)

	// Output:
	// 24
}

// ExampleDel 删除例子.
func ExampleDel() {
	num := tag.Add(1<<3, 1<<4)
	num = tag.Del(num, 1<<3, 1<<5)

	fmt.Println(num)

	// Output:
	// 16
}

// ExampleHas 包含例子.
func ExampleHas() {
	num := tag.Add(1<<3, 1<<4)
	query1 := tag.Add(1 << 3)
	query2 := tag.Add(1<<3, 1<<4)
	query3 := tag.Add(1 << 2)

	fmt.Println(tag.Has(num, query1))
	fmt.Println(tag.Has(num, query1, query2))
	fmt.Println(tag.Has(num, query3))
	fmt.Println(tag.Has(num, query3, query2))

	// Output:
	// true
	// true
	// false
	// true
}

// ExampleHit 命中例子.
func ExampleHit() {
	num := tag.Add(1<<3, 1<<4)
	query1 := tag.Add(1 << 3)
	query2 := tag.Add(1<<3, 1<<4)
	query3 := tag.Add(1 << 2)

	fmt.Println(tag.Hit(num, query1))
	fmt.Println(tag.Hit(num, query1, query2))
	fmt.Println(tag.Hit(num, query3))
	fmt.Println(tag.Hit(num, query3, query2))

	// Output:
	// true
	// true
	// false
	// false
}
