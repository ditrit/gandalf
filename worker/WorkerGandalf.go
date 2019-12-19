package worker	

import (
    "fmt"
	"os"
	"gandalfgo/client"  
)


type WorkerGandalf struct {
	results 			chan ResponseMessage
	commandsRoutine 	map[string][]CommandRoutine
	eventsRoutine 		map[string][]EventRoutine
	workerConfiguration WorkerConfiguration
	clientGandalf 		ClientGandalf
}

func NewWorkerGandalf(path string) (workerGandalf *WorkerGandalf) {
	workerGandalf = new(WorkerGandalf)

	workerGandalf.workerConfiguration := WorkerConfiguration.loadConfiguration(path)

	workerGandalf.results := make(chan message.CommandResponse)
	workerGandalf. loadFunctions()

	workerGandalf.clientGandalf = ClientGandalf.New(workerConfiguration.identity, workerConfiguration.senderCommandConnection, workerConfiguration.senderEventConnection, 
		workerConfiguration.receiverCommandConnection, workerConfiguration.receiverEventConnection,
		 commandsRoutine map[string][]CommandRoutine, eventsRoutine map[string][]EventRoutine, results chan ResponseMessage)
}

unc (wg WorkerGandalf) run() {
	go wg.clientGandalf.run()
}

func (wg WorkerGandalf) loadFunctions() {
	wg.loadCommands()
	wg.loadEvents()
}

func (wg WorkerGandalf) loadCommands() {
	//TODO
}


func (wg WorkerGandalf) loadEvents() {
	//TODO
}
