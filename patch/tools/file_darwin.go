//go:build darwin

package tools

import (
	"fmt"
	"os"
	"syscall"
)

func GetFileLastChangeTime(info os.FileInfo) int64 {
	t := info.Sys().(*syscall.Stat_t).Mtimespec.Nano()
	fmt.Println(t, info.ModTime().Nanosecond())
	return t
}
