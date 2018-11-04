package masternode

import (
	"syscall"
	"time"

	"github.com/StackExchange/wmi"
)

var kernel = syscall.NewLazyDLL("Kernel32.dll")

type memoryStatusEx struct {
	cbSize                  uint32
	dwMemoryLoad            uint32
	ullTotalPhys            uint64 // in bytes
	ullAvailPhys            uint64
	ullTotalPageFile        uint64
	ullAvailPageFile        uint64
	ullTotalVirtual         uint64
	ullAvailVirtual         uint64
	ullAvailExtendedVirtual uint64
}

type Storage struct {
	Name       string
	FileSystem string
	Total      uint64
	Free       uint64
}

type storageInfo struct {
	Name       string
	Size       uint64
	FreeSpace  uint64
	FileSystem string
}

func GetStartTime() string {
	GetTickCount := kernel.NewProc("GetTickCount")
	r, _, _ := GetTickCount.Call()
	if r == 0 {
		return ""
	}
	ms := time.Duration(r * 1000 * 1000)
	return ms.String()
}

func GetFreeStorage() uint64 {
	var storageInfo []storageInfo
	//var loaclStorageList []Storage
	err := wmi.Query("Select * from Win32_LogicalDisk", &storageInfo)
	if err != nil {
		return 0
	}

	totalFree := uint64(0)

	for _, storage := range storageInfo {
		totalFree = totalFree + storage.Size
	}
	return totalFree
}
