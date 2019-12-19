package main

import (
	"flag"
	"fmt"
	"gandalf-go/cluster"
	"gandalf-go/worker"
)

func main() {
	/* 	mod := flag.String("mod", "", "cluster, aggregator, connector, worker")
	   	path := flag.String("path")

	   	switch mod {
	   	case "cluster":
	   		cluster.ClusterGandalf.New(path)
	   	case "aggregator":
	   		aggregator.AggregatorGandalf.New(path)
	   	case "connector":
	   		connector.ConnectorGandalf.New(path)
	   	case "worker":
	   		worker.WorkerGandalf.New(path)
	   	default:
	   		//ERROR
		   }*/

	var isReceive bool

	flag.BoolVar(&isReceive, "r", false, "")
	flag.BoolVar(&isReceive, "recv", false, "")
	flag.Parse()

	if isReceive {
		clusterGandalf1 := cluster.NewClusterGandalf("/home/orness/go/src/gandalf-go/tset/cluster/cluster1.json")
		go clusterGandalf1.Run()
		fmt.Print("%s", "Cluster1")
		/* 	clusterGandalf2 := cluster.NewClusterGandalf("/home/orness/go/src/gandalf-go/tset/cluster/cluster2.json")
		go clusterGandalf2.Run()
		fmt.Print("%s", "Cluster2")
		aggregatorGandalf1 := aggregator.NewAggregatorGandalf("/home/orness/go/src/gandalf-go/tset/aggregator/aggregator1.json")
		go aggregatorGandalf1.Run()
		fmt.Print("%s", "Aggregator1")
		aggregatorGandalf2 := aggregator.NewAggregatorGandalf("/home/orness/go/src/gandalf-go/tset/aggregator/aggregator2.json")
		go aggregatorGandalf2.Run()
		fmt.Print("%s", "Aggregator2")
		connectorGandalf1 := connector.NewConnectorGandalf("/home/orness/go/src/gandalf-go/tset/connector/connector1.json")
		go connectorGandalf1.Run()
		fmt.Print("%s", "Connector1")
		connectorGandalf2 := connector.NewConnectorGandalf("/home/orness/go/src/gandalf-go/tset/connector/connector2.json")
		go connectorGandalf2.Run()
		fmt.Print("%s", "Connector2")
		workerGandalf := worker.NewWorkerGandalf("/home/orness/go/src/gandalf-go/tset/worker/worker1.json")
		go workerGandalf.Run()
		fmt.Print("%s", "Worker1") */
	} else {
		workerGandalf := worker.NewWorkerGandalf("/home/orness/go/src/gandalf-go/tset/worker/worker2.json")
		go workerGandalf.Run()
		fmt.Print("%s", "Worker2")
		workerGandalf.ClientGandalf.SenderGandalf.SenderCommandRoutine.SendCommandSync("context", "timeout", "uuid", "connectorType", "commandType", "command", "payload")
	}
}
