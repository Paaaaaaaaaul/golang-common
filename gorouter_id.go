package common

import (
	"bytes"
	"runtime"
	"strconv"
)

//获取id
func GetGorouterID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

//获取id标志
func GetGorouterIDFlag() uint64 {
	return GetGorouterID() + 10000000
}
