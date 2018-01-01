package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"net"
	"os"
	"syscall"
)

const (
	inFileName  = ".in"
	outFileName = ".out"
	host        = "127.0.0.1"
	port        = 5566

	// SpliceFMove represent SPLICE_F_MOVE
	SpliceFMove = 0x01
	// SpliceFMore represent SPLICE_F_MORE
	SpliceFMore = 0x04
)

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

func doSplice(inFd int, outFd int) {
	var p [2]int
	var st syscall.Stat_t

	err := syscall.Pipe(p[:])
	if err != nil {
		panic(err)
	}
	defer syscall.Close(p[0])
	defer syscall.Close(p[1])

	err = syscall.Fstat(inFd, &st)
	if err != nil {
		panic(err)
	}

	inFileSize := st.Size
	count := 8192
	var offset int64

	for inFileSize > 0 {
		n, err := syscall.Splice(inFd, &offset, p[1], nil, count, SpliceFMove|SpliceFMore)
		if err != nil {
			panic(err)
		}

		_, err = syscall.Splice(p[0], nil, outFd, nil, count, SpliceFMove|SpliceFMore)
		if err != nil {
			panic(err)
		}
		inFileSize -= n
	}
}

func prepareFile(fileName string) {
	inFile, err := os.Open("/dev/urandom")
	if err != nil {
		panic(err)
	}

	defer inFile.Close()

	outFile, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	defer outFile.Close()

	count := 1000 * 128
	for i := 0; i < count; i++ {
		var buf [1024]byte

		n, err := inFile.Read(buf[:])
		if err != nil {
			panic(err)
		}

		_, err = outFile.Write(buf[:n])
		if err != nil {
			panic(err)
		}
	}
}

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

	// clean outFileName
	defer func() {
		err := os.Remove(outFileName)
		if err != nil {
			log.Fatal(err)
		}
	}()

	inFd := int(inFile.Fd())
	outFd := int(outFile.Fd())

	doSplice(inFd, outFd)

	err = outFile.Sync()
	if err != nil {
		panic(err)
	}

	eq := checksum(inFileName, outFileName)
	if !eq {
		panic("checksum not equal")
	}
}

func server(ok chan bool) {

	inFd, err := syscall.Open(inFileName, syscall.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	defer syscall.Close(inFd)

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

	doSplice(inFd, cFd)
}

func client(ok chan bool) {

	var mode uint32 = syscall.S_IRUSR | syscall.S_IWUSR | syscall.S_IRGRP | syscall.S_IROTH //0644

	outFd, err := syscall.Open(outFileName, syscall.O_CREAT|syscall.O_WRONLY, mode)
	if err != nil {
		panic(err)
	}
	defer syscall.Close(outFd)

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

	for {
		var buf [8192]byte

		n, err := syscall.Read(fd, buf[:])
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

func main() {

	prepareFile(inFileName)

	// clean the test files
	defer func() {
		err := os.Remove(inFileName)
		if err != nil && !os.IsNotExist(err) {
			log.Fatal(err)
		}

		err = os.Remove(outFileName)
		if err != nil && !os.IsNotExist(err) {
			log.Fatal(err)
		}
	}()

	copyFile(inFileName, outFileName)

	ok := make(chan bool)

	go server(ok)
	<-ok
	go client(ok)
	<-ok

	eq := checksum(inFileName, outFileName)
	if !eq {
		panic("checksum not equal")
	}

}
