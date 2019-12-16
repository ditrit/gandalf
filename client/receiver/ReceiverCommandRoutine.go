package receiver

import(
	"gandalfgo/message"
	"github.com/pebbe/zmq4"
)

type ReceiverCommandRoutine struct {
	context							zmq4.Context
	results 						chan message.ResponseMessage
	workerCommandReceive 			zmq4.Socket
	receiverCommandConnection 		string
	workerEventReceive 				zmq4.Socket
	workerEventReceiveConnection	string
	identity 						string
	commandsRoutine 				map[string][]message.CommandRoutine					
}

func (r ReceiverCommandRoutine) New(identity, receiverCommandConnection string, commandsRoutine map[string][]CommandFunction, results chan message.ResponseMessage) {
	r.identity = identity
	r.receiverCommandConnection = receiverCommandConnection
	r.commandsRoutine = commandsRoutine
	r.results = results

	r.context, _ = zmq4.NewContext()
	r.workerCommandReceive = r.context.NewDealer(receiverCommandConnection)
	r.workerCommandReceive.SetIdentity(r.identity)
	fmt.Printf("workerCommandReceive connect : " + receiverCommandConnection)

	r.loadCommandRoutines()

	result, err := r.validationFunctions()
	if err != nil {
		panic(err)
	}
	go r.run()
}

func (r ReceiverCommandRoutine) run() {

	go r.sendResults()

	pi := zmq4.PollItems{
		zmq4.PollItem{Socket: workerCommandReceive, Events: zmq4.POLLIN},

	var command = [][]byte{}
	var event = [][]byte{}

	for {
		r.sendReadyCommand()

		pi.Poll(-1)

		switch {
		case pi[0].REvents&zmq4.POLLIN != 0:

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

func (r ReceiverCommandRoutine) sendResults() {
	for {
		reply, err <- r.replys
		if err != nil {
			
		} 
		go reply.sendWith(workerCommandReceive)
	}
}