package types

import (
	"fmt"

	"github.com/xuender/kit/base"
)

const _units = "BKMGTP"

func Storage(size uint64) string {
	var (
		tmp  uint64
		unit int
	)

	for {
		tmp = size / base.Kilo

		if tmp == 0 {
			return fmt.Sprintf("%d%c", size%base.Kilo, _units[unit])
		}

		size = tmp
		unit++
	}
}
