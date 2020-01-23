package tset

import (
	"fmt"
	"gandalf-go/worker"
)

type WorkerReceiver struct {
	WorkerGandalf *worker.WorkerGandalf
}

func NewWorkerReceiver(path string) (workerReceiver *WorkerReceiver) {
	workerReceiver = new(WorkerReceiver)
	workerReceiver.WorkerGandalf = worker.NewWorkerGandalf(path)
	return
}

func (wr WorkerReceiver) Run() {
	//commandMessage := wr.WorkerGandalf.ClientGandalfGrpc.WaitCommand("test")
	commandMessage := wr.WorkerGandalf.ClientGandalfGrpc.WaitEvent("toto", "toto")
	fmt.Println("RECEIVE")
	fmt.Println(commandMessage)

}
