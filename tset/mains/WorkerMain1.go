package mains

import (
	"gandalf-go/worker"
)

func main() {
	workerGandalf := worker.NewWorkerGandalf("/home/orness/go/src/gandalf-go/tset/worker/worker1.json")
	go workerGandalf.run()
}
