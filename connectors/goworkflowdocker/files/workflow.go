package main

import (
	"fmt"

	"github.com/ditrit/gandalf/libraries/goclient"
)

func Workflow(clientGandalf *goclient.ClientGandalf) {
	fmt.Println("Start 4")

	fmt.Println("SEND COMMMAND CREATE_REPOSITORY")
	payload := `{"Token":"","Name":"TestConnector","Description":"TestConnector","Private":true}`
	commandMessageUUID := clientGandalf.SendCommand("vcs.CREATE_REPOSITORY", map[string]string{"payload": payload})
	fmt.Println(commandMessageUUID)

}
