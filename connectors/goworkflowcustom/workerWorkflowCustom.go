package main

import (
	"fmt"

	"github.com/ditrit/gandalf/connectors/goworkflowcustom/workers"

	goworkflow "github.com/ditrit/gandalf/connectors/goworkflow"

	goclient "github.com/ditrit/gandalf/libraries/goclient"
)

func main() {

	var commands = []string{}
	var version = int64(1)

	workerWorkflow := goworkflow.NewWorkerWorkflow(version, commands)
	fmt.Println("workerWorkflow.Upload")
	fmt.Println(workerWorkflow.Upload)
	fmt.Println("workerWorkflow.Upload")

	workerWorkflow.Upload = Upload
	fmt.Println("RUN CUSTOM")
	fmt.Println("workerWorkflow.Upload")
	fmt.Println(workerWorkflow.Upload)
	fmt.Println("workerWorkflow.Upload")
	workerWorkflow.Run()
}

//Upload
func Upload(clientGandalf *goclient.ClientGandalf, version int64) {
	/* var configuration Configuration
	mydir, _ := os.Getwd()
	file, _ := os.Open(mydir + "/test.json")
	decoder := json.NewDecoder(file)
	decoder.Decode(&configuration) */
	//done := make(chan bool)
	fmt.Println("UPLOAD")
	/* 	input := bufio.NewScanner(os.Stdin)
	   	input.Scan()
	   	fmt.Println(input.Text())
	   	fmt.Println("THX") */

	workerMail := workers.NewWorkerUpload(clientGandalf)
	go workerMail.Run()
	//<-done

}

type Configuration struct {
	Identity    string
	Connections []string
}
