package receiver

import(
	"fmt"
	"gandalfgo/message"
	"gandalfgo/worker/routine"
	"github.com/pebbe/zmq4"
)

type ReceiverCommandRoutine struct {
	Context							*zmq4.Context
	Replys 							chan message.CommandMessageReply
	WorkerCommandReceive 			*zmq4.Socket
	ReceiverCommandConnection 		string
	WorkerEventReceive 				*zmq4.Socket
	WorkerEventReceiveConnection	string
	Identity 						string
	CommandsRoutine 				map[string][]routine.CommandRoutine					
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
	fmt.Printf("workerCommandReceive connect : " + receiverCommandConnection)

	receiverCommandRoutine.loadCommandRoutines()

	result, err := receiverCommandRoutine.validationFunctions()
	if err != nil {
		panic(err)
	}
	go receiverCommandRoutine.run()
}

func (r ReceiverCommandRoutine) run() {

	go r.sendResults()

	poller := zmq4.NewPoller()
	poller.Add(r.WorkerCommandReceive, zmq4.POLLIN)

	command := [][]byte{}

	for {
		r.sendReadyCommand()

		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {

			switch currentSocket := socket.Socket; currentSocket {
			case r.WorkerCommandReceive:

				command, err := currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				err = r.processCommandReceive(command)
				if err != nil {
					panic(err)
				}
			}
		}
	}
	fmt.Println("done")
}

func (r ReceiverCommandRoutine) loadCommandRoutines() (result bool, err error) {
	//TODO
	return
}


func (r ReceiverCommandRoutine) validationFunctions() (result bool, err error) {
	r.sendValidationFunctions()
	for {
		command, err := WorkerCommandReceive.RecvMessageBytes(0)
		if err != nil {
			panic(err)
		}
	}
	result = command 
	return
}

func (r ReceiverCommandRoutine) sendValidationFunctions()  {
	//COMMAND
	functionkeys := make([]string, 0, len(commandsRoutine))
    for key := range commandsRoutine {
        functionkeys = append(functionkeys, key)
	}
	commandFunction := message.NewCommandFunction(keys)
	go commandFunction.sendWith(r.WorkerCommandReceive)
}

func (r ReceiverCommandRoutine) sendReadyCommand() () {
	commandReady := message.NewCommandReady()
	go commandReady.sendWith(r.WorkerCommandReceive)
}

func (r ReceiverCommandRoutine) processCommandReceive(command [][]byte) () {
	commandMessage := message.DecodeCommandMessage(command[1])
	commandRoutine, err := r.getCommandRoutine(commandMessage.Command)
	if err != nil {
		
	}
	go commandRoutine.execute(commandMessage, results)
}

func (r ReceiverCommandRoutine) getCommandRoutine(command string) (commandRoutine routine.CommandRoutine, err error) {
	if commandRoutine, ok := r.CommandsRoutine[command]; ok {
		return commandRoutine
	}
}

func (r ReceiverCommandRoutine) sendResults() {
	for {
		reply <- r.Replys
		if err != nil {
			
		} 
		go reply.sendWith(r.WorkerCommandReceive)
	}
}