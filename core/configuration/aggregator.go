/*
Copyright © 2020 DitRit community <contact@ditrit.io>
This file is part of Gandalf
*/

// Package cmd manages commands and configuration
package configuration

import (
	"fmt"

	"github.com/ditrit/gandalf/core/aggregator"
	cmodels "github.com/ditrit/gandalf/core/configuration/models"
	"github.com/ditrit/gandalf/verdeter"

	"github.com/spf13/viper"
)

// aggregatorCmd represents the aggregator command
var aggregatorCfg = verdeter.NewConfigCmd(
	"aggregator",
	"Launch gandalf in 'aggregator' mode.",
	`Gandalf is launched as an aggregator instance.`,
	func(cfg *verdeter.ConfigCmd, args []string) {
		fmt.Println("aggregator called")
		fmt.Printf("tenant = '%s'\n", viper.GetString("tenant"))
		fmt.Println("cluster to connect = " + viper.GetString("cluster"))
		done := make(chan bool)
		viper.WriteConfig()

		configurationAggregator := cmodels.NewConfigurationAggregator()
		aggregator.AggregatorMemberInit(configurationAggregator)
		<-done
	})

func init() {
	rootCfg.AddConfig(aggregatorCfg)

	aggregatorCfg.SetRequired("lname")

	aggregatorCfg.Key("api_port", verdeter.IsInt, "", "Port to bind (default is 9199 + offset if defined)")
	aggregatorCfg.SetCheck("api_port", verdeter.CheckTCPHighPort)
	aggregatorCfg.SetComputedValue("api_port",
		func() interface{} {
			return 9199 + verdeter.GetOffset()
		})

	aggregatorCfg.Key("tenant", verdeter.IsStr, "t", "name of the tenant name of the aggregator")
	aggregatorCfg.SetCheck("tenant", verdeter.CheckNotEmpty)
	aggregatorCfg.SetRequired("tenant")
	aggregatorCfg.SetNormalize("tenant", verdeter.TrimToLower)

	aggregatorCfg.Key("cluster", verdeter.IsStr, "c", "remote address of one of the cluster members to link")
	aggregatorCfg.SetCheck("cluster", verdeter.CheckNotEmpty)
	aggregatorCfg.SetRequired("cluster")
	aggregatorCfg.SetNormalize("cluster", verdeter.TrimToLower)

	aggregatorCfg.SetRequired("secret")

}
