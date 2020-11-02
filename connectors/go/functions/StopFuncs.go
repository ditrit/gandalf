package functions

import (
	gomodels "github.com/ditrit/gandalf/connectors/go/models"
	goclient "github.com/ditrit/gandalf/libraries/goclient"
)

//SendCommands
func Stop(clientGandalf *goclient.ClientGandalf, major, minor int64, workerState *gomodels.WorkerState) {
	validate := clientGandalf.SendStop(major, minor)

	if validate.GetValid() {
		workerState.SetStoppingWorkerState()
	}
}
