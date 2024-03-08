package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

// https://github.com/Ne0nd0g/sliver/blob/d96bfc2b09921783155e1453cff8ada712917fdb/sliver/limits/limits.go
func PlatformLimits() {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	isDebuggerPresent := kernel32.NewProc("IsDebuggerPresent")
	var nargs uintptr = 0
	ret, _, _ := isDebuggerPresent.Call(nargs)
	log.Printf("IsDebuggerPresent = %#v\n", int32(ret))
	if int32(ret) != 0 {
		os.Exit(1)
	}
}

func main() {
	PlatformLimits()

	fmt.Printf("shellcode ...")
}
