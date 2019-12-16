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
	r.workerCaptureCommandReceive = r.context.NewSocket(zmq4.DEALER)
	r.workerCaptureCommandReceive.SetIdentity(r.identity)
	r.workerCaptureCommandReceive.Connect(r.workerCaptureCommandReceiveConnection)
	fmt.Printf("workerCaptureCommandReceive connect : " + workerCaptureCommandReceiveConnection)

	r.workerCaptureEventReceiveConnection = workerCaptureEventReceiveConnection
	r.workerCaptureEventReceive = r.context.NewSocket(zmq4.SUB)
	r.workerCaptureEventReceive.SetIdentity(r.identity)
	r.workerCaptureEventReceive.Connect(r.workerCaptureEventReceiveConnection)
	fmt.Printf("workerCaptureEventReceive connect : " + workerCaptureEventReceiveConnection)
}

func (r ClusterCaptureWorkerRoutine) close() {
	r.workerCaptureCommandReceive.close()
	r.workerCaptureEventReceive.close()
	r.context.close()
}

func (r ClusterCaptureWorkerRoutine) run() {
	poller := zmq4.NewPoller()
	poller.Add(r.workerCaptureCommandReceive, zmq4.POLLIN)
	poller.Add(r.workerCaptureEventReceive, zmq4.POLLIN)

	command := [][]byte{}
	event := [][]byte{}
	

	for {
		r.sendReadyCommand()

		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {

			switch currentSocket := socket.Socket; currentSocket {
			case workerCaptureCommandReceive:

				command, err := currentSocket.RecvMessage()
				if err != nil {
					panic(err)
				}
				err = r.processCommand(command)
				if err != nil {
					panic(err)
				}

			case workerCaptureEventReceive:

				event, err := currentSocket.RecvMessage()
				if err != nil {
					panic(err)
				}
				err = r.processEvent(event)
				if err != nil {
					panic(err)
				}
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
