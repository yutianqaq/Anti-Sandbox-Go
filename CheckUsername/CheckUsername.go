package main

import (
	"fmt"
	"os"
	"strings"
)

func CheckUsername() {
	hostname, err := os.Hostname()
	if err == nil && strings.ToLower(hostname) != strings.ToLower("Windows 11") {
		os.Exit(1)
	}
}

func main() {
	ok, err := CheckUsername()
	if err == nil && !ok {
		os.Exit(1)
	}

	fmt.Printf("shellcode ...")
}
