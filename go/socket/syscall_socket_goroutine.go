package main

import (
	"fmt"
	"net"
	"os"
	"syscall"
)

// socket related info
const (
	PORT = 5566
	HOST = "127.0.0.1"
)

func echo(fd int) {

	var buf [1024]byte

	defer syscall.Close(fd)

	for {
		nbytes, err := syscall.Read(fd, buf[:])
		if err != nil {
			fmt.Println(err)
			break
		}

		if nbytes <= 0 {
			break
		}

		_, err = syscall.Write(fd, buf[:nbytes])
		if err != nil {
			fmt.Println(err)
			break
		}
	}

	fmt.Println("connection close")
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		panic(err)
	}
	defer syscall.Close(fd)

	err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
	if err != nil {
		panic(err)
	}

	addr := syscall.SockaddrInet4{Port: PORT}
	copy(addr.Addr[:], net.ParseIP(HOST).To4())

	err = syscall.Bind(fd, &addr)
	if err != nil {
		panic(err)
	}

	err = syscall.Listen(fd, 10)
	if err != nil {
		panic(err)
	}

	for {
		cfd, _, err := syscall.Accept(fd)
		if err != nil {
			fmt.Println(err)
			continue
		}

		go echo(cfd)
	}

}
