package receiver

type ReceiverEventRoutine struct {
	workerEventReceive zmq.Sock
	receiverEventConnection string
	identity string
	eventsRoutine map[string][]EventFunction					
}

func (r ReceiverEventRoutine) New(identity, receiverEventConnection string) err error {
	r.identity = identity
	r.receiverEventConnection = receiverEventConnection
	
	r.workerEventReceive = zmq.NewSub(receiverEventConnection)
	r.workerEventReceive.Identity(r.identity)
	fmt.Printf("workerEventReceive connect : " + receiverEventConnection)

	r.loadEventRoutines()

	result, err := r.validationFunctions()
	if err != nil {
		panic(err)
	}

	go r.run()

}

func (r ReceiverEventRoutine) loadEventRoutines() err error {
	//TODO
}

func (r ReceiverEventRoutine) run() err error {

	pi := zmq.PollItems{
		zmq.PollItem{Socket: workerEventReceive, Events: zmq.POLLIN}

	var event = [][]byte{}

	for {

		_, _ = zmq.Poll(pi, -1)

		switch {

		case pi[0].REvents&zmq.POLLIN != 0:

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
	commandFunction.sendWith(r.workerEventReceive)
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
