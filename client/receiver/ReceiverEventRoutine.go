package receiver

import(
	"fmt"
	"errors"
	"gandalfgo/message"
	"gandalfgo/worker/routine"
	"github.com/pebbe/zmq4"
)

type ReceiverEventRoutine struct {
	Context							*zmq4.Context
	WorkerEventReceive 				*zmq4.Socket
	ReceiverEventConnection 		string
	Identity 						string
	EventsRoutine 					map[string][]routine.EventRoutine					
}

func NewReceiverEventRoutine(identity, receiverEventConnection string, eventsRoutine map[string][]routine.EventRoutine) (receiverEventRoutine *ReceiverEventRoutine) {
	receiverEventRoutine = new(ReceiverEventRoutine)
	
	receiverEventRoutine.Identity = identity
	receiverEventRoutine.ReceiverEventConnection = receiverEventConnection
	receiverEventRoutine.EventsRoutine = eventsRoutine

	receiverEventRoutine.Context, _ = zmq4.NewContext()
	receiverEventRoutine.WorkerEventReceive, _ = receiverEventRoutine.Context.NewSocket(zmq4.SUB)
	receiverEventRoutine.WorkerEventReceive.SetIdentity(receiverEventRoutine.Identity)
	receiverEventRoutine.WorkerEventReceive.Connect(receiverEventRoutine.Identity)
	fmt.Printf("workerEventReceive connect : " + receiverEventConnection)

	receiverEventRoutine.loadEventRoutines()

	result, err := receiverEventRoutine.validationFunctions()
	if err != nil {
		panic(err)
	}
	if result {
		go receiverEventRoutine.run()
	}

	return
}

func (r ReceiverEventRoutine) run() {

	poller := zmq4.NewPoller()
	poller.Add(r.WorkerEventReceive, zmq4.POLLIN)

	event := [][]byte{}
	err := errors.New("")

	for {

		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {

			switch currentSocket := socket.Socket; currentSocket {
			case r.WorkerEventReceive:

				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				err = r.processEventReceive(event)
				if err != nil {
					panic(err)
				}
			}
		}
	}
	fmt.Println("done")
}


func (r ReceiverEventRoutine) loadEventRoutines() {
	//TODO
}


func (r ReceiverEventRoutine) validationFunctions() (result bool, err error) {
	r.sendValidationFunctions()
	event := [][]byte{}
	for {
		event, err = r.WorkerEventReceive.RecvMessageBytes(0)
		if err != nil {
			panic(err)
		}
	}
	reply, err := message.DecodeCommandFunctionReply(event[1])
	result = reply.Validation 
	return
}

func (r ReceiverEventRoutine) sendValidationFunctions()  {
	//EVENT
	functionkeys := make([]string, 0, len(r.EventsRoutine))
    for key := range r.EventsRoutine {
        functionkeys = append(functionkeys, key)
	}
	commandFunction := message.NewCommandFunction(functionkeys)
	go commandFunction.SendWith(r.WorkerEventReceive)
}

func (r ReceiverEventRoutine) processEventReceive(event [][]byte) (err error) {
	eventMessage, _ := message.DecodeEventMessage(event[1])
	eventRoutine := r.getEventRoutine(eventMessage.Event)
	if err != nil {

	}
	go eventRoutine.ExecuteEvent(eventMessage)

	return 
}

func (r ReceiverEventRoutine) getEventRoutine(event string) (eventRoutine routine.EventRoutine) {
	if eventRoutine, ok := r.EventsRoutine[event]; ok {
		return eventRoutine[0]
	}
	return
}
