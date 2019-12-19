package listener

type ListenerGandalf struct {
	listenerEventRoutine   *ListenerEventRoutine
	listenerCommandRoutine *ListenerCommandRoutine
}

func NewListenerGandalf() (listenerGandalf ListenerGandalf) {
	listenerGandalf = new(ListenerGandalf)
	//identity, workerCommandReceiveC2WConnection, workerEventReceiveC2WConnection string, topics *string
	//LOAD CONF
	listenerGandalf.listenerEventRoutine = ListenerEventRoutine.NewListenerEventRoutine()
	listenerGandalf.listenerCommandRoutine = ListenerCommandRoutine.NewListenerCommandRoutine()

	//go listenerGandalf.listenerEventRoutine.run()
	//go listenerGandalf.listenerCommandRoutine.run()
}

func (lg ListenerGandalf) run() {
	go lg.listenerEventRoutine.run()
	go lg.listenerCommandRoutine.run()
}
