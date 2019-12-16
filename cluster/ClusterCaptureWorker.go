package cluster

import (
	"fmt"
    "net/http"
	"github.com/pebbe/zmq4"
)

type ClusterCaptureWorkerRoutine struct {
	context									  zmq4.Context
	workerCaptureCommandReceive           zmq4.Socket
	workerCaptureCommandReceiveConnection string
	workerCaptureEventReceive            zmq4.Socket
	workerCaptureEventReceiveConnection   string
	identity                                  string
}

func (r ClusterCaptureWorkerRoutine) New(identity, workerCaptureCommandReceiveConnection, workerCaptureEventReceiveConnection string, topics []string) {
	r.identity = identity

	r.context, _ = zmq4.NewContext()
	r.workerCaptureCommandReceiveConnection = workerCaptureCommandReceiveConnection
	r.workerCaptureCommandReceive = r.context.NewDealer(workerCaptureCommandReceiveConnection)
	r.workerCaptureCommandReceive.SetIdentity(r.identity)
	fmt.Printf("workerCaptureCommandReceive connect : " + workerCaptureCommandReceiveConnection)

	r.workerCaptureEventReceiveConnection = workerCaptureEventReceiveConnection
	r.workerCaptureEventReceive = r.context.NewSub(workerCaptureEventReceiveConnection)
	r.workerCaptureEventReceive.SetIdentity(r.identity)
	fmt.Printf("workerCaptureEventReceive connect : " + workerCaptureEventReceiveConnection)
}

func (r ClusterCaptureWorkerRoutine) close() {
	r.workerCaptureCommandReceive.close()
	r.workerCaptureEventReceive.close()
	r.context.close()
}

func (r ClusterCaptureWorkerRoutine) run() {
	pi := zmq4.PollItems{
		zmq4.PollItem{Socket: workerCaptureCommandReceive, Events: zmq4.POLLIN},
		zmq4.PollItem{Socket: workerCaptureEventReceive, Events: zmq4.POLLIN}}

	command := [][]byte{}
	event := [][]byte{}

	for {
		r.sendReadyCommand()

		pi.Poll(pi, -1)

		switch {
		case pi[0].REvents&zmq4.POLLIN != 0:

			command, err := pi[0].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processCommand(command)
			if err != nil {
				panic(err)
			}

		case pi[1].REvents&zmq4.POLLIN != 0:

			event, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processEvent(event)
			if err != nil {
				panic(err)
			}
		}
	}
	fmt.Println("done")
}

func (r ClusterCaptureWorkerRoutine) processCommand(command [][]byte) {
	command = r.updateHeaderCommand(command)
    response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(command[1]))
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    }
}

func (r ClusterCaptureWorkerRoutine) updateHeaderCommand(command [][]byte) {
}

func (r ClusterCaptureWorkerRoutine) processEvent(event [][]byte) {
	event = r.updateHeaderEvent(event)
    response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(event[0]))
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    }
}

func (r ClusterCaptureWorkerRoutine) updateHeaderEvent(event [][]byte) {
}
