/*
Copyright © 2020 DitRit community <contact@ditrit.io>
This file is part of Gandalf
*/

// Package configuration manages commands and configuration
package configuration

import (
	"fmt"
	"strings"

	"github.com/ditrit/gandalf/verdeter"

	cmodels "github.com/ditrit/gandalf/core/configuration/models"
	"github.com/ditrit/gandalf/core/connector"

	"github.com/spf13/viper"
)

var connectorCfg = verdeter.NewConfigCmd(
	"connector",
	"Launch gandalf in 'connector' mode.",
	`Gandalf is launched as a connector instance.`,
	func(cfg *verdeter.ConfigCmd, args []string) {
		fmt.Println("connector called")

		done := make(chan bool)
		configurationConnector := cmodels.NewConfigurationConnector(cfg)
		connector.ConnectorMemberInit(configurationConnector)
		//go oauth2.NewOAuth2Client()
		<-done
	})

func init() {
	startCfg.AddConfig(connectorCfg)

	//	connectorCfg.SetRequired("lname")

	connectorCfg.LKey("aggregator", verdeter.IsStr, "a", "remote address of one of the cluster members to link")
	connectorCfg.SetCheck("aggregator", verdeter.CheckNotEmpty)
	connectorCfg.SetRequired("aggregator")
	connectorCfg.SetNormalize("aggregator", verdeter.TrimToLower)

	connectorCfg.LKey("class", verdeter.IsStr, "c", "the type of connector (bus, csv, orchestrator, etc.)")
	connectorCfg.SetCheck("class", verdeter.CheckNotEmpty)
	connectorCfg.SetRequired("class")
	connectorCfg.SetNormalize("product", verdeter.TrimToLower)

	connectorCfg.LKey("product", verdeter.IsStr, "p", "the type of connector (bus, csv, orchestrator, etc.)")
	connectorCfg.SetCheck("product", verdeter.CheckNotEmpty)
	connectorCfg.SetRequired("product")
	connectorCfg.SetNormalize("product", verdeter.TrimToLower)

	//connectorCfg.LKey("workers", verdeter.IsStr, "w", "path for the workers configuration (absolute or relative to the certificates directory)")
	//connectorCfg.SetDefault("workers", "/tmp/")
	/* 	connectorCfg.SetCheck("workers", func(val interface{}) bool {
		valStr, ok := val.(string)
		if ok {
			return verdeter.CreateWritableDirectory(viper.GetString("config_dir") + "workers")
		}
		return false
	}) */
	connectorCfg.SetComputedValue("gandalf_var",
		func() interface{} {
			path := "/var/lib/gandalf/"
			ok := verdeter.CreateWritableDirectory(path)
			if ok {
				viper.Set("gandalf_var", path)
				return path
			}
			path = verdeter.GetHomeDirectory() + "/gandalf/"
			ok = verdeter.CreateWritableDirectory(path)
			if ok {
				viper.Set("gandalf_var", path)
				return path
			}
			fmt.Println("Error: can't create gandalf data directory")
			return nil
		})

	connectorCfg.SetComputedValue("workers",
		func() interface{} {
			//viper.Set("gandalf_var", connectorCfg.GetComputedValue()["gandalf_var"])
			ok := verdeter.CreateWritableDirectory(viper.GetString("config_dir") + "workers")
			if !ok {
				fmt.Println("Error: can't create workers subdirectory into config directory")
				return nil
			}
			fmt.Println(viper.GetString("gandalf_var"))
			workersVarDir := viper.GetString("gandalf_var") + "workers"
			ok = verdeter.CreateWritableDirectory(workersVarDir)
			if !ok {
				fmt.Println("Error: can't create workers subdirectory into config directory")
				return nil
			}
			viper.Set("workers_vardir", workersVarDir)
			return viper.GetString("config_dir") + "workers"
		})

	//TODO REVOIR DEFAULT
	//connectorCfg.Key("grpc_dir", verdeter.IsStr, "g", "path for the sockets directory (absolute or relative to the certificates directory)")
	//connectorCfg.SetDefault("grpc_dir", "/tmp/")
	connectorCfg.SetComputedValue("var_run_dir",
		func() interface{} {
			ok := verdeter.CreateWritableDirectory("/var/run/gandalf/")
			if ok {
				viper.Set("var_run_dir", "/var/run/gandalf/")
				return "/var/run/gandalf/"
			}
			ok = verdeter.CreateWritableDirectory("/tmp/gandalf/run/")
			if ok {
				viper.Set("var_run_dir", "/tmp/gandalf/run/")
				return "/tmp/gandalf/run/"
			}
			fmt.Println("Error: can't create gRPC run directory")
			return nil
		})

	//connectorCfg.Key("grpc_bind", isStr, "", "GRPC address to bind (default is [grpc_dir]_[class]_[product]_[hash])")
	connectorCfg.SetComputedValue("grpc_bind",
		func() interface{} {
			return viper.GetString("var_run_dir") + viper.GetString("lname") + "_" + viper.GetString("class") + "_" + viper.GetString("product") //+ "_" + utils.GenerateHash(viper.GetString("lname"))
		})

	connectorCfg.LKey("workers_url", verdeter.IsStr, "u", "workers URL")
	connectorCfg.SetDefault("workers_url", "https://github.com/ditrit/workers/raw/master")

	connectorCfg.LKey("versions", verdeter.IsStr, "v", "worker versions")

	connectorCfg.LKey("update_mode", verdeter.IsStr, "", "update mode (manual|auto|planed)")
	connectorCfg.SetDefault("update_mode", "manual")
	connectorCfg.SetCheck("update_mode", func(val interface{}) bool {
		strVal := strings.ToLower(strings.TrimSpace(val.(string)))
		return map[string]bool{"manual": true, "auto": true, "planed": true}[strVal]
	})

	connectorCfg.LKey("update_time", verdeter.IsStr, "", "time for planed update mode")

	connectorCfg.SetRequired("secret")
}
