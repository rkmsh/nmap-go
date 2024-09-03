package main

import (
	"fmt"
	"nmap-go/internal/detect"
	"nmap-go/internal/ping"
	"nmap-go/internal/portscan"
)

func main() {
	host := "127.0.0.1"
	if ping.Ping(host) {
		fmt.Println("Host is up")
	} else {
		fmt.Println("Host is down")
		return
	}

	ports := []int{80, 443, 8080}
	openPorts := portscan.ScanPorts(host, ports)
	for port, isOpen := range openPorts {
		if isOpen {
			fmt.Printf("Port %d is open\n", port)
			banner := detect.GrabBanner(host, port)
			if banner != "" {
				fmt.Printf("Banner: %s\n", banner)
			}
		} else {
			fmt.Printf("Port %d is closed\n", port)
		}
	}
}
