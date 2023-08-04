package times_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/base"
	"github.com/xuender/kit/times"
)

func TestSleep(t *testing.T) {
	t.Parallel()

	assert.Greater(t, times.Sleep(1502), time.Second)
	assert.Greater(t, times.Sleep(502), time.Second)
}

// nolint: paralleltest
func TestInScope(t *testing.T) {
	ass := assert.New(t)
	now, _ := time.ParseInLocation("2006-01-02 15:04:05", "2023-01-01 11:15:01", time.Local)
	patches := gomonkey.ApplyFuncReturn(time.Now, now)

	defer patches.Reset()

	ass.False(times.InScope(1116, 1003))
	ass.True(times.InScope(201, 1530))
	ass.False(times.InScope(201, 530))
	ass.False(times.InScope(2001, 2310))
	ass.False(times.InScope(701, 1114))
	ass.True(times.InScope(701, 1115))
	ass.False(times.InScope(1116, 1214))
	ass.True(times.InScope(1115, 1514))

	ass.True(times.InScope(2001, 1901))
	ass.False(times.InScope(2001, 2130))
	ass.False(times.InScope(2101, 1103))

	ass.True(times.InScope(201, 101))
	ass.False(times.InScope(2001, 901))
}

type value struct {
	count int
	lock  sync.RWMutex
}

func (p *value) Get() int {
	p.lock.RLock()
	defer p.lock.RUnlock()

	return p.count
}

func (p *value) Inc() {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.count++
}

func add(ctx context.Context) {
	valu, _ := ctx.Value(base.None).(*value)
	valu.Inc()
}

// nolint: paralleltest
func TestBetween(t *testing.T) {
	// defer goleak.VerifyNone(t)
	ass := assert.New(t)
	now, _ := time.ParseInLocation("2006-01-02 15:04:05", "2023-01-01 11:15:01", time.Local)
	patches1 := gomonkey.ApplyFuncReturn(time.Now, now)

	defer patches1.Reset()

	val := &value{}
	ctx1, canel1 := context.WithCancel(context.Background())

	go times.Between(context.WithValue(ctx1, base.None, val), 1001, 1114, add)

	time.Sleep(time.Second)

	canel1()

	ctx2, canel2 := context.WithCancel(context.Background())

	go times.Between(context.WithValue(ctx2, base.None, val), 1001, 1214, add)

	time.Sleep(time.Second)

	canel2()

	ass.Equal(1, val.Get())
}
