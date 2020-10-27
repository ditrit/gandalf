package workers

import (
	"encoding/json"
	"fmt"

	"github.com/ditrit/gandalf/connectors/goutilscustom/mail"

	goclient "github.com/ditrit/gandalf/libraries/goclient"
	models "github.com/ditrit/gandalf/libraries/goclient/models"
)

type WorkerMail struct {
	clientGandalf *goclient.ClientGandalf
	major         int64
	minor         int64
	clientMail    *mail.MailClient
	address       string
	port          string
}

func NewWorkerMail(address, port string, clientGandalf *goclient.ClientGandalf, major, minor int64) *WorkerMail {
	workerMail := new(WorkerMail)
	workerMail.address = address
	workerMail.port = port
	workerMail.clientGandalf = clientGandalf
	workerMail.major = major
	workerMail.minor = minor

	return workerMail
}

func (r WorkerMail) Run() {
	done := make(chan bool)
	go r.SendAuthMail()
	<-done
}

func (r WorkerMail) SendAuthMail() {
	id := r.clientGandalf.CreateIteratorCommand()
	for true {
		command := r.clientGandalf.WaitCommand("SEND_AUTH_MAIL", id, r.major)

		var mailPayload mail.MailPayload
		err := json.Unmarshal([]byte(command.GetPayload()), &mailPayload)

		if err == nil {
			r.clientMail = mail.NewMailClient(r.address, r.port)

			auth := r.clientMail.Auth(mailPayload.Username, mailPayload.Password, r.address)

			result := r.clientMail.SendAuthMail(mailPayload.Sender, mailPayload.Body, mailPayload.Receivers, auth)
			if result {
				r.clientGandalf.SendReply(command.GetCommand(), "SUCCES", command.GetUUID(), models.NewOptions("", ""))
			} else {
				r.clientGandalf.SendReply(command.GetCommand(), "FAIL", command.GetUUID(), models.NewOptions("", ""))
			}
		}
	}
}

func (r WorkerMail) SendAuthMailTest(sender, body, login, password, email string, receivers []string) (result bool) {
	result = true

	r.clientMail = mail.NewMailClient(r.address, r.port)
	auth := r.clientMail.Auth(login, password, r.address)
	fmt.Println(login)
	fmt.Println(password)
	fmt.Println(r.address)
	fmt.Println(auth)
	result = r.clientMail.SendAuthMail(sender, body, receivers, auth)

	return result
}
