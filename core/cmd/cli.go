/*
Copyright © 2020 DitRit community <contact@ditrit.io>
This file is part of Gandalf
*/

// Package cmd manages commands and configuration
package cmd

import (
	"fmt"

	"github.com/spf13/viper"
)

// cliCmd represents the cli command
var cliCfg = NewConfigCmd(
	"cli",
	"Launch gandalf in 'cli' mode.",
	`Gandalf is launched as CLI (Command Line Interface) to interact with a Gandalf system.`,
	func(cfg *ConfigCmd, args []string) {
		fmt.Println("cli called")
		fmt.Println("IsSet('config') =", viper.IsSet("config"), ", value('config') = ", viper.Get("config"))
	})

func init() {
	rootCfg.AddConfig(cliCfg)
}
