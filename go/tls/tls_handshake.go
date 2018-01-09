package main

import (
	"./tls"

	"bufio"
	"encoding/hex"
	"log"
	"net"
)

func checkMapAll(m map[uint8]bool) bool {
	for _, v := range m {
		if v == false {
			return false
		}
	}
	return true
}

// Usage
// 	$ openssl genrsa -out key.pem 2048
//	$ openssl req -new -x509 -sha256 -key key.pem -out cert.pem -days 365 \
//		-subj "/C=TW/ST=Taiwan/L=Taipei/O=OrgName/OU=OrgUnitName/CN=example.com/"
//	$ openssl s_server -key key.pem -cert cert.pem -accept 443 -www &
// 	$ ./tls_handshake
func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	host := "localhost:443"

	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)

	cliMsg := tls.ClientHelloField()
	tlsMsg := tls.TLSRecord(cliMsg)

	log.Printf("\n%s\n", hex.Dump(tlsMsg))

	// send client Hello
	w.Write(tlsMsg)
	w.Flush()

	// recv server hello
	m := map[uint8]bool{
		tls.ServerHello:       false,
		tls.Certificate:       false,
		tls.ServerKeyExchange: false,
		tls.ServerHelloDone:   false,
	}

	rec := make([]tls.Record, 0)

	for !checkMapAll(m) {

		buf := make([]byte, 4096)
		n, err := r.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		rec = append(rec, tls.ParseTLSRecord(buf[:n])...)

		for _, r := range rec {
			p := tls.GetRecordProtocol(&r)
			switch p {
			case tls.ServerHello:
				m[tls.ServerHello] = true
				tls.ParseServerHello(r.Data[:])
			case tls.Certificate:
				m[tls.Certificate] = true
				tls.ParseCertificate(r.Data[:])
			case tls.ServerKeyExchange:
				m[tls.ServerKeyExchange] = true
				tls.ParseServerKeyExchange(r.Data[:])
			case tls.ServerHelloDone:
				m[tls.ServerHelloDone] = true
				tls.ParseServerHelloDone(r.Data[:])
			}
		}
	}

	// client key exchange
}
