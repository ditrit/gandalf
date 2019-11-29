package connector

type ConnectorWorkerMap struct {
	workerCommandsMap        map[string]List
	workerCommandSendFileMap map[string]string
}

func (cw ConnectorWorkerMap) new() {
	cw.workerCommandSendFileMap = ""
	cw.workerCommandsMap = ""
}

func (cw ConnectorWorkerMap) close() {
}

func (cw ConnectorWorkerMap) reconnectToProxy() {

}
