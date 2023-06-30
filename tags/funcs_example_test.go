package tags_test

import (
	"fmt"

	"github.com/xuender/kit/tags"
)

// ExampleAdd 增加例子.
func ExampleAdd() {
	tag := tags.Add(1<<3, 1<<4)

	fmt.Println(tag)

	// Output:
	// 24
}

// ExampleDel 删除例子.
func ExampleDel() {
	tag := tags.Add(1<<3, 1<<4)
	tag = tags.Del(tag, 1<<3, 1<<5)

	fmt.Println(tag)

	// Output:
	// 16
}

// ExampleHas 包含例子.
func ExampleHas() {
	tag := tags.Add(1<<3, 1<<4)
	query1 := tags.Add(1 << 3)
	query2 := tags.Add(1<<3, 1<<4)
	query3 := tags.Add(1 << 2)

	fmt.Println(tags.Has(tag, query1))
	fmt.Println(tags.Has(tag, query1, query2))
	fmt.Println(tags.Has(tag, query3))
	fmt.Println(tags.Has(tag, query3, query2))

	// Output:
	// true
	// true
	// false
	// true
}

// ExampleHit 命中例子.
func ExampleHit() {
	tag := tags.Add(1<<3, 1<<4)
	query1 := tags.Add(1 << 3)
	query2 := tags.Add(1<<3, 1<<4)
	query3 := tags.Add(1 << 2)

	fmt.Println(tags.Hit(tag, query1))
	fmt.Println(tags.Hit(tag, query1, query2))
	fmt.Println(tags.Hit(tag, query3))
	fmt.Println(tags.Hit(tag, query3, query2))

	// Output:
	// true
	// true
	// false
	// false
}
