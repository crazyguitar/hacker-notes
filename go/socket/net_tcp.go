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
	service     = ":5566"
)

func prepareSrcFile() {
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

	var buf [1024]byte
	count := 1000 * 128

	for i := 0; i < count; i++ {
		n, err := inFile.Read(buf[:])
		if err != nil {
			panic(err)
		}
		_, err = outFile.Write(buf[:n])
		if err != nil {
			panic(err)
		}
	}
	log.Println("prepare source file done")
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

func doSendfile(inFd int, outFd int) {
	var offset int64
	count := 8192

	for {
		n, err := syscall.Sendfile(outFd, inFd, &offset, count)
		if err != nil {
			panic(err)
		}
		if n <= 0 {
			break
		}
	}
}

func client(ok chan bool) {
	outFile, err := os.Create(outFileName)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	conn, err := net.Dial("tcp", "127.0.0.1"+service)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	var buf [4096]byte

	for {
		n, err := conn.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		_, err = outFile.Write(buf[:n])
		if err != nil {
			panic(err)
		}
	}
	ok <- true
}

func server(ok chan bool) {
	addr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		panic(err)
	}

	ln, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	ok <- true

	count := 1

	for i := 0; i < count; i++ {
		conn, err := ln.AcceptTCP()
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		connFile, err := conn.File()
		if err != nil {
			panic(err)
		}
		defer connFile.Close()

		// get the connection fd
		connFd := int(connFile.Fd())

		inFile, err := os.Open(inFileName)
		if err != nil {
			panic(err)
		}
		defer inFile.Close()

		// get the source file fd
		inFd := int(inFile.Fd())

		doSendfile(inFd, connFd)
	}
}

func main() {

	prepareSrcFile()

	// cleanup
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
