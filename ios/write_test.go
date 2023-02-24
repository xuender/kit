package ios_test

import (
	"os"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/ios"
)

func TestWrite(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	ignore := ios.IgnoreWriter{}
	num, err := ios.Write(ignore, []byte("xxx"), []byte("1111"))
	ass.Equal(7, num)
	ass.Nil(err)

	patches := gomonkey.ApplyMethodReturn(ignore, "Write", nil, os.ErrClosed)
	defer patches.Reset()

	_, err = ios.Write(ignore, []byte("xxx"))
	ass.NotNil(err)
}
