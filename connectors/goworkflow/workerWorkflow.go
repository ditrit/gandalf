package goworkflow

import (
	"fmt"

	worker "github.com/ditrit/gandalf/connectors/go"
	goclient "github.com/ditrit/gandalf/libraries/goclient"
)

type WorkerWorkflow interface {
	Upload(clientGandalf *goclient.ClientGandalf, version int64)
}

//Worker
type workerWorkflow struct {
	worker *worker.Worker

	Upload func(clientGandalf *goclient.ClientGandalf, version int64)
}

func NewWorkerWorkflow(version int64, commandes []string) *workerWorkflow {
	workerWorkflow := new(workerWorkflow)
	workerWorkflow.worker = worker.NewWorker(version, commandes)
	workerWorkflow.worker.Execute = workerWorkflow.Execute

	return workerWorkflow
}

func (ww workerWorkflow) Execute() {
	fmt.Println("EXECUTE")
	ww.Upload(ww.worker.GetClientGandalf(), ww.worker.GetVersion())
}

func (ww workerWorkflow) Run() {
	fmt.Println("RUN")
	ww.worker.Run()
}
