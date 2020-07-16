package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ditrit/gandalf/connectors/goworkflowcustom/workers"

	goworkflow "github.com/ditrit/gandalf/connectors/goworkflow"

	goclient "github.com/ditrit/gandalf/libraries/goclient"
)

//main : main
func main() {

	var commands = []string{}
	var version = int64(1)

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(input.Text())

	workerWorkflow := goworkflow.NewWorkerWorkflow(version, commands)
	workerWorkflow.Upload = Upload
	workerWorkflow.Run()
}

//Upload : Upload
func Upload(clientGandalf *goclient.ClientGandalf, version int64) {
	/* var configuration Configuration
	mydir, _ := os.Getwd()
	file, _ := os.Open(mydir + "/test.json")
	decoder := json.NewDecoder(file)
	decoder.Decode(&configuration) */
	//done := make(chan bool)

	workerMail := workers.NewWorkerUpload(clientGandalf)
	go workerMail.Run()
	//<-done

}

//Configuration : Configuration
type Configuration struct {
	Identity    string
	Connections []string
}
