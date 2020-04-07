//Package test :
//File workerReceiver.go
package test

import (
	"fmt"
	"gandalf-go/core/worker"
)

//WorkerReceiver :
type WorkerReceiver struct {
	WorkerGandalf *worker.WorkerGandalf
}

//NewWorkerReceiver :
func NewWorkerReceiver(path string) (workerReceiver *WorkerReceiver) {
	workerReceiver = new(WorkerReceiver)
	workerReceiver.WorkerGandalf = worker.NewWorkerGandalf(path)

	return
}

//Run :
func (wr WorkerReceiver) Run() {
	//commandMessage := wr.WorkerGandalf.ClientGandalfGrpc.WaitCommand("test")
	commandMessage := wr.WorkerGandalf.ClientGandalfGrpc.WaitEvent("toto", "toto")

	fmt.Println("RECEIVE")
	fmt.Println(commandMessage)
}
