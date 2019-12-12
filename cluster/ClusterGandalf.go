package cluster

import (
    "fmt"
    "os"
)

type ClusterGandalf struct {
	clusterConfiguration 		ClusterConfiguration
	clusterCommandRoutine       ClusterCommandRoutine
	clusterEventRoutine         ClusterEventRoutine
	clusterCaptureWorkerRoutine ClusterCaptureWorkerRoutine
}

func (cg ClusterGandalf) New(path string) {
	path := os.Args[0]
	clusterConfiguration := ClusterConfiguration.loadConfiguration(path)

	cg.clusterCommandRoutine = ClusterCommandRoutine.new(clusterConfiguration.identity, clusterConfiguration.clusterCommandSendConnection, clusterConfiguration.clusterCommandReceiveConnection, clusterConfiguration.clusterCommandCaptureConnection)
	cg.clusterEventRoutine = ClusterEventRoutine.new(clusterConfiguration.identity, clusterConfiguration.clusterEventSendConnection, clusterConfiguration.clusterEventReceiveConnection, clusterConfiguration.clusterEventCaptureConnection)
	cg.clusterCaptureWorkerRoutine = ClusterCaptureWorkerRoutine.new(clusterConfiguration.identity, clusterConfiguration.workerCaptureCommandReceiveCL2WConnection, clusterConfiguration.workerCaptureEventReceiveC2WConnection)

	go cg.clusterCommandRoutine.run()
	go cg.clusterEventRoutine.run()
	go cg.clusterCaptureWorkerRoutine.run()
}
