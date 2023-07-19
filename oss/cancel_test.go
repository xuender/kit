package oss_test

import (
	"context"
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/xuender/kit/oss"
)

func ExampleCancel() {
	cancel := func() {
		fmt.Println("cancel")
	}

	go oss.Cancel(cancel)

	time.Sleep(time.Millisecond * 5)

	_ = syscall.Kill(os.Getpid(), syscall.SIGINT)

	time.Sleep(time.Millisecond * 5)

	// Output:
	// cancel
}

func ExampleCancelContext() {
	ctx, _ := oss.CancelContext(context.Background())

	go func() {
		<-ctx.Done()

		fmt.Println("cancel")
	}()

	time.Sleep(time.Millisecond * 5)

	_ = syscall.Kill(os.Getpid(), syscall.SIGINT)

	time.Sleep(time.Millisecond * 5)

	// Output:
	// cancel
}
