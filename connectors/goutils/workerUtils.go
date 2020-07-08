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

func NewWorkerWorkflow(version int64, commandes []string) *WorkerUtils {
	workerUtils := new(WorkerUtils)
	workerUtils.worker = worker.NewWorker(version, commandes)
	workerUtils.worker.Execute = workerUtils.Execute

	return workerUtils
}

func (wu WorkerUtils) Execute() {
	wu.worker.Run()
	wu.CreateApplication(wu.worker.clientGandalf, wu.worker.version)
	wu.CreateForm(wu.worker.clientGandalf, wu.worker.version)
	wu.SendAuthMail(wu.worker.clientGandalf, wu.worker.version)
}
