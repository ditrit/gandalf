package mains

import (
	"gandalf-go/connector"
)

func main() {
	connectorGandalf := connector.NewConnectorGandalf("/home/orness/go/src/gandalf-go/tset/connector/connector1.json")
	go connectorGandalf.run()
}
