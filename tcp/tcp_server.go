package main

import (
	"crypto/tls"
	"log"
	"net"
	"strings"
)

var (
	config_server tls.Config
)

func serverTcp(connect string) {

	cert, _ := tls.LoadX509KeyPair("./cert.pem", "./key.pem")
	config_server = tls.Config{
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
		go handleTlsConnection(conn)
	}
}

func handleTlsConnection(unenc_conn net.Conn) {
	conn := tls.Server(unenc_conn, &config_server)
	defer conn.Close()
	buffer := make([]byte, 1024)
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
			conn.Close()
			return
		case "ADDR":
			log.Printf("Server : Unencrypting connection from %s", conn.RemoteAddr())
		default:
			log.Printf("server : echoing : %s", response)
		}
	}
}
