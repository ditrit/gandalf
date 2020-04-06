package test

import "fmt"

var sendIndex = 0

type WorkerCliSend struct {
	clients     []*ClientGrpcTest
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
	for _, connection := range connections {
		workerCliSend.clients = append(workerCliSend.clients, NewClientGrpcTest(identity, connection))
	}
	//workerCliSend.client = NewClientGrpcTest(identity, connection)

	return workerCliSend
}

func (r WorkerCliSend) Run() {
	client := r.clients[getSendIndex(r.clients)]
	if r.messageType == "cmd" {
		commandUUID := client.SendCommand("100000", "test", r.value, r.payload)
		if commandUUID != nil {
			id := client.CreateIteratorEvent()
			for {
				event := client.WaitTopic(commandUUID.GetUUID(), id)
				fmt.Println(event)

				if event.GetEvent() == "SUCCES" || event.GetEvent() == "FAIL" {
					fmt.Println(event.GetPayload())
					break
				}
			}
		}
	} else if r.messageType == "evt" {
		client.SendEvent(r.topic, "100000", r.value, r.payload)
	}

	//r.client.SendEvent("test", "10000", "test", "test", "test")
}

func getSendIndex(conns []*ClientGrpcTest) int {
	aux := sendIndex
	sendIndex++
	if sendIndex >= len(conns) {
		sendIndex = 0
	}
	return aux
}
