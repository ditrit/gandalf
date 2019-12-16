package cluster

import (
	"fmt"
    "net/http"
	"github.com/pebbe/zmq4"
)

type ClusterCaptureWorkerRoutine struct {
	context									  *zmq4.Context
	workerCaptureCommandReceiveCL2W           *zmq4.Socket
	workerCaptureCommandReceiveCL2WConnection string
	workerCaptureEventReceiveCL2W             *zmq4.Socket
	workerCaptureEventReceiveCL2WConnection   string
	identity                                  string
}

func (r ClusterCaptureWorkerRoutine) New(identity, workerCaptureCommandReceiveCL2WConnection, workerCaptureEventReceiveC2WConnection string, topics *string) {
	r.identity = identity

	context, _ := zmq4.NewContext()
	r.workerCaptureCommandReceiveCL2WConnection = workerCaptureCommandReceiveCL2WConnection
	r.workerCaptureCommandReceiveCL2W = context.NewDealer(workerCaptureCommandReceiveCL2WConnection)
	r.workerCaptureCommandReceiveCL2W.Identity(r.identity)
	fmt.Printf("workerCaptureCommandReceiveCL2W connect : " + workerCaptureCommandReceiveCL2WConnection)

	r.workerCaptureEventReceiveCL2WConnection = workerCaptureEventReceiveCL2WConnection
	r.workerCaptureEventReceiveC2W = context.NewSub(workerCaptureEventReceiveCL2WConnection)
	r.workerCaptureEventReceiveC2W.Identity(r.identity)
	fmt.Printf("workerCaptureEventReceiveC2W connect : " + workerCaptureEventReceiveCL2WConnection)
}

func (r ClusterCaptureWorkerRoutine) close() {
	r.workerCaptureCommandReceiveC2W.close()
	r.workerCaptureEventReceiveC2W.close()
	r.Context.close()
}

func (r ClusterCaptureWorkerRoutine) run() {
	pi := zmq4.PollItems{
		zmq4.PollItem{Socket: workerCaptureCommandReceiveCL2W, Events: zmq4.POLLIN},
		zmq4.PollItem{Socket: workerCaptureEventReceiveC2W, Events: zmq4.POLLIN}}

	var command = [][]byte{}
	var event = [][]byte{}

	for {
		r.sendReadyCommand()

		_, _ = zmq4.Poll(pi, -1)

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
