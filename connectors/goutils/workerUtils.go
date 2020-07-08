package goutils

import (
	"gandalf/libraries/goclient"

	"github.com/ditrit/gandalf/connectors/go/worker"
)

//Worker
type WorkerWorkflow struct {
	worker worker.Worker

	CreateApplication func(clientGandalf *goclient.ClientGandalf)
	CreateForm        func(clientGandalf *goclient.ClientGandalf)
	SendAuthMail      func(clientGandalf *goclient.ClientGandalf)
}

func NewWorkerWorkflow(version int64, commandes []string) *WorkerWorkflow {
	workerWorkflow := new(WorkerWorkflow)
	workerWorkflow.worker = worker.NewWorker(version, commandes)

	return workerWorkflow
}
