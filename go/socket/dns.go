package main

import (
	"fmt"
	"net"
)

// test domain
var domains = []string{
	"www.google.com:443",
	"www.amazon.com:443",
	"www.github.com:443",
	"www.facebook.com:443",
	"localhost:80",
}

func main() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	for _, domain := range domains {
		tcpAddr, err := net.ResolveTCPAddr("tcp4", domain)
		if err != nil {
			panic(err)
		}
		fmt.Println(tcpAddr.IP)
	}
}
