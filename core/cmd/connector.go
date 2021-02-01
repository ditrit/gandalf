/*
Copyright © 2020 DitRit community <contact@ditrit.io>
This file is part of Gandalf
*/

// Package cmd manages commands and configuration
package cmd

import (
	"fmt"
	"strings"
)

var connectorCfg = NewConfigCmd(
	"connector",
	"Launch gandalf in 'connector' mode.",
	`Gandalf is launched as a connector instance.`,
	func(cfg *ConfigCmd, args []string) {
		fmt.Println("connector called")
	})

func init() {
	rootCfg.AddConfig(connectorCfg)

	connectorCfg.Key("aggregator", isStr, "a", "remote address of one of the cluster members to link")
	connectorCfg.SetCheck("aggregator", CheckNotEmpty)
	connectorCfg.SetRequired("aggregator")
	connectorCfg.SetNormalize("aggregator", TrimToLower)

	connectorCfg.Key("class", isStr, "c", "the type of connector (bus, csv, orchestrator, etc.)")
	connectorCfg.SetCheck("class", CheckNotEmpty)
	connectorCfg.SetRequired("class")
	connectorCfg.SetNormalize("product", TrimToLower)

	connectorCfg.Key("product", isStr, "p", "the type of connector (bus, csv, orchestrator, etc.)")
	connectorCfg.SetCheck("product", CheckNotEmpty)
	connectorCfg.SetRequired("product")
	connectorCfg.SetNormalize("product", TrimToLower)

	connectorCfg.Key("workers", isStr, "w", "path for the workers configuration (absolute or relative to the certificates directory)")
	connectorCfg.SetDefault("workers", "workers")

	//TODO REVOIR
	connectorCfg.Key("grpc_dir", isStr, "w", "path for the sockets directory (absolute or relative to the certificates directory)")
	connectorCfg.SetDefault("grpc_dir", "/tmp")

	//TODO CALCULATED
	connectorCfg.Key("grpc_bind", isStr, "w", "GRPC address to bind (default is [grpc_dir]_[class]_[product]_[hash])")
	connectorCfg.SetDefault("grpc_bind", "")

	connectorCfg.Key("workers_url", isStr, "w", "workers URL")
	connectorCfg.SetDefault("workers_url", "https://github.com/ditrit/workers/raw/master")

	connectorCfg.Key("versions", isStr, "w", "worker versions")

	//END TODO
	connectorCfg.Key("update_mode", isStr, "", "update mode (manual|auto|planed)")
	connectorCfg.SetDefault("update_mode", "manual")
	connectorCfg.SetCheck("update_mode", func(val interface{}) bool {
		strVal := strings.ToLower(strings.TrimSpace(val.(string)))
		return map[string]bool{"manual": true, "auto": true, "planed": true}[strVal]
	})

	connectorCfg.Key("update_time", isStr, "", "time for planed update mode")

}
