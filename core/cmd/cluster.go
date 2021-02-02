/*
Copyright © 2020 DitRit community <contact@ditrit.io>
This file is part of Gandalf
*/

// Package cmd manages commands and configuration
package cmd

import (
	"fmt"

	"github.com/ditrit/gandalf/core/cluster"
	"github.com/ditrit/gandalf/core/cmd/models"

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

	clusterCfg.Key("join", isStr, "j", "remote address (of an already existing member of cluster) to join")
	clusterCfg.SetCheck("join", CheckNotEmpty)
	clusterCfg.SetNormalize("join", TrimToLower)

	clusterCfg.Key("db_path", isStr, "", "path for the gandalf database (absolute or relative to the configuration directory)")
	clusterCfg.SetCheck("db_path", CheckNotEmpty)
	//TEST
	clusterCfg.SetDefault("db_path", "/home/romainfairant/gandalf")

	clusterCfg.Key("db_nodename", isStr, "", "name of the gandalf node")
	clusterCfg.SetCheck("db_nodename", CheckNotEmpty)
	clusterCfg.SetDefault("db_nodename", "node1")

	clusterCfg.Key("db_port", isInt, "", "Address to bind (default is *:9199)")
	clusterCfg.SetDefault("db_port", 9199+GetOffset())

	/* 	connectorCfg.SetComputedValue("db_port",
	func() interface{} {
		return 9199 + GetOffset()
	}) */

	clusterCfg.Key("db_http_port", isInt, "", "Address to bind (default is *:9299)")
	clusterCfg.SetDefault("db_http_port", 9299+GetOffset())

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
