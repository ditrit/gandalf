package test

import (
	"fmt"
	"libraries/goclient"
	"strconv"
	"time"
)

type WorkerCliReceive struct {
	client      *goclient.ClientGandalf
	messageType string
	value       string
	topic       string
}

func NewWorkerCliReceive(identity, messageType, value, topic string, connections []string) *WorkerCliReceive {
	workerCliReceive := new(WorkerCliReceive)
	workerCliReceive.messageType = messageType
	workerCliReceive.value = value
	workerCliReceive.topic = topic
	workerCliReceive.client = goclient.NewClientGandalf(identity, connections)

	//workerCliReceive.clients = append() NewClientGrpcTest(identity, connection)

	return workerCliReceive
}

func (r WorkerCliReceive) Run() {
	if r.messageType == "cmd" {
		id := r.client.CreateIteratorCommand()
		fmt.Println(id)
		command := r.client.WaitCommand(r.value, id)
		fmt.Println(command)
		//id := r.client.CreateIteratorEvent()
		//event := r.client.WaitEvent("test", "test", id)
		for i := 1; i <= 10; i++ {
			r.client.SendEvent(command.GetUUID(), "10000", strconv.Itoa(i*10), "test")
			time.Sleep(time.Duration(500) * time.Millisecond)

		}
		r.client.SendEvent(command.GetUUID(), "10000", "SUCCES", "test")

	} else if r.messageType == "evt" {
		id := r.client.CreateIteratorEvent()
		event := r.client.WaitEvent(r.value, r.topic, id)
		fmt.Println(event)
	}

}
