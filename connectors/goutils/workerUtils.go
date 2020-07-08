package goutils

import (
	"gandalf/libraries/goclient"

	"github.com/ditrit/gandalf/connectors/go/worker"
)

//Worker
type WorkerUtils struct {
	worker worker.Worker

	CreateApplication func(clientGandalf *goclient.ClientGandalf)
	CreateForm        func(clientGandalf *goclient.ClientGandalf)
	SendAuthMail      func(clientGandalf *goclient.ClientGandalf)
}

func NewWorkerWorkflow(version int64, commandes []string) *WorkerWorkflow {
	workerUtils := new(WorkerUtils)
	workerUtils.worker = worker.NewWorker(version, commandes)

	return workerUtils
}

func (wu WorkerUtils) Run() {
	wu.worker.Run()
	wu.CreateApplication(wu.worker.clientGandalf)
	wu.CreateForm(wu.worker.clientGandalf)
	wu.SendAuthMail(wu.worker.clientGandalf)
}
