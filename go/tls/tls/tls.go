package tls

const (

	// TLS/SSL version
	SSLv3  uint16 = 0x0300
	TLSv10 uint16 = 0x0301
	TLSv11 uint16 = 0x0302
	TLSv12 uint16 = 0x0303

	// Content Type
	ChangeCipherSpec uint8 = 20
	Alert            uint8 = 21
	Handshake        uint8 = 22
	PpplicationData  uint8 = 23

	// Handshake Type
	HelloRequest       uint8 = 0
	ClientHello        uint8 = 1
	ServerHello        uint8 = 2
	Certificate        uint8 = 11
	ServerKeyExchange  uint8 = 12
	CertificateRequest uint8 = 13
	ServerHelloDone    uint8 = 14
	CertificateVerify  uint8 = 15
	ClientKeyExchange  uint8 = 16
	Finished           uint8 = 20
)
