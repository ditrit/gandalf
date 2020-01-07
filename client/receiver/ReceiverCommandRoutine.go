package receiver

import (
	"errors"
	"fmt"
	"gandalf-go/message"
	"gandalf-go/worker/routine"

	"github.com/pebbe/zmq4"
)

type ReceiverCommandRoutine struct {
	Context                      *zmq4.Context
	Replys                       chan message.CommandMessageReply
	WorkerCommandReceive         *zmq4.Socket
	ReceiverCommandConnection    string
	WorkerEventReceive           *zmq4.Socket
	WorkerEventReceiveConnection string
	Identity                     string
	CommandsRoutine              map[string][]routine.CommandRoutine
}

func NewReceiverCommandRoutine(identity, receiverCommandConnection string, commandsRoutine map[string][]routine.CommandRoutine, results chan message.CommandMessageReply) (receiverCommandRoutine *ReceiverCommandRoutine) {
	receiverCommandRoutine = new(ReceiverCommandRoutine)

	receiverCommandRoutine.Identity = identity
	receiverCommandRoutine.ReceiverCommandConnection = receiverCommandConnection
	receiverCommandRoutine.CommandsRoutine = commandsRoutine
	receiverCommandRoutine.Replys = make(chan message.CommandMessageReply)

	receiverCommandRoutine.Context, _ = zmq4.NewContext()
	receiverCommandRoutine.WorkerCommandReceive, _ = receiverCommandRoutine.Context.NewSocket(zmq4.DEALER)
	receiverCommandRoutine.WorkerCommandReceive.SetIdentity(receiverCommandRoutine.Identity)
	receiverCommandRoutine.WorkerCommandReceive.Connect(receiverCommandRoutine.ReceiverCommandConnection)
	fmt.Println("workerCommandReceive connect : " + receiverCommandConnection)

	receiverCommandRoutine.loadCommandRoutines()

	result := true
	/* 	result, err := receiverCommandRoutine.validationFunctions()
	   	if err != nil {
	   		panic(err)
	   	} */
	if result {
		go receiverCommandRoutine.run()
	}

	return
}

func (r ReceiverCommandRoutine) run() {

	go r.sendResults()

	poller := zmq4.NewPoller()
	poller.Add(r.WorkerCommandReceive, zmq4.POLLIN)

	command := [][]byte{}
	err := errors.New("")

	for {
		fmt.Println("Running ReceiverCommandRoutine")
		r.sendReadyCommand()
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {

			switch currentSocket := socket.Socket; currentSocket {
			case r.WorkerCommandReceive:

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processCommandReceive(command)
			}
		}
	}
	fmt.Println("done")
}

func (r ReceiverCommandRoutine) loadCommandRoutines() {
	//TODO
}

func (r ReceiverCommandRoutine) validationFunctions() (result bool, err error) {
	r.sendValidationFunctions()
	command := [][]byte{}

	for {
		command, err = r.WorkerCommandReceive.RecvMessageBytes(0)
		if err != nil {
			panic(err)
		}
	}
	reply, err := message.DecodeCommandFunctionReply(command[1])
	result = reply.Validation
	return
}

func (r ReceiverCommandRoutine) sendValidationFunctions() {
	//COMMAND
	functionkeys := make([]string, 0, len(r.CommandsRoutine))
	for key := range r.CommandsRoutine {
		functionkeys = append(functionkeys, key)
	}
	commandFunction := message.NewCommandFunction(functionkeys)
	go commandFunction.SendWith(r.WorkerCommandReceive)
}

func (r ReceiverCommandRoutine) sendReadyCommand() {
	commandReady := message.NewCommandMessageReady()
	go commandReady.SendWith(r.WorkerCommandReceive)
}

func (r ReceiverCommandRoutine) processCommandReceive(command [][]byte) {
	commandMessage, _ := message.DecodeCommandMessage(command[1])
	commandRoutine := r.getCommandRoutine(commandMessage.Command)

	go commandRoutine.ExecuteCommand(commandMessage, r.Replys)
}

func (r ReceiverCommandRoutine) getCommandRoutine(command string) (commandRoutine routine.CommandRoutine) {
	if commandRoutine, ok := r.CommandsRoutine[command]; ok {
		return commandRoutine[0]
	}
	return
}

func (r ReceiverCommandRoutine) sendResults() {
	err := errors.New("")
	for {
		reply := <-r.Replys
		if err != nil {

		}
		go reply.SendCommandReplyWith(r.WorkerCommandReceive)
	}
}
