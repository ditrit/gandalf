package worker

import (
	"fmt"

	zmq "github.com/zeromq/goczmq"
)

type Routine struct {
	workerCommandReceiveC2W           zmq.Sock
	workerCommandReceiveC2WConnection string
	workerEventReceiveC2W             zmq.Sock
	workerEventReceiveC2WConnection   string
	identity                          string
	topics                            *string
	commandStateManager               CommandStateManager
	mapUUIDCommandStates              map[string][]string
	mapUUIDState                      map[string]*ReferenceState
	mapCommand                        map[string]*CommandFunction
	mapEvent                          map[string]*EventFunction
}

func (r Routine) new(identity, workerCommandReceiveC2WConnection, workerEventReceiveC2WConnection string, topics *string) {
	r.Identity = identity

	r.workerCommandReceiveC2WConnection = workerCommandReceiveC2WConnection
	r.workerCommandReceiveC2W = zmq.NewDealer(workerCommandReceiveC2WConnection)
	r.workerCommandReceiveC2W.Identity(w.identity)
	fmt.Printf("workerCommandReceiveC2W connect : " + workerCommandReceiveC2WConnection)

	r.workerEventReceiveC2WConnection = workerEventReceiveC2WConnection
	r.workerEventReceiveC2W = zmq.NewSub(workerEventReceiveC2WConnection)
	r.workerEventReceiveC2W.Identity(w.identity)
	fmt.Printf("workerEventReceiveC2W connect : " + workerEventReceiveC2WConnection)

	r.topics = topics
	r.commandStateManager = new(CommandStateManager)

	ga.mapUUIDCommandStates = make(map[string][]string)
	ga.mapUUIDState = make(map[string]*ReferenceState)

	ga.mapCommandFunction = make(map[string]*CommandFunction)
	ga.mapEventFunction = make(map[string]*EventFunction)
}

func (r Routine) close() {
	r.WorkerCommandFrontEndReceive.close()
	r.WorkerEventFrontEndReceive.close()
	r.Context.close()
}

func (r Routine) sendReadyCommand() {

}

func (r Routine) sendCommandState(request goczmq.Message, state, payload string) {
	//response := [][]byte{}
}

func (r Routine) run() {
	pi := zmq.PollItems{
		zmq.PollItem{Socket: workerCommandReceiveC2W, Events: zmq.POLLIN},
		zmq.PollItem{Socket: workerEventReceiveC2W, Events: zmq.POLLIN}}

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
			//PROCESS COMMAND
			err = routerSock.SendMessage(msg)
			if err != nil {
				panic(err)
			}

		case pi[1].REvents&zmq.POLLIN != 0:

			event, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			//PROCESS EVENT
			err = routerSock.SendMessage(msg)
			if err != nil {
				panic(err)
			}
		}
	}

	fmt.Println("done")

}

func (r Routine) processRoutingWorkerCommand(command [][]byte) {

	r.executeWorkerCommandFunction(command)
	//TODO message pack
}

func (r Routine) processRoutingSubscriberCommand(event [][]byte) {
	r.executeWorkerEventFunction(event)
	//TODO message pack
}

func (r Routine) updateHeaderFrontEndWorker(command [][]byte) {

}

func (r Routine) reconnectToConnector() {
	if r.workerCommandFrontEndReceive != nil {
		r.workerCommandFrontEndReceive.Destroy()
	}
	if r.workerEventFrontEndReceive != nil {
		r.workerEventFrontEndReceive.Destroy()
	}
	r.init(r.identity, r.workerCommandFrontEndReceive, r.workerEventFrontEndReceive)
	r.WorkerZMQ.sendReadyCommand()
}

func (r Routine) GetMapCommandByName(name string) *CommandFunction {
	return ga.mapCommandFunction[name]
}

func (r Routine) GetMapEventByName(name string) *EventFunction {
	return ga.mapEventFunction[name]
}

func (r Routine) executeWorkerCommandFunction(commandExecute [][]byte) {
	r.GetMapCommandByName("CommandPrint").executeCommand();
}

func (r Routine) executeWorkerEventFunction(eventExecute [][]byte) {
	r.GetMapEventByName("EventPrint").executeEvent());
}

func (r Routine) GetMapUUIDCommandStates() map[string]List {
	return ga.mapUUIDCommandStates
}

func (r Routine) SetMapUUIDCommandStates(mapUUIDCommandStates map[string]List) {
	return ga.mapUUIDCommandStates = mapUUIDCommandStates
}

func (r Routine) GetMapUUIDState() map[string]ReferenceState {
	return ga.mapUUIDState
}

func (r Routine) SetMapUUIDState(mapUUIDState map[string]ReferenceState) {
	return ga.mapUUIDState = mapUUIDState
}

func (r Routine) New() {
	ga.mapUUIDCommandStates = make(map[string]List)
	ga.mapUUIDState = make(map[string]ReferenceState)
}

func (r Routine) GetMapUUIDCommandStatesByUUID(uuid string) []string {
	return ga.mapUUIDCommandStates[uuid]
}

func (r Routine) GetMapUUIDStateByUUID(uuid string) *ReferenceState {
	return ga.mapUUIDState[uuid]
}

