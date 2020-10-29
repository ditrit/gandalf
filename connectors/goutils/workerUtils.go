package goutils

import (
	goclient "github.com/ditrit/gandalf/libraries/goclient"

	worker "github.com/ditrit/gandalf/connectors/go"
)

//WorkerUtils : WorkerUtils
type WorkerUtils interface {
	CreateApplication(clientGandalf *goclient.ClientGandalf, major, minor int64)
	CreateForm(clientGandalf *goclient.ClientGandalf, major, minor int64)
	SendAuthMail(clientGandalf *goclient.ClientGandalf, major, minor int64)
}

//workerUtils : workerUtils
type workerUtils struct {
	worker *worker.Worker

	CreateApplication func(clientGandalf *goclient.ClientGandalf, major, minor int64) int
	CreateForm        func(clientGandalf *goclient.ClientGandalf, major, minor int64) int
	SendAuthMail      func(clientGandalf *goclient.ClientGandalf, major, minor int64) int
}

//NewWorkerUtils : NewWorkerUtils
func NewWorkerUtils(major, minor int64, commandes []string) *workerUtils {
	workerUtils := new(workerUtils)
	workerUtils.worker = worker.NewWorker(major, minor, commandes)
	//workerUtils.worker.Execute = workerUtils.Execute

	return workerUtils
}

//Run : Run
func (wu workerUtils) Run() {

	wu.worker.CommandesFuncs["CREATE_FORM"] = wu.CreateForm
	wu.worker.CommandesFuncs["SEND_AUTH_MAIL"] = wu.SendAuthMail

	wu.worker.Run()
}
