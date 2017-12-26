package main

import (
	"log"
	"net"
	"os"
	"syscall"
)

// related socket info
const (
	PORT      = 5566
	HOST      = "127.0.0.1"
	FDSETSIZE = 256
)

// FdSet just do FD_SET
func FdSet(p *syscall.FdSet, i int) {
	p.Bits[i/64] |= 1 << uint(i) % 64
}

// FdClr just do FD_CLR
func FdClr(p *syscall.FdSet, i int) {
	p.Bits[i/64] &= ^(1 << uint(i) % 64)
}

// FdIsSet just do FD_ISSET
func FdIsSet(p *syscall.FdSet, i int) bool {
	return (p.Bits[i/64] & (1 << uint(i) % 64)) != 0
}

// FdZero just do FD_ZERO
func FdZero(p *syscall.FdSet) {
	for i := range p.Bits {
		p.Bits[i] = 0
	}
}

// Usage:
//
// $ syscall_select &
// $ nc localhost 5566
// Hello
// Hello
func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

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

	wfds := syscall.FdSet{}
	timeout := syscall.Timeval{Sec: 5, Usec: 0}

	FdZero(&wfds)
	FdSet(&wfds, fd)

	defer FdClr(&wfds, fd)

	for {
		rfds := wfds

		_, err := syscall.Select(FDSETSIZE, &rfds, nil, nil, &timeout)
		if err != nil {
			log.Println(err)
			continue
		}

		for i := 0; i < FDSETSIZE; i++ {
			if !FdIsSet(&rfds, i) {
				continue
			}

			if i == fd {
				cfd, _, err := syscall.Accept(fd)
				if err != nil {
					log.Println(err)
					continue
				}
				FdSet(&wfds, cfd)
			} else {
				var buf [1024]byte
				nbytes, err := syscall.Read(i, buf[:])
				if err != nil {
					log.Println(err)
					FdClr(&wfds, i)
					syscall.Close(i)
					continue
				}
				if nbytes <= 0 {
					FdClr(&wfds, i)
					syscall.Close(i)
					continue
				}
				_, err = syscall.Write(i, buf[:nbytes])
				if err != nil {
					log.Println(err)
					FdClr(&wfds, i)
					syscall.Close(i)
					continue
				}

			}
		}
	}
}
