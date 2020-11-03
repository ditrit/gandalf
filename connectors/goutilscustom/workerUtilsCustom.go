package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	worker "github.com/ditrit/gandalf/connectors/go"
	"github.com/ditrit/shoset/msg"

	"github.com/ditrit/gandalf/connectors/goutilscustom/form"
	"github.com/ditrit/gandalf/connectors/goutilscustom/mail"
	"github.com/ditrit/gandalf/libraries/goclient"
)

//main : main
func main() {

	var major = int64(2)
	var minor = int64(2)

	fmt.Println("VERSION")
	fmt.Println(major)
	fmt.Println(minor)

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(input.Text())

	worker := worker.NewWorker(major, minor)
	worker.RegisterCommandsFuncs("CREATE_FORM", CreateForm)
	worker.RegisterCommandsFuncs("SEND_AUTH_MAIL", SendAuthMail)

	worker.Run()
}

func SendAuthMail(clientGandalf *goclient.ClientGandalf, major, minor int64, command msg.Command) int {

	var configuration Configuration
	mydir, _ := os.Getwd()
	file, _ := os.Open(mydir + "/test.json")
	decoder := json.NewDecoder(file)
	decoder.Decode(&configuration)

	var mailPayload mail.MailPayload
	err := json.Unmarshal([]byte(command.GetPayload()), &mailPayload)

	if err == nil {
		clientmail := mail.NewMailClient(configuration.Address, configuration.Port)

		auth := clientmail.Auth(mailPayload.Username, mailPayload.Password, configuration.Address)

		clientmail.SendAuthMail(mailPayload.Sender, mailPayload.Body, mailPayload.Receivers, auth)

		return 0
	}
	return 1

}

func CreateForm(clientGandalf *goclient.ClientGandalf, major, minor int64, command msg.Command) int {

	var formPayload form.FormPayload
	err := json.Unmarshal([]byte(command.GetPayload()), &formPayload)

	if err == nil {
		clientFormServer := form.NewFormServer(command.GetUUID(), formPayload, clientGandalf)
		go clientFormServer.Run()

		return 0
	}
	return 1

}

/* //CreateApplication : CreateApplication
func CreateApplication(clientGandalf *goclient.ClientGandalf, major, minor int64) {
	var configuration Configuration
	mydir, _ := os.Getwd()
	file, _ := os.Open(mydir + "/test.json")
	decoder := json.NewDecoder(file)
	decoder.Decode(&configuration)

	workerApp := workers.NewWorkerApplication(clientGandalf, major, minor)
	//go workerApp.Run()
}

//CreateForm : CreateForm
func CreateForm(clientGandalf *goclient.ClientGandalf, major, minor int64) {
	var configuration Configuration
	mydir, _ := os.Getwd()
	file, _ := os.Open(mydir + "/test.json")
	decoder := json.NewDecoder(file)
	decoder.Decode(&configuration)

	workerForm := workers.NewWorkerForm(clientGandalf, major, minor)
	//go workerForm.Run()
}

//SendAuthMail : SendAuthMail
func SendAuthMail(clientGandalf *goclient.ClientGandalf, major, minor int64) {
	var configuration Configuration
	mydir, _ := os.Getwd()
	file, _ := os.Open(mydir + "/test.json")
	decoder := json.NewDecoder(file)
	decoder.Decode(&configuration)

	workerMail := workers.NewWorkerMail(configuration.Address, configuration.Port, clientGandalf, major, minor)
	//go workerMail.Run()
} */

//Configuration : Configuration
type Configuration struct {
	Address string
	Port    string
	Contact string
	Pwd     string
	Mail    string
}
