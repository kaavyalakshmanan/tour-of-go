package main

import (
"fmt"
"strings"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	ipAddrStringBuilder := []string{}
	for _, num := range ip {
		ipAddrStringBuilder = append(ipAddrStringBuilder, fmt.Sprint(int(num)))	
	}
	return strings.Join(ipAddrStringBuilder, ".")
	
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
