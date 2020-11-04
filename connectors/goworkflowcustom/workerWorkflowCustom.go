package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/ditrit/gandalf/libraries/goclient/models"

	worker "github.com/ditrit/gandalf/connectors/go"
)

//main : main
func main() {

	var major = int64(1)
	var minor = int64(0)

	fmt.Println("VERSION")
	fmt.Println(major)
	fmt.Println(minor)

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(input.Text())

	worker := worker.NewWorker(major, minor)
	clientGandalf := worker.Start()

	id := clientGandalf.CreateIteratorEvent()

	for true {

		fmt.Println("SEND COMMMAND CREATE_FORM")
		payload := `{"Fields":[{"Name":"ID","HtmlType":"TextField","Value":"Id"}]}`
		commandMessageUUID := clientGandalf.SendCommand("Utils.CREATE_FORM", models.NewOptions("", payload))
		formUUID := commandMessageUUID.GetUUID()

		fmt.Println("formUUID")
		fmt.Println(formUUID)
		event := clientGandalf.WaitReplyByEvent("CREATE_FORM", "SUCCES", formUUID, id)
		fmt.Println("event")
		fmt.Println(event)

		time.Sleep(5 * time.Second)
	}

	//workerUpload := workers.NewWorkerUpload(clientGandalf)
	//go workerUpload.Run()
}

/* //Upload : Upload
func Upload(clientGandalf *goclient.ClientGandalf, version int64) {
	/* var configuration Configuration
	mydir, _ := os.Getwd()
	file, _ := os.Open(mydir + "/test.json")
	decoder := json.NewDecoder(file)
	decoder.Decode(&configuration) */
//done := make(chan bool)

//workerMail := workers.NewWorkerUpload(clientGandalf)
//go workerMail.Run()
//<-done

//} */

//Configuration : Configuration
type Configuration struct {
	Identity    string
	Connections []string
}
