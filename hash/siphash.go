package hash

import (
	"encoding/binary"
	"encoding/hex"
	"hash"

	"github.com/dchest/siphash"
	"github.com/xuender/kit/base"
)

const (
	key0 = 506097522914230528
	key1 = 1084818905618843912
)

// SipHash32 32位哈希.
func SipHash32(data []byte) uint32 {
	sum := make([]byte, base.Eight)
	binary.BigEndian.PutUint64(sum, SipHash64(data))

	return binary.BigEndian.Uint32(sum[0:base.Four])
}

// SipHash128 哈希.
func SipHash128(data []byte) (uint64, uint64) {
	return siphash.Hash128(key0, key1, data)
}

// SipHash64 哈希，和Google Guava的sipHash24相同.
func SipHash64(data []byte) uint64 {
	return siphash.Hash(key0, key1, data)
}

// SipHashNumber 兼容JS, 53位长度.
func SipHashNumber(data []byte) uint64 {
	return JSNumber(siphash.Hash(key0, key1, data))
}

// SipHashHex 字符串.
func SipHashHex(data []byte) string {
	var (
		sum = siphash.Hash(key0, key1, data)
		b8  = make([]byte, base.Eight)
	)

	binary.LittleEndian.PutUint64(b8, sum)

	return hex.EncodeToString(b8)
}

func NewSipHash64() hash.Hash64 {
	key := make([]byte, base.Sixteen)
	binary.LittleEndian.PutUint64(key, key0)
	binary.LittleEndian.PutUint64(key[base.Eight:], key1)

	return siphash.New(key)
}
