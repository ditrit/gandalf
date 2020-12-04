package main

import (
	"fmt"
	"time"

	"github.com/ditrit/gandalf/libraries/goclient/models"

	goclient "github.com/ditrit/gandalf/libraries/goclient"
)

func main() {
	fmt.Println("START")
	testWorkflow()
	//demo()
	//testgitlab()
	fmt.Println("STOP")
}

func testWorkflow() {

	clientGandalf := goclient.NewClientGandalf("identity", "timeout", []string{"socket"})

	id := clientGandalf.CreateIteratorEvent()

	fmt.Println("SEND COMMMAND CREATE_FORM")
	payload := `{"Fields":[{"Name":"ID","HtmlType":"TextField","Value":"Id"}]}`
	commandMessageUUID := clientGandalf.SendCommand("Utils.CREATE_FORM", models.NewOptions("", payload))
	formUUID := commandMessageUUID.GetUUID()
	fmt.Println(formUUID)

	time.Sleep(5 * time.Second)

	fmt.Println("SEND COMMMAND ADMIN_UPDATE")
	commandMessageUUIDupdate := clientGandalf.SendAdminCommand("Utils.ADMIN_UPDATE", models.NewOptions("", `""`))
	updateUUID := commandMessageUUIDupdate.GetUUID()
	fmt.Println(updateUUID)
	event := clientGandalf.WaitReplyByEvent("ADMIN_UPDATE", "SUCCES", updateUUID, id)
	fmt.Println(event)

	time.Sleep(5 * time.Second)

	fmt.Println("SEND COMMMAND CREATE_FORM")
	payload = `{"Fields":[{"Name":"ID","HtmlType":"TextField","Value":"Id"}]}`
	commandMessageUUID = clientGandalf.SendCommand("Utils.CREATE_FORM", models.NewOptions("", payload))
	formUUID = commandMessageUUID.GetUUID()
}

/*
func testgitlab() {
	var configuration Configuration
	mydir, _ := os.Getwd()
	file, _ := os.Open(mydir + "/demoWorkflow.json")
	decoder := json.NewDecoder(file)
	decoder.Decode(&configuration)
	client := goclient.NewClientGandalf(configuration.Identity, configuration.Timeout, configuration.Connections)

	id := client.CreateIteratorEvent()

	payload := `{"Name":"` + "test" + `","Team":"` + "test" + `","TemplateName":` + "test" + `"}`

	commandMessageUUID := client.SendCommand("Gitlab.CREATE_PROJECT", models.NewOptions("", payload))
	projectUUID := commandMessageUUID.GetUUID()
	event := client.WaitReplyByEvent("CREATE_PROJECT", "SUCCES", projectUUID, id)
	fmt.Println(event)

}

func demo() {
	var configuration Configuration
	mydir, _ := os.Getwd()
	file, _ := os.Open(mydir + "/demoWorkflow.json")
	decoder := json.NewDecoder(file)
	decoder.Decode(&configuration)
	client := goclient.NewClientGandalf(configuration.Identity, configuration.Timeout, configuration.Connections)

	//WAIT APP
	id := client.CreateIteratorEvent()
	event := client.WaitEvent("Application", "NEW_APP", id)
	fmt.Println(event)
	if event.GetEvent() == "NEW_APP" {
		//CREATE FORM
		//CreateProject":{"required":["Name","Team","TemplateName"]
		payload := `{"Fields":[{"Name":"ID","HtmlType":"TextField","Value":"Id"}]}`

		//commandMessageUUID := client.SendCommand("Utils.CREATE_FORM", models.NewOptions("", payload))
		commandMessageUUID := client.SendCommand("Utils.CREATE_FORM", models.NewOptions("", payload))
		formUUID := commandMessageUUID.GetUUID()
		event = client.WaitReplyByEvent("CREATE_FORM", "SUCCES", formUUID, id)
		fmt.Println(event)

		if event.GetEvent() == "SUCCES" {
			//SEND FORM
			receivers, _ := json.Marshal(configuration.Receivers)

			payload = `{"Sender":"` + configuration.Sender + `","Body":"` + configuration.Body + " " + event.GetPayload() + `","Receivers":` + string(receivers) + `,"Username":"` + configuration.Username + `",
			"Password":"` + configuration.Password + `","Host":"` + configuration.Host + `"}`

			commandMessageUUID = client.SendCommand("Utils.SEND_AUTH_MAIL", models.NewOptions("", payload))
			event = client.WaitReplyByEvent("SEND_AUTH_MAIL", "SUCCES", commandMessageUUID.GetUUID(), id)
			fmt.Println(event)

			if event.GetEvent() == "SUCCES" {

				event = client.WaitReplyByEvent("VALIDATION_FORM", "SUCCES", formUUID, id)
				fmt.Println(event)
				if event.GetEvent() == "SUCCES" {
					//START AZURE

					payload = `{"ResourceGroupName":"` + configuration.ResourceGroupName + `","ResourceGroupLocation":"` + configuration.ResourceGroupLocation + `",
				"DeploymentName":"` + configuration.DeploymentName + `","TemplateFile":"` + configuration.TemplateFile + `",
				"ParametersFile":"` + configuration.ParametersFile + `"}`

					commandMessageUUID = client.SendCommand("Azure.CREATE_VM_BY_JSON", models.NewOptions("", payload))
					event = client.WaitReplyByEvent("CREATE_VM_BY_JSON", "SUCCES", commandMessageUUID.GetUUID(), id)
					fmt.Println(event)
				}

			}
		}
	}

}

type Configuration struct {
	Identity              string
	Connections           []string
	Timeout               string
	Sender                string
	Body                  string
	Receivers             []string
	Username              string
	Password              string
	Host                  string
	ResourceGroupName     string
	ResourceGroupLocation string
	DeploymentName        string
	TemplateFile          string
	ParametersFile        string
}
*/
