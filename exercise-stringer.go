package main

import "fmt"

type IPAddr [4]byte

func (ip_addr IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip_addr[0], ip_addr[1], ip_addr[2], ip_addr[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

// $ go run methods18.go
// loopback: 127.0.0.1
// googleDNS: 8.8.8.8
