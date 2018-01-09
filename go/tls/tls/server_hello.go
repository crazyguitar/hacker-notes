package tls

import (
	"encoding/binary"
	"encoding/hex"
	"log"
)

type RandomField struct {
	gmtUnixTime uint32
	randomBytes [28]byte
}

type ProtocolServerHello struct {
	Version           uint16
	Random            RandomField
	SessionIDLength   uint8
	CipherSuite       uint16
	CompressionMethod uint8
	Extensions        []byte
}

type DHServerParams struct {
	PLength      uint16
	P            []byte
	GLength      uint16
	G            []byte
	PubkeyLen    uint16
	Pubkey       []byte
	SignatureLen uint16
	Signature    []byte
}

func ParseServerHello(pkt []byte) (p *ProtocolServerHello) {

	_ = pkt[0]   // pkt[0] is handshake protocol
	_ = pkt[1:4] // pkt[1:4] is length

	version := binary.BigEndian.Uint16(pkt[4:6]) // pkt[4:6] is TLS version
	random := RandomField{
		gmtUnixTime: binary.BigEndian.Uint32(pkt[6:10]),
	}
	copy(random.randomBytes[:], pkt[10:38])

	sessionIDLength := uint8(pkt[38])
	cipherSuite := binary.BigEndian.Uint16(pkt[39:41])
	compressionMethon := uint8(pkt[41])
	extLen := binary.BigEndian.Uint16(pkt[42:44])
	extensions := pkt[44 : 44+extLen]

	return &ProtocolServerHello{
		Version:           version,
		Random:            random,
		SessionIDLength:   sessionIDLength,
		CipherSuite:       cipherSuite,
		CompressionMethod: compressionMethon,
		Extensions:        extensions,
	}
}

func ParseServerKeyExchange(pkt []byte) {
	_ = pkt[0] // pkt[0] is handshake protocol

	buf := make([]byte, 4)
	copy(buf[1:], pkt[1:4])
	length := binary.BigEndian.Uint32(buf) // pkt[1:4] is length
	params := pkt[4:length]                // pkt[4:length] is the pparameters of key exchange

	log.Printf("\n%s\n", hex.Dump(params))

}

func ParseServerHelloDone(pkt []byte) {

	_ = pkt[0]   // pkt[0] is handshake protocol
	_ = pkt[1:4] // pkt[1:4] is legnth
}
