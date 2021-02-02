/*
Copyright © 2020 DitRit community <contact@ditrit.io>
This file is part of Gandalf
*/

// Package cmd manages commands and configuration
package cmd

import (
	"fmt"
	"gandalf/core/cluster"
	"gandalf/core/cmd/models"

	"github.com/spf13/viper"
)

// clusterCmd represents the cluster command
var clusterCfg = NewConfigCmd(
	"cluster",
	"Launch gandalf in 'cluster' mode.",
	`Gandalf is launched as a cluster member of a Gandalf system.`,
	func(cfg *ConfigCmd, args []string) {
		fmt.Println("cluster called")

		port := viper.GetInt("port")
		fmt.Printf("computed port : %d\n", port)

		offset := GetOffset()
		fmt.Printf("computed offset : %d\n", offset)

		done := make(chan bool)
		configurationCluster := models.NewConfigurationCluster()

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

	clusterCfg.Key("join", isStr, "j", "remote address (of an already existing member of cluster) to join")
	clusterCfg.SetCheck("join", CheckNotEmpty)
	clusterCfg.SetNormalize("join", TrimToLower)

	clusterCfg.Key("db_path", isStr, "", "path for the gandalf database (absolute or relative to the configuration directory)")
	clusterCfg.SetCheck("db_path", CheckNotEmpty)
	clusterCfg.SetDefault("db_path", "db")

	clusterCfg.Key("db_nodename", isStr, "", "name of the gandalf node")
	clusterCfg.SetCheck("db_nodename", CheckNotEmpty)
	clusterCfg.SetDefault("db_nodename", "node1")

	//TODO REPLACE CALCULATED
	//clusterCfg.Key("db_bind", isStr, "", "Database address to bind (default is *:10099)")
	//clusterCfg.SetDefault("db_bind", "*:10099")
	connectorCfg.SetComputedValue("db_bind",
		func() interface{} {
			return 9199 + GetOffset()
		})

	//clusterCfg.Key("db_http_bind", isStr, "", "Database HTTP address to bind (default is *:11099)")
	//clusterCfg.SetDefault("db_http_bind", "*:11099")
	connectorCfg.SetComputedValue("db_http_bind",
		func() interface{} {
			return 9299 + GetOffset()
		})

	clusterCfg.SetConstraint("a secret can not be set for cluster initialization (no join provided)",
		func() bool {
			return viper.IsSet("join") || !viper.IsSet("secret")
		})

	clusterCfg.SetConstraint("secret key is required to add a cluster member",
		func() bool {
			return !viper.IsSet("join") || viper.IsSet("secret")
		})

	clusterCfg.SetComputedValue("port",
		func() interface{} {
			return 9099 + GetOffset()
		})

}
