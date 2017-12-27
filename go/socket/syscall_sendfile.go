package main

import (
	"fmt"
	"net"
	"os"
	"syscall"
)

const (
	inFileName  = ".in"
	outFileName = ".out"
	host        = "127.0.0.1"
	port        = 5566
)

func doSendfile(inFileName string, outFileName string) {

	inFile, err := os.Open(inFileName)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	outFile, err := os.Create(outFileName)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	defer func() {
		err := os.Remove(outFileName)
		if err != nil && !os.IsNotExist(err) {
			fmt.Println(err)
		}
	}()

	inFd := int(inFile.Fd())
	outFd := int(outFile.Fd())

	var offset int64
	count := 8192

	for {
		n, err := syscall.Sendfile(outFd, inFd, &offset, count)
		if err != nil {
			panic(err)
		}
		if n == 0 {
			break
		}
	}
}

func server() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	inFile, err := os.Open(inFileName)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	st, err := inFile.Stat()
	if err != nil {
		panic(err)
	}

	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		panic(err)
	}

	err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
	if err != nil {
		panic(err)
	}

	addr := syscall.SockaddrInet4{Port: port}
	copy(addr.Addr[:], net.ParseIP(host).To4())

	err = syscall.Bind(fd, &addr)
	if err != nil {
		panic(err)
	}

	err = syscall.Listen(fd, 1)
	if err != nil {
		panic(err)
	}

	cFd, _, err := syscall.Accept(fd)
	if err != nil {
		panic(err)
	}

	var offset int64
	count := 8192
	fSize := st.Size()
	inFd := int(inFile.Fd())

	for fSize > 0 {
		n, err := syscall.Sendfile(cFd, inFd, &offset, count)
		if err != nil {
			panic(err)
		}
		fSize -= int64(n)
	}
}

func main() {

	inFile, err := os.Open("/dev/urandom")
	if err != nil {
		panic(err)
	}

	defer inFile.Close()

	outFile, err := os.Create(inFileName)
	if err != nil {
		panic(err)
	}

	defer outFile.Close()

	defer func() {
		err := os.Remove(inFileName)
		if err != nil && !os.IsNotExist(err) {
			fmt.Println(err)
		}
	}()

	buf := make([]byte, 1024)
	count := 1000 * 512

	for i := 0; i < count; i++ {
		n, err := inFile.Read(buf)
		if err != nil {
			panic(err)
		}
		if n == 0 {
			break
		}
		_, err = outFile.Write(buf[:n])
		if err != nil {
			panic(err)
		}
	}
	outFile.Close()
	outFile = nil

	doSendfile(inFileName, outFileName)
	server()
}
