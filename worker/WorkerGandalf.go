package worker

type WorkerGandalf struct {	
	workerConfiguration WorkerConfiguration
	workerRoutine 		WorkerRoutine
}

func (wg WorkerGandalf) main() {
	path := ""
	workerConfiguration := WorkerConfiguration.loadConfiguration(path)

	wg.workerRoutine = WorkerRoutine.new(workerConfiguration.identity, workerConfiguration.workerCommandReceiveConnection, workerConfiguration.workerEventReceiveConnection)

	go wg.workerRoutine.run()
}


