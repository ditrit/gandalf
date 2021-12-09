package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ditrit/gandalf/libraries/goclient/models"

	worker "github.com/ditrit/gandalf/connectors/go"
)

//main : main
func main() {

	var major = int64(1)
	var minor = int64(5)

	fmt.Println("VERSION")
	fmt.Println(major)
	fmt.Println(minor)

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(input.Text())

	worker := worker.NewWorker(major, minor)
	worker.Start()

	payload := `{"Fields":[{"Name":"ID","HtmlType":"TextField","Value":"Id"}]}`
	commandMessageUUID := worker.GetClientGandalf().SendCommand("Utils.CREATE_FORM", models.NewOptions("", payload))
	formUUID := commandMessageUUID.GetUUID()

	id := worker.GetClientGandalf().CreateIteratorEvent()
	event := worker.GetClientGandalf().WaitReplyByEvent("CREATE_FORM", "SUCCES", formUUID, id)
	fmt.Println(event)
}
