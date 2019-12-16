package receiver

import(
	"gandalfgo/message"
	"github.com/alecthomas/gozmq"
)

type ReceiverCommandRoutine struct {
	context							*gozmq.Context
	results 						chan ResponseMessage
	workerCommandReceive 			gozmq.Socket
	receiverCommandConnection 		string
	workerEventReceive 				gozmq.Socket
	workerEventReceiveConnection	string
	identity 						string
	commandsRoutine 				map[string][]CommandFunction					
}

func (r ReceiverCommandRoutine) New(identity, receiverCommandConnection string, commandsRoutine map[string][]CommandFunction, results chan) err error {
	r.identity = identity
	r.receiverCommandConnection = receiverCommandConnection
	r.commandsRoutine = commandsRoutine
	r.results = results

	r.context, _ := gozmq.NewContext()
	r.workerCommandReceive = r.context.NewDealer(receiverCommandConnection)
	r.workerCommandReceive.Identity(r.identity)
	fmt.Printf("workerCommandReceive connect : " + receiverCommandConnection)

	r.loadCommandRoutines()

	result, err := r.validationFunctions()
	if err != nil {
		panic(err)
	}
	go r.run()
}

func (r ReceiverCommandRoutine) run() err error {

	go r.sendResults()

	pi := zmq.PollItems{
		zmq.PollItem{Socket: workerCommandReceive, Events: zmq.POLLIN},

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
		}
	}
	fmt.Println("done")
}

func (r ReceiverCommandRoutine) validationFunctions() (result bool, err error) {
	r.sendValidationFunctions()
	for {
		command, err := workerCommandReceive.RecvMessage()
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
	commandFunction := CommandFunction.New(keys)
	go commandFunction.sendWith(r.workerCommandReceive)
}

func (r ReceiverCommandRoutine) sendReadyCommand() () {
	commandReady := CommandReady.New()
	go commandReady.sendWith(r.workerCommandReceive)
}

func (r ReceiverCommandRoutine) processCommandReceive(command [][]byte) () {
	commandMessage := message.CommandMessage.decodeCommand(command[1])
	commandRoutine, err := r.getCommandRoutine(commandMessage.command)
	if err != nil {
		
	}
	go commandRoutine.execute(commandMessage, results)
}

func (r ReceiverCommandRoutine) getCommandRoutine(command string) (commandRoutine CommandRoutine, err error) {
	if commandRoutine, ok := r.commandsRoutine[command]; ok {
		return commandRoutine
	}
}

func (r ReceiverCommandRoutine) sendResults() err error {
	for {
		reply, err <- r.replys
		if err != nil {
			
		} 
		go reply.sendWith(workerCommandReceive)
	}
}