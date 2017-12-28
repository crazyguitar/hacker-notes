package main

import (
	"log"
	"net/http"
)

const (
	port     = "4433"
	certFile = "cert.pem"
	keyFile  = "key.pem"
)

// usage:
//	$ openssl genrsa -out server.key 2048
//	$ openssl req -new -x509 -sha256 -key key.pem -out cert.pem -days 365 \
//		-subj "/C=TW/ST=Taiwan/L=Taipei/O=OrgName/OU=OrgUnitName/CN=example.com/"
// 	$ ./https &
func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))

	// start https service
	err := http.ListenAndServeTLS(":"+port, certFile, keyFile, nil)
	if err != nil {
		log.Fatal(err)
	}
}
