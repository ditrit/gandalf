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

func (wg WorkerGandalf) New(path string) {
	path := path
	workerConfiguration := WorkerConfiguration.loadConfiguration(path)

	wg.results := make(chan message.CommandResponse)
	wg. loadFunctions()

	wg.clientGandalf = ClientGandalf.New(workerConfiguration.identity, workerConfiguration.senderCommandConnection, workerConfiguration.senderEventConnection, 
		workerConfiguration.receiverCommandConnection, workerConfiguration.receiverEventConnection,
		 commandsRoutine map[string][]CommandRoutine, eventsRoutine map[string][]EventRoutine, results chan ResponseMessage)
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
