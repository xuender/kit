package ios_test

import (
	"io"
	"os"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xuender/kit/ios"
)

func TestWrite(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	req := require.New(t)
	num, err := ios.Write(io.Discard, []byte("xxx"), []byte("1111"))
	ass.Equal(7, num)
	req.NoError(err)

	patches := gomonkey.ApplyMethodReturn(io.Discard, "Write", nil, os.ErrClosed)
	defer patches.Reset()

	_, err = ios.Write(io.Discard, []byte("xxx"))
	req.Error(err)
}
