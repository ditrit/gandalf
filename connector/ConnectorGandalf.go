package connector

type ConnectorGandalf struct {
	connectorCommandRoutine     ConnectorCommandRoutine
	connectorEventRoutine       ConnectorEventRoutine
	connectorCommandsMap        map[string][]string
	connectorCommandSendFileMap map[string]string
}

func (cg ConnectorGandalf) New() err error {
	cg.connectorCommandsMap = make(map[string][]string)
	cg.connectorCommandSendFileMap = make(map[string]string)
	cg.connectorCommandRoutine = ConnectorCommandRoutine.new()
	cg.connectorEventRoutine = ConnectorEventRoutine.new()

	//RUN
	go cg.connectorCommandRoutine.run()
	go cg.connectorEventRoutine.run()
}

func (cg ConnectorGandalf) getWorkerCommands(worker string) (workerCommand []string, err error) {
	return cg.connectorCommandsMap[worker]
}

func (cg ConnectorGandalf) addWorkerCommands(worker, command string) err error {
	var sizeList = len(cg.connectorCommandsMap[worker])
	cg.connectorCommandsMap[worker][sizeList] = command

}

func (cg ConnectorGandalf) getWorkerCommandSendFile(worker string) (workerCommandFile string, err error) {
	return cg.connectorCommandSendFileMap[worker]
}

func (cg ConnectorGandalf) addWorkerCommandSendFile(worker, commandSendFile string) err error {
	cg.connectorCommandSendFileMap[worker] = commandSendFile
}
