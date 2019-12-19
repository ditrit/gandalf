package mains

import (
	"gandalf-go/aggregator"
)

func main() {
	aggregatorGandalf := aggregator.NewAggregatorGandalf("/home/orness/go/src/gandalf-go/tset/aggregator/aggregator2.json")
	go aggregatorGandalf.run()
}
