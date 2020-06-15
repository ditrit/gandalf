//Package main :
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/ditrit/gandalf-core/aggregator"
	"github.com/ditrit/gandalf-core/cluster"
	"github.com/ditrit/gandalf-core/connector"
	"github.com/ditrit/gandalf-core/database"

	net "github.com/ditrit/shoset"
)

func main() {

	/* 	KeysConfig := map[string]models.KeyConfig{
	"foo": 1,
	"bar": 2}
	*/
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
		fmt.Printf("  		  log path     \n")
		fmt.Printf("  		  db path     \n")
		fmt.Printf("    mode : aggregator\n")
		fmt.Printf("        arguments:\n")
		fmt.Printf("  	      logical name	  \n")
		fmt.Printf("  		  bind address    \n")
		fmt.Printf("  		  link address     \n")
		fmt.Printf("  		  log path     \n")
		fmt.Printf("    mode : connector\n")
		fmt.Printf("        arguments:\n")
		fmt.Printf("  	      logical name	  \n")
		fmt.Printf("  		  bind address    \n")
		fmt.Printf("  		  bind grpc address    \n")
		fmt.Printf("  		  link address     \n")
		fmt.Printf("  		  log path     \n")
		fmt.Printf("  		  timeout max     \n")
	}

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
	}
}
