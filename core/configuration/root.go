// Package configuration is part of Gandalf
package configuration

import (
	"github.com/ditrit/gandalf/verdeter"
	"github.com/spf13/viper"
)

var rootCfg = verdeter.NewConfigCmd(
	"gandalf",

	"Gandalf is a tool to easily assemble DevOps software factories",

	`Gandalf stands for Gandalf is A Natural Devops Application Life-cycle Framework.
	Gandalf components and multi language abstract workflow primitives allow you to build or modify in few minutes a DevOps software factory in an efficient, highly secured, enterprise grade way.
	Gandalf philosophy is not to replace or to be a additional layer on existing tools. It only provides a way to easily assemble tools and make them efficiently communicate.`,

	func(cfg *verdeter.ConfigCmd, args []string) {
		mode := viper.GetString("mode")
		if mode == "cli" {
			cfg.CallSubRun(mode)
		}
		startCfg.CallSubRun(mode)
	})

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCfg.Execute()
}

func init() {
	// cobra.OnInitialize(initConfig)
	rootCfg.Initialize()

	rootCfg.GKey("config_dir", verdeter.IsStr, "", "Path to the config directory")
	rootCfg.SetNormalize("config_dir", func(val interface{}) interface{} {
		strval, ok := val.(string)
		if ok {
			if strval != "" {
				lastChar := strval[len(strval)-1:]
				if lastChar != "/" {
					return strval + "/"
				}
				return strval
			}
		}
		return nil
	})

	rootCfg.GKey("config_file", verdeter.IsStr, "", "Path to the config file")

	rootCfg.GKey("log_dir", verdeter.IsStr, "", "directory to store gandalf logfile")
	//rootCfg.SetDefault("log_dir", "/var/log/")
	rootCfg.SetDefault("log_dir", "/var/log/gandalf/")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	verdeter.InitConfig(rootCfg)
}
