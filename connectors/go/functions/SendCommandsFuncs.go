package functions

import (
	"fmt"

	goclient "github.com/ditrit/gandalf/libraries/goclient"
)

//SendCommands
func SendCommands(clientGandalf *goclient.ClientGandalf, major, minor int64, commandes []string) bool {
	fmt.Println("SEND COMMAND LIST WORKER")
	validate := clientGandalf.SendCommandList(major, minor, commandes)

	return validate.GetValid()
}
