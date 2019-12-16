package listener

type ListenerGandalf struct {
	listenerEventRoutine   ListenerEventRoutine
	listenerCommandRoutine ListenerCommandRoutine
}

func (lg ListenerGandalf) main() {
	//identity, workerCommandReceiveC2WConnection, workerEventReceiveC2WConnection string, topics *string
	//LOAD CONF
	cg.listenerEventRoutine = ListenerEventRoutine.new()
	cg.listenerCommandRoutine = ListenerCommandRoutine.new()

	go cg.listenerEventRoutine.run()
	go cg.listenerCommandRoutine.run()
}
