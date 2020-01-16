package main

import (
	"flag"
	"fmt"
	"gandalf-go/aggregator"
	"gandalf-go/cluster"
	"gandalf-go/connector"
	"gandalf-go/tset"
	"gandalf-go/worker"
	"time"

	"github.com/pebbe/zmq4"
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
		fmt.Println("Cluster " + config)
	case "aggregator":
		aggregatorGandalf := aggregator.NewAggregatorGandalf(config)
		aggregatorGandalf.Run()
		fmt.Println("Aggregator " + config)
	case "connector":
		connectorGandalf := connector.NewConnectorGandalf(config)
		connectorGandalf.Run()
		fmt.Println("Connector " + config)
	case "worker":
		workerGandalf := worker.NewWorkerGandalf(config)
		workerGandalf.Run()
		fmt.Println("Worker " + config)
	case "workerTestSend":
		tset.NewWorkerSender(config).Run()
		//toto.WorkerGandalf.ClientGandalf.SendCommand("toto", "100000000000000", "toto", "toto", "toto", "toto", "toto")
		fmt.Println("BOOP")
		//toto.WorkerGandalf.ClientGandalf.SendEvent("toto", "100", "toto", "toto", "toto")
		//time.Sleep(time.Second * 5)

		fmt.Println("WorkerSend " + config)

		//clientT := sender.NewSenderGandalf("toto", "tcp://127.0.0.1:9241", "tcp://127.0.0.1:9251")
		//clientT.SenderCommandRoutine.SendCommand("context", "timeout", "uuid", "connectorType", "commandType", "command", "payload")
		//time.Sleep(time.Second * 5)
		//clientT.SenderEventRoutine.SendEvent("topic", "timeout", "uuid", "event", "payload")
	case "workerTestReceive":
		tset.NewWorkerReceiver(config).Run()
		/* commandsRoutine := make(map[string][]routine.CommandRoutine)
		command := new(function.FunctionTest)
		fmt.Println("BEFORE")
		fmt.Println(commandsRoutine["command"])
		commandsRoutine["command"] = append(commandsRoutine["command"], command)
		fmt.Println("AFTER")
		fmt.Println(commandsRoutine["command"])

		eventsRoutine := make(map[string][]routine.EventRoutine)
		fmt.Println("BEFORE")
		fmt.Println(eventsRoutine["command"])
		event := new(function.FunctionTest)
		eventsRoutine["command"] = append(eventsRoutine["command"], event)
		fmt.Println("AFTER")
		fmt.Println(eventsRoutine["command"])

		receiveT := worker.NewWorkerGandalfRoutine(config, commandsRoutine, eventsRoutine)
		receiveT.Run() */
	case "Pub":
		//  Prepare our publisher
		publisher, _ := zmq4.NewSocket(zmq4.PUB)
		defer publisher.Close()
		publisher.Connect("tcp://localhost:5563")
		time.Sleep(time.Second)

		for {
			//  Write two messages, each with an envelope and content
			publisher.SendBytes([]byte("A"), zmq4.SNDMORE)
			publisher.SendBytes([]byte("We don't want to see this"), 0)
			publisher.SendBytes([]byte("B"), zmq4.SNDMORE)
			publisher.SendBytes([]byte("We would like to see this"), 0)

			contents, _ := publisher.RecvMessageBytes(0)
			fmt.Println(contents)
			time.Sleep(time.Second)
		}

	case "Sub":
		//  Prepare our subscriber
		subscriber, _ := zmq4.NewSocket(zmq4.SUB)
		defer subscriber.Close()
		subscriber.Bind("tcp://*:5563")
		subscriber.SetSubscribe("A")

		time.Sleep(time.Second)

		//subscriber.SendBytes([]byte{0x01}, 0)
		for {
			//  Read envelope with address
			//address, _ := subscriber.RecvBytes(0)
			//  Read message contents
			contents, _ := subscriber.RecvMessageBytes(0)
			fmt.Println(contents)
		}
	}
}
