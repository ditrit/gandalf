package main

import (
	"flag"
	"fmt"
	"gandalf-go/cluster"
	"gandalf-go/worker"
	"gandalf-go/aggregator"
	"gandalf-go/connector"

	"gandalf-go/client/sender"
)

func main() {

	var mode string
	var config string

	flag.StringVar(&mode, "m", "", "")
	flag.StringVar(&mode, "mode", "", "")
	flag.StringVar(&config, "c", "", "")
	flag.StringVar(&config, "config", "", "")
	flag.Parse()

	switch mode {
	case "cluster":
		clusterGandalf := cluster.NewClusterGandalf(config)
		clusterGandalf.Run()
		fmt.Print("%s", "Cluster " + config)
	case "aggregator":
		aggregatorGandalf := aggregator.NewAggregatorGandalf(config)
		fmt.Print("%s", "Running")
		aggregatorGandalf.Run()
		fmt.Print("%s", "Aggregator " + config)
	case "connector":
		connectorGandalf := connector.NewConnectorGandalf(config)
		connectorGandalf.Run()
		fmt.Print("%s", "Connector " + config)
	case "worker":
		workerGandalf := worker.NewWorkerGandalf(config)
		workerGandalf.Run()
		fmt.Print("%s", "Worker " + config) 
	case "workerTest":
	
		fmt.Print("%s", "Worker " + config)

		clientT := sender.NewSenderGandalf("toto", "tcp://127.0.0.1:9141", "127.0.0.1:9151")
		clientT.SenderCommandRoutine.SendCommandSync("context", "timeout", "uuid", "connectorType", "commandType", "command", "payload")
	}
}
