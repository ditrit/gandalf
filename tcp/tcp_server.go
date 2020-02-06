//Package tcp :
//File tcp_server.go
package main

import (
	"crypto/tls"
	"log"
	"net"
	"strings"
)

var (
	configServer tls.Config
)

// ServerTCP : Type socket Serveur
// type ServerTCP struct {
// 	socket net.Conn
// }
// TODO : check if this struct is useless

//serverTCP :
func serverTCP(connect string) {
	cert, _ := tls.LoadX509KeyPair("./cert.pem", "./key.pem")
	configServer = tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	listener, err := net.Listen("tcp", connect)

	if err != nil {
		log.Printf("server: listen %s", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("server: accept %s", err)
			break
		}

		go handleTLSConnection(conn)
	}
}

//handleTLSConnection :
func handleTLSConnection(unencConn net.Conn) {
	conn := tls.Server(unencConn, &configServer)
	buffer := make([]byte, 1024)

	defer conn.Close()

	conn.Handshake()

	for {
		bytesRead, err := conn.Read(buffer)

		if err != nil {
			log.Printf("server: conn.read %s", err)
			return
		}

		response := strings.TrimSpace(string(buffer[0:bytesRead]))

		_, err = conn.Write([]byte(response + "\n"))

		if err != nil {
			log.Printf("Server: conn.write %s", err)
			return
		}

		switch response {
		case "STOP":
			log.Printf("Server : exiting TCP server!")

			return
		case "ADDR":
			log.Printf("Server : Unencrypting connection from %s", conn.RemoteAddr())
		default:
			log.Printf("server : echoing : %s", response)
		}
	}
}
