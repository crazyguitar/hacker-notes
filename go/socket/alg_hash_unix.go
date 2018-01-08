package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"

	"golang.org/x/sys/unix"
)

// Sha1Hash use AF_ALG to calculate sha1 hash
func Sha1Hash(b []byte) {

	fd, err := unix.Socket(unix.AF_ALG, unix.SOCK_SEQPACKET, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer unix.Close(fd)

	addr := &unix.SockaddrALG{Type: "hash", Name: "sha256"}

	err = unix.Bind(fd, addr)
	if err != nil {
		log.Fatal(err)
	}

	hashfd, _, errno := unix.Syscall(unix.SYS_ACCEPT, uintptr(fd), 0, 0)
	if errno != 0 {
		log.Fatal(err)
	}

	err = unix.Sendto(int(hashfd), b, unix.MSG_MORE, addr)
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 32)
	_, err = unix.Read(int(hashfd), buf[:])
	if err != nil {
		log.Fatal(err)
	}

	// show the hash (calculate via AF_ALG socket)
	log.Println(hex.EncodeToString(buf))

}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	b := []byte("Hello AF_ALG")
	Sha1Hash(b)

	// show the hash (calculate via crypto/sha1)
	hash := sha256.Sum256(b)
	log.Println(hex.EncodeToString(hash[:]))
}
