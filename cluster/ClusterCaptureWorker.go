//Package cluster :
//File ClusterCaptureWorker.go
package cluster

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/pebbe/zmq4"
)

//ClusterCaptureWorkerRoutine :
type ClusterCaptureWorkerRoutine struct {
	Context                               *zmq4.Context
	WorkerCaptureCommandReceive           *zmq4.Socket
	WorkerCaptureCommandReceiveConnection string
	WorkerCaptureEventReceive             *zmq4.Socket
	WorkerCaptureEventReceiveConnection   string
	Identity                              string
}

//NewClusterCaptureWorkerRoutine :
func NewClusterCaptureWorkerRoutine(identity, workerCaptureCommandReceiveConnection, workerCaptureEventReceiveConnection string, topics []string) (clusterCaptureWorkerRoutine *ClusterCaptureWorkerRoutine) {
	clusterCaptureWorkerRoutine = new(ClusterCaptureWorkerRoutine)

	clusterCaptureWorkerRoutine.Identity = identity

	clusterCaptureWorkerRoutine.Context, _ = zmq4.NewContext()
	clusterCaptureWorkerRoutine.WorkerCaptureCommandReceiveConnection = workerCaptureCommandReceiveConnection
	clusterCaptureWorkerRoutine.WorkerCaptureCommandReceive, _ = clusterCaptureWorkerRoutine.Context.NewSocket(zmq4.DEALER)
	clusterCaptureWorkerRoutine.WorkerCaptureCommandReceive.SetIdentity(clusterCaptureWorkerRoutine.Identity)
	clusterCaptureWorkerRoutine.WorkerCaptureCommandReceive.Connect(clusterCaptureWorkerRoutine.WorkerCaptureCommandReceiveConnection)
	fmt.Println("workerCaptureCommandReceive connect : " + workerCaptureCommandReceiveConnection)

	clusterCaptureWorkerRoutine.WorkerCaptureEventReceiveConnection = workerCaptureEventReceiveConnection
	clusterCaptureWorkerRoutine.WorkerCaptureEventReceive, _ = clusterCaptureWorkerRoutine.Context.NewSocket(zmq4.SUB)
	clusterCaptureWorkerRoutine.WorkerCaptureEventReceive.SetIdentity(clusterCaptureWorkerRoutine.Identity)
	clusterCaptureWorkerRoutine.WorkerCaptureEventReceive.Connect(clusterCaptureWorkerRoutine.WorkerCaptureEventReceiveConnection)
	fmt.Println("workerCaptureEventReceive connect : " + workerCaptureEventReceiveConnection)

	return
}

//close :
func (r ClusterCaptureWorkerRoutine) close() {
	r.WorkerCaptureCommandReceive.Close()
	r.WorkerCaptureEventReceive.Close()
	r.Context.Term()
}

//run :
func (r ClusterCaptureWorkerRoutine) run() {
	poller := zmq4.NewPoller()
	poller.Add(r.WorkerCaptureCommandReceive, zmq4.POLLIN)
	poller.Add(r.WorkerCaptureEventReceive, zmq4.POLLIN)

	for {
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {
			switch currentSocket := socket.Socket; currentSocket {
			case r.WorkerCaptureCommandReceive:
				command, err := currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}

				r.processCommand(command)

			case r.WorkerCaptureEventReceive:
				event, err := currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}

				r.processEvent(event)
			}
		}
	}
}

//processCommand :
func (r ClusterCaptureWorkerRoutine) processCommand(command [][]byte) {
	resp, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(command[1]))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %#v\n", err)
	}
	defer resp.Body.Close()
}

//processEvent :
func (r ClusterCaptureWorkerRoutine) processEvent(event [][]byte) {
	resp, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(event[0]))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	defer resp.Body.Close()
}
