package gandalfgo

import (
	"gandalfgo/aggregator"
	"gandalfgo/client"
	"gandalfgo/cluster"
	"gandalfgo/connector"
	"gandalfgo/worker"
)

func main() {
	mod := flag.String("mod", "", "cluster, aggregator, connector, worker")
	path := flag.String("path")

	switch mod {
	case "cluster":
		ClusterGandalf.New(path)
	case "aggregator":
		AggregatorGandalf.New(path)
	case "connector":
		ConnectorGandalf.New(path)
	case "worker":
		WorkerGandalf.New(path)
	default:
		//ERROR
	}

}
