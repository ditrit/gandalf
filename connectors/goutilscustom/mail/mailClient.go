package mail

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"log"
	"net/smtp"
)

type MailClient struct {
	address string
	port    string
	rootUrl string
	client  *smtp.Client
}

func NewMailClient(address, port string) *MailClient {
	mailClient := new(MailClient)
	mailClient.address = address
	mailClient.port = port
	mailClient.rootUrl = mailClient.address + mailClient.port

	return mailClient
}

func (m MailClient) Auth(login, password, address string) smtp.Auth {
	auth := smtp.PlainAuth(
		"",
		login,
		password,
		address,
	)
	return auth
}

func (m MailClient) SendAuthMail(sender, body string, receivers []string, auth smtp.Auth) (result bool) {
	result = true

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         m.address,
	}

	conn, err := tls.Dial("tcp", m.rootUrl, tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	m.client, err = smtp.NewClient(conn, m.address)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = m.client.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = m.client.Mail(sender); err != nil {
		log.Panic(err)
	}

	for _, receiver := range receivers {
		if err = m.client.Rcpt(receiver); err != nil {
			log.Panic(err)
		}
	}

	//MESSAGE
	header := make(map[string]string)
	header["Subject"] = "Validation"
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	// Data
	w, err := m.client.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	m.client.Quit()

	return result
}
