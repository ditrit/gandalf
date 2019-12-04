package cluster

type ClusterGandalf struct {
	clusterCommandRoutine       ClusterCommandRoutine
	clusterEventRoutine         ClusterEventRoutine
	clusterCaptureWorkerRoutine ClusterCaptureWorkerRoutine
}

func (cg ClusterGandalf) main() {
	//identity, workerCommandReceiveC2WConnection, workerEventReceiveC2WConnection string, topics *string
	//LOAD CONF
	cg.clusterCommandRoutine = ClusterCommandRoutine.new()
	cg.clusterEventRoutine = ClusterEventRoutine.new()
	cg.clusterCaptureWorkerRoutine = ClusterCaptureWorkerRoutine.new()

	go cg.clusterCommandRoutine.run()
	go cg.clusterEventRoutine.run()
	go cg.clusterCaptureWorkerRoutine.run()
}
