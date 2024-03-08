package main

import (
	"net"
	"os"
)

func CheckInterface() (bool, int) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return false, 0
	}
	// println(len(interfaces))
	total := len(interfaces)

	for _, iface := range interfaces {
		hwAddr := iface.HardwareAddr
		// vmware 00:0c
		if len(hwAddr) >= 5 && hwAddr.String()[:5] == "00:0c" {
			return false, total
		}
	}

	return true, total
}

func main() {

	f, t := CheckInterface()
	if !f || t < 3 {
		os.Exit(1)
	}

	print("shellcode...")
}
