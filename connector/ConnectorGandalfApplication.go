package connector

type ConnectorGandalfApplication struct {
	workerCommandsMap        map[string]List
	workerCommandSendFileMap map[string]string
}

func (cg ConnectorGandalfApplication) new() {
	cw.workerCommandsMap = make(map[string]List)
	cw.workerCommandSendFileMap = make(map[string]string)
}

func (cg ConnectorGandalfApplication) getWorkerCommands(worker string) List {
	return cg.workerCommandsMap.get(worker)
}

func (cg ConnectorGandalfApplication) addWorkerCommands(worker, command string) {

}

func (cg ConnectorGandalfApplication) getWorkerCommandSendFile(worker string) string {
	return cg.workerCommandSendFileMap.get(worker)
}

func (cg ConnectorGandalfApplication) addWorkerCommandSendFile(worker, commandSendFile string) {

}
