package connector

import (
	"fmt"
    "os"
)

type ConnectorGandalf struct {
	connectorConfiguration 		ConnectorConfiguration
	connectorCommandRoutine     ConnectorCommandRoutine
	connectorEventRoutine       ConnectorEventRoutine
	connectorCommandsMap        map[string][]string
	connectorCommandSendFileMap map[string]string
}

func (cg ConnectorGandalf) New(path string) err error {
	path := os.Args[0]
	connectorConfiguration := ConnectorConfiguration.loadConfiguration(path)

	cg.connectorCommandsMap = make(map[string][]string)
	cg.connectorCommandSendFileMap = make(map[string]string)

	cg.connectorCommandRoutine = ConnectorCommandRoutine.New(connectorConfiguration.identity, connectorConfiguration.connectorCommandSendA2WConnection, connectorConfiguration.connectorCommandReceiveA2WConnection, connectorConfiguration.connectorCommandSendW2AConnection, connectorConfiguration.connectorCommandReceiveW2AConnection)
	cg.connectorEventRoutine = ConnectorEventRoutine.New(connectorConfiguration.identity, connectorConfiguration.connectorEventSendA2WConnection, connectorConfiguration.connectorEventReceiveA2WConnection, connectorConfiguration.connectorEventSendW2AConnection, connectorConfiguration.connectorEventReceiveW2AConnection)

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
