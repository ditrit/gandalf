package listener

type ListenerGandalf struct {
	listenerEventRoutine   *ListenerEventRoutine
	listenerCommandRoutine *ListenerCommandRoutine
}

func NewListenerGandalf(identity, listenerCommandReceiveConnection, listenerEventReceiveConnection string) (listenerGandalf *ListenerGandalf) {
	listenerGandalf = new(ListenerGandalf)
	//identity, workerCommandReceiveC2WConnection, workerEventReceiveC2WConnection string, topics *string
	//LOAD CONF
	listenerGandalf.listenerEventRoutine = NewListenerEventRoutine(identity, listenerEventReceiveConnection)
	listenerGandalf.listenerCommandRoutine = NewListenerCommandRoutine(identity, listenerCommandReceiveConnection)

	//go listenerGandalf.listenerEventRoutine.run()
	//go listenerGandalf.listenerCommandRoutine.run()
	return
}

func (lg ListenerGandalf) run() {
	go lg.listenerEventRoutine.run()
	go lg.listenerCommandRoutine.run()
}
