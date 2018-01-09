package tls

import (
	"crypto/x509"
	"encoding/binary"
	"log"
)

func ParseCertificate(pkt []byte) {

	_ = pkt[0] // pkt[0] is handshake Type

	buf := make([]byte, 4)
	copy(buf[1:], pkt[1:4])
	_ = binary.BigEndian.Uint32(buf[:]) // length 24 bit

	copy(buf[1:], pkt[4:7])
	certsLength := binary.BigEndian.Uint32(buf[:]) // 24 bit
	certs := pkt[7:certsLength]

	offset := uint32(0)

	arr := [](*x509.Certificate){}

	for offset < certsLength {
		copy(buf[1:], certs[offset:offset+3])
		cl := binary.BigEndian.Uint32(buf[:])

		offset += 3
		asn1data := certs[offset : offset+cl]
		cert, err := x509.ParseCertificate(asn1data)
		if err != nil {
			log.Fatal(err)
		}

		// append certificate
		arr = append(arr, cert)

		offset += cl
	}
}
