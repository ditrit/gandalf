package functions

import (
	goclient "github.com/ditrit/gandalf/libraries/goclient"
)

//SendCommands
func SendCommands(clientGandalf *goclient.ClientGandalf, major int64, commandes []string) {
	clientGandalf.SendCommandList(major, commandes)
}
