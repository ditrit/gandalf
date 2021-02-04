/*
Copyright © 2020 DitRit community <contact@ditrit.io>
This file is part of Gandalf
*/

// Package cmd manages commands and configuration
package configuration

import (
	"fmt"

	"github.com/ditrit/gandalf/core/configuration/config"

	"github.com/ditrit/gandalf/core/cluster"
	cmodels "github.com/ditrit/gandalf/core/configuration/models"

	"github.com/spf13/viper"
)

// clusterCmd represents the cluster command
var clusterCfg = config.NewConfigCmd(
	"cluster",
	"Launch gandalf in 'cluster' mode.",
	`Gandalf is launched as a cluster member of a Gandalf system.`,
	func(cfg *config.ConfigCmd, args []string) {
		fmt.Println("cluster called")

		port := viper.GetInt("port")
		fmt.Printf("computed port : %d\n", port)

		offset := config.GetOffset()
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

	clusterCfg.Key("join", config.IsStr, "j", "remote address (of an already existing member of cluster) to join")
	clusterCfg.SetCheck("join", CheckNotEmpty)
	clusterCfg.SetNormalize("join", TrimToLower)

	clusterCfg.Key("api_port", config.IsInt, "", "Address to bind (default is *:9199)")
	clusterCfg.SetDefault("api_port", 9199+config.GetOffset())

	clusterCfg.Key("db_path", config.IsStr, "", "path for the gandalf database (absolute or relative to the configuration directory)")
	clusterCfg.SetCheck("db_path", CheckNotEmpty)
	clusterCfg.SetDefault("db_path", "/var/lib/cockroach/")

	clusterCfg.Key("db_nodename", config.IsStr, "", "name of the gandalf node")
	clusterCfg.SetCheck("db_nodename", CheckNotEmpty)
	clusterCfg.SetDefault("db_nodename", "node1")

	clusterCfg.Key("db_port", config.IsInt, "", "Address to bind (default is *:9299)")
	clusterCfg.SetDefault("db_port", 9299+config.GetOffset())

	/* 	connectorCfg.SetComputedValue("db_port",
	func() interface{} {
		return 9199 + GetOffset()
	}) */

	clusterCfg.Key("db_http_port", config.IsInt, "", "Address to bind (default is *:9399)")
	clusterCfg.SetDefault("db_http_port", 9399+config.GetOffset())

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
