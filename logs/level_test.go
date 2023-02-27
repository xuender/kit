package logs_test

import (
	"os"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/kit/logs"
)

// nolint: paralleltest
func TestDefaultLevel(t *testing.T) {
	patches := gomonkey.ApplyFuncReturn(os.Getenv, "1")

	ass := assert.New(t)

	ass.Equal(logs.Debug, logs.DefaultLevel())
	patches.Reset()

	patches2 := gomonkey.ApplyFuncReturn(os.Getenv, "erRor")
	defer patches2.Reset()

	ass.Equal(logs.Error, logs.DefaultLevel())
}
