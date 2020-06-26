package goazure

import "connectors/goazure/workers"

func main() {
	done := make(chan bool)
	workerCompute := workers.NewWorkerCompute(identity, connections)
	go workerCompute.Run()
	<-done
}
