package main

import (
	"context"
	"time"

	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/oss"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	group := errgroup.Group{}
	group.Go(sleep(4))
	group.Go(sleep(7))
	group.Go(func() error {
		logs.D.Println("ctx")
		<-ctx.Done()
		logs.D.Println("ctx")
		return nil
	})
	group.Go(oss.Cancel(cancel))

	logs.D.Println("xxx")
	group.Wait()
	logs.D.Println("xxx")
}

func sleep(num int) func() error {
	return func() error {
		logs.D.Println(num)
		time.Sleep(time.Second * time.Duration(num))
		logs.D.Println(num)

		return nil
	}
}
