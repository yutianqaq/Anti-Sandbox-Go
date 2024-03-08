package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

// https://github.com/Ne0nd0g/sliver/blob/d96bfc2b09921783155e1453cff8ada712917fdb/sliver/limits/limits.go
func isDomainJoined() (bool, error) {
	var domain *uint16
	var status uint32
	err := syscall.NetGetJoinInformation(nil, &domain, &status)
	if err != nil {
		return false, err
	}
	syscall.NetApiBufferFree((*byte)(unsafe.Pointer(domain)))
	return status == syscall.NetSetupDomainName, nil
}

func main() {
	ok, err := isDomainJoined()
	if err == nil && !ok {
		os.Exit(1)
	}

	fmt.Printf("shellcode ...")
}
