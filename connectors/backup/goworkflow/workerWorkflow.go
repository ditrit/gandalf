package main

import (
	"encoding/json"
	"os"

	"github.com/ditrit/gandalf/connectors/goworkflow/workers"

	worker "github.com/ditrit/gandalf/connectors/go"

	goclient "github.com/ditrit/gandalf/libraries/goclient"
)

func main() {

	var commands = []string{}
	var version = int64(1)

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
	//done := make(chan bool)
	workerMail := workers.NewWorkerUpload(clientGandalf)
	go workerMail.Run()
	//<-done

}

type Configuration struct {
	Identity    string
	Connections []string
}
