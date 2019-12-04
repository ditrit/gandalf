package connector

type ConnectorGandalf struct {
	connectorCommandRoutine     ConnectorCommandRoutine
	connectorEventRoutine       ConnectorEventRoutine
	connectorCommandsMap        map[string][]string
	connectorCommandSendFileMap map[string]string
}

func (cg ConnectorGandalf) new() {
	cg.connectorCommandsMap = make(map[string][]string)
	cg.connectorCommandSendFileMap = make(map[string]string)
	cg.connectorCommandRoutine = ConnectorCommandRoutine.new()
	cg.connectorEventRoutine = ConnectorEventRoutine.new()

	//RUN
	go cg.connectorCommandRoutine.run()
	go cg.connectorEventRoutine.run()
}

func (cg ConnectorGandalf) getWorkerCommands(worker string) []string {
	return cg.connectorCommandsMap[worker]
}

func (cg ConnectorGandalf) addWorkerCommands(worker, command string) {
	var sizeList = len(cg.connectorCommandsMap[worker])
	cg.connectorCommandsMap[worker][sizeList] = command

}

func (cg ConnectorGandalf) getWorkerCommandSendFile(worker string) string {
	return cg.connectorCommandSendFileMap[worker]
}

func (cg ConnectorGandalf) addWorkerCommandSendFile(worker, commandSendFile string) {
	cg.connectorCommandSendFileMap[worker] = commandSendFile
}
