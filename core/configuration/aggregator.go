/*
Copyright © 2020 DitRit community <contact@ditrit.io>
This file is part of Gandalf
*/

// Package configuration manages commands and configuration
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

		fmt.Println("Aggregator call done")

		<-done
	})

func init() {
	startCfg.AddConfig(aggregatorCfg)

	//	aggregatorCfg.SetRequired("lname")

	aggregatorCfg.LKey("api_port", verdeter.IsInt, "", "Port to bind (default is 9199 + offset if defined)")
	aggregatorCfg.SetCheck("api_port", verdeter.CheckTCPHighPort)
	aggregatorCfg.SetComputedValue("api_port",
		func() interface{} {
			return 9199 + verdeter.GetOffset()
		})

	aggregatorCfg.LKey("api_path", verdeter.IsStr, "", "path for the gandalf API files (absolute or relative to the configuration directory)")
	aggregatorCfg.SetCheck("api_path", func(val interface{}) bool {
		valStr, ok := val.(string)
		if ok {
			return verdeter.CreateWritableDirectory(valStr)
		}
		return false
	})
	aggregatorCfg.SetComputedValue("api_path",
		func() interface{} {
			ok := verdeter.CreateWritableDirectory("/var/lib/gandalf/files/")
			if ok {
				return "/var/lib/gandalf/files/"
			}
			dbDir := verdeter.GetHomeDirectory() + "/gandalf/files/"
			ok = verdeter.CreateWritableDirectory(dbDir)
			if ok {
				return dbDir
			}
			fmt.Println("Error: can't use or write into API directory")
			return nil
		})

	aggregatorCfg.LKey("tenant", verdeter.IsStr, "t", "name of the tenant name of the aggregator")
	aggregatorCfg.SetCheck("tenant", verdeter.CheckNotEmpty)
	aggregatorCfg.SetRequired("tenant")
	aggregatorCfg.SetNormalize("tenant", verdeter.TrimToLower)

	aggregatorCfg.LKey("cluster", verdeter.IsStr, "c", "remote address of one of the cluster members to link")
	aggregatorCfg.SetCheck("cluster", verdeter.CheckNotEmpty)
	aggregatorCfg.SetRequired("cluster")
	aggregatorCfg.SetNormalize("cluster", verdeter.TrimToLower)

	aggregatorCfg.SetRequired("secret")

}
