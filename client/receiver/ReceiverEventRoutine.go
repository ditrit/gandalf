package receiver

import(
	"gandalfgo/message"
	"github.com/pebbe/zmq4"
)

type ReceiverEventRoutine struct {
	context							zmq4.Context
	workerEventReceive 				zmq4.Socket
	receiverEventConnection 		string
	identity 						string
	eventsRoutine 					map[string][]EventFunction					
}

func (r ReceiverEventRoutine) New(identity, receiverEventConnection string, eventsRoutine map[string][]EventFunction) {
	r.identity = identity
	r.receiverEventConnection = receiverEventConnection
	r.eventsRoutine = eventsRoutine

	r.context, _ := zmq4.NewContext()
	r.workerEventReceive = r.context.NewSub(receiverEventConnection)
	r.workerEventReceive.Identity(r.identity)
	fmt.Printf("workerEventReceive connect : " + receiverEventConnection)

	r.loadEventRoutines()

	result, err := r.validationFunctions()
	if err != nil {
		panic(err)
	}
	go r.run()
}

func (r ReceiverEventRoutine) run() {

	pi := zmq4.PollItems{
		zmq4.PollItem{Socket: workerEventReceive, Events: zmq4.POLLIN}

	var event = [][]byte{}

	for {

		pi.Poll(-1)

		switch {

		case pi[0].REvents&zmq4.POLLIN != 0:

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

func (r ReceiverEventRoutine) validationFunctions() (result bool, err error) {
	r.sendValidationFunctions()
	for {
		event, err := workerEventReceive.RecvMessage()
		if err != nil {
			panic(err)
		}
	}
	result = event
	return
}

func (r ReceiverEventRoutine) sendValidationFunctions()  
	//EVENT
	functionkeys := make([]string, 0, len(eventsRoutine))
    for key := range eventsRoutine {
        functionkeys = append(functionkeys, key)
	}
	commandFunction := CommandFunction.New(functionkeys)
	go commandFunction.sendWith(r.workerEventReceive)
}

func (r ReceiverEventRoutine) processEventReceive(event [][]byte) () {
	eventMessage := message.EventMessage.decodeEvent(event[1])
	eventRoutine, err := r.getEventRoutine(eventMessage.event)
	if err != nil {

	}
	go eventRoutine.execute(eventMessage)
}

func (r ReceiverEventRoutine) getEventRoutine(event string) (eventRoutine EventRoutine, err error) {
	if eventRoutine, ok := r.eventsRoutine[command]; ok {
		return eventRoutine
	}
}
