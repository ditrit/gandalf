package gandalfgo

import (
	"flag"
	"gandalf-go/aggregator"
	"gandalf-go/cluster"
	"gandalf-go/connector"
	"gandalf-go/worker"
)

func main() {
	mod := flag.String("mod", "", "cluster, aggregator, connector, worker")
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
	}

}
