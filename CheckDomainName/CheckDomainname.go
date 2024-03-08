package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

var (
	modkernel32            = syscall.NewLazyDLL("kernel32.dll")
	procGetComputerNameExA = modkernel32.NewProc("GetComputerNameExA")
)

func GetComputerNameExA() {
	name_type := 2
	var ret uintptr
	size := uint32(0)

	// syscall.Syscall(getComputerNameExA, 3, uintptr(name_type), 0, uintptr(unsafe.Pointer(&size)))
	syscall.Syscall(procGetComputerNameExA.Addr(), 3, uintptr(name_type), 0, uintptr(unsafe.Pointer(&size)))

	buffer := make([]byte, size)
	ret, _, _ = syscall.Syscall(procGetComputerNameExA.Addr(), 3, uintptr(name_type), uintptr(unsafe.Pointer(&buffer[0])), uintptr(unsafe.Pointer(&size)))
	if ret == 0 {
		os.Exit(1)
	}

	if string(buffer[:size]) != "havoc.local" {
		os.Exit(1)
	}
}

func main() {
	// DC01.havoc.local
	// 1 - DC01
	// 2 - havoc.local
	// 3 - DC01.havoc.local

	GetComputerNameExA()

	fmt.Printf("shellcode ...")
}
