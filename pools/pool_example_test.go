package pools_test

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/samber/lo"
	"github.com/xuender/kit/pools"
)

func Example() {
	pool := pools.New(10, func(value, num int) string {
		time.Sleep(time.Millisecond)

		return fmt.Sprintf("%d: %d*2=%d", num, value, value*2)
	})
	defer pool.Close()

	outputs := pool.Post(lo.Range(100))

	fmt.Println(len(outputs))

	// Output:
	// 100
}

func ExamplePool_context() {
	pool := pools.New(10, func(input lo.Tuple2[context.Context, int], _ int) int {
		time.Sleep(time.Millisecond)

		return input.B * input.B
	})
	defer pool.Close()

	inputs := lo.Map(lo.Range(100), func(num, _ int) lo.Tuple2[context.Context, int] {
		return lo.T2(context.Background(), num)
	})
	outputs := pool.Post(inputs)

	fmt.Println(len(outputs))

	// Output:
	// 100
}

func ExamplePool_Run() {
	pool := pools.New(10, func(value, _ int) string {
		time.Sleep(time.Millisecond)

		return fmt.Sprintf("%d*2=%d", value, value*2)
	})
	defer pool.Close()

	fmt.Println(pool.Run(3))

	// Output:
	// 3*2=6
}

func ExamplePool_error() {
	pool := pools.New(10, func(value, _ int) lo.Tuple2[int, error] {
		time.Sleep(time.Millisecond)

		if value == 0 {
			// nolint
			return lo.T2(0, errors.New("divide by zero"))
		}

		return lo.T2[int, error](100/value, nil)
	})
	defer pool.Close()

	outputs := pool.Post(lo.Range(100))

	fmt.Println(len(outputs))

	// Output:
	// 100
}
