//Package worker :
//File WorkerGandalf.go
package worker

import (
	"fmt"
	"gandalf-go/client"
)

//WorkerGandalf :
type WorkerGandalf struct {
	WorkerConfiguration *WorkerConfiguration
	ClientGandalfGrpc   *client.ClientGandalfGrpc
}

//NewWorkerGandalf :
func NewWorkerGandalf(path string) (workerGandalf *WorkerGandalf) {
	workerGandalf = new(WorkerGandalf)

	workerGandalf.WorkerConfiguration, _ = LoadConfiguration(path)
	//workerGandalf.loadFunctions()

	workerGandalf.ClientGandalfGrpc = client.NewClientGandalfGrpc(workerGandalf.WorkerConfiguration.Identity,
		workerGandalf.WorkerConfiguration.SenderCommandConnection, workerGandalf.WorkerConfiguration.SenderEventConnection,
		workerGandalf.WorkerConfiguration.WaiterCommandConnection, workerGandalf.WorkerConfiguration.WaiterEventConnection)

	return
}

//Run :
func (wg WorkerGandalf) Run() {
	for {
		//GESTION CHANNEL
		fmt.Println("Im running into noting TODO me")
	}
}

/* //TODO REVOIR
func (wg WorkerGandalf) loadFunctions() {
	wg.loadCommands()
	wg.loadEvents()
}

func (wg WorkerGandalf) loadCommand(command string, commandRoutine routine.CommandRoutine) {
	//wg.CommandsRoutine["receive"] = tset.NewFunctionTest()
	//wg.CommandsRoutine["send"] = tset.NewFunctionTestSend()
}

func (wg WorkerGandalf) loadCommands() {
	//wg.CommandsRoutine["receive"] = tset.NewFunctionTest()
	//wg.CommandsRoutine["send"] = tset.NewFunctionTestSend()
}

func (wg WorkerGandalf) loadEvent(event string, eventRoutine routine.EventRoutine) {

}

func (wg WorkerGandalf) loadEvents() {
	//TODO
}
*/
