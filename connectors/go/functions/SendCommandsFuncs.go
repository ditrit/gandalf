package functions

import (
	goclient "github.com/ditrit/gandalf/libraries/goclient"
)

//SendCommands
func SendCommands(clientGandalf *goclient.ClientGandalf, major, minor int64, commandes []string) {
	clientGandalf.SendCommandList(major, minor, commandes)
}
