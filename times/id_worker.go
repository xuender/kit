package times

import (
	"hash/fnv"
	"sync"
	"time"

	"github.com/xuender/kit/base"
)

const (
	_start      int64 = 1_682_000_000
	_seqDefault int64 = 20
)

type IDWorker struct {
	machine int64
	last    time.Time
	seq     int64
	seqMask int64
	seqLen  int64
	mutex   sync.Mutex
}

// NewIDWorkerByMachine 根据机器信息创建 IDWorker.
func NewIDWorkerByMachine(machine, machineLength int64) *IDWorker {
	sequenceLength := _seqDefault

	if machineLength > 0 {
		if machineLength > base.Sixteen {
			machineLength = base.Sixteen
		}

		if machineLength >= base.Ten {
			sequenceLength = base.Sixteen
		}

		machine &= (-1 ^ (-1 << machineLength))
		machine <<= (machineLength + sequenceLength)
	} else {
		machine = 0
	}

	return &IDWorker{
		machine: machine,
		last:    time.Now(),
		seq:     0,
		seqMask: -1 ^ (-1 << (sequenceLength + machineLength)),
		seqLen:  sequenceLength,
		mutex:   sync.Mutex{},
	}
}

// NewIDWorkerByKey 根据字符串创建 IDWorker.
func NewIDWorkerByKey(key string) *IDWorker {
	hash := fnv.New64()
	hash.Write([]byte(key))

	return NewIDWorkerByMachine(int64(hash.Sum64()), _seqDefault)
}

// NewIDWorker 创建非分布式的 IDWorker.
func NewIDWorker() *IDWorker {
	return NewIDWorkerByMachine(0, 0)
}

// ID 生成ID.
func (p *IDWorker) ID() int64 {
	uid, err := p.IDAndError()
	if err != nil {
		panic(err)
	}

	return uid
}

// IDAndError 生成ID和错误信息.
func (p *IDWorker) IDAndError() (int64, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	now := time.Now()
	if now.Unix() < p.last.Unix() {
		return 0, ErrLast
	}

	if now.Unix() > p.last.Unix() {
		p.last = now
		p.seq = 0
	} else {
		p.seq = (p.seq + 1) & p.seqMask

		if p.seq == 0 {
			time.Sleep(time.Second - time.Since(p.last))
			p.last = time.Now()
		}
	}

	return ((p.last.Unix() - _start) << p.seqLen) | p.machine | p.seq, nil
}
