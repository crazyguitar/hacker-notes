package main

import (
	"encoding/binary"
	"encoding/hex"
	"log"
	"syscall"
	"unsafe"
)

const (
	// EthPALL represent ETH_P_ALL which defined in linux/if_ether.h
	EthPALL = 0x0003

	// EthPIP represent ETH_P_IP which defined in linux/if_ether.h
	EthPIP = 0x0800
)

type ethHdr struct {
	hDest   [6]byte
	hSource [6]byte
	hProto  uint16
}

// IP refer the linux header file netinet/ip.h
type IP struct {
	ipHl  uint8  // Header length: 4 bits
	ipV   uint8  // Version: 4 bits
	ipTos uint8  // Type of service
	ipLen int16  // Total length
	ipID  uint16 // Identification
	ipOff int16  // Fragment offset field
	ipTTL uint8  // Time to live
	ipP   uint8  // Protocol
	ipSum uint16 // Checksum
	ipSrc uint32 // Source address
	ipDst uint32 // Dest address
}

// Htons convert host short to net short
func Htons(i uint16) uint16 {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, i)
	return *(*uint16)(unsafe.Pointer(&b[0]))
}

func decodeEth(buf []byte) (hdr ethHdr) {
	hdr = ethHdr{}
	copy(hdr.hDest[:], buf[0:6])
	copy(hdr.hSource[:], buf[6:12])
	hdr.hProto = binary.BigEndian.Uint16(buf[12:14])
	return
}

func decodeIP(buf []byte) (ip IP) {
	ip = IP{}
	b := buf[0]

	ip.ipHl = b >> 4
	ip.ipV = b & uint8(0x0f)
	ip.ipTos = buf[1]
	ip.ipLen = int16(binary.BigEndian.Uint16(buf[2:4]))
	ip.ipID = binary.BigEndian.Uint16(buf[4:6])
	ip.ipOff = int16(binary.BigEndian.Uint16(buf[6:8]))
	ip.ipTTL = buf[8]
	ip.ipP = buf[9]
	ip.ipSum = binary.BigEndian.Uint16(buf[10:12])
	ip.ipSrc = binary.BigEndian.Uint32(buf[12:16])
	ip.ipDst = binary.BigEndian.Uint32(buf[16:20])

	return
}

func printIPInfo(ip IP) {
	log.Println("------------------- IP_FRAME ------------------")
	log.Printf("Header length:    %#04x", ip.ipHl)
	log.Printf("Version:          %#04x", ip.ipV)
	log.Printf("Service Field:    %#04x", ip.ipTos)
	log.Printf("Total length:     %#04x", ip.ipLen)
	log.Printf("Identification:   %#04x", ip.ipID)
	log.Printf("Fragment offset:  %#04x", ip.ipOff)
	log.Printf("Time to live:     %#04x", ip.ipTTL)
	log.Printf("Protocol:         %#04x", ip.ipP)
	log.Printf("Checksum:         %#04x", ip.ipSum)
	log.Printf("Source address:   %d.%d.%d.%d", 0xf&(ip.ipSrc>>24), 0xf&(ip.ipSrc>>16), 0xf&(ip.ipSrc>>8), 0xf&ip.ipSrc)
	log.Printf("Dest address:     %d.%d.%d.%d", 0xf&(ip.ipDst>>24), 0xf&(ip.ipDst>>16), 0xf&(ip.ipDst>>8), 0xf&ip.ipDst)
}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, int(Htons(uint16(EthPALL))))
	if err != nil {
		log.Fatal(err)
	}
	defer syscall.Close(fd)

	for {
		var buf [8192]byte
		_, _, err := syscall.Recvfrom(fd, buf[:], 0)
		if err != nil {
			log.Println(err)
			continue
		}
		ethhdr := decodeEth(buf[0:14])

		if ethhdr.hProto != EthPIP {
			continue
		}

		ip := decodeIP(buf[14:34])

		printIPInfo(ip)
		log.Printf("\n%s\n", hex.Dump(buf[0:34]))
	}
}
