package connector

type ConnectorGandalf struct {
	ConnectorConfiguration      *ConnectorConfiguration
	ConnectorCommandRoutine     *ConnectorCommandRoutine
	ConnectorEventRoutine       *ConnectorEventRoutine
	ConnectorCommandsMap        map[string][]string
	ConnectorCommandSendFileMap map[string]string
}

func NewConnectorGandalf(path string) (connectorGandalf *ConnectorGandalf) {
	connectorGandalf = new(ConnectorGandalf)

	connectorGandalf.ConnectorConfiguration, _ = LoadConfiguration(path)

	connectorGandalf.ConnectorCommandsMap = make(map[string][]string)
	connectorGandalf.ConnectorCommandSendFileMap = make(map[string]string)

	connectorGandalf.ConnectorCommandRoutine = NewConnectorCommandRoutine(connectorGandalf.ConnectorConfiguration.Identity, connectorGandalf.ConnectorConfiguration.ConnectorCommandSendToWorkerConnection, connectorGandalf.ConnectorConfiguration.ConnectorCommandReceiveFromWorkerConnection, connectorGandalf.ConnectorConfiguration.ConnectorCommandReceiveFromAggregatorConnections, connectorGandalf.ConnectorConfiguration.ConnectorCommandSendToAggregatorConnections)
	connectorGandalf.ConnectorEventRoutine = NewConnectorEventRoutine(connectorGandalf.ConnectorConfiguration.Identity, connectorGandalf.ConnectorConfiguration.ConnectorEventSendToWorkerConnection, connectorGandalf.ConnectorConfiguration.ConnectorEventReceiveFromAggregatorConnection, connectorGandalf.ConnectorConfiguration.ConnectorEventSendToAggregatorConnection, connectorGandalf.ConnectorConfiguration.ConnectorEventReceiveFromWorkerConnection)

	//RUN
	//go connectorGandalf.ConnectorCommandRoutine.run()
	//go connectorGandalf.ConnectorEventRoutine.run()

	return
}

func (cg ConnectorGandalf) run() {
	go cg.ConnectorCommandRoutine.run()
	go cg.ConnectorEventRoutine.run()
}

func (cg ConnectorGandalf) getWorkerCommands(worker string) (workerCommand []string) {
	return cg.ConnectorCommandsMap[worker]
}

func (cg ConnectorGandalf) addWorkerCommands(worker, command string) {
	var sizeList = len(cg.ConnectorCommandsMap[worker])
	cg.ConnectorCommandsMap[worker][sizeList] = command

}

func (cg ConnectorGandalf) getWorkerCommandSendFile(worker string) (workerCommandFile string) {
	return cg.ConnectorCommandSendFileMap[worker]
}

func (cg ConnectorGandalf) addWorkerCommandSendFile(worker, commandSendFile string) {
	cg.ConnectorCommandSendFileMap[worker] = commandSendFile
}
