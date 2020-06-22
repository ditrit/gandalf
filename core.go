//Package main :
package main

import (
	"flag"
	"fmt"
	"github.com/ditrit/gandalf-core/aggregator"
	"github.com/ditrit/gandalf-core/cluster"
	"github.com/ditrit/gandalf-core/connector"
	"github.com/ditrit/gandalf-core/database"
	"github.com/ditrit/gandalf-core/configuration"
	"os"
	net "github.com/ditrit/shoset"
	"strconv"
)

func main() {

	configuration.ConfigMain()

	gandalfLogicalName , err := configuration.GetStringConfig("logical_name")
	if err != nil {
		log.Fatalf("No logical name: %v",err)
	}
	gandalfLogPath, err := configuration.GetStringConfig("gandalf_log")
	if err != nil{
		log.Fatalf("No valid log path : %v",err)
	}
	gandalfBindAddress ,err := configuration.GetStringConfig("bind_address")
	if err != nil {
		log.Fatalf("No valid bind address : %v",err)
	}
	gandalfType, err := configuration.GetStringConfig("gandalf_type")
	if err == nil {

<<<<<<< HEAD
	flag.BoolVar(&debug, "d", false, "")
	flag.BoolVar(&debug, "debug", false, "")
	flag.StringVar(&config, "c", "", "")
	flag.StringVar(&config, "config", "", "")
	flag.Parse()
	args := flag.Args()

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

						home, _ := os.UserHomeDir()
						LogPath := home + "/gandalf/logs/cluster/"

						if len(args) >= 6 {
							LogPath = args[5]
						}

						dbPath := home + "/gandalf/database/"
						if len(args) >= 7 {
							dbPath = args[6]
						}

						//CREATE CLUSTER
						fmt.Println("Running Gandalf with:")
						fmt.Println("  Mode : " + mode)
						fmt.Println("  Logical Name : " + LogicalName)
						fmt.Println("  Bind Address : " + BindAdd)
						fmt.Println("  Log Path : " + LogPath)
						fmt.Println("  Db Path : " + dbPath)
						fmt.Println("  Config : " + config)

						cluster.ClusterMemberInit(LogicalName, BindAdd, dbPath, LogPath)

						add, _ := net.DeltaAddress(BindAdd, 1000)
						go database.DatabaseMemberInit(add, dbPath, 1)
						//database.List([]string{add})

						<-done
					} else {
						flag.Usage()
					}

					//break

				case "join": //join
					if len(args) >= 5 {
						done := make(chan bool)

						LogicalName := args[2]
						BindAdd := args[3]
						JoinAdd := args[4]

						home, _ := os.UserHomeDir()
						LogPath := home + "/gandalf/logs/cluster/"

						if len(args) >= 6 {
							LogPath = args[5]
						}

						dbPath := home + "/gandalf/database/"
						if len(args) >= 7 {
							dbPath = args[6]
						}

						//CREATE CLUSTER
						fmt.Println("Running Gandalf with:")
						fmt.Println("  Mode : " + mode)
						fmt.Println("  Logical Name : " + LogicalName)
						fmt.Println("  Bind Address : " + BindAdd)
						fmt.Println("  Join Address : " + JoinAdd)
						fmt.Println("  Log Path : " + LogPath)
						fmt.Println("  Db Path : " + dbPath)
						fmt.Println("  Config : " + config)

						member := cluster.ClusterMemberJoin(LogicalName, BindAdd, JoinAdd, dbPath, LogPath)

						add, _ := net.DeltaAddress(BindAdd, 1000)
						id := len(*member.Store)

						go database.DatabaseMemberInit(add, dbPath, id)

						err := database.AddNodesToLeader(id, add, *member.Store)
						fmt.Println(err)

						<-done
					} else {
						flag.Usage()
					}

					//break

				default:
					break
=======
		switch gandalfType {
		case "cluster":
			gandalfDBPath, err := configuration.GetStringConfig("gandalf_db")
			if err != nil {
				log.Fatalf("No valid database path : %v",err)
			}
			gandalfJoin, err := configuration.GetStringConfig("cluster_join")
			if err == nil {
				if gandalfJoin == "" {
					done := make(chan bool)
					cluster.ClusterMemberInit(gandalfLogicalName, gandalfBindAddress, gandalfLogPath)
					add, _ := net.DeltaAddress(gandalfBindAddress, 1000)
					go database.DatabaseMemberInit(add, gandalfDBPath, 1)
					<- done
				} else {
					done := make(chan bool)
					member := cluster.ClusterMemberJoin(gandalfLogicalName, gandalfBindAddress, gandalfJoin, gandalfLogPath)
					add, _ := net.DeltaAddress(gandalfBindAddress, 1000)
					id := len(*member.Store)

					go database.DatabaseMemberInit(add, gandalfDBPath, id)

					_ = database.AddNodesToLeader(id, add, *member.Store)
					<- done
>>>>>>> link core et configuration
				}
			}
			break
		case "aggregator":
<<<<<<< HEAD
			if len(args) >= 5 {
				done := make(chan bool)

				LogicalName := args[1]
				Tenant := args[2]
				BindAdd := args[3]
				LinkAdd := args[4]

				home, _ := os.UserHomeDir()
				LogPath := home + "/gandalf/logs/aggregator/"

				if len(args) >= 6 {
					LogPath = args[5]
				}

				//CREATE AGGREGATOR
				fmt.Println("Running Gandalf with:")
				fmt.Println("  Logical Name : " + LogicalName)
				fmt.Println("  Tenant : " + Tenant)
				fmt.Println("  Bind Address : " + BindAdd)
				fmt.Println("  Link Address : " + LinkAdd)
				fmt.Println("  Log Path : " + LogPath)
				fmt.Println("  Config : " + config)

				aggregator.AggregatorMemberInit(LogicalName, Tenant, BindAdd, LinkAdd, LogPath)

				<-done
			}

			//break

		case "connector":
			if len(args) >= 7 {
				done := make(chan bool)

				LogicalName := args[1]
				Tenant := args[2]
				BindAdd := args[3]
				GrpcBindAdd := args[4]
				LinkAdd := args[5]
				ConnectorType := args[6]

				TargetAdd := ""

				if len(args) >= 8 {
					TargetAdd = args[7]
				}

				home, _ := os.UserHomeDir()
				WorkerPath := home + "/gandalf/workers/" + ConnectorType + "/"

				if len(args) >= 9 {
					WorkerPath = args[8]
				}

				LogPath := home + "/gandalf/logs/connector/"

				if len(args) >= 10 {
					LogPath = args[9]
				}

				TimeoutMax := int64(100000)

				if len(args) >= 11 {
					TimeoutMax, _ = strconv.ParseInt(args[10], 10, 64)
				}

				//CREATE CONNECTOR
				fmt.Println("Running Gandalf with:")
				fmt.Println("  Logical Name : " + LogicalName)
				fmt.Println("  Tenant : " + Tenant)
				fmt.Println("  Bind Address : " + BindAdd)
				fmt.Println("  Grpc Bind Address : " + GrpcBindAdd)
				fmt.Println("  Link Address : " + LinkAdd)
				fmt.Println("  Connector Type : " + ConnectorType)
				fmt.Println("  Target Address : " + TargetAdd)
				fmt.Println("  Worker Path : " + WorkerPath)
				fmt.Println("  Log Path : " + LogPath)
				fmt.Printf("  Timeout Max : %d \n", TimeoutMax)
				fmt.Println("  Config : " + config)

				connector.ConnectorMemberInit(LogicalName, Tenant, BindAdd, GrpcBindAdd, LinkAdd, ConnectorType, TargetAdd, WorkerPath, LogPath, TimeoutMax)

				<-done
			}

			//break

		case "test":
			if len(args) >= 1 {
				command := args[1]
				switch command {
				case "list":
					fmt.Println("LIST")
					database.List([]string{"127.0.0.1:10000", "127.0.0.1:10001", "127.0.0.1:10002"})

					//break

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
	}  */

	configuration.ConfigMain()

=======
			gandalfTenant,err := configuration.GetStringConfig("tenant")
			if err !=nil {
				log.Fatalf("no valid tenant : %v",err)
			}
			gandalfClusterLink,err := configuration.GetStringConfig("clusters")
			if err != nil {
				log.Fatalf("no valid cluster address: %v",err)
			}
			done := make(chan bool)
			aggregator.AggregatorMemberInit(gandalfLogicalName, gandalfTenant, gandalfBindAddress, gandalfClusterLink, gandalfLogPath)
			<-done
			break
		case "connector":
			gandalfTenant,err := configuration.GetStringConfig("tenant")
			if err !=nil {
				log.Fatalf("no valid tenant : %v",err)
			}
			gandalfGRPCBindAddress, err := configuration.GetStringConfig("grpc_bind_address")
			if err !=nil {
				log.Fatalf("no valid tenant : %v",err)
			}
			gandalfAggregatorLink,err := configuration.GetStringConfig("aggregators")
			if err !=nil {
				log.Fatalf("no valid tenant : %v",err)
			}
			gandalfMaxTimeout,err := configuration.GetIntegerConfig("max_timeout")
			if err !=nil {
				log.Fatalf("no valid tenant : %v",err)
			}
			done := make(chan bool)
			connector.ConnectorMemberInit(gandalfLogicalName, gandalfTenant, gandalfBindAddress, gandalfGRPCBindAddress, gandalfAggregatorLink, gandalfLogPath, int64(gandalfMaxTimeout))
			<- done
			break

		default:
			break
		}
	}
	/*
	//CREATE CLUSTER
	fmt.Println("Running Gandalf with:")
	fmt.Println("  Mode : " + mode)
	fmt.Println("  Logical Name : " + LogicalName)
	fmt.Println("  Bind Address : " + BindAdd)
	fmt.Println("  Log Path : " + LogPath)
	fmt.Println("  Db Path : " + dbPath)
	fmt.Println("  Config : " + config)

	<-done

	//CREATE CLUSTER
	fmt.Println("Running Gandalf with:")
	fmt.Println("  Mode : " + mode)
	fmt.Println("  Logical Name : " + LogicalName)
	fmt.Println("  Bind Address : " + BindAdd)
	fmt.Println("  Join Address : " + JoinAdd)
	fmt.Println("  Log Path : " + LogPath)
	fmt.Println("  Db Path : " + dbPath)
	fmt.Println("  Config : " + config)

	<-done

	//CREATE AGGREGATOR
	fmt.Println("Running Gandalf with:")
	fmt.Println("  Logical Name : " + LogicalName)
	fmt.Println("  Tenant : " + Tenant)
	fmt.Println("  Bind Address : " + BindAdd)
	fmt.Println("  Link Address : " + LinkAdd)
	fmt.Println("  Log Path : " + LogPath)
	fmt.Println("  Config : " + config)

	aggregator.AggregatorMemberInit(LogicalName, Tenant, BindAdd, LinkAdd, LogPath)

	<-done
	//CREATE CONNECTOR
	fmt.Println("Running Gandalf with:")
	fmt.Println("  Logical Name : " + LogicalName)
	fmt.Println("  Tenant : " + Tenant)
	fmt.Println("  Bind Address : " + BindAdd)
	fmt.Println("  Grpc Bind Address : " + GrpcBindAdd)
	fmt.Println("  Link Address : " + LinkAdd)
	fmt.Println("  Log Path : " + LogPath)
	fmt.Printf("   Timeout Max : %d \n", TimeoutMax)
	fmt.Println("  Config : " + config)

	connector.ConnectorMemberInit(LogicalName, Tenant, BindAdd, GrpcBindAdd, LinkAdd, LogPath, TimeoutMax)

	<-done
 	*/
>>>>>>> link core et configuration
}
