package test

import (
	"fmt"
	"strconv"
	"time"
)

var receiveIndex = 0

type WorkerCliReceive struct {
	clients     []*ClientGrpcTest
	messageType string
	value       string
	topic       string
}

func NewWorkerCliReceive(identity, messageType, value, topic string, connections []string) *WorkerCliReceive {
	workerCliReceive := new(WorkerCliReceive)
	workerCliReceive.messageType = messageType
	workerCliReceive.value = value
	workerCliReceive.topic = topic
	for _, connection := range connections {
		workerCliReceive.clients = append(workerCliReceive.clients, NewClientGrpcTest(identity, connection))
	}
	//workerCliReceive.clients = append() NewClientGrpcTest(identity, connection)

	return workerCliReceive
}

func (r WorkerCliReceive) Run() {
	client := r.clients[getReceiveIndex(r.clients)]
	if r.messageType == "cmd" {
		id := client.CreateIteratorCommand()
		command := client.WaitCommand(r.value, id)
		fmt.Println(command)
		//id := r.client.CreateIteratorEvent()
		//event := r.client.WaitEvent("test", "test", id)
		for i := 1; i <= 10; i++ {
			client.SendEvent(command.GetUUID(), "10000", strconv.Itoa(i*10), "test")
			time.Sleep(time.Duration(500) * time.Millisecond)

		}
		client.SendEvent(command.GetUUID(), "10000", "SUCCES", "test")

	} else if r.messageType == "evt" {
		id := client.CreateIteratorEvent()
		event := client.WaitEvent(r.value, r.topic, id)
		fmt.Println(event)
	}

}

func getReceiveIndex(conns []*ClientGrpcTest) int {
	aux := receiveIndex
	receiveIndex++
	if receiveIndex >= len(conns) {
		receiveIndex = 0
	}
	return aux
}
