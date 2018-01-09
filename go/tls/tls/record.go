package tls

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
)

type Record struct {
	ContentType     uint8
	ProtocolVersion uint16
	Length          uint16
	Data            []byte
}

func (r Record) String() string {
	return fmt.Sprintf("Content Type: %v\nVersion: %v\nRecord:\n%v", r.ContentType, r.ProtocolVersion, hex.Dump(r.Data))
}

// TLSRecord marshall the packet of TLS record layer
func TLSRecord(protocol []byte) []byte {

	// Record layer
	record := Record{
		ContentType:     Handshake,
		ProtocolVersion: TLSv10,
		Length:          uint16(len(protocol)),
	}

	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.BigEndian, uint8(record.ContentType))
	if err != nil {
		log.Fatal(err)
	}
	err = binary.Write(buf, binary.BigEndian, uint16(record.ProtocolVersion))
	if err != nil {
		log.Fatal(err)
	}
	err = binary.Write(buf, binary.BigEndian, uint16(record.Length))
	if err != nil {
		log.Fatal(err)
	}

	out := buf.Bytes()

	out = append(out, protocol...)

	return out
}

// ParseTLSRecord parse the tls record
func ParseTLSRecord(pkt []byte) []Record {

	offset := 0

	rec := make([]Record, 0)

	for offset < len(pkt) {
		r := Record{
			ContentType:     uint8(pkt[offset]),
			ProtocolVersion: binary.BigEndian.Uint16(pkt[offset+1 : offset+3]),
			Length:          binary.BigEndian.Uint16(pkt[offset+3 : offset+5]),
		}

		offset += 5
		r.Data = pkt[offset : offset+int(r.Length)]
		offset += int(r.Length)

		rec = append(rec, r)
	}

	return rec
}

// GetRecordProtocol retrive the protocol type
func GetRecordProtocol(r *Record) uint8 {
	return uint8(r.Data[0])
}
