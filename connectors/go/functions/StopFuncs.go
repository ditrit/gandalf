package functions

import (
	"fmt"
	"time"

	gomodels "github.com/ditrit/gandalf/connectors/go/models"
	goclient "github.com/ditrit/gandalf/libraries/goclient"
)

//SendCommands
func Stop(clientGandalf *goclient.ClientGandalf, major, minor int64, workerState *gomodels.WorkerState) {

	for workerState.GetState() == 0 {
		fmt.Println("STATE")
		fmt.Println(workerState.GetState())
		validate := clientGandalf.SendStop(major, minor)
		fmt.Println("VALIDATE STOP")
		fmt.Println(validate.GetValid())
		if !validate.GetValid() {
			workerState.SetStoppingWorkerState()
		}

		time.Sleep(1 * time.Second)
	}

}
