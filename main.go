package main

import (
	"flag"
	"fmt"
	"gandalf-go/aggregator"
	"gandalf-go/cluster"
	"gandalf-go/connector"
	"gandalf-go/message"
	"gandalf-go/worker"
	"gandalf-go/worker/routine"
	"log"
	"time"

	"gandalf-go/client/sender"
	"gandalf-go/tset/function"

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

		fmt.Println("WorkerSend " + config)

		clientT := sender.NewSenderGandalf("toto", "tcp://127.0.0.1:9141", "tcp://127.0.0.1:9151")
		//clientT.SenderCommandRoutine.SendCommandSyncTEST("context", "timeout", "uuid", "connectorType", "commandType", "command", "payload")
		time.Sleep(time.Second * 5)
		clientT.SenderEventRoutine.SendEvent("topic", "timeout", "uuid", "event", "payload")
	case "workerTestReceive":
		commandsRoutine := make(map[string][]routine.CommandRoutine)
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
		receiveT.Run()
	case "Pub":
		//  Prepare our publisher
		publisher, _ := zmq4.NewSocket(zmq4.XPUB)
		defer publisher.Close()
		publisher.Connect("tcp://127.0.0.1:9251")

		/* subscriber, _ := zmq4.NewSocket(zmq4.XSUB)
		defer subscriber.Close()
		subscriber.Connect("tcp://127.0.0.1:9250")
		subscriber.SendBytes([]byte{0x01}, 0) //SUBSCRIBE ALL
		*/
		time.Sleep(time.Second * 5)

		eventMessage := message.NewEventMessage("topic", "timeout", "uuid", "event", "payload")
		for {
			go eventMessage.SendEventWith(publisher)
			/* 		publisher.Send("A", zmq4.SNDMORE)
			   		publisher.Send("WTF! 1 ", 0) */

			time.Sleep(time.Second * 5)
		}

	case "Sub":
		//  Prepare our publisher
		subscriber, _ := zmq4.NewSocket(zmq4.XSUB)
		defer subscriber.Close()
		err := subscriber.Bind("tcp://*:9251")
		subscriber.SendBytes([]byte{0x01}, 0) //SUBSCRIBE ALL
		if err != nil {
			log.Fatal(err)
		}
		_, err = subscriber.SendBytes([]byte{0x01}, 0)
		toto, _ := subscriber.RecvBytes(0)
		tata, _ := subscriber.RecvBytes(0)
		fmt.Println(string(toto))
		fmt.Println(string(tata))
	}
}
