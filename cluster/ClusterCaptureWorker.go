package cluster

import (
	"fmt"
    "net/http"
	"github.com/pebbe/zmq4"
)

type ClusterCaptureWorkerRoutine struct {
	Context									  zmq4.Context
	WorkerCaptureCommandReceive           	zmq4.Socket
	WorkerCaptureCommandReceiveConnection 		string
	WorkerCaptureEventReceive            zmq4.Socket
	WorkerCaptureEventReceiveConnection   string
	Identity                                  string
}

func NewClusterCaptureWorkerRoutine(identity, workerCaptureCommandReceiveConnection, workerCaptureEventReceiveConnection string, topics []string) (clusterCaptureWorkerRoutine *ClusterCaptureWorkerRoutine) {
	clusterCaptureWorkerRoutine = new(ClusterCaptureWorkerRoutine)
	
	clusterCaptureWorkerRoutine.identity = identity

	clusterCaptureWorkerRoutine.context, _ = zmq4.NewContext()
	clusterCaptureWorkerRoutine.workerCaptureCommandReceiveConnection = workerCaptureCommandReceiveConnection
	clusterCaptureWorkerRoutine.workerCaptureCommandReceive, _ = clusterCaptureWorkerRoutine.context.NewSocket(zmq4.DEALER)
	clusterCaptureWorkerRoutine.workerCaptureCommandReceive.SetIdentity(clusterCaptureWorkerRoutine.identity)
	clusterCaptureWorkerRoutine.workerCaptureCommandReceive.Connect(clusterCaptureWorkerRoutine.workerCaptureCommandReceiveConnection)
	fmt.Printf("workerCaptureCommandReceive connect : " + workerCaptureCommandReceiveConnection)

	clusterCaptureWorkerRoutine.workerCaptureEventReceiveConnection = workerCaptureEventReceiveConnection
	clusterCaptureWorkerRoutine.workerCaptureEventReceive, _ = clusterCaptureWorkerRoutine.context.NewSocket(zmq4.SUB)
	clusterCaptureWorkerRoutine.workerCaptureEventReceive.SetIdentity(clusterCaptureWorkerRoutine.identity)
	clusterCaptureWorkerRoutine.workerCaptureEventReceive.Connect(clusterCaptureWorkerRoutine.workerCaptureEventReceiveConnection)
	fmt.Printf("workerCaptureEventReceive connect : " + workerCaptureEventReceiveConnection)

	return
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
