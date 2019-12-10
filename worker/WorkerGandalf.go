package worker

type WorkerGandalf struct {	
	workerConfiguration WorkerConfiguration
	workerRoutine 		WorkerRoutine
}

func (wg WorkerGandalf) main() {
	path := ""
	workerConfiguration := WorkerConfiguration.loadConfiguration(path)

	wg.workerRoutine = WorkerRoutine.New(workerConfiguration.identity, workerConfiguration.workerCommandReceiveConnection, workerConfiguration.workerEventReceiveConnection)
}


