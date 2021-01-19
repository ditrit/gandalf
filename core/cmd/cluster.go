/*
Copyright © 2020 DitRit community <contact@ditrit.io>
This file is part of Gandalf
*/

// Package cmd manages commands and configuration
package cmd

import (
	"fmt"
)

// clusterCmd represents the cluster command
var clusterCfg = NewConfigCmd(
	"cluster",
	"Launch gandalf in 'cluster' mode.",
	`Gandalf is launched as a cluster member of a Gandalf system.`,
	func(cfg *ConfigCmd, args []string) {
		fmt.Println("cluster called")
	})

func init() {
	rootCfg.AddConfig(clusterCfg)

	clusterCfg.Key("join", isStr, "j", "remote address (of an already existing member of cluster) to join")
	clusterCfg.SetCheck("join", CheckNotEmpty)
	clusterCfg.SetNormalize("join", TrimToLower)

	clusterCfg.Key("db_path", isStr, "", "path for the gandalf database (absolute or relative to the configuration directory)")
	clusterCfg.SetDefault("db_path", "db")

	clusterCfg.Key("db_nodename", isStr, "", "name of the gandalf node")
	clusterCfg.SetDefault("db_nodename", "node1")

}
