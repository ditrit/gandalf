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
	workerUtils.Execute = Execute

	workerUtils.Run()
}

//
func Execute(clientGandalf *goclient.ClientGandalf, version int64) {
	var configuration Configuration
	mydir, _ := os.Getwd()
	file, _ := os.Open(mydir + "/test.json")
	decoder := json.NewDecoder(file)
	decoder.Decode(&configuration)

	workerForm := workers.NewWorkerForm(clientGandalf, version)
	go workerForm.Run()

	workerMail := workers.NewWorkerMail(configuration.Address, configuration.Port, clientGandalf, version)
	go workerMail.Run()

	workerApp := workers.NewWorkerApplication(clientGandalf, version)
	go workerApp.Run()
}

type Configuration struct {
	Address string
	Port    string
	Contact string
	Pwd     string
	Mail    string
}
