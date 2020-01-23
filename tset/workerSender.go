package tset

import (
	"fmt"
	"gandalf-go/worker"
	"time"
)

type WorkerSender struct {
	WorkerGandalf *worker.WorkerGandalf
}

func NewWorkerSender(path string) (workerSender *WorkerSender) {
	workerSender = new(WorkerSender)
	workerSender.WorkerGandalf = worker.NewWorkerGandalf(path)
	return
}

func (ws WorkerSender) Run() {
	for {
		fmt.Println("SEND")
		//go ws.WorkerGandalf.ClientGandalfGrpc.SendCommand("test", "100000000", "test", "test", "test", "test", "test")
		time.Sleep(time.Second * 10)

		go ws.WorkerGandalf.ClientGandalfGrpc.SendEvent("toto", "10000", "toto", "toto", "toto")
	}

}
