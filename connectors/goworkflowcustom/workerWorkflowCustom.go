package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/ditrit/gandalf/connectors/goworkflowcustom/server"
	"github.com/ditrit/gandalf/libraries/goclient"

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
	fmt.Println("Start")

	worker.Start()
	fmt.Println("Start 2")

	clientGandalf := worker.GetClientGandalf()
	fmt.Println("clientGandalf")
	fmt.Println(clientGandalf)
	fmt.Println("Start 3")

	testGithub(clientGandalf)
	//testUtils(clientGandalf)

	//toto := server.NewWorkflowServer(clientGandalf)
	//toto.Run()

}

func testDemo(clientGandalf *goclient.ClientGandalf) {
	toto := server.NewWorkflowServer(clientGandalf)
	toto.Run()

}

func testDocker(clientGandalf *goclient.ClientGandalf) {

	fmt.Println("SEND COMMMAND REGISTER")
	payload := `{"name":"test", "content": "{test}"}`
	commandMessageUUID := clientGandalf.SendCommand("Workflow.REGISTER", map[string]string{"payload": payload})
	fmt.Println(commandMessageUUID)

	time.Sleep(5 * time.Second)

	fmt.Println("SEND COMMMAND EXECUTE")
	payload = `{"name":"test"}`
	commandMessageUUID = clientGandalf.SendCommand("Workflow.EXECUTE", map[string]string{"payload": payload})
	fmt.Println(commandMessageUUID)

}

func testUtils(clientGandalf *goclient.ClientGandalf) {

	id := clientGandalf.CreateIteratorEvent()
	fmt.Println("Start 4")

	fmt.Println("SEND COMMMAND CREATE_FORM")
	payload := `{"Fields":[{"Name":"ID","HtmlType":"TextField","Value":"Id"}]}`
	commandMessageUUID := clientGandalf.SendCommand("Utils.CREATE_FORM", map[string]string{"payload": payload})
	fmt.Println(commandMessageUUID)

	time.Sleep(5 * time.Second)

	fmt.Println("SEND COMMMAND ADMIN_UPDATE")
	commandMessageUUIDupdate := clientGandalf.SendAdminCommand("Utils.ADMIN_UPDATE", map[string]string{})
	fmt.Println(commandMessageUUIDupdate)
	event := clientGandalf.WaitReplyByEvent("ADMIN_UPDATE", "SUCCES", commandMessageUUIDupdate, id)
	fmt.Println(event)

	time.Sleep(5 * time.Second)

	fmt.Println("SEND COMMMAND CREATE_FORM")
	payload = `{"Fields":[{"Name":"ID","HtmlType":"TextField","Value":"Id"}]}`
	commandMessageUUID = clientGandalf.SendCommand("Utils.CREATE_FORM", map[string]string{"payload": payload})
	fmt.Println(commandMessageUUID)

}

func testGithub(clientGandalf *goclient.ClientGandalf) {

	fmt.Println("Start 4")

	fmt.Println("SEND COMMMAND CREATE_REPOSITORY")
	payload := `{"Username":"","Password":"","Token":"","Name":"TestConnector","Description":"TestConnector","Private":true}`
	commandMessageUUID := clientGandalf.SendCommand("vcs.CREATE_REPOSITORY", map[string]string{"payload": payload})
	fmt.Println(commandMessageUUID)

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
