package goworkflow

import (
	"gandalf/libraries/goclient"

	"github.com/ditrit/gandalf/connectors/go/worker"
)

//Worker
type WorkerWorkflow struct {
	worker worker.Worker

	Upload func(clientGandalf *goclient.ClientGandalf)
}

func NewWorkerWorkflow(version int64, commandes []string) *WorkerWorkflow {
	workerWorkflow := new(WorkerWorkflow)
	workerWorkflow.worker = worker.NewWorker(version, commandes)

	return workerWorkflow
}
