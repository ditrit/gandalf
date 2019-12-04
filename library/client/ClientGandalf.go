package client

type ClientGandalf struct {
	clientCommandRoutine ClientCommandRoutine
	clientEventRoutine   ClientEventRoutine
}

func (cg ClientGandalf) main() {
	//identity, workerCommandReceiveC2WConnection, workerEventReceiveC2WConnection string, topics *string
	//LOAD CONF
	cg.clientCommandRoutine = ClientCommandRoutine.new()
	cg.clientEventRoutine = ClientEventRoutine.new()
}
