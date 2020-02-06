//Package tcp :
//File tcp_client.go
package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"
)

var (
	configClient tls.Config
)

//initTLS :
func initTLS(connect string) (*tls.Conn, error) {
	CAPool := x509.NewCertPool()
	serverCert, err := ioutil.ReadFile("./cert.pem")

	if err != nil {
		//log.Print("initTLs : Could not load server certificate!")
		return nil, err
	}

	CAPool.AppendCertsFromPEM(serverCert)

	configClient = tls.Config{
		RootCAs:            CAPool,
		InsecureSkipVerify: true, //nolint: gosec
	}

	unencConn, err := net.Dial("tcp", connect)
	if err != nil {
		//log.Printf("initTLs : net.Dial %s", err)
		return nil, err
	}

	conn := tls.Client(unencConn, &configClient)
	err = conn.Handshake()

	if err != nil {
		//log.Printf("initTLs : tls handshake %s", err)
		conn.Close()
		return nil, err
	}

	return conn, nil
}

//clientTCP :
func clientTCP(connect string) {
	var buffer = make([]byte, 1024)

	for {
		conn, err := initTLS(connect)

		if err != nil {
			log.Printf("clientTcp : %s", err)
			break // TODO : define behavior
		}

		defer conn.Close()

		for {
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			fmt.Println(">> " + text)

			_, err := conn.Write([]byte(text))
			if err != nil {
				//log.Printf("Client : conn.Write %s", err)
				break
			}

			bytesRead, err := conn.Read(buffer)
			if err != nil {
				//log.Printf("Client: conn.Read %s", err)
				break
			}

			response := string(buffer[0:bytesRead])
			fmt.Print("->: " + response)

			if response == "STOP" {
				fmt.Println("TCP client exiting...")
				return
			}
		}

		time.Sleep(100 * time.Millisecond)
		fmt.Println("trying to reconnect")
	}
}
