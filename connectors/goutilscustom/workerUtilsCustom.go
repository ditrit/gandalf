package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/ditrit/gandalf/connectors/goutilscustom/workers"

	goutils "github.com/ditrit/gandalf/connectors/goutils"

	goclient "github.com/ditrit/gandalf/libraries/goclient"
)

//main : main
func main() {

	var commands = []string{"SEND_AUTH_MAIL", "CREATE_FORM"}
	var major = int64(2)
	var minor = int64(0)

	fmt.Println("VERSION")
	fmt.Println(major)
	fmt.Println(minor)

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(input.Text())

	workerUtils := goutils.NewWorkerUtils(major, minor, commands)
	workerUtils.CreateApplication = CreateApplication
	workerUtils.CreateForm = CreateForm
	workerUtils.SendAuthMail = SendAuthMail

	workerUtils.Run()
}

//CreateApplication : CreateApplication
func CreateApplication(clientGandalf *goclient.ClientGandalf, major, minor int64) {
	var configuration Configuration
	mydir, _ := os.Getwd()
	file, _ := os.Open(mydir + "/test.json")
	decoder := json.NewDecoder(file)
	decoder.Decode(&configuration)

	workerApp := workers.NewWorkerApplication(clientGandalf, major, minor)
	go workerApp.Run()
}

//CreateForm : CreateForm
func CreateForm(clientGandalf *goclient.ClientGandalf, major, minor int64) {
	var configuration Configuration
	mydir, _ := os.Getwd()
	file, _ := os.Open(mydir + "/test.json")
	decoder := json.NewDecoder(file)
	decoder.Decode(&configuration)

	workerForm := workers.NewWorkerForm(clientGandalf, major, minor)
	go workerForm.Run()
}

//SendAuthMail : SendAuthMail
func SendAuthMail(clientGandalf *goclient.ClientGandalf, major, minor int64) {
	var configuration Configuration
	mydir, _ := os.Getwd()
	file, _ := os.Open(mydir + "/test.json")
	decoder := json.NewDecoder(file)
	decoder.Decode(&configuration)

	workerMail := workers.NewWorkerMail(configuration.Address, configuration.Port, clientGandalf, major, minor)
	go workerMail.Run()
}

//Configuration : Configuration
type Configuration struct {
	Address string
	Port    string
	Contact string
	Pwd     string
	Mail    string
}
