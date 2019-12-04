package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"
)

var (
	config_client tls.Config
)

func initTls(connect string) (*tls.Conn, error) {

	CA_Pool := x509.NewCertPool()
	serverCert, err := ioutil.ReadFile("./cert.pem")
	if err != nil {
		//log.Print("initTLs : Could not load server certificate!")
		return nil, err
	}
	CA_Pool.AppendCertsFromPEM(serverCert)

	config_client = tls.Config{
		RootCAs:            CA_Pool,
		InsecureSkipVerify: true,
	}

	unenc_conn, err := net.Dial("tcp", connect)
	if err != nil {
		//log.Printf("initTLs : net.Dial %s", err)
		return nil, err
	}

	conn := tls.Client(unenc_conn, &config_client)
	err = conn.Handshake()
	if err != nil {
		//log.Printf("initTLs : tls handshake %s", err)
		conn.Close()
		return nil, err
	}

	return conn, nil
}

func clientTcp(connect string) {

	var buffer = make([]byte, 1024)

	for {
		conn, err := initTls(connect)
		defer conn.Close()

		if err != nil {
			//log.Printf("clientTcp : %s", err)
		} else {
			for {
				reader := bufio.NewReader(os.Stdin)
				fmt.Print(">> ")
				text, _ := reader.ReadString('\n')
				fmt.Print(text + "\n")

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
		}

		time.Sleep(100 * time.Millisecond)
		//fmt.Print("trying to reconnect")
	}
}
