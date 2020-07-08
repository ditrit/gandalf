package main

import (
	"encoding/json"
	"os"

	"github.com/ditrit/gandalf/connectors/goutils/workers"

	worker "github.com/ditrit/gandalf/connectors/go"

	goclient "github.com/ditrit/gandalf/libraries/goclient"
)

func main() {

	var commands = []string{"SEND_AUTH_MAIL", "CREATE_FORM"}
	var version = int64(2)

	workerUtils := worker.NewWorker(version, commands)
	//workerUtils.Execute = Execute

	workerUtils.CreateApplication = CreateApplication
	workerUtils.CreateForm = CreateForm
	workerUtils.SendAuthMail = SendAuthMail

	workerUtils.Run()
}

//CreateApplication
func CreateApplication(clientGandalf *goclient.ClientGandalf, version int64) {
	var configuration Configuration
	mydir, _ := os.Getwd()
	file, _ := os.Open(mydir + "/test.json")
	decoder := json.NewDecoder(file)
	decoder.Decode(&configuration)

	workerApp := workers.NewWorkerApplication(clientGandalf, version)
	go workerApp.Run()
}

//CreateForm
func CreateForm(clientGandalf *goclient.ClientGandalf, version int64) {
	var configuration Configuration
	mydir, _ := os.Getwd()
	file, _ := os.Open(mydir + "/test.json")
	decoder := json.NewDecoder(file)
	decoder.Decode(&configuration)

	workerForm := workers.NewWorkerForm(clientGandalf, version)
	go workerForm.Run()
}

//SendAuthMail
func SendAuthMail(clientGandalf *goclient.ClientGandalf, version int64) {
	var configuration Configuration
	mydir, _ := os.Getwd()
	file, _ := os.Open(mydir + "/test.json")
	decoder := json.NewDecoder(file)
	decoder.Decode(&configuration)

	workerMail := workers.NewWorkerMail(configuration.Address, configuration.Port, clientGandalf, version)
	go workerMail.Run()
}

type Configuration struct {
	Address string
	Port    string
	Contact string
	Pwd     string
	Mail    string
}
