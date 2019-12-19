package worker	

import (
    "fmt"
	"os"
	"gandalf-go/client"  
	"gandalf-go/worker/routine"  
)


type WorkerGandalf struct {
	results 			chan ResponseMessage
	commandsRoutine 	map[string][]routine.CommandRoutine
	eventsRoutine 		map[string][]routine.EventRoutine
	workerConfiguration WorkerConfiguration
	clientGandalf 		ClientGandalf
}

func NewWorkerGandalf(path string) (workerGandalf *WorkerGandalf) {
	workerGandalf = new(WorkerGandalf)

	workerGandalf.workerConfiguration := WorkerConfiguration.loadConfiguration(path)

	workerGandalf.commandsRoutine := make(map[string][]CommandRoutine)
	workerGandalf.eventsRoutine := make(map[string][]EventRoutine)
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
