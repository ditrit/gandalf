/*
Copyright © 2020 DitRit community <contact@ditrit.io>
This file is part of Gandalf
*/

// Package cmd manages commands and configuration
package configuration

import (
	"fmt"

	"github.com/ditrit/gandalf/verdeter"

	"github.com/ditrit/gandalf/core/cluster"
	cmodels "github.com/ditrit/gandalf/core/configuration/models"

	"github.com/spf13/viper"
)

// clusterCmd represents the cluster command
var clusterCfg = verdeter.NewConfigCmd(
	"cluster",
	"Launch gandalf in 'cluster' mode.",
	`Gandalf is launched as a cluster member of a Gandalf system.`,
	func(cfg *verdeter.ConfigCmd, args []string) {
		fmt.Println("cluster called")

		offset := verdeter.GetOffset()
		fmt.Printf("computed offset : %d\n", offset)

		done := make(chan bool)
		configurationCluster := cmodels.NewConfigurationCluster()
		//fmt.Println(viper.GetString("bind"))
		//fmt.Println(configurationCluster.GetBindAddress())
		if !viper.IsSet("join") {
			fmt.Printf("calling ClusterMemberInit\n")
			cluster.ClusterMemberInit(configurationCluster)
		} else {
			fmt.Printf("calling ClusterMemberJoin\n")
			cluster.ClusterMemberJoin(configurationCluster)
		}
		fmt.Printf("Cluster call done\n")
		<-done
	})

func init() {
	rootCfg.AddConfig(clusterCfg)

	//clusterCfg.SetRequired("lname")
	clusterCfg.SetDefault("lname", "cluster")

	clusterCfg.Key("join", verdeter.IsStr, "j", "remote address (of an already existing member of cluster) to join")
	clusterCfg.SetCheck("join", verdeter.CheckNotEmpty)
	clusterCfg.SetNormalize("join", verdeter.TrimToLower)

	clusterCfg.Key("api_port", verdeter.IsInt, "", "Port to bind (default is 9199 + offset if defined)")
	//clusterCfg.SetDefault("api_port", 9199+verdeter.GetOffset())
	clusterCfg.SetCheck("api_port", verdeter.CheckTCPHighPort)
	clusterCfg.SetComputedValue("api_port",
		func() interface{} {
			return 9199 + verdeter.GetOffset()
		})

	clusterCfg.Key("db_path", verdeter.IsStr, "", "path for the gandalf database (absolute or relative to the configuration directory)")
	clusterCfg.SetCheck("db_path", func(val interface{}) bool {
		valStr, ok := val.(string)
		if ok {
			return verdeter.CreateWritableDirectory(valStr)
		}
		return false
	})
	clusterCfg.SetComputedValue("db_path",
		func() interface{} {
			ok := verdeter.CreateWritableDirectory("/var/lib/cockroach/")
			if ok {
				return "/var/lib/cockroach/"
			}
			dbDir := verdeter.GetHomeDirectory() + "/cockroach/"
			ok = verdeter.CreateWritableDirectory(dbDir)
			if ok {
				return dbDir
			}
			fmt.Println("Error: can't use or write into database directory")
			return nil
		})
	//clusterCfg.SetDefault("db_path", "/var/lib/cockroach/")

	clusterCfg.Key("db_nodename", verdeter.IsStr, "", "name of the gandalf node")
	clusterCfg.SetCheck("db_nodename", verdeter.CheckNotEmpty)
	//clusterCfg.SetDefault("db_nodename", "node1")
	clusterCfg.SetComputedValue("db_nodename",
		func() interface{} {
			return fmt.Sprint("node", verdeter.GetOffset())
		})

	clusterCfg.Key("db_port", verdeter.IsInt, "", "Port to bind (default is 9299 + offset if defined)")
	//clusterCfg.SetDefault("db_port", 9299)
	clusterCfg.SetCheck("db_port", verdeter.CheckTCPHighPort)
	clusterCfg.SetComputedValue("db_port",
		func() interface{} {
			return 9299 + verdeter.GetOffset()
		})

	clusterCfg.Key("db_http_port", verdeter.IsInt, "", "Port to bind (default is 9399 + offset if defined)")
	//clusterCfg.SetDefault("db_http_port", 9399+verdeter.GetOffset())
	clusterCfg.SetCheck("db_http_port", verdeter.CheckTCPHighPort)
	clusterCfg.SetComputedValue("db_http_port",
		func() interface{} {
			return 9399 + verdeter.GetOffset()
		})
	/* 	connectorCfg.SetComputedValue("db_http_port",
	func() interface{} {
		return 9299 + GetOffset()
	}) */

	clusterCfg.SetConstraint("a secret can not be set for cluster initialization (no join provided)",
		func() bool {
			return viper.IsSet("join") || !viper.IsSet("secret")
		})

	clusterCfg.SetConstraint("secret key is required to add a cluster member",
		func() bool {
			return !viper.IsSet("join") || viper.IsSet("secret")
		})
}
