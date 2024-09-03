package portscan

import (
	"fmt"
	"net"
	"time"
)

func ScanPorts(host string, ports []int) map[int]bool {
	openPorts := make(map[int]bool)
	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", address, time.Second)
		if err != nil {
			openPorts[port] = false
			continue
		}
		conn.Close()
		openPorts[port] = true
	}
	return openPorts
}