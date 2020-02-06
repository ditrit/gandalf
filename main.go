//Package main :
//File main.go
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"gandalf-go/aggregator"
	"gandalf-go/cluster"
	"gandalf-go/connector"
	"gandalf-go/database"
	"gandalf-go/tset"
	"gandalf-go/worker"

	"github.com/pebbe/zmq4"
)

//nolint: funlen, gocyclo
func main() {
	var (
		mode   string
		config string
	)

	commandLine := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	flag.StringVar(&mode, "m", "", "")
	flag.StringVar(&mode, "mode", "", "")
	flag.StringVar(&config, "c", "", "")
	flag.StringVar(&config, "config", "", "")
	flag.Parse()

	fmt.Println("Running Gandalf with:")
	fmt.Println("  Mode : " + mode)
	fmt.Println("  Config : " + config)

	switch mode {
	case "cluster":
		clusterGandalf := cluster.NewClusterGandalf(config)
		clusterGandalf.Run()
	case "aggregator":
		aggregatorGandalf := aggregator.NewAggregatorGandalf(config)
		aggregatorGandalf.Run()
	case "connector":
		connectorGandalf := connector.NewConnectorGandalf(config)
		connectorGandalf.Run()
	case "worker":
		workerGandalf := worker.NewWorkerGandalf(config)
		workerGandalf.Run()
	case "database":
		databaseClusterGandalf := database.NewDatabaseClusterGandalf(config)
		databaseClusterGandalf.Run()
	case "workerTestSend":
		tset.NewWorkerSender(config).Run()
		//toto.WorkerGandalf.ClientGandalf.SendCommand("toto", "100000000000000", "toto", "toto", "toto", "toto", "toto")
		fmt.Println("BOOP")
		//toto.WorkerGandalf.ClientGandalf.SendEvent("toto", "100", "toto", "toto", "toto")
		//time.Sleep(time.Second * 5)

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

		_ = publisher.Connect("tcp://localhost:5563")

		time.Sleep(time.Second)

		for {
			//  Write two messages, each with an envelope and content
			_, _ = publisher.SendBytes([]byte("A"), zmq4.SNDMORE)
			_, _ = publisher.SendBytes([]byte("We don't want to see this"), 0)
			_, _ = publisher.SendBytes([]byte("B"), zmq4.SNDMORE)
			_, _ = publisher.SendBytes([]byte("We would like to see this"), 0)

			contents, _ := publisher.RecvMessageBytes(0)
			fmt.Println(contents)
			time.Sleep(time.Second)
		}

	case "Sub":
		//  Prepare our subscriber
		subscriber, _ := zmq4.NewSocket(zmq4.SUB)
		defer subscriber.Close()
		_ = subscriber.Bind("tcp://*:5563")
		_ = subscriber.SetSubscribe("A")

		time.Sleep(time.Second)

		//subscriber.SendBytes([]byte{0x01}, 0)
		for {
			//  Read envelope with address
			//address, _ := subscriber.RecvBytes(0)
			//  Read message contents
			contents, _ := subscriber.RecvMessageBytes(0)
			fmt.Println(contents)
		}

	default:
		fmt.Fprintf(commandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
}
