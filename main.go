package main

import (
	"fmt"
	"syscall"
	"time"
)

func main() {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	getTickCount64 := kernel32.NewProc("GetTickCount64")

	uptimeMilliseconds, _, _ := getTickCount64.Call()
	uptime := time.Duration(uptimeMilliseconds) * time.Millisecond

	fmt.Printf("System Uptime: %s\n", uptime)
}
