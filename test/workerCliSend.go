package test

import (
	"fmt"
	"libraries/goclient"
)

type WorkerCliSend struct {
	client      *goclient.ClientGandalf
	messageType string
	value       string
	topic       string
	payload     string
}

func NewWorkerCliSend(identity, messageType, value, payload, topic string, connections []string) *WorkerCliSend {
	workerCliSend := new(WorkerCliSend)
	workerCliSend.messageType = messageType
	workerCliSend.value = value
	workerCliSend.topic = topic
	workerCliSend.payload = payload
	workerCliSend.client = goclient.NewClientGandalf(identity, connections)

	//workerCliSend.client = NewClientGrpcTest(identity, connection)

	return workerCliSend
}

func (r WorkerCliSend) Run() {
	if r.messageType == "cmd" {
		commandUUID := r.client.SendCommand("100000", "test", r.value, r.payload)
		fmt.Println("commandUUID")
		fmt.Println(commandUUID)
		fmt.Println("commandUUID")
		if commandUUID != nil {
			id := r.client.CreateIteratorEvent()
			fmt.Println(id)
			for {

				event := r.client.WaitTopic(commandUUID.GetUUID(), id)
				fmt.Println(event)

				if event.GetEvent() == "SUCCES" || event.GetEvent() == "FAIL" {
					fmt.Println(event.GetPayload())
					break
				}
			}
		}
	} else if r.messageType == "evt" {
		r.client.SendEvent(r.topic, "100000", r.value, r.payload)
	}

	//r.client.SendEvent("test", "10000", "test", "test", "test")
}
