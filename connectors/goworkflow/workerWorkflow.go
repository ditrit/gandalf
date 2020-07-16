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
	currentWorkerWorkflow := new(workerWorkflow)
	currentWorkerWorkflow.worker = worker.NewWorker(version, commandes)
	currentWorkerWorkflow.worker.Execute = workerWorkflow.Execute

	return currentWorkerWorkflow
}

/* //GetClientGandalf
func (ww workerWorkflow) GetWorker() *goclient.ClientGandalf {
	return ww.worker
} */

func (ww workerWorkflow) Execute() {
	fmt.Println("EXECUTE UPLOAD")
	fmt.Println("UPLOAD")
	fmt.Println(ww.Upload)
	fmt.Println("UPLOAD")
	ww.Upload(ww.worker.GetClientGandalf(), ww.worker.GetVersion())
	fmt.Println("END EXECUTE")
}

func (ww workerWorkflow) Run() {
	fmt.Println("RUN")
	fmt.Println("UPLOAD")
	fmt.Println(ww.Upload)
	fmt.Println("UPLOAD")
	ww.worker.Run()
}
