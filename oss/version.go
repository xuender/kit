package oss

/*
#include<stdint.h>
#include<string.h>
void getCompileDateTime(uint8_t  dt[12],uint8_t tm[9]){
  strcpy(dt, __DATE__); //Mmm dd yyyy
  strcpy(tm,__TIME__);  //hh:mm:ss
}
*/
import "C"

import (
	"bytes"
	"time"
	"unsafe"

	"github.com/xuender/kit/base"
)

// nolint: gochecknoglobals
var (
	// Version 版本号.
	Version = ""
	// BuildTime 编译时间.
	BuildTime = ""
)

// GetBuildTime 获取编译时间.
func GetBuildTime() string {
	dbs := make([]byte, base.Fifteen)
	tbs := make([]byte, base.Ten)

	C.getCompileDateTime((*C.uint8_t)(unsafe.Pointer(&dbs[0])), (*C.uint8_t)(unsafe.Pointer(&tbs[0])))

	dts, tms := string(bytes.Trim(dbs, "\x00")), string(bytes.Trim(tbs, "\x00"))
	date, _ := time.Parse("Jan 02 200615:04:05", dts+tms)

	return date.Format("2006-01-02 15:04:05")
}

// nolint: gochecknoinits
func init() {
	BuildTime = GetBuildTime()
}
