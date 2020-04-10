package connector

import (
	"context"
	"errors"
	pb "garcimore/grpc"
	"garcimore/utils"
	"log"
	"net"
	"shoset/msg"
	sn "shoset/net"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

var sendIndex = 0

type ConnectorGrpc struct {
	GrpcConnection string
	Shoset         sn.Shoset
	//MapWorkerIterators map[string][]*msg.Iterator
	MapIterators      map[string]*msg.Iterator
	CommandChannel    chan msg.Message
	EventChannel      chan msg.Message
	ValidationChannel chan msg.Message
	timeoutMax        int64
}

func NewConnectorGrpc(GrpcConnection string, timeoutMax int64, shoset *sn.Shoset) (connectorGrpc ConnectorGrpc, err error) {
	connectorGrpc.Shoset = *shoset
	connectorGrpc.GrpcConnection = GrpcConnection
	connectorGrpc.timeoutMax = timeoutMax
	//connectorGrpc.MapWorkerIterators = make(map[string][]*msg.Iterator)
	connectorGrpc.MapIterators = make(map[string]*msg.Iterator)
	connectorGrpc.CommandChannel = make(chan msg.Message)
	connectorGrpc.EventChannel = make(chan msg.Message)
	connectorGrpc.ValidationChannel = make(chan msg.Message)

	return
}

//GRPC
//startGrpcServer :
func (r ConnectorGrpc) startGrpcServer() {
	lis, err := net.Listen("tcp", r.GrpcConnection)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	connectorGrpcServer := grpc.NewServer()

	pb.RegisterConnectorCommandServer(connectorGrpcServer, &r)
	pb.RegisterConnectorEventServer(connectorGrpcServer, &r)
	connectorGrpcServer.Serve(lis)
}

//SendCommandMessage :
func (r ConnectorGrpc) SendCommandMessage(ctx context.Context, in *pb.CommandMessage) (commandMessageUUID *pb.CommandMessageUUID, err error) {
	cmd := pb.CommandFromGrpc(in)
	cmd.Tenant = r.Shoset.Context["tenant"].(string)
	shosets := utils.GetByType(r.Shoset.ConnsByAddr, "a")
	if len(shosets) != 0 {
		if cmd.GetTimeout() > r.timeoutMax {
			cmd.Timeout = r.timeoutMax
		}

		iteratorMessage, _ := r.CreateIteratorEvent(ctx, new(pb.Empty))
		iterator := r.MapIterators[iteratorMessage.GetId()]
		go r.runIterator(iteratorMessage.GetId(), cmd.GetUUID(), "validation", iterator, r.ValidationChannel)

		notSend := true
		for notSend {

			index := getSendIndex(shosets)
			shosets[index].SendMessage(cmd)
			log.Printf("%s : send command %s to %s\n", r.Shoset.GetBindAddr(), cmd.GetCommand(), shosets[index])

			timeoutSend := time.Duration((int(cmd.GetTimeout()) / len(shosets)))

			messageChannel := <-r.ValidationChannel
			log.Printf("%s : receive validation event for command %s to %s\n", r.Shoset.GetBindAddr(), cmd.GetCommand(), shosets[index])

			if messageChannel != nil {
				notSend = false
				break
			}
			time.Sleep(timeoutSend)
		}
		if notSend {
			return nil, nil
		}
		commandMessageUUID = &pb.CommandMessageUUID{UUID: cmd.UUID}
	} else {
		log.Println("Can't find aggregators to send")
		err = errors.New("Can't find aggregators to send")
	}
	return commandMessageUUID, nil
}

//WaitCommandMessage :
func (r ConnectorGrpc) WaitCommandMessage(ctx context.Context, in *pb.CommandMessageWait) (commandMessage *pb.CommandMessage, err error) {

	iterator := r.MapIterators[in.GetIteratorId()]

	go r.runIterator(in.GetIteratorId(), in.GetValue(), "cmd", iterator, r.CommandChannel)
	messageChannel := <-r.CommandChannel
	commandMessage = pb.CommandToGrpc(messageChannel.(msg.Command))

	return
}

//SendEventMessage :
func (r ConnectorGrpc) SendEventMessage(ctx context.Context, in *pb.EventMessage) (*pb.Empty, error) {
	evt := pb.EventFromGrpc(in)
	evt.Tenant = r.Shoset.Context["tenant"].(string)
	thisOne := r.Shoset.GetBindAddr()

	r.Shoset.ConnsByAddr.Iterate(
		func(key string, val *sn.ShosetConn) {
			if key != r.Shoset.GetBindAddr() && key != thisOne && val.ShosetType == "a" {
				val.SendMessage(evt)
				log.Printf("%s : send event %s to %s\n", thisOne, evt.GetEvent(), val)
			}
		},
	)

	return &pb.Empty{}, nil
}

//WaitEventMessage :
func (r ConnectorGrpc) WaitEventMessage(ctx context.Context, in *pb.EventMessageWait) (messageEvent *pb.EventMessage, err error) {

	iterator := r.MapIterators[in.GetIteratorId()]

	go r.runIterator(in.GetIteratorId(), in.GetEvent(), "evt", iterator, r.EventChannel)

	messageChannel := <-r.EventChannel
	messageEvent = pb.EventToGrpc(messageChannel.(msg.Event))

	return
}

//WaitEventMessage :
func (r ConnectorGrpc) WaitTopicMessage(ctx context.Context, in *pb.TopicMessageWait) (messageEvent *pb.EventMessage, err error) {

	iterator := r.MapIterators[in.GetIteratorId()]

	go r.runIterator(in.GetIteratorId(), in.GetTopic(), "topic", iterator, r.EventChannel)

	messageChannel := <-r.EventChannel
	messageEvent = pb.EventToGrpc(messageChannel.(msg.Event))

	return
}

//TODO REFACTORING
//CreateIteratorCommand :
func (r ConnectorGrpc) CreateIteratorCommand(ctx context.Context, in *pb.Empty) (iteratorMessage *pb.IteratorMessage, err error) {
	iterator := msg.NewIterator(r.Shoset.Queue["cmd"])
	index := uuid.New()
	//r.MapWorkerIterators[index.String()] = append(r.MapWorkerIterators[index.String()], iterator)
	r.MapIterators[index.String()] = iterator

	iteratorMessage = &pb.IteratorMessage{Id: index.String()}

	return
}

//CreateIteratorEvent :
func (r ConnectorGrpc) CreateIteratorEvent(ctx context.Context, in *pb.Empty) (iteratorMessage *pb.IteratorMessage, err error) {
	iterator := msg.NewIterator(r.Shoset.Queue["evt"])
	index := uuid.New()
	//r.MapWorkerIterators[index.String()] = append(r.MapWorkerIterators[index.String()], iterator)
	r.MapIterators[index.String()] = iterator

	iteratorMessage = &pb.IteratorMessage{Id: index.String()}
	return
}

func (r ConnectorGrpc) runIterator(iteratorId, value, msgtype string, iterator *msg.Iterator, channel chan msg.Message) {

	for {
		//fmt.Println("ITERATOR QUEUE")
		//iterator.PrintQueue()
		messageIterator := iterator.Get()

		if messageIterator != nil {
			if msgtype == "cmd" {

				message := (messageIterator.GetMessage()).(msg.Command)

				if value == message.GetCommand() {
					channel <- message

					break
				}
			} else if msgtype == "evt" {
				message := (messageIterator.GetMessage()).(msg.Event)

				if value == message.Event {
					channel <- message

					break
				}
			} else if msgtype == "topic" {
				message := (messageIterator.GetMessage()).(msg.Event)

				if value == message.Topic {
					channel <- message

					break
				}
			} else if msgtype == "validation" {
				message := (messageIterator.GetMessage()).(msg.Event)

				if value == message.ReferencesUUID && message.Event == "TAKEN" {
					channel <- message
					break
				}
			}
		}
		time.Sleep(time.Duration(2000) * time.Millisecond)
	}
	//delete(r.MapIterators, iteratorId)
}

func getSendIndex(conns []*sn.ShosetConn) int {
	aux := sendIndex
	sendIndex++
	if sendIndex >= len(conns) {
		sendIndex = 0
	}
	return aux
}
