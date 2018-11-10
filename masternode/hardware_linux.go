package masternode

import (
	"os"
	"syscall"
)

func GetFreeStorage() uint64 {
	var stat syscall.Statfs_t
	wd, _ := os.Getwd()
	syscall.Statfs(wd, &stat)

	// Available blocks * size per block = available space in bytes
	return stat.Bavail * uint64(stat.Bsize)
}
