package mains

import (
	"gandalf-go/cluster"
)

func main() {
	clusterGandalf := cluster.NewClusterGandalf("/home/orness/go/src/gandalf-go/tset/cluster/cluster1.json")
	go clusterGandalf.run()
}
