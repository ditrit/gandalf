package cluster

import "fmt"

type ClusterGandalf struct {
	clusterConfiguration        *ClusterConfiguration
	clusterCommandRoutine       *ClusterCommandRoutine
	clusterEventRoutine         *ClusterEventRoutine
	clusterCaptureWorkerRoutine *ClusterCaptureWorkerRoutine
}

func NewClusterGandalf(path string) (clusterGandalf *ClusterGandalf) {
	clusterGandalf = new(ClusterGandalf)
	//clusterGandalf = ClusterGandalf{}

	clusterGandalf.clusterConfiguration, _ = LoadConfiguration(path)
	fmt.Print(clusterGandalf.clusterConfiguration)
	fmt.Print("totototo")
	fmt.Print(clusterGandalf.clusterConfiguration.ClusterCommandSendConnection)

	clusterGandalf.clusterCommandRoutine = NewClusterCommandRoutine(clusterGandalf.clusterConfiguration.Identity, clusterGandalf.clusterConfiguration.ClusterCommandSendConnection, clusterGandalf.clusterConfiguration.ClusterCommandReceiveConnection, clusterGandalf.clusterConfiguration.ClusterCommandCaptureConnection)
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
}
