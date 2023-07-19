package oss

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// nolint: gochecknoglobals
var _cancelSignals = []os.Signal{
	// 挂起 终端连接断开
	syscall.SIGHUP,
	// 中断 Ctrl+C
	syscall.SIGINT,
	// 结束
	syscall.SIGTERM,
	// 退出 (Ctrl+/)
	syscall.SIGQUIT,
	// 杀死
	syscall.SIGKILL,
}

func Cancel(cancel context.CancelFunc) {
	osc := make(chan os.Signal, 1)

	signal.Notify(osc, _cancelSignals...)

	<-osc
	cancel()
}

func CancelFunc(cancel context.CancelFunc) func() error {
	return func() error {
		Cancel(cancel)

		return nil
	}
}

func CancelContext(ctx context.Context) (context.Context, context.CancelFunc) {
	return signal.NotifyContext(context.Background(), _cancelSignals...)
}
