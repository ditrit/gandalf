package main

import (
	"core/aggregator"
	"core/cluster"
	"core/connector"
	"core/database"
	"core/demo"
	"flag"
	"fmt"
	"os"
	"shoset/net"
	"strconv"
)

func main() {

	var (
		debug  bool
		config string
	)
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("  gandalf mode command [options]")
		fmt.Printf("  mode : cluster, aggragor, connector, agent\n")
		fmt.Printf("    mode : cluster\n")
		fmt.Printf("	  cluster command : init, join\n")
		fmt.Printf("        arguments:\n")
		fmt.Printf("  	      logical name	  \n")
		fmt.Printf("  		  bind address    \n")
		fmt.Printf("  		  join address     \n")
		fmt.Printf("    mode : aggregator\n")
		fmt.Printf("        arguments:\n")
		fmt.Printf("  	      logical name	  \n")
		fmt.Printf("  		  bind address    \n")
		fmt.Printf("  		  link address     \n")
		fmt.Printf("    mode : connector\n")
		fmt.Printf("        arguments:\n")
		fmt.Printf("  	      logical name	  \n")
		fmt.Printf("  		  bind address    \n")
		fmt.Printf("  		  bind grpc address    \n")
		fmt.Printf("  		  link address     \n")
		fmt.Printf("  		  timeout max     \n")
	}

	flag.BoolVar(&debug, "d", false, "")
	flag.BoolVar(&debug, "debug", false, "")
	flag.StringVar(&config, "c", "", "")
	flag.StringVar(&config, "config", "", "")
	flag.Parse()
	args := flag.Args()

	fmt.Println("debug")
	fmt.Println(debug)

	if len(args) >= 1 {
		mode := args[0]
		switch mode {
		case "cluster":
			if len(args) >= 2 {
				command := args[1]

				switch command {
				case "init":
					if len(args) >= 4 {
						done := make(chan bool)

						LogicalName := args[2]
						BindAdd := args[3]
						//CREATE CLUSTER
						fmt.Println("Running Gandalf with:")
						fmt.Println("  Mode : " + mode)
						fmt.Println("  Logical Name : " + LogicalName)
						fmt.Println("  Bind Address : " + BindAdd)
						fmt.Println("  Config : " + config)

						cluster.ClusterMemberInit(LogicalName, BindAdd)

						add, _ := net.DeltaAddress(BindAdd, 1000)
						go database.DatabaseMemberInit(add, 1)
						database.List([]string{add})

						<-done
					} else {
						flag.Usage()
					}
					break
				case "join": //join
					if len(args) >= 5 {
						done := make(chan bool)

						LogicalName := args[2]
						BindAdd := args[3]
						JoinAdd := args[4]
						//CREATE CLUSTER
						fmt.Println("Running Gandalf with:")
						fmt.Println("  Mode : " + mode)
						fmt.Println("  Logical Name : " + LogicalName)
						fmt.Println("  Bind Address : " + BindAdd)
						fmt.Println("  Join Address : " + JoinAdd)
						fmt.Println("  Config : " + config)

						member := cluster.ClusterMemberJoin(LogicalName, BindAdd, JoinAdd)

						add, _ := net.DeltaAddress(BindAdd, 1000)
						id := len(*member.Store)

						go database.DatabaseMemberInit(add, id)

						err := database.AddNodesToLeader(id, add, *member.Store)
						fmt.Println(err)
						database.List([]string{add})

						<-done

					} else {
						flag.Usage()
					}
					break
				default:
					break
				}
			} else {
				flag.Usage()
			}
		case "aggregator":
			if len(args) >= 5 {
				done := make(chan bool)

				LogicalName := args[1]
				Tenant := args[2]
				BindAdd := args[3]
				LinkAdd := args[4]

				//CREATE AGGREGATOR
				fmt.Println("Running Gandalf with:")
				fmt.Println("  Logical Name : " + LogicalName)
				fmt.Println("  Tenant : " + Tenant)
				fmt.Println("  Bind Address : " + BindAdd)
				fmt.Println("  Link Address : " + LinkAdd)
				fmt.Println("  Config : " + config)

				aggregator.AggregatorMemberInit(LogicalName, Tenant, BindAdd, LinkAdd)

				<-done
			}
			break
		case "connector":
			TimeoutMax := int64(100000)
			if len(args) >= 6 {
				done := make(chan bool)

				LogicalName := args[1]
				Tenant := args[2]
				BindAdd := args[3]
				GrpcBindAdd := args[4]
				LinkAdd := args[5]

				if len(args) >= 7 {
					TimeoutMax, _ = strconv.ParseInt(args[6], 10, 64)
				}

				//CREATE CONNECTOR
				fmt.Println("Running Gandalf with:")
				fmt.Println("  Logical Name : " + LogicalName)
				fmt.Println("  Tenant : " + Tenant)
				fmt.Println("  Bind Address : " + BindAdd)
				fmt.Println("  Grpc Bind Address : " + GrpcBindAdd)
				fmt.Println("  Link Address : " + LinkAdd)
				fmt.Println("  Timeout Max : " + string(TimeoutMax))
				fmt.Println("  Config : " + config)

				connector.ConnectorMemberInit(LogicalName, Tenant, BindAdd, GrpcBindAdd, LinkAdd, TimeoutMax)

				<-done

			}
			break
		case "test":
			if len(args) >= 1 {
				command := args[1]
				switch command {
				case "send":
					if len(args) >= 5 {
						messageType := args[2]
						value := args[3]
						payload := args[4]
						var topic = ""

						if len(args) >= 6 {
							topic = args[5]
						}

						done := make(chan bool)
						tutu := demo.NewWorkerCliSend("test", messageType, value, payload, topic, []string{"127.0.0.1:7010", "127.0.0.1:7011"})
						go tutu.Run()
						<-done
					}

					break
				case "receive":
					if len(args) >= 4 {
						messageType := args[2]
						value := args[3]
						var topic = ""

						if len(args) >= 5 {
							topic = args[4]
						}

						done := make(chan bool)
						tutu := demo.NewWorkerCliReceive("test", messageType, value, topic, []string{"127.0.0.1:7110", "127.0.0.1:7111"})
						go tutu.Run()
						<-done
					}
					break
				default:
					break
				}
			} else {
				flag.Usage()
			}
		default:
			break
		}

	} else {
		flag.Usage()
	}
}
