package tls

import (
	"bytes"
	"encoding/binary"
	"log"
	"math/rand"
	"time"
)

func randomField() []byte {

	random := struct {
		gmtUnixTime uint32
		randomBytes [28]byte
	}{}

	random.gmtUnixTime = uint32(time.Now().Unix())
	rand.Read(random.randomBytes[:])

	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.BigEndian, uint32(random.gmtUnixTime))
	if err != nil {
		log.Fatal(err)
	}

	out := buf.Bytes()

	out = append(out, random.randomBytes[:]...)

	return out
}

func cipherSuitesField() []byte {

	cipherSuites := []uint16{
		TLS_DHE_RSA_WITH_AES_256_CBC_SHA,
		TLS_DHE_DSS_WITH_AES_256_CBC_SHA,
		TLS_RSA_WITH_AES_256_CBC_SHA,
		TLS_DHE_RSA_WITH_3DES_EDE_CBC_SHA,
		TLS_DHE_DSS_WITH_3DES_EDE_CBC_SHA,
		TLS_RSA_WITH_3DES_EDE_CBC_SHA,
		TLS_DHE_RSA_WITH_AES_128_CBC_SHA,
		TLS_DHE_DSS_WITH_AES_128_CBC_SHA,
		TLS_DHE_RSA_WITH_SEED_CBC_SHA,
		TLS_DHE_DSS_WITH_SEED_CBC_SHA,
		TLS_RSA_WITH_SEED_CBC_SHA,
		TLS_RSA_WITH_RC4_128_SHA,
		TLS_RSA_WITH_RC4_128_MD5,
		TLS_DHE_RSA_WITH_DES_CBC_SHA,
		TLS_DHE_RSA_WITH_DES_CBC_SHA,
		TLS_DHE_DSS_WITH_DES_CBC_SHA,
		TLS_RSA_WITH_DES_CBC_SHA,
		TLS_DHE_RSA_EXPORT_WITH_DES40_CBC_SHA,
		TLS_DHE_DSS_EXPORT_WITH_DES40_CBC_SHA,
		TLS_RSA_EXPORT_WITH_DES40_CBC_SHA,
		TLS_RSA_EXPORT_WITH_RC2_CBC_40_MD5,
		TLS_RSA_EXPORT_WITH_RC4_40_MD5,
	}

	buf := new(bytes.Buffer)

	for _, v := range cipherSuites {

		err := binary.Write(buf, binary.BigEndian, uint16(v))
		if err != nil {
			log.Fatal(err)
		}
	}

	out := buf.Bytes()
	return out
}

func extSessionTicketTLSField() []byte {

	field := struct {
		msgType uint16
		data    []byte
	}{
		msgType: uint16(0x0023),
	}

	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.BigEndian, uint16(field.msgType))
	if err != nil {
		log.Fatal(err)
	}

	err = binary.Write(buf, binary.BigEndian, uint16(len(field.data)))
	if err != nil {
		log.Fatal(err)
	}

	out := buf.Bytes()

	out = append(out, field.data[:]...)
	return out
}

// ClientHelloField marshall the packet of Client Handshake Protocol
func ClientHelloField() []byte {

	rField := randomField()              // Random
	sField := []byte{0x0}                // Session ID
	cField := cipherSuitesField()        // Cipher Suite
	mField := []byte{0x0}                // Compression Method
	eField := extSessionTicketTLSField() // Extension: SessionTicket TLS

	cFieldLenBuf := make([]byte, 2) // Cipher Suite length
	mFieldLenBuf := make([]byte, 1) // Compression Field length
	eFieldLenBuf := make([]byte, 2) // Extension Field length

	binary.BigEndian.PutUint16(cFieldLenBuf, uint16(len(cField)))
	binary.BigEndian.PutUint16(eFieldLenBuf, uint16(len(eField)))
	mFieldLenBuf[0] = uint8(len(mField))

	pkt := make([]byte, 0)

	pkt = append(pkt, rField[:]...)
	pkt = append(pkt, sField[:]...)
	pkt = append(pkt, cFieldLenBuf[:]...)
	pkt = append(pkt, cField[:]...)
	pkt = append(pkt, mFieldLenBuf[:]...)
	pkt = append(pkt, mField[:]...)
	pkt = append(pkt, eFieldLenBuf[:]...)
	pkt = append(pkt, eField[:]...)

	length := uint32(len(pkt)) + uint32(2) // payload len + version len

	// Handshake Protocol
	handshake := struct {
		msgType uint8
		length  uint32 // only 3 bytes will be used
		version uint16
	}{
		msgType: ClientHello,
		length:  length,
		version: TLSv10,
	}

	out := make([]byte, 0)

	// append msg type
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, uint8(handshake.msgType))
	if err != nil {
		log.Fatal(err)
	}
	out = append(out, buf.Bytes()[:]...)

	// append handshake protocol length
	buf = new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, uint32(handshake.length))
	if err != nil {
		log.Fatal(err)
	}
	out = append(out, buf.Bytes()[1:]...) // only 24 bits

	// append tls version
	buf = new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, uint16(handshake.version))
	if err != nil {
		log.Fatal(err)
	}
	out = append(out, buf.Bytes()[:]...)

	// append handshake protocol
	out = append(out, pkt...)

	return out
}
