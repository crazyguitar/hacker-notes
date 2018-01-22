package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
	"os"
	"strconv"
	"syscall"
	"time"
)

const (
	// ICMP type
	icmpEchoReply      uint8 = 0
	icmpDestUnreach    uint8 = 3
	icmpSourceQuench   uint8 = 4
	icmpRedirect       uint8 = 5
	icmpEcho           uint8 = 8
	icmpTimeExceeded   uint8 = 11
	icmpParameterProb  uint8 = 12
	icmpTimestamp      uint8 = 13
	icmpTimestampReply uint8 = 14
	icmpInfoRequest    uint8 = 15
	icmpInfoReply      uint8 = 16
	icmpAddress        uint8 = 17
	icmpAddressReply   uint8 = 18

	// ICMP codes for unreach
	icmpNetUnreach    uint8 = 0
	icmpHostUnreach   uint8 = 1
	icmpProtUnreach   uint8 = 2
	icmpPortUnreach   uint8 = 3
	icmpFragNeeded    uint8 = 4
	icmpSRFailed      uint8 = 5
	icmpNetUnknown    uint8 = 6
	icmpHostUnknown   uint8 = 7
	icmpHostIsolated  uint8 = 8
	icmpNetAno        uint8 = 9
	icmpHostAno       uint8 = 10
	icmpNetUnrTos     uint8 = 11
	icmpHostUnrTos    uint8 = 12
	icmpPktFiltered   uint8 = 13
	icmpPrecViolation uint8 = 14
	icmpPrecCutOff    uint8 = 15

	// ICMP codes for redirect
	icmpRedirNet     uint8 = 0
	icmpRedirHost    uint8 = 1
	icmpRedirNetTos  uint8 = 2
	icmpRedirHostTos uint8 = 3

	// ICMP codes for time exceeded
	icmpExcTTL      uint8 = 0
	icmpExcFragTime uint8 = 1
)

type icmpHdrEcho struct {
	icmpType uint8
	icmpCode uint8
	checksum uint16
	id       uint16
	sequence uint16
}

func inCksum(b []byte, csum uint16) uint16 {

	sum := uint32(csum)

	bound := (len(b) / 2) * 2
	count := 0

	for count < bound {
		v := uint16(b[count+1])*256 + uint16(b[count])
		sum += uint32(v)
		count += 2
	}

	if bound < len(b) {
		sum += uint32(b[len(b)-1])
	}

	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)

	ans := ^sum
	ans &= 0xffff

	ans = (ans >> 8) | (ans << 8 & 0xff00)

	return uint16(ans)

}

func getTimestamp() []byte {

	tv := &syscall.Timeval{}

	err := syscall.Gettimeofday(tv)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = binary.Write(buf, binary.BigEndian, uint32(tv.Sec))
	if err != nil {
		log.Fatal(err)
	}
	err = binary.Write(buf, binary.BigEndian, uint32(tv.Usec))
	if err != nil {
		log.Fatal(err)
	}

	return buf.Bytes()
}

func genData(size int) []byte {

	out := []byte{}

	for i := 0; i < size; i++ {
		out = append(out, uint8(0x08+i))
	}

	return out
}

func icmpEchoHeaderMarshal(typ uint8, code uint8, sum uint16, id uint16, seq uint16) []byte {

	buf := new(bytes.Buffer)

	// Type
	err := binary.Write(buf, binary.BigEndian, typ)
	if err != nil {
		log.Fatal(err)
	}

	// Code
	err = binary.Write(buf, binary.BigEndian, code)
	if err != nil {
		log.Fatal(err)
	}

	// Checksum
	err = binary.Write(buf, binary.BigEndian, sum)
	if err != nil {
		log.Fatal(err)
	}

	// ID
	err = binary.Write(buf, binary.BigEndian, id)
	if err != nil {
		log.Fatal(err)
	}

	// Sequence
	err = binary.Write(buf, binary.BigEndian, seq)
	if err != nil {
		log.Fatal(err)
	}

	hdr := buf.Bytes()

	return hdr
}

func icmpPktParse(pkt []byte) (hdr *icmpHdrEcho, t *syscall.Timeval) {

	typ := pkt[0]
	code := pkt[1]
	sum := binary.BigEndian.Uint16(pkt[2:4])
	id := binary.BigEndian.Uint16(pkt[4:6])
	seq := binary.BigEndian.Uint16(pkt[6:8])

	hdr = &icmpHdrEcho{
		icmpType: typ,
		icmpCode: code,
		checksum: sum,
		id:       id,
		sequence: seq,
	}

	t = &syscall.Timeval{
		Sec:  int64(binary.BigEndian.Uint32(pkt[8:12])),
		Usec: int64(binary.BigEndian.Uint32(pkt[12:16])),
	}

	return
}

func icmpEchoMarshal(id uint16) []byte {

	sum := uint16(0)
	seq := uint16(0)

	hdr := icmpEchoHeaderMarshal(icmpEcho, icmpEchoReply, sum, id, seq)

	// Timestamp
	ts := getTimestamp()

	// Data
	data := genData(48)

	pkt := make([]byte, 0)
	pkt = append(pkt, hdr...)
	pkt = append(pkt, ts...)
	pkt = append(pkt, data...)

	sum = inCksum(pkt, 0)
	hdr = icmpEchoHeaderMarshal(icmpEcho, icmpEchoReply, sum, id, seq)

	pkt = make([]byte, 0)
	pkt = append(pkt, hdr...)
	pkt = append(pkt, ts...)
	pkt = append(pkt, data...)

	return pkt
}

func fdSet(p *syscall.FdSet, i int) {
	p.Bits[i/64] |= 1 << uint(i) % 64
}

func fdIsSet(p *syscall.FdSet, i int) bool {
	return (p.Bits[i/64] & (1 << uint(i) % 64)) != 0
}

func fdZero(p *syscall.FdSet) {
	for i := range p.Bits {
		p.Bits[i] = 0
	}
}

func sendOnePing(fd int, id uint16, ip net.IP) {

	addr := &syscall.SockaddrInet4{}

	copy(addr.Addr[:], ip.To4()[:])

	pkt := icmpEchoMarshal(id)

	err := syscall.Sendto(fd, pkt, 0, addr)
	if err != nil {
		log.Fatal(err)
	}

}

func recvOnePing(fd int, id uint16, timeout *syscall.Timeval) {

	buf := make([]byte, 65535)

	rfds := &syscall.FdSet{}

	recvtime := &syscall.Timeval{}

	for {

		fdZero(rfds)
		fdSet(rfds, fd)

		_, err := syscall.Select(fd+1, rfds, nil, nil, timeout)
		if err != nil {
			log.Fatal(err)
		}

		// timeout
		if !fdIsSet(rfds, fd) {
			break
		}

		err = syscall.Gettimeofday(recvtime)
		if err != nil {
			log.Fatal(err)
		}

		n, addr, err := syscall.Recvfrom(fd, buf, 0)
		if err != nil {
			log.Fatal(err)
		}

		icmpPkt := buf[20:n]

		hdr, t := icmpPktParse(icmpPkt[:n])
		if hdr.icmpType != icmpEchoReply && hdr.id != id {
			continue
		}

		delay := float64(recvtime.Sec-t.Sec) + float64(recvtime.Usec-t.Usec)/1000000.0

		a := addr.(*syscall.SockaddrInet4)
		log.Printf("from (%v.%v.%v.%v) time=%f sec\n", a.Addr[0], a.Addr[1], a.Addr[2], a.Addr[3], delay)

		break

	}

}

// Usage
//	$ ./ping github.com 2
// 	2018/01/22 23:27:14 ping.go:291: from (192.30.255.112) time=0.151557 sec
// 	2018/01/22 23:27:16 ping.go:291: from (192.30.255.112) time=0.150595 sec
func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if len(os.Args) != 3 {
		log.Fatal("usage: ping host count")
	}

	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	if err != nil {
		log.Fatal(err)
	}
	defer syscall.Close(fd)

	host := os.Args[1]

	// convert hostname to ip address
	addrs, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	// retrive ipv4 address
	var ipv4Addr net.IP
	for _, ip := range addrs {
		if ip.To4() == nil {
			continue
		}
		ipv4Addr = ip
		break
	}

	id := uint16(syscall.Getpid())
	timeout := &syscall.Timeval{Sec: int64(3)}

	count, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < count; i++ {
		sendOnePing(fd, id, ipv4Addr)
		recvOnePing(fd, id, timeout)

		time.Sleep(1 * time.Second)
	}
}
