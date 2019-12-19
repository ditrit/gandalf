package mains

import (
	"gandalf-go/aggregator"
)

func main() {
	aggregatorGandalf := aggregator.NewAggregatorGandalf("/home/orness/go/src/gandalf-go/tset/aggregator/aggregator1.json")
	go aggregatorGandalf.run()
}
