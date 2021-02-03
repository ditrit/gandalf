/*
Copyright © 2020 DitRit community <contact@ditrit.io>
This file is part of Gandalf
*/

// Package cmd manages commands and configuration
package configuration

import (
	"fmt"

	"github.com/ditrit/gandalf/core/configuration/config"

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
	})

func init() {
	rootCfg.AddConfig(cliCfg)
}
