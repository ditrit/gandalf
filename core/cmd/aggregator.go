/*
Copyright © 2020 DitRit community <contact@ditrit.io>
This file is part of Gandalf
*/

// Package cmd manages commands and configuration
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// aggregatorCmd represents the aggregator command
var aggregatorCfg = NewConfigCmd(
	"aggregator",
	"Launch gandalf in 'aggregator' mode.",
	`Gandalf is launched as an aggregator instance.`,
	func(cfg *ConfigCmd, args []string) {
		certDir := viper.GetString("cert_dir")
		fmt.Printf("cert_dir value = %s", certDir)

		fmt.Println("aggregator called")
		fmt.Printf("tenant = '%s'\n", viper.GetString("tenant"))
		fmt.Println("cluster to connect = " + viper.GetString("cluster"))
	})

func init() {
	rootCfg.AddConfig(aggregatorCfg)

	aggregatorCfg.Key("tenant", isStr, "t", "name of the tenant name of the aggregator")
	aggregatorCfg.SetCheck("tenant", CheckNotEmpty)
	aggregatorCfg.SetRequired("tenant")
	aggregatorCfg.SetNormalize("tenant", TrimToLower)

	aggregatorCfg.Key("cluster", isStr, "c", "remote address of one of the cluster members to link")
	aggregatorCfg.SetCheck("cluster", CheckNotEmpty)
	aggregatorCfg.SetRequired("cluster")
	aggregatorCfg.SetNormalize("cluster",
		func(value interface{}) interface{} {
			str := value.(string)
			return strings.ToLower(strings.TrimSpace(str))
		})
}
