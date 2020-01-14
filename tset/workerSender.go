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
		go ws.WorkerGandalf.ClientGandalf.SendCommand("toto", "100", "toto", "toto", "toto", "toto", "toto")

		time.Sleep(time.Second * 5)
	}

}
