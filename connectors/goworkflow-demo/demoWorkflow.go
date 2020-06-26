package main

import (
	"encoding/json"
	"fmt"
	goclient "libraries/gandalf-libraries-goclient"
	"libraries/gandalf-libraries-goclient/models"
	"os"
)

func main() {
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

		payload := `{"Fields":[{"Name":"ID","HtmlType":"TextField","Value":"Id"}]}`
		fmt.Println("TATA")

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
