/*
Copyright © 2020 DitRit community <contact@ditrit.io>
This file is part of Gandalf
*/

// Package cmd manages commands and configuration
package configuration

import (
	"fmt"
	"strings"

	"github.com/ditrit/gandalf/core/cli"
	"github.com/ditrit/gandalf/core/configuration/config"
	cmodels "github.com/ditrit/gandalf/core/configuration/models"

	"github.com/spf13/viper"
)

// cliCmd represents the cli command
var cliCfg = config.NewConfigCmd(
	"cli",
	"Launch gandalf in 'cli' mode.",
	`Gandalf is launched as CLI (Command Line Interface) to interact with a Gandalf system.`,
	func(cfg *config.ConfigCmd, args []string) {
		fmt.Println("cli called")
		fmt.Println("IsSet('config') =", viper.IsSet("config"), ", value('config') = ", viper.Get("config"))
		configurationCli := cmodels.NewConfigurationCli()
		cli.Cli(configurationCli)
	})

func init() {
	rootCfg.AddConfig(cliCfg)

	cliCfg.Key("api_port", config.IsInt, "", "Address to bind (default is *:9199)")
	cliCfg.SetDefault("api_port", 9199+config.GetOffset())

	cliCfg.Key("database_mode", config.IsStr, "", "database mode (gandalf|tenant)")
	cliCfg.SetDefault("database_mode", "gandalf")
	cliCfg.SetCheck("database_mode", func(val interface{}) bool {
		strVal := strings.ToLower(strings.TrimSpace(val.(string)))
		return map[string]bool{"gandalf": true, "tenant": true}[strVal]
	})

	cliCfg.Key("tenant", config.IsStr, "", "database mode (gandalf|tenant)")
	cliCfg.SetConstraint("a tenant should be provided if database_mode == tenant",
		func() bool {
			return viper.IsSet("database_mode") && viper.GetString("database_mode") == "tenant" && viper.IsSet("tenant")
		})

	cliCfg.Key("model", config.IsStr, "", "models  gandalf(authentication|cluster|tenant|role|user) || tenant(authentication|aggregator|connector|role|user)")
	cliCfg.SetCheck("model", func(val interface{}) bool {
		strVal := strings.ToLower(strings.TrimSpace(val.(string)))
		return map[string]bool{"authentication": true, "cluster": true, "tenant": true, "role": true, "user": true, "aggregator": true, "connector": true}[strVal]
	})
	//TODO REVOIR
	/* 	cliCfg.SetConstraint("a id should be provided if command == (delete|update|read)",
	func() bool {
		return viper.IsSet("database_mode") || viper.GetString("database_mode") == "gandalf" || viper.GetString("command") == "update" || viper.GetString("command") == "read" || viper.IsSet("id")
	}) */

	cliCfg.Key("command", config.IsStr, "", "command  (list|read|create|update|delete|upload)")
	cliCfg.SetCheck("command", func(val interface{}) bool {
		strVal := strings.ToLower(strings.TrimSpace(val.(string)))
		return map[string]bool{"list": true, "read": true, "create": true, "update": true, "delete": true, "upload": true}[strVal]
	})

	cliCfg.Key("token", config.IsStr, "", "")
	cliCfg.SetConstraint("a token should be provided if model != authenticaion",
		func() bool {
			return viper.IsSet("model") && viper.GetString("model") != "authentication" && viper.IsSet("token")
		})

	cliCfg.Key("id", config.IsStr, "", "id")
	cliCfg.SetConstraint("a id should be provided if command == (delete|update|read)",
		func() bool {
			return viper.IsSet("command") && (viper.GetString("command") == "delete" || viper.GetString("command") == "update" || viper.GetString("command") == "read") && viper.IsSet("id")
		})

	cliCfg.Key("value", config.IsStr, "", "json")
	cliCfg.SetConstraint("a value should be provided if command == (create|update)",
		func() bool {
			return viper.IsSet("command") && (viper.GetString("command") == "create" || viper.GetString("command") == "update") && viper.IsSet("value")
		})
}
