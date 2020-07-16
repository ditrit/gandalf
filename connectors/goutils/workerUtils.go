package goutils

import (
	goclient "github.com/ditrit/gandalf/libraries/goclient"

	worker "github.com/ditrit/gandalf/connectors/go"
)

//WorkerUtils : WorkerUtils
type WorkerUtils interface {
	CreateApplication(clientGandalf *goclient.ClientGandalf, version int64)
	CreateForm(clientGandalf *goclient.ClientGandalf, version int64)
	SendAuthMail(clientGandalf *goclient.ClientGandalf, version int64)
}

//workerUtils : workerUtils
type workerUtils struct {
	worker *worker.Worker

	CreateApplication func(clientGandalf *goclient.ClientGandalf, version int64)
	CreateForm        func(clientGandalf *goclient.ClientGandalf, version int64)
	SendAuthMail      func(clientGandalf *goclient.ClientGandalf, version int64)
}

//NewWorkerUtils : NewWorkerUtils
func NewWorkerUtils(version int64, commandes []string) *workerUtils {
	workerUtils := new(workerUtils)
	workerUtils.worker = worker.NewWorker(version, commandes)
	//workerUtils.worker.Execute = workerUtils.Execute

	return workerUtils
}

//Run : Run
func (wu workerUtils) Run() {
	wu.worker.Run()

	done := make(chan bool)
	wu.CreateApplication(wu.worker.GetClientGandalf(), wu.worker.GetVersion())
	wu.CreateForm(wu.worker.GetClientGandalf(), wu.worker.GetVersion())
	wu.SendAuthMail(wu.worker.GetClientGandalf(), wu.worker.GetVersion())
	<-done
}
