package main

import (
	"encoding/binary"
	"encoding/hex"
	"log"
	"syscall"
	"unsafe"
)

const (
	// IntMax represent the size of int
	IntMax = int(unsafe.Sizeof(0))
	// EthPALL represent the socket protocol of ETH_P_ALL
	EthPALL = 0x0003
	// EthARP represent define ETH_P_ARP
	EthARP = 0x0806
)

type ethHdr struct {
	hDest   [6]byte
	hSource [6]byte
	hProto  uint16
}

type arpHdr struct {
	arHrd uint16
	arPro uint16
	arHln uint8
	arPln uint8
	arOp  uint16
}

type arpPayload struct {
	arSha [6]byte
	arSip [4]byte
	arTha [6]byte
	arTip [4]byte
}

// ARP represent the struct of ARP packet
type ARP struct {
	arpHdr
	arpPayload
}

func isLittleEdian() (ret bool) {
	i := 0x1
	b := (*[IntMax]byte)(unsafe.Pointer(&i))
	ret = true
	if b[0] == 0 {
		ret = false
	}
	return
}

func htons(hostshort uint16) (netshort uint16) {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, hostshort)
	if isLittleEdian() {
		netshort = binary.LittleEndian.Uint16(b[:])
	} else {
		netshort = binary.BigEndian.Uint16(b[:])
	}
	return
}

func decodeEth(buf []byte) (hdr ethHdr) {
	hdr = ethHdr{}
	copy(hdr.hDest[:], buf[0:6])
	copy(hdr.hSource[:], buf[6:12])
	hdr.hProto = binary.BigEndian.Uint16(buf[12:14])
	return
}

func printEthInfo(hdr ethHdr) {
	log.Println("---------------- ETHERNET_FRAME ----------------")
	log.Println("Dest MAC:        " + hex.EncodeToString(hdr.hDest[:]))
	log.Println("Source MAC:      " + hex.EncodeToString(hdr.hSource[:]))
	log.Printf("Type:            %#04x\n", hdr.hProto)
}

func decodeARP(buf []byte) (arp ARP) {
	arp = ARP{}
	// header
	arp.arHrd = binary.BigEndian.Uint16(buf[0:2])
	arp.arPro = binary.BigEndian.Uint16(buf[2:4])
	arp.arHln = uint8(buf[4])
	arp.arPln = uint8(buf[5])
	arp.arOp = binary.BigEndian.Uint16(buf[6:8])

	// payload
	copy(arp.arSha[:], buf[8:14])
	copy(arp.arSip[:], buf[14:18])
	copy(arp.arTha[:], buf[18:24])
	copy(arp.arTip[:], buf[24:28])
	return
}

func printARPInfo(arp ARP) {
	log.Println("------------------- ARP_FRAME ------------------")
	log.Printf("Hardware type:   %#04x", arp.arHrd)
	log.Printf("Protocol type:   %#04x", arp.arPro)
	log.Printf("Hardware size:   %#02x", arp.arHln)
	log.Printf("Protocol size:   %#02x", arp.arPln)
	log.Printf("Opcode:          %#04x", arp.arOp)
	log.Printf("Source Mac:      %s\n", hex.EncodeToString(arp.arSha[:]))
	log.Printf("Source IP:       %d.%d.%d.%d\n", arp.arSip[0], arp.arSip[1], arp.arSip[2], arp.arSip[3])
	log.Printf("Target Mac:      %s\n", hex.EncodeToString(arp.arTha[:]))
	log.Printf("Target IP:       %d.%d.%d.%d\n", arp.arTip[0], arp.arTip[1], arp.arTip[2], arp.arTip[3])
}

// usage:
//	$ sudo ./arp_sniffer &
// 	$ sudo arping -c 1 10.0.0.1
func main() {

	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, int(htons(uint16(EthPALL))))
	if err != nil {
		panic(err)
	}
	defer syscall.Close(fd)

	for {
		var buf [42]byte
		_, _, err := syscall.Recvfrom(fd, buf[:], 0)
		if err != nil {
			log.Println(err)
			continue
		}

		ethhdr := decodeEth(buf[0:14])

		if ethhdr.hProto != EthARP {
			continue
		}

		arp := decodeARP(buf[14:42])

		printEthInfo(ethhdr)
		printARPInfo(arp)
	}
}
