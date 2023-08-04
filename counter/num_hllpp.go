package counter

import (
	"encoding/binary"
	"math"

	"github.com/retailnext/hllpp"
	"github.com/xuender/kit/base"
	"golang.org/x/exp/constraints"
)

type NumHLLPP[T constraints.Integer | constraints.Float] struct {
	Hllpp *hllpp.HLLPP
}

func NewNumHLLPP[T constraints.Integer | constraints.Float]() *NumHLLPP[T] {
	return &NumHLLPP[T]{Hllpp: hllpp.New()}
}

func (p *NumHLLPP[T]) Add(num T) {
	p.Hllpp.Add(Number2Bytes(num))
}

func (p *NumHLLPP[T]) Count() uint64 {
	return p.Hllpp.Count()
}

func (p *NumHLLPP[T]) Marshal() []byte {
	return p.Hllpp.Marshal()
}

// nolint
func Unmarshal[T constraints.Integer | constraints.Float](data []byte) (*NumHLLPP[T], error) {
	hll, err := hllpp.Unmarshal(data)
	if err != nil {
		return nil, err
	}

	return &NumHLLPP[T]{
		Hllpp: hll,
	}, nil
}

func Number2Bytes[T constraints.Integer | constraints.Float](num T) []byte {
	switch (any)(num).(type) {
	case float32, float64:
		return Number2Bytes(math.Float64bits(float64(num)))
	}

	data := make([]byte, base.Eight)
	binary.LittleEndian.PutUint64(data, uint64(num))

	for i := base.Eight; i > 0; i-- {
		if data[i-1] > 0 {
			return data[:i]
		}
	}

	return nil
}
