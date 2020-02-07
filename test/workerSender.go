//Package test :
//File workerSender.go
package test

import (
	"fmt"
	"gandalf-go/core/worker"
	"time"
)

//WorkerSender :
type WorkerSender struct {
	WorkerGandalf *worker.WorkerGandalf
}

//NewWorkerSender :
func NewWorkerSender(path string) (workerSender *WorkerSender) {
	workerSender = new(WorkerSender)
	workerSender.WorkerGandalf = worker.NewWorkerGandalf(path)

	return
}

//Run :
func (ws WorkerSender) Run() {
	for {
		fmt.Println("SEND")
		//go ws.WorkerGandalf.ClientGandalfGrpc.SendCommand("test", "100000000", "test", "test", "test", "test", "test")
		time.Sleep(time.Second * 10)

		go ws.WorkerGandalf.ClientGandalfGrpc.SendEvent("toto", "10000", "toto", "toto", "toto")
	}
}
