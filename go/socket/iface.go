package main

import (
	"fmt"
	"net"
)

func main() {
	ifaces, err := net.Interfaces()
	if err != nil {
		return
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}

		ifname := i.Name
		ifIndex := i.Index
		ifMac := i.HardwareAddr.String()

		for _, a := range addrs {
			switch v := a.(type) {
			case *net.IPNet:
				fmt.Printf("%v (%d) : %s (%s) mac: %s\n", ifname, ifIndex, v, v.IP.DefaultMask(), ifMac)
			default:
				fmt.Printf("I don't know about type %T!\n", v)
			}

		}
	}
}
