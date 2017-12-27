package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
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

func copyFile(inFileName string, outFileName string) {

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

	doSendfile(inFd, outFd)

	eq := checksum(inFileName, outFileName)
	if !eq {
		panic("check sum not equal")
	}
}

func doSendfile(inFd int, outFd int) {

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

func server(ok chan bool) {

	inFile, err := os.Open(inFileName)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		panic(err)
	}
	defer syscall.Close(fd)

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

	ok <- true

	cFd, _, err := syscall.Accept(fd)
	if err != nil {
		panic(err)
	}
	defer syscall.Close(cFd)

	inFd := int(inFile.Fd())
	doSendfile(inFd, cFd)
}

func client(ok chan bool) {

	outFile, err := os.Create(outFileName)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		panic(err)
	}
	defer syscall.Close(fd)

	addr := syscall.SockaddrInet4{Port: port}
	copy(addr.Addr[:], net.ParseIP(host).To4())

	err = syscall.Connect(fd, &addr)
	if err != nil {
		panic(err)
	}

	outFd := int(outFile.Fd())
	var buf [4096]byte

	for {
		n, _ := syscall.Read(fd, buf[:])
		if err != nil {
			panic(err)
		}
		if n <= 0 {
			break
		}
		_, err = syscall.Write(outFd, buf[:n])
		if err != nil {
			panic(err)
		}
	}
	ok <- true
}

func checksum(file1 string, file2 string) (eq bool) {

	f1, err := os.Open(file1)
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	f2, err := os.Open(file2)
	if err != nil {
		panic(err)
	}
	defer f2.Close()

	h1 := sha256.New()
	if _, err := io.Copy(h1, f1); err != nil {
		panic(err)
	}

	h2 := sha256.New()
	if _, err := io.Copy(h2, f2); err != nil {
		panic(err)
	}

	s1 := hex.EncodeToString(h1.Sum(nil))
	s2 := hex.EncodeToString(h2.Sum(nil))
	return s1 == s2

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

		err = os.Remove(outFileName)
		if err != nil && !os.IsNotExist(err) {
			fmt.Println(err)
		}
	}()

	buf := make([]byte, 1024)
	count := 1000 * 128

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

	outFile.Sync()

	ok := make(chan bool)

	copyFile(inFileName, outFileName)

	go server(ok)
	<-ok
	go client(ok)
	<-ok

	eq := checksum(inFileName, outFileName)
	if !eq {
		panic("check sum not equal")
	}
}
