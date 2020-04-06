package test

import "fmt"

type WorkerSend struct {
	client *ClientGrpcTest
}

func NewWorkerSend(identity, connection string) *WorkerSend {
	workerSend := new(WorkerSend)
	workerSend.client = NewClientGrpcTest(identity, connection)

	return workerSend
}

func (r WorkerSend) Run() {
	commandUUID := r.client.SendCommand("100000000", "test", "test", "test")
	if commandUUID != nil {
		id := r.client.CreateIteratorEvent()
		for {
			event := r.client.WaitTopic(commandUUID.GetUUID(), id)
			fmt.Println(event)

			if event.GetEvent() == "SUCCES" || event.GetEvent() == "FAIL" {
				fmt.Println(event.GetPayload())
				break
			}
		}
	}

	//r.client.SendEvent("test", "10000", "test", "test", "test")
}
