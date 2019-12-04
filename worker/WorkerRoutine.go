package worker

import (
	"fmt"

	zmq "github.com/zeromq/goczmq"
)

type WorkerRoutine struct {
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

func (r WorkerRoutine) new(identity, workerCommandReceiveC2WConnection, workerEventReceiveC2WConnection string, topics *string) {
	r.identity = identity

	r.workerCommandReceiveC2WConnection = workerCommandReceiveC2WConnection
	r.workerCommandReceiveC2W = zmq.NewDealer(workerCommandReceiveC2WConnection)
	r.workerCommandReceiveC2W.Identity(r.identity)
	fmt.Printf("workerCommandReceiveC2W connect : " + workerCommandReceiveC2WConnection)

	r.workerEventReceiveC2WConnection = workerEventReceiveC2WConnection
	r.workerEventReceiveC2W = zmq.NewSub(workerEventReceiveC2WConnection)
	r.workerEventReceiveC2W.Identity(r.identity)
	fmt.Printf("workerEventReceiveC2W connect : " + workerEventReceiveC2WConnection)

	r.topics = topics
	r.commandStateManager = new(CommandStateManager)

	ga.mapUUIDCommandStates = make(map[string][]string)
	ga.mapUUIDState = make(map[string]*ReferenceState)

	ga.mapCommandFunction = make(map[string]*CommandFunction)
	ga.mapEventFunction = make(map[string]*EventFunction)
}

func (r WorkerRoutine) close() {
	r.WorkerCommandFrontEndReceive.close()
	r.WorkerEventFrontEndReceive.close()
	r.Context.close()
}

func (r WorkerRoutine) sendReadyCommand() {

}

func (r WorkerRoutine) sendCommandState(request goczmq.Message, state, payload string) {
	//response := [][]byte{}
}

func (r WorkerRoutine) run() {
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
			err = r.processRoutingWorkerCommand(command)
			if err != nil {
				panic(err)
			}

		case pi[1].REvents&zmq.POLLIN != 0:

			event, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processRoutingSubscriberCommand(event)
			if err != nil {
				panic(err)
			}
		}
	}
	fmt.Println("done")
}

func (r WorkerRoutine) processRoutingWorkerCommand(command [][]byte) {
	command = r.updateHeaderWorkerCommand(command)
	r.executeWorkerCommandFunction(command)
	//TODO message pack
}

func (r WorkerRoutine) updateHeaderWorkerCommand(command [][]byte) {

}


func (r WorkerRoutine) processRoutingSubscriberCommand(event [][]byte) {
	event = r.updateHeaderWorkerEvent(event)
	r.executeWorkerEventFunction(event)
	//TODO message pack
}

func (r WorkerRoutine) updateHeaderWorkerEvent(command [][]byte) {

}

func (r WorkerRoutine) reconnectToConnector() {
	if r.workerCommandFrontEndReceive != nil {
		r.workerCommandFrontEndReceive.Destroy()
	}
	if r.workerEventFrontEndReceive != nil {
		r.workerEventFrontEndReceive.Destroy()
	}
	r.init(r.identity, r.workerCommandFrontEndReceive, r.workerEventFrontEndReceive)
	r.WorkerZMQ.sendReadyCommand()
}

func (r WorkerRoutine) GetMapCommandByName(name string) *CommandFunction {
	return ga.mapCommandFunction[name]
}

func (r WorkerRoutine) GetMapEventByName(name string) *EventFunction {
	return ga.mapEventFunction[name]
}

func (r WorkerRoutine) executeWorkerCommandFunction(commandExecute [][]byte) {
	r.GetMapCommandByName("CommandPrint").executeCommand();
}

func (r WorkerRoutine) executeWorkerEventFunction(eventExecute [][]byte) {
	r.GetMapEventByName("EventPrint").executeEvent());
}

func (r WorkerRoutine) GetMapUUIDCommandStates() map[string]List {
	return ga.mapUUIDCommandStates
}

func (r WorkerRoutine) SetMapUUIDCommandStates(mapUUIDCommandStates map[string]List) {
	return ga.mapUUIDCommandStates = mapUUIDCommandStates
}

func (r WorkerRoutine) GetMapUUIDState() map[string]ReferenceState {
	return ga.mapUUIDState
}

func (r WorkerRoutine) SetMapUUIDState(mapUUIDState map[string]ReferenceState) {
	return ga.mapUUIDState = mapUUIDState
}

func (r WorkerRoutine) New() {
	ga.mapUUIDCommandStates = make(map[string]List)
	ga.mapUUIDState = make(map[string]ReferenceState)
}

func (r WorkerRoutine) GetMapUUIDCommandStatesByUUID(uuid string) []string {
	return ga.mapUUIDCommandStates[uuid]
}

func (r WorkerRoutine) GetMapUUIDStateByUUID(uuid string) *ReferenceState {
	return ga.mapUUIDState[uuid]
}

