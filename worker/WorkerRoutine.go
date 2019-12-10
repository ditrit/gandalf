package worker

import (
	"fmt"
	"message"
	"os/exec"
	zmq "github.com/zeromq/goczmq"
)

const pathRoutine string = "/enabled_worker/"

type WorkerRoutine struct {
	replys chan ResponseMessage
	workerCommandReceive zmq.Sock
	workerCommandReceiveConnection string
	workerEventReceive zmq.Sock
	workerEventReceiveConnection string
	identity string
	commandsRoutine map[string][]CommandFunction					
	eventsRoutine map[string][]EventFunction					
}

func (r WorkerRoutine) New(identity, workerCommandReceiveConnection, workerEventReceiveConnection string) err error {
	r.identity = identity
	r.workerCommandReceiveConnection = workerCommandReceiveConnection
	r.workerEventReceiveConnection = workerEventReceiveConnection
	results := make(chan message.CommandResponse)

	r.workerCommandReceive = zmq.NewDealer(workerCommandReceiveConnection)
	r.workerCommandReceive.Identity(r.identity)
	fmt.Printf("workerCommandReceive connect : " + workerCommandReceiveConnection)
	
	r.workerEventReceive = zmq.NewSub(workerEventReceiveConnection)
	r.workerEventReceive.Identity(r.identity)
	fmt.Printf("workerEventReceive connect : " + workerEventReceiveConnection)

	r.loadCommandRoutines()
	r.loadEventRoutines()

}

func (r WorkerRoutine) loadCommandRoutines() err error {
	//TODO
	//CHANNEL ADD
}

func (r WorkerRoutine) loadEventRoutines() err error {
	//TODO
}

func (r WorkerRoutine) run() err error {



	go r.sendResults()

	pi := zmq.PollItems{
		zmq.PollItem{Socket: workerCommandReceive, Events: zmq.POLLIN},
		zmq.PollItem{Socket: workerEventReceive, Events: zmq.POLLIN}

	var command = [][]byte{}
	var event = [][]byte{}

	for {
		r.sendReadyCommand()

		_, _ = zmq.Poll(pi, -1)

		switch {
		case pi[0].REvents&zmq.POLLIN != 0:

			command, err := pi[0].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processCommandReceive(command)
			if err != nil {
				panic(err)
			}

		case pi[1].REvents&zmq.POLLIN != 0:

			event, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processEventReceive(event)
			if err != nil {
				panic(err)
			}
		}
	}
	fmt.Println("done")
}

func (r WorkerRoutine) validationCommandsEvents() () {
	r.sendCommandsEvents()
	for {
		command, err := workerCommandReceive.RecvMessage()
	}
}

func (r WorkerRoutine) sendCommandsEvents() () {
	commandCommandsEvents := CommandCommandsEvents.New()
	commandCommandsEvents.sendWith(r.workerCommandReceive)
}

func (r WorkerRoutine) sendReadyCommand() () {
	commandReady := CommandReady.New()
	commandReady.sendWith(r.workerCommandReceive)
}

func (r WorkerRoutine) processCommandReceive(command [][]byte) () {
	commandMessage := message.CommandMessage.decodeCommand(command[1])
	commandRoutine, err := r.getCommandRoutine(commandMessage.command)
	if err != nil {
		
	}
	go commandRoutine.execute(commandMessage, results)
}

func (r WorkerRoutine) getCommandRoutine(command string) (commandRoutine CommandRoutine, err error) {
	if commandRoutine, ok := r.commandsRoutine[command]; ok {
		return commandRoutine
	}
}

func (r WorkerRoutine) processEventReceive(event [][]byte) () {
	eventMessage := message.EventMessage.decodeEvent(event[1])
	eventRoutine, err := r.getEventRoutine(eventMessage.event)
	if err != nil {

	}
	go eventRoutine.execute(eventMessage)
}

func (r WorkerRoutine) getEventRoutine(event string) (eventRoutine EventRoutine, err error) {
	if eventRoutine, ok := r.eventsRoutine[command]; ok {
		return eventRoutine
	}
}

func (r WorkerRoutine) sendResults() err error {
	for {
		reply, err <- r.replys
		if err != nil {
			
		} 
		reply.sendWith(workerCommandReceive)
	}
}