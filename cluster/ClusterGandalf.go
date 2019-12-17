package cluster

type ClusterGandalf struct {
	clusterConfiguration 		*ClusterConfiguration
	clusterCommandRoutine       *ClusterCommandRoutine
	clusterEventRoutine         *ClusterEventRoutine
	clusterCaptureWorkerRoutine *ClusterCaptureWorkerRoutine
}

func (cg ClusterGandalf) New(path string) {
	cg.clusterConfiguration, _ = LoadConfiguration(path)

	cg.clusterCommandRoutine = NewClusterCommandRoutine(cg.clusterConfiguration.Identity, cg.clusterConfiguration.ClusterCommandSendConnection, cg.clusterConfiguration.ClusterCommandReceiveConnection, cg.clusterConfiguration.ClusterCommandCaptureConnection)
	cg.clusterEventRoutine = NewClusterEventRoutine(cg.clusterConfiguration.Identity, cg.clusterConfiguration.ClusterEventSendConnection, cg.clusterConfiguration.ClusterEventReceiveConnection, cg.clusterConfiguration.ClusterEventCaptureConnection)
	cg.clusterCaptureWorkerRoutine = NewClusterCaptureWorkerRoutine(cg.clusterConfiguration.Identity, cg.clusterConfiguration.WorkerCaptureCommandReceiveConnection, cg.clusterConfiguration.WorkerCaptureEventReceiveConnection, cg.clusterConfiguration.Topics)

	go cg.clusterCommandRoutine.run()
	go cg.clusterEventRoutine.run()
	go cg.clusterCaptureWorkerRoutine.run()
}
