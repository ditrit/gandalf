//Package main :
package main

import (
	"fmt"
	"log"
	"os"
	"github.com/ditrit/gandalf/core/aggregator"
	"github.com/ditrit/gandalf/core/cluster"
	"github.com/ditrit/gandalf/core/configuration"
	"github.com/ditrit/gandalf/core/connector"
	"github.com/ditrit/gandalf/core/database"
	net "github.com/ditrit/shoset"
)

func main() {

	configuration.ConfigMain(os.Args[0],os.Args[1:])

	gandalfLogicalName, err := configuration.GetStringConfig("logical_name")
	if err != nil {
		log.Fatalf("No logical name: %v", err)
	}
	gandalfLogPath, err := configuration.GetStringConfig("gandalf_log")
	if err != nil {
		log.Fatalf("No valid log path : %v", err)
	}
	gandalfBindAddress, err := configuration.GetStringConfig("bind_address")
	if err != nil {
		log.Fatalf("No valid bind address : %v", err)
	}
	gandalfType, err := configuration.GetStringConfig("gandalf_type")
	if err == nil {

		switch gandalfType {
		case "cluster":
			gandalfDBPath, err := configuration.GetStringConfig("gandalf_db")
			if err != nil {
				log.Fatalf("No valid database path : %v", err)
			}
			gandalfJoin, err := configuration.GetStringConfig("cluster_join")
			if err == nil {
				if gandalfJoin == "" {
					done := make(chan bool)
					//CREATE CLUSTER
					fmt.Println("Running Gandalf for a " + gandalfType + " with :")
					fmt.Println("  Logical Name : " + gandalfLogicalName)
					fmt.Println("  Bind Address : " + gandalfBindAddress)
					fmt.Println("  Log Path : " + gandalfLogPath)
					fmt.Println("  Db Path : " + gandalfDBPath)
					cluster.ClusterMemberInit(gandalfLogicalName, gandalfBindAddress, gandalfDBPath, gandalfLogPath)
					add, _ := net.DeltaAddress(gandalfBindAddress, 1000)
					go database.DatabaseMemberInit(add, gandalfDBPath, 1)
					<-done
				} else {
					done := make(chan bool)
					//CREATE CLUSTER
					fmt.Println("Running Gandalf for a " + gandalfType + " with :")
					fmt.Println("  Logical Name : " + gandalfLogicalName)
					fmt.Println("  Bind Address : " + gandalfBindAddress)
					fmt.Println("  Join Address : " + gandalfJoin)
					fmt.Println("  Log Path : " + gandalfLogPath)
					fmt.Println("  Db Path : " + gandalfDBPath)
					member := cluster.ClusterMemberJoin(gandalfLogicalName, gandalfBindAddress, gandalfJoin, gandalfDBPath, gandalfLogPath)
					fmt.Println(member)
					add, _ := net.DeltaAddress(gandalfBindAddress, 1000)
					id := len(*member.Store)

					go database.DatabaseMemberInit(add, gandalfDBPath, id)

					_ = database.AddNodesToLeader(id, add, *member.Store)
					<-done
				}
			}
			break
		case "aggregator":
			//CREATE AGGREGATOR
			gandalfTenant, err := configuration.GetStringConfig("tenant")
			if err != nil {
				log.Fatalf("no valid tenant : %v", err)
			}
			gandalfClusterLink, err := configuration.GetStringConfig("clusters")
			if err != nil {
				log.Fatalf("no valid cluster address: %v", err)
			}
			fmt.Println("Running Gandalf for a " + gandalfType + " with :")
			fmt.Println("  Logical Name : " + gandalfLogicalName)
			fmt.Println("  Tenant : " + gandalfTenant)
			fmt.Println("  Bind Address : " + gandalfBindAddress)
			fmt.Println("  Link Address : " + gandalfClusterLink)
			fmt.Println("  Log Path : " + gandalfLogPath)
			done := make(chan bool)
			aggregator.AggregatorMemberInit(gandalfLogicalName, gandalfTenant, gandalfBindAddress, gandalfClusterLink, gandalfLogPath)
			<-done
			break
		case "connector":
			gandalfTenant, err := configuration.GetStringConfig("tenant")
			if err != nil {
				log.Fatalf("Invalid tenant : %v", err)
			}
			gandalfGRPCBindAddress, err := configuration.GetStringConfig("grpc_bind_address")
			if err != nil {
				log.Fatalf("Invalid  bind address : %v", err)
			}
			gandalfAggregatorLink, err := configuration.GetStringConfig("aggregators")
			if err != nil {
				log.Fatalf("Invalid aggregator address to link to : %v", err)
			}
			gandalfConnectorType, err := configuration.GetStringConfig("connector_type")
			if err != nil {
				log.Fatalf("Invalid connector type : %v", err)
			}
			gandalfProduct, err := configuration.GetStringConfig("product_name")
			if err != nil {
				log.Fatalf("Invalid product: %v", err)
			}
			gandalfProductUrl, err := configuration.GetStringConfig("product_url")
			if err != nil {
				log.Fatalf("Invalid product url : %v", err)
			}
			gandalfWorkers, err := configuration.GetStringConfig("workers")
			if err != nil {
				log.Fatalf("Invalid workers path: %v", err)
			}
			gandalfMaxTimeout, err := configuration.GetIntegerConfig("max_timeout")
			if err != nil {
				log.Fatalf("Invalid maximum timeout : %v", err)
			}
			gandalfVersionsString, err := configuration.GetStringConfig("versions")
			if err != nil {
				log.Fatalf("Invalid versions : %v", err)
			}
			gandalfVersions,err := configuration.GetVersionsList(gandalfVersionsString)
			if err != nil {
				log.Fatalf("Invalid versions : %v", err)
			}

			//CREATE CONNECTOR
			fmt.Println("Running Gandalf for a " + gandalfType + " with :")
			fmt.Println("  Logical Name : " + gandalfLogicalName)
			fmt.Println("  Tenant : " + gandalfTenant)
			fmt.Println("  Bind Address : " + gandalfBindAddress)
			fmt.Println("  Grpc Bind Address : " + gandalfGRPCBindAddress)
			fmt.Println("  Link Address : " + gandalfAggregatorLink)
			fmt.Println("  Connector Type : " + gandalfConnectorType)
			fmt.Println("  Product : " + gandalfProduct)
			fmt.Println("  Product Url : " + gandalfProductUrl)
			fmt.Println("  Workers Path : " + gandalfWorkers)
			fmt.Println("  Log Path : " + gandalfLogPath)
			fmt.Println("  Maximum timeout :", gandalfMaxTimeout)
			done := make(chan bool)

			connector.ConnectorMemberInit(gandalfLogicalName, gandalfTenant, gandalfBindAddress, gandalfGRPCBindAddress, gandalfAggregatorLink, gandalfConnectorType, gandalfProductUrl, gandalfWorkers, gandalfLogPath, int64(gandalfMaxTimeout), gandalfVersions)
			<-done
			break

		default:
			break
		}
	}

}
