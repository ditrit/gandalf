package connector

type ConnectorGandalf struct {
	connectorConfiguration 		*ConnectorConfiguration
	connectorCommandRoutine     *ConnectorCommandRoutine
	connectorEventRoutine       *ConnectorEventRoutine
	connectorCommandsMap        map[string][]string
	connectorCommandSendFileMap map[string]string
}

func NewConnectorGandalf(path string) (connectorGandalf *ConnectorGandalf) {
	connectorConfiguration = LoadConfiguration(path)

	cg.connectorCommandsMap = make(map[string][]string)
	cg.connectorCommandSendFileMap = make(map[string]string)

	cg.connectorCommandRoutine = NewConnectorCommandRoutine(connectorConfiguration.Identity, connectorConfiguration.ConnectorCommandSendToWorkerConnection, connectorConfiguration.ConnectorCommandReceiveFromWorkerConnection, connectorConfiguration.ConnectorCommandReceiveFromAggregatorConnections, connectorConfiguration.ConnectorCommandSendToAggregatorConnections)
	cg.connectorEventRoutine = NewConnectorEventRoutine(connectorConfiguration.Identity, connectorConfiguration.ConnectorEventSendToWorkerConnection, connectorConfiguration.ConnectorEventReceiveFromAggregatorConnection, connectorConfiguration.ConnectorEventSendToAggregatorConnection, connectorConfiguration.ConnectorEventReceiveFromWorkerConnection)

	//RUN
	go cg.connectorCommandRoutine.run()
	go cg.connectorEventRoutine.run()
	
	return
}

func (cg ConnectorGandalf) getWorkerCommands(worker string) (workerCommand []string, err error) {
	return cg.connectorCommandsMap[worker]
}

func (cg ConnectorGandalf) addWorkerCommands(worker, command string) (err error) {
	var sizeList = len(cg.connectorCommandsMap[worker])
	cg.connectorCommandsMap[worker][sizeList] = command

}

func (cg ConnectorGandalf) getWorkerCommandSendFile(worker string) (workerCommandFile string, err error) {
	return cg.connectorCommandSendFileMap[worker]
}

func (cg ConnectorGandalf) addWorkerCommandSendFile(worker, commandSendFile string) (err error) {
	cg.connectorCommandSendFileMap[worker] = commandSendFile
}
