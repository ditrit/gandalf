package cluster

import "fmt"

type ClusterGandalf struct {
	clusterConfiguration        *ClusterConfiguration
	clusterCommandRoutine       *ClusterCommandRoutine
	clusterEventRoutine         *ClusterEventRoutine
	clusterCaptureWorkerRoutine *ClusterCaptureWorkerRoutine
	clusterStopChannel          chan int
}

func NewClusterGandalf(path string) (clusterGandalf *ClusterGandalf) {
	clusterGandalf = new(ClusterGandalf)
	clusterGandalf.clusterStopChannel = make(chan int)

	clusterGandalf.clusterConfiguration, _ = LoadConfiguration(path)

	clusterGandalf.clusterCommandRoutine = NewClusterCommandRoutine(clusterGandalf.clusterConfiguration.Identity, clusterGandalf.clusterConfiguration.ClusterCommandSendConnection, clusterGandalf.clusterConfiguration.ClusterCommandReceiveConnection, clusterGandalf.clusterConfiguration.ClusterCommandCaptureConnection, clusterGandalf.clusterConfiguration.DatabaseClusterConnections)
	clusterGandalf.clusterEventRoutine = NewClusterEventRoutine(clusterGandalf.clusterConfiguration.Identity, clusterGandalf.clusterConfiguration.ClusterEventSendConnection, clusterGandalf.clusterConfiguration.ClusterEventReceiveConnection, clusterGandalf.clusterConfiguration.ClusterEventCaptureConnection)
	clusterGandalf.clusterCaptureWorkerRoutine = NewClusterCaptureWorkerRoutine(clusterGandalf.clusterConfiguration.Identity, clusterGandalf.clusterConfiguration.WorkerCaptureCommandReceiveConnection, clusterGandalf.clusterConfiguration.WorkerCaptureEventReceiveConnection, clusterGandalf.clusterConfiguration.Topics)

	//go clusterGandalf.clusterCommandRoutine.run()
	//go clusterGandalf.clusterEventRoutine.run()
	//go clusterGandalf.clusterCaptureWorkerRoutine.run()
	return
}

func (cg ClusterGandalf) Run() {
	go cg.clusterCommandRoutine.run()
	go cg.clusterEventRoutine.run()
	go cg.clusterCaptureWorkerRoutine.run()

	defer cg.clusterCommandRoutine.close()
	defer cg.clusterEventRoutine.close()
	defer cg.clusterCaptureWorkerRoutine.close()

	<-cg.clusterStopChannel
	fmt.Println("quit")
}

func (cg ClusterGandalf) Stop() {
	cg.clusterStopChannel <- 0
}
