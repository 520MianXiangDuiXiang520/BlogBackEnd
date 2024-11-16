//go:build linux

package tools

import (
	"fmt"
	"os"
	"syscall"
)

func GetFileLastChangeTime(info os.FileInfo) int64 {
	t := info.Sys().(*syscall.Stat_t).Mtim.Nano()
	fmt.Println(t, info.ModTime().Nanosecond())
	return t
}
