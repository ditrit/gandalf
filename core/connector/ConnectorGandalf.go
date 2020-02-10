//Package connector :
//File ConnectorGandalf.go
package connector

import "fmt"

//ConnectorGandalf :
type ConnectorGandalf struct {
	connectorConfiguration      *ConnectorConfiguration
	connectorCommandRoutine     *ConnectorCommandRoutine
	connectorEventRoutine       *ConnectorEventRoutine
	connectorCommandsMap        map[string][]string
	connectorCommandSendFileMap map[string]string
	connectorStopChannel        chan int
}

//NewConnectorGandalf :
func NewConnectorGandalf(path string) (connectorGandalf *ConnectorGandalf) {
	connectorGandalf = new(ConnectorGandalf)
	connectorGandalf.connectorStopChannel = make(chan int)

	connectorGandalf.connectorConfiguration, _ = NewConnectorConfiguration(path)

	connectorGandalf.connectorCommandsMap = make(map[string][]string)
	connectorGandalf.connectorCommandSendFileMap = make(map[string]string)

	connectorGandalf.connectorCommandRoutine = NewConnectorCommandRoutine(
		connectorGandalf.connectorConfiguration.Identity,
		connectorGandalf.connectorConfiguration.ConnectorCommandWorkerConnection,
		connectorGandalf.connectorConfiguration.ConnectorCommandReceiveFromAggregatorConnections,
		connectorGandalf.connectorConfiguration.ConnectorCommandSendToAggregatorConnections)
	connectorGandalf.connectorEventRoutine = NewConnectorEventRoutine(
		connectorGandalf.connectorConfiguration.Identity,
		connectorGandalf.connectorConfiguration.ConnectorEventWorkerConnection,
		connectorGandalf.connectorConfiguration.ConnectorEventReceiveFromAggregatorConnections,
		connectorGandalf.connectorConfiguration.ConnectorEventSendToAggregatorConnections)

	//RUN
	//go connectorGandalf.ConnectorCommandRoutine.run()
	//go connectorGandalf.ConnectorEventRoutine.run()

	return
}

//Run :
func (cg ConnectorGandalf) Run() {
	go cg.connectorCommandRoutine.run()
	go cg.connectorCommandRoutine.startGrpcServer()
	go cg.connectorEventRoutine.run()
	go cg.connectorEventRoutine.startGrpcServer()

	<-cg.connectorStopChannel
	fmt.Println("quit")

	cg.connectorCommandRoutine.close()
	cg.connectorEventRoutine.close()
}

//Stop :
func (cg ConnectorGandalf) Stop() {
	cg.connectorStopChannel <- 0
}

//getWorkerCommands :
func (cg ConnectorGandalf) getWorkerCommands(worker string) (workerCommand []string) {
	return cg.connectorCommandsMap[worker]
}

//addWorkerCommands :
func (cg *ConnectorGandalf) addWorkerCommands(worker, command string) {
	cg.connectorCommandsMap[worker] = append(cg.connectorCommandsMap[worker], command)
}

//getWorkerCommandSendFile :
func (cg ConnectorGandalf) getWorkerCommandSendFile(worker string) (workerCommandFile string) {
	return cg.connectorCommandSendFileMap[worker]
}

//addWorkerCommandSendFile :
func (cg ConnectorGandalf) addWorkerCommandSendFile(worker, commandSendFile string) {
	cg.connectorCommandSendFileMap[worker] = commandSendFile
}
