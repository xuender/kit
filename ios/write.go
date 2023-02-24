package ios

import (
	"io"
)

// Write 写入数组.
func Write(writer io.Writer, datas ...[]byte) (int, error) {
	count := 0

	for _, data := range datas {
		num, err := writer.Write(data)
		if err != nil {
			return count, err
		}

		count += num
	}

	return count, nil
}
