//Package grpc :
package grpc

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"time"

	"github.com/ditrit/gandalf/core/connector/utils"

	cmodels "github.com/ditrit/gandalf/core/configuration/models"
	pb "github.com/ditrit/gandalf/libraries/gogrpc"

	"github.com/ditrit/gandalf/core/models"

	sn "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

var grpcSendIndex = 0

// ConnectorGrpc : ConnectorGrpc struct.
type ConnectorGrpc struct {
	pb.UnimplementedConnectorServer
	pb.UnimplementedConnectorCommandServer
	pb.UnimplementedConnectorEventServer
	GrpcConnection string
	Shoset         sn.Shoset
	//MapWorkerIterators map[string][]*msg.Iterator
	MapIterators      map[string]*msg.Iterator
	MapCommandChannel map[string]chan msg.Message
	EventChannel      chan msg.Message
	ValidationChannel chan msg.Message
	timeoutMax        int64
}

// NewConnectorGrpc : ConnectorGrpc constructor.
func NewConnectorGrpc(grpcConnection string, timeoutMax int64, shoset *sn.Shoset) (connectorGrpc ConnectorGrpc, err error) {
	connectorGrpc.Shoset = *shoset
	connectorGrpc.GrpcConnection = grpcConnection
	connectorGrpc.timeoutMax = timeoutMax
	//connectorGrpc.MapWorkerIterators = make(map[string][]*msg.Iterator)
	connectorGrpc.MapIterators = make(map[string]*msg.Iterator)
	connectorGrpc.MapCommandChannel = make(map[string]chan msg.Message)
	connectorGrpc.EventChannel = make(chan msg.Message)
	connectorGrpc.ValidationChannel = make(chan msg.Message)

	return
}

// StartGrpcServer : ConnectorGrpc start.
func (r ConnectorGrpc) StartGrpcServer() {

	if err := os.RemoveAll(r.GrpcConnection); err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("unix", r.GrpcConnection)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	connectorGrpcServer := grpc.NewServer()

	pb.RegisterConnectorServer(connectorGrpcServer, &r)
	pb.RegisterConnectorCommandServer(connectorGrpcServer, &r)
	pb.RegisterConnectorEventServer(connectorGrpcServer, &r)
	connectorGrpcServer.Serve(lis)
}

//SendCommandList : Connector send command list function.
func (r ConnectorGrpc) SendCommandList(ctx context.Context, in *pb.CommandList) (validate *pb.Validate, err error) {
	log.Println("Handle send command list")
	validation := false

	/* 	mapVersionConnectorCommands := r.Shoset.Context["mapVersionConnectorCommands"].(map[int8][]string)
	   	if mapVersionConnectorCommands != nil {
	   		mapVersionConnectorCommands[int8(in.GetMajor())] = append(mapVersionConnectorCommands[int8(in.GetMajor())], in.GetCommands()...)
	   	} */

	mapPivots := r.Shoset.Context["mapPivots"].(map[string][]*models.Pivot)
	mapProductConnector := r.Shoset.Context["mapProductConnector"].(map[string][]*models.ProductConnector)

	if mapPivots != nil && mapProductConnector != nil {
		configurationConnector := r.Shoset.Context["configuration"].(*cmodels.ConfigurationConnector)
		connectorType := configurationConnector.GetConnectorType()
		product := configurationConnector.GetProduct()
		pivot := utils.GetPivotByVersion(int8(in.GetMajor()), int8(in.GetMinor()), mapPivots[connectorType])
		productConnector := utils.GetConnectorProductByVersion(int8(in.GetMajor()), int8(in.GetMinor()), mapProductConnector[product])
		if pivot != nil && productConnector != nil {
			var configCommands []string
			for _, pivotCommand := range pivot.CommandTypes {
				configCommands = append(configCommands, pivotCommand.Name)
			}
			for _, productConnectorCommand := range productConnector.CommandTypes {
				configCommands = append(configCommands, productConnectorCommand.Name)
			}

			result := true
			for _, command := range configCommands {
				r, _ := regexp.Compile("ADMIN_*")
				if !r.MatchString(command) {
					currentResult := false
					for _, icommand := range in.GetCommands() {
						if command == icommand {
							currentResult = true
						}
					}
					result = result && currentResult
				}

			}
			validation = result

		} else {
			log.Printf("Can't get connector configuration with connector type %s, and version %s", connectorType, int8(in.GetMajor()))
		}

	} else {
		log.Printf("Connectors configuration not found")
	}

	activeWorkers := r.Shoset.Context["mapActiveWorkers"].(map[models.Version]bool)
	activeWorkers[models.Version{Major: int8(in.GetMajor()), Minor: int8(in.GetMinor())}] = validation
	r.Shoset.Context["mapActiveWorkers"] = activeWorkers

	return &pb.Validate{Valid: validation}, nil
}

//SendStop : Connector send stop.
func (r ConnectorGrpc) SendStop(ctx context.Context, in *pb.Stop) (validate *pb.Validate, err error) {
	log.Println("Handle send command list")

	activeWorkers := r.Shoset.Context["mapActiveWorkers"].(map[models.Version]bool)
	/* 	for activeWorkers[models.Version{Major: int8(in.GetMajor()), Minor: int8(in.GetMinor())}] {
		time.Sleep(5 * time.Second)
	} */
	return &pb.Validate{Valid: activeWorkers[models.Version{Major: int8(in.GetMajor()), Minor: int8(in.GetMinor())}]}, nil

}

//SendCommandMessage : Connector send command function.
func (r ConnectorGrpc) SendCommandMessage(ctx context.Context, in *pb.CommandMessage) (commandMessageUUID *pb.CommandMessageUUID, err error) {
	log.Println("Handle send command")

	cmd := pb.CommandFromGrpc(in)
	//connectorType := r.Shoset.Context["connectorType"].(string)

	validate := false
	mapPivots := r.Shoset.Context["mapPivots"].(map[string][]*models.Pivot)
	if mapPivots != nil {

		if cmd.GetContext()["isAdmin"].(bool) {
			connectorType := "Admin"
			if connectorType != "" {
				var pivot *models.Pivot
				if listPivot, ok := mapPivots[connectorType]; ok {

					pivot = utils.GetPivotByVersion(1, 0, listPivot)

					//connectorTypeConfig := utils.GetConnectorTypeConfigByVersion(int64(cmd.GetMajor()), listConnectorTypeConfig)
					if pivot != nil {
						commandType := utils.GetConnectorCommandType(cmd.GetCommand(), pivot.CommandTypes)
						if commandType.Name != "" {
							validate = utils.ValidatePayload(cmd.GetPayload(), commandType.Schema)
						} else {
							log.Println("Connector pivot command type not found")
						}
					} else {
						log.Println("Connector pivot by version not found")
					}
				} else {
					log.Printf("Connector pivot by type %s not found \n", connectorType)
				}
			} else {
				log.Println("Connector type not found")
			}
		} else {
			configurationConnector := r.Shoset.Context["configuration"].(*cmodels.ConfigurationConnector)
			connectorType := configurationConnector.GetConnectorType()
			if connectorType != "" {
				var pivot *models.Pivot
				if listPivot, ok := mapPivots[connectorType]; ok {
					if cmd.Major == 0 {
						versions := r.Shoset.Context["versions"].([]models.Version)
						if versions != nil {
							maxVersion := utils.GetMaxVersion(versions)
							cmd.Major = maxVersion.Major
							//connectorTypeConfig := utils.GetConnectorTypeConfigByVersion(int64(cmd.GetMajor()), listConnectorTypeConfig)
							pivot = utils.GetPivotByVersion(maxVersion.Major, maxVersion.Minor, listPivot)

						} else {
							log.Println("Versions not found")
						}
					} else {
						pivot = utils.GetPivotByVersion(cmd.Major, cmd.Minor, listPivot)
					}

					//connectorTypeConfig := utils.GetConnectorTypeConfigByVersion(int64(cmd.GetMajor()), listConnectorTypeConfig)
					if pivot != nil {
						commandType := utils.GetConnectorCommandType(cmd.GetCommand(), pivot.CommandTypes)
						if commandType.Name != "" {
							validate = utils.ValidatePayload(cmd.GetPayload(), commandType.Schema)
						} else {
							mapProductConnector := r.Shoset.Context["mapProductConnector"].(map[string][]*models.ProductConnector)
							if mapProductConnector != nil {
								product := configurationConnector.GetProduct()
								if product != "" {
									var productConnector *models.ProductConnector
									if listProductConnector, ok := mapProductConnector[product]; ok {
										if cmd.Major == 0 {
											versions := r.Shoset.Context["versions"].([]models.Version)
											if versions != nil {
												maxVersion := utils.GetMaxVersion(versions)
												cmd.Major = maxVersion.Major
												//connectorTypeConfig := utils.GetConnectorTypeConfigByVersion(int64(cmd.GetMajor()), listConnectorTypeConfig)
												productConnector = utils.GetConnectorProductByVersion(maxVersion.Major, maxVersion.Minor, listProductConnector)

											} else {
												log.Println("Versions not found")
											}
										} else {
											productConnector = utils.GetConnectorProductByVersion(cmd.Major, cmd.Minor, listProductConnector)
										}
										if productConnector != nil {
											commandType := utils.GetConnectorCommandType(cmd.GetCommand(), productConnector.CommandTypes)
											if commandType.Name != "" {
												validate = utils.ValidatePayload(cmd.GetPayload(), commandType.Schema)
											} else {
												log.Println("Command type not found")
											}
										} else {
											log.Println("Product Connector by version not found")
										}
									} else {
										log.Printf("Product connector by product %s not found \n", product)
									}
								} else {
									log.Println("Product not found")
								}
							} else {
								log.Println("Map product connector not found")
							}
						}
					} else {
						log.Println("Pivot by version not found")
					}
				} else {
					log.Printf("Pivot by type %s not found \n", connectorType)
				}
			} else {
				log.Println("Connector type not found")
			}
		}
	} else {
		log.Println("Map pivot not found")

	}

	if validate {
		cmd.Tenant = r.Shoset.Context["tenant"].(string)
		shosets := sn.GetByType(r.Shoset.ConnsByAddr, "a")

		if len(shosets) != 0 {
			if cmd.GetTimeout() > r.timeoutMax {
				cmd.Timeout = r.timeoutMax
			}

			iteratorMessage, _ := r.CreateIteratorEvent(ctx, new(pb.Empty))
			iterator := r.MapIterators[iteratorMessage.GetId()]

			go r.runIteratorEvent(cmd.GetCommand(), "ON_GOING", cmd.GetUUID(), iterator, r.ValidationChannel)

			notSend := true
			for notSend {
				index := getGrpcSendIndex(shosets)
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
			log.Println("can't find aggregators to send")
			err = errors.New("can't find aggregators to send")
		}
	} else {
		log.Println("wrong payload command")
		err = errors.New("wrong payload command")
	}

	return commandMessageUUID, err
}

//WaitCommandMessage : Connector wait command function.
func (r ConnectorGrpc) WaitCommandMessage(ctx context.Context, in *pb.CommandMessageWait) (commandMessage *pb.CommandMessage, err error) {
	log.Println("Handle wait command")

	iterator := r.MapIterators[in.GetIteratorId()]

	go r.runIteratorCommand(in.GetValue(), int8(in.GetMajor()), iterator, r.MapCommandChannel[in.GetIteratorId()])

	messageChannel := <-r.MapCommandChannel[in.GetIteratorId()]
	commandMessage = pb.CommandToGrpc(messageChannel.(msg.Command))

	return
}

//SendEventMessage : Connector send event function.
func (r ConnectorGrpc) SendEventMessage(ctx context.Context, in *pb.EventMessage) (empty *pb.Empty, err error) {
	log.Println("Handle send event")
	evt := pb.EventFromGrpc(in)
	evt.Tenant = r.Shoset.Context["tenant"].(string)
	thisOne := r.Shoset.GetBindAddr()
	validate := false

	if evt.GetReferenceUUID() == "" {
		mapPivots := r.Shoset.Context["mapPivots"].(map[string][]*models.Pivot)
		if mapPivots != nil {

			configurationConnector := r.Shoset.Context["configuration"].(*cmodels.ConfigurationConnector)
			connectorType := configurationConnector.GetConnectorType()
			if connectorType != "" {
				var pivot *models.Pivot
				if listPivot, ok := mapPivots[connectorType]; ok {
					if evt.Major == 0 {
						versions := r.Shoset.Context["versions"].([]models.Version)
						if versions != nil {
							maxVersion := utils.GetMaxVersion(versions)
							evt.Major = maxVersion.Major
							//connectorTypeConfig := utils.GetConnectorTypeConfigByVersion(int64(cmd.GetMajor()), listConnectorTypeConfig)
							pivot = utils.GetPivotByVersion(maxVersion.Major, maxVersion.Minor, listPivot)

						} else {
							log.Println("Versions not found")
						}
					} else {
						pivot = utils.GetPivotByVersion(evt.Major, evt.Minor, listPivot)
					}

					//connectorTypeConfig := utils.GetConnectorTypeConfigByVersion(int64(cmd.GetMajor()), listConnectorTypeConfig)
					if pivot != nil {
						commandType := utils.GetConnectorEventType(evt.GetEvent(), pivot.EventTypes)
						if commandType.Name != "" {
							validate = utils.ValidatePayload(evt.GetPayload(), commandType.Schema)
						} else {
							mapProductConnector := r.Shoset.Context["mapProductConnector"].(map[string][]*models.ProductConnector)
							if mapProductConnector != nil {
								product := configurationConnector.GetProduct()
								if product != "" {
									var productConnector *models.ProductConnector
									if listProductConnector, ok := mapProductConnector[product]; ok {
										if evt.Major == 0 {
											versions := r.Shoset.Context["versions"].([]models.Version)
											if versions != nil {
												maxVersion := utils.GetMaxVersion(versions)
												evt.Major = maxVersion.Major
												//connectorTypeConfig := utils.GetConnectorTypeConfigByVersion(int64(cmd.GetMajor()), listConnectorTypeConfig)
												productConnector = utils.GetConnectorProductByVersion(maxVersion.Major, maxVersion.Minor, listProductConnector)

											} else {
												log.Println("Versions not found")
											}
										} else {
											productConnector = utils.GetConnectorProductByVersion(evt.Major, evt.Minor, listProductConnector)
										}
										if productConnector != nil {
											commandType := utils.GetConnectorEventType(evt.GetEvent(), productConnector.EventTypes)
											if commandType.Name != "" {
												validate = utils.ValidatePayload(evt.GetPayload(), commandType.Schema)
											} else {
												log.Println("Event type not found")
											}
										} else {
											log.Println("Product Connector by version not found")
										}
									} else {
										log.Printf("Product connector by product %s not found \n", product)
									}
								} else {
									log.Println("Product not found")
								}
							} else {
								log.Println("Map product connector not found")
							}
						}
					} else {
						log.Println("Pivot by version not found")
					}
				} else {
					log.Printf("Pivot by type %s not found \n", connectorType)
				}
			} else {
				log.Println("Connector type not found")
			}
		} else {
			log.Println("Map pivot not found")

		}
	}

	if validate {
		r.Shoset.ConnsByAddr.Iterate(
			func(key string, val *sn.ShosetConn) {
				if key != r.Shoset.GetBindAddr() && key != thisOne && val.ShosetType == "a" {
					val.SendMessage(evt)
					log.Printf("%s : send event %s to %s\n", thisOne, evt.GetEvent(), val)
				}
			},
		)
	} else {
		log.Println("wrong payload command")
		err = errors.New("wrong payload command")
	}

	return &pb.Empty{}, err
}

//WaitEventMessage : Connector wait event function.
func (r ConnectorGrpc) WaitEventMessage(ctx context.Context, in *pb.EventMessageWait) (messageEvent *pb.EventMessage, err error) {
	log.Println("Handle wait event")

	iterator := r.MapIterators[in.GetIteratorId()]

	go r.runIteratorEvent(in.GetTopic(), in.GetEvent(), in.GetReferenceUUID(), iterator, r.EventChannel)

	messageChannel := <-r.EventChannel
	messageEvent = pb.EventToGrpc(messageChannel.(msg.Event))

	return
}

//WaitTopicMessage : Connector wait event by topic function.
func (r ConnectorGrpc) WaitTopicMessage(ctx context.Context, in *pb.TopicMessageWait) (messageEvent *pb.EventMessage, err error) {
	log.Println("Handle wait event by topic")

	iterator := r.MapIterators[in.GetIteratorId()]

	go r.runIteratorTopic(in.GetTopic(), in.GetReferenceUUID(), iterator, r.EventChannel)

	messageChannel := <-r.EventChannel
	messageEvent = pb.EventToGrpc(messageChannel.(msg.Event))

	return
}

//TODO REFACTORING

//CreateIteratorCommand : Connector create command iterator function.
func (r ConnectorGrpc) CreateIteratorCommand(ctx context.Context, in *pb.Empty) (iteratorMessage *pb.IteratorMessage, err error) {
	log.Println("Handle create iterator command")

	iterator := msg.NewIterator(r.Shoset.Queue["cmd"])
	index := uuid.New()
	log.Printf("Create new iterator command: %s", index)

	//r.MapWorkerIterators[index.String()] = append(r.MapWorkerIterators[index.String()], iterator)
	r.MapIterators[index.String()] = iterator
	r.MapCommandChannel[index.String()] = make(chan msg.Message)
	iteratorMessage = &pb.IteratorMessage{Id: index.String()}

	return
}

//CreateIteratorEvent : Connector create event iterator function.
func (r ConnectorGrpc) CreateIteratorEvent(ctx context.Context, in *pb.Empty) (iteratorMessage *pb.IteratorMessage, err error) {
	log.Println("Handle create iterator event")
	fmt.Println("CREATE ITE EVENT")
	iterator := msg.NewIterator(r.Shoset.Queue["evt"])
	index := uuid.New()
	log.Printf("Create new iterator event: %s", index)

	//r.MapWorkerIterators[index.String()] = append(r.MapWorkerIterators[index.String()], iterator)
	r.MapIterators[index.String()] = iterator

	iteratorMessage = &pb.IteratorMessage{Id: index.String()}

	return
}

//TODO REVOIR
// runIterator : Iterator run function.
func (r ConnectorGrpc) runIteratorCommand(command string, major int8, iterator *msg.Iterator, channel chan msg.Message) {
	log.Printf("Run iterator command on command %s", command)

	for {

		//fmt.Println("messageIterator" + command)
		//iterator.PrintQueue()

		messageIterator := iterator.Get()

		if messageIterator != nil {
			message := (messageIterator.GetMessage()).(msg.Command)

			if command == message.GetCommand() {
				versionMajor := message.GetMajor()
				if major == 0 || (major != 0 && versionMajor == major) {
					log.Println("Get iterator command")
					log.Println(message)

					channel <- message

					break
				}
			}
		}
		time.Sleep(time.Duration(500) * time.Millisecond)
	}
	//delete(r.MapIterators, iteratorId)
}

// runIteratorEvent : Iterator run function.
func (r ConnectorGrpc) runIteratorEvent(topic, event, referenceUUID string, iterator *msg.Iterator, channel chan msg.Message) {
	log.Printf("Run iterator event on topic %s, event %s, ref %s", topic, event, referenceUUID)

	for {
		messageIterator := iterator.Get()
		//iterator.PrintQueue()

		if messageIterator != nil {
			message := (messageIterator.GetMessage()).(msg.Event)
			if topic == message.Topic {
				if event == message.Event {
					if referenceUUID != "" {
						if referenceUUID == message.GetReferenceUUID() {
							log.Println("Get iterator event by ref")
							log.Println(message)

							channel <- message

							break
						}
					} else {
						log.Println("Get iterator event")
						log.Println(message)

						channel <- message

						break
					}
				}
			}
		}
		time.Sleep(time.Duration(500) * time.Millisecond)
	}
	//delete(r.MapIterators, iteratorId)
}

// runIteratorTopic : Iterator run function.
func (r ConnectorGrpc) runIteratorTopic(topic, referenceUUID string, iterator *msg.Iterator, channel chan msg.Message) {
	log.Printf("Run iterator topic on topic %s ref %s", topic, referenceUUID)

	for {
		messageIterator := iterator.Get()

		if messageIterator != nil {
			message := (messageIterator.GetMessage()).(msg.Event)

			if topic == message.Topic {

				if referenceUUID != "" {
					if referenceUUID == message.GetReferenceUUID() {
						log.Println("Get iterator event by topic and ref")
						log.Println(message)

						channel <- message

						break
					}
				} else {
					log.Println("Get iterator event by topic")
					log.Println(message)

					channel <- message

					break
				}
			}

		}
		time.Sleep(time.Duration(500) * time.Millisecond)
	}
	//delete(r.MapIterators, iteratorId)
}

// getGrpcSendIndex : Connector getGrpcSendIndex function.
func getGrpcSendIndex(conns []*sn.ShosetConn) int {
	aux := grpcSendIndex
	grpcSendIndex++

	if grpcSendIndex >= len(conns) {
		grpcSendIndex = 0
	}

	return aux
}
