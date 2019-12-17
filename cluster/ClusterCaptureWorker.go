package cluster

import (
	"fmt"
	"net/http"
	"bytes"
	"errors"
	"github.com/pebbe/zmq4"
)

type ClusterCaptureWorkerRoutine struct {
	Context									  	*zmq4.Context
	WorkerCaptureCommandReceive           		*zmq4.Socket
	WorkerCaptureCommandReceiveConnection 		string
	WorkerCaptureEventReceive            		*zmq4.Socket
	WorkerCaptureEventReceiveConnection   		string
	Identity                                  	string
}

func NewClusterCaptureWorkerRoutine(identity, workerCaptureCommandReceiveConnection, workerCaptureEventReceiveConnection string, topics []string) (clusterCaptureWorkerRoutine *ClusterCaptureWorkerRoutine) {
	clusterCaptureWorkerRoutine = new(ClusterCaptureWorkerRoutine)
	
	clusterCaptureWorkerRoutine.Identity = identity

	clusterCaptureWorkerRoutine.Context, _ = zmq4.NewContext()
	clusterCaptureWorkerRoutine.WorkerCaptureCommandReceiveConnection = workerCaptureCommandReceiveConnection
	clusterCaptureWorkerRoutine.WorkerCaptureCommandReceive, _ = clusterCaptureWorkerRoutine.Context.NewSocket(zmq4.DEALER)
	clusterCaptureWorkerRoutine.WorkerCaptureCommandReceive.SetIdentity(clusterCaptureWorkerRoutine.Identity)
	clusterCaptureWorkerRoutine.WorkerCaptureCommandReceive.Connect(clusterCaptureWorkerRoutine.WorkerCaptureCommandReceiveConnection)
	fmt.Printf("workerCaptureCommandReceive connect : " + workerCaptureCommandReceiveConnection)

	clusterCaptureWorkerRoutine.WorkerCaptureEventReceiveConnection = workerCaptureEventReceiveConnection
	clusterCaptureWorkerRoutine.WorkerCaptureEventReceive, _ = clusterCaptureWorkerRoutine.Context.NewSocket(zmq4.SUB)
	clusterCaptureWorkerRoutine.WorkerCaptureEventReceive.SetIdentity(clusterCaptureWorkerRoutine.Identity)
	clusterCaptureWorkerRoutine.WorkerCaptureEventReceive.Connect(clusterCaptureWorkerRoutine.WorkerCaptureEventReceiveConnection)
	fmt.Printf("workerCaptureEventReceive connect : " + workerCaptureEventReceiveConnection)

	return
}

func (r ClusterCaptureWorkerRoutine) close() {
	r.WorkerCaptureCommandReceive.Close()
	r.WorkerCaptureEventReceive.Close()
	r.Context.Term()
}

func (r ClusterCaptureWorkerRoutine) run() {
	poller := zmq4.NewPoller()
	poller.Add(r.WorkerCaptureCommandReceive, zmq4.POLLIN)
	poller.Add(r.WorkerCaptureEventReceive, zmq4.POLLIN)

	command := [][]byte{}
	event := [][]byte{}
	err := errors.New("")

	for {

		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {

			switch currentSocket := socket.Socket; currentSocket {
			case r.WorkerCaptureCommandReceive:

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				err = r.processCommand(command)
				if err != nil {
					panic(err)
				}

			case r.WorkerCaptureEventReceive:

				event, err = currentSocket.RecvMessageBytes(0)
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

func (r ClusterCaptureWorkerRoutine) processCommand(command [][]byte) (err error) {
    _, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(command[1]))
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	return
}

func (r ClusterCaptureWorkerRoutine) processEvent(event [][]byte) (err error) {
    _, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(event[0]))
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	return
}

