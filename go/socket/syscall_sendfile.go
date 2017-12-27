package main

import (
	"fmt"
	"os"
	"syscall"
)

const (
	inFileName  = ".in"
	outFileName = ".out"
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
}
