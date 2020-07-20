package goworkflow

import (
	worker "github.com/ditrit/gandalf/connectors/go"
	goclient "github.com/ditrit/gandalf/libraries/goclient"
)

//WorkerWorkflow : WorkerWorkflow
type WorkerWorkflow interface {
	Upload(clientGandalf *goclient.ClientGandalf, version int64)
}

//workerWorkflow : workerWorkflow
type workerWorkflow struct {
	worker *worker.Worker

	Upload func(clientGandalf *goclient.ClientGandalf, version int64)
}

//NewWorkerWorkflow : NewWorkerWorkflow
func NewWorkerWorkflow(version int64, commandes []string) *workerWorkflow {
	currentWorkerWorkflow := new(workerWorkflow)
	currentWorkerWorkflow.worker = worker.NewWorker(version, commandes)
	//currentWorkerWorkflow.worker.Execute = workerWorkflow.Execute

	return currentWorkerWorkflow
}

//Run : Run
func (ww workerWorkflow) Run() {
	ww.worker.Run()

	done := make(chan bool)
	//START WORKER ADMIN
	ww.Upload(ww.worker.GetClientGandalf(), ww.worker.GetVersion())
	<-done
}
