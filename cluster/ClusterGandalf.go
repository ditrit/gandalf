package cluster

type ClusterGandalf struct {
	clusterConfiguration 		*ClusterConfiguration
	clusterCommandRoutine       *ClusterCommandRoutine
	clusterEventRoutine         *ClusterEventRoutine
	clusterCaptureWorkerRoutine *ClusterCaptureWorkerRoutine
}

func NewClusterGandalf(path string) (clusterGandalf ClusterGandalf) {
	clusterGandalf = new(ClusterGandalf)

	clusterGandalf.clusterConfiguration, _ = LoadConfiguration(path)

	clusterGandalf.clusterCommandRoutine = NewClusterCommandRoutine(clusterGandalf.clusterConfiguration.Identity, clusterGandalf.clusterConfiguration.ClusterCommandSendConnection, clusterGandalf.clusterConfiguration.ClusterCommandReceiveConnection, clusterGandalf.clusterConfiguration.ClusterCommandCaptureConnection)
	clusterGandalf.clusterEventRoutine = NewClusterEventRoutine(clusterGandalf.clusterConfiguration.Identity, clusterGandalf.clusterConfiguration.ClusterEventSendConnection, clusterGandalf.clusterConfiguration.ClusterEventReceiveConnection, clusterGandalf.clusterConfiguration.ClusterEventCaptureConnection)
	clusterGandalf.clusterCaptureWorkerRoutine = NewClusterCaptureWorkerRoutine(clusterGandalf.clusterConfiguration.Identity, clusterGandalf.clusterConfiguration.WorkerCaptureCommandReceiveConnection, clusterGandalf.clusterConfiguration.WorkerCaptureEventReceiveConnection, clusterGandalf.clusterConfiguration.Topics)

	go clusterGandalf.clusterCommandRoutine.run()
	go clusterGandalf.clusterEventRoutine.run()
	go clusterGandalf.clusterCaptureWorkerRoutine.run()
}
