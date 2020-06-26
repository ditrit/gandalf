package functions

import (
	goclient "github.com/ditrit/gandalf/libraries/goclient"
)

//SendCommands
func SendCommands(clientGandalf *goclient.ClientGandalf, version int64, commandes []string) {
	clientGandalf.SendCommandList(version, commandes)
}
