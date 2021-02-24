// Package cmd is part of Gandalf
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
		cfg.CallSubRun(mode)
	})

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCfg.Execute()
}

func init() {
	// cobra.OnInitialize(initConfig)
	rootCfg.Initialize()

	rootCfg.Key("offset", verdeter.IsInt, "", "Offset used in case of multiple Gandals instances hosted on the same host")

	// flags common to all commands
	rootCfg.Key("lname", verdeter.IsStr, "l", "logical name (non empty value required)")
	rootCfg.SetCheck("lname", verdeter.CheckNotEmpty)
	rootCfg.SetNormalize("lname", verdeter.TrimToLower)

	rootCfg.Key("config_dir", verdeter.IsStr, "", "Path to the config directory")
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

	rootCfg.Key("config_file", verdeter.IsStr, "", "Path to the config file")

	rootCfg.Key("secret", verdeter.IsStr, "", "Path to the secret (absolute or relative to the configuration directory)")
	rootCfg.SetCheck("secret", verdeter.CheckNotEmpty)

	rootCfg.Key("max_timeout", verdeter.IsInt, "", "maximum timeout of the connector")
	rootCfg.SetDefault("max_timeout", 100)

	rootCfg.Key("bind", verdeter.IsStr, "", "Address to bind (default is 127.0.0.1)")
	rootCfg.SetDefault("bind", "127.0.0.1")
	rootCfg.SetNormalize("bind", verdeter.TrimToLower)
	//If no offset use localhost else if local address is unique return it
	rootCfg.SetComputedValue("bind",
		func() interface{} {
			if verdeter.GetOffset() > 0 {
				return "127.0.0.1"

			}
			address, err := verdeter.GetUniqueInterface()
			if err != nil {
				return nil
			}
			return address

		})

	rootCfg.Key("port", verdeter.IsInt, "", "Port to bind (default is 9099 + offset if defined)")
	//rootCfg.SetDefault("port", 9099+verdeter.GetOffset())
	rootCfg.SetCheck("port", verdeter.CheckTCPHighPort)
	rootCfg.SetComputedValue("port",
		func() interface{} {
			return 9099 + verdeter.GetOffset()
		})

	rootCfg.Key("cert_dir", verdeter.IsStr, "", "path of the certificates directory (absolute or relative to the configuration directory)")
	rootCfg.SetDefault("cert_dir", "/etc/gandalf/certs/")
	rootCfg.SetComputedValue("cert_dir",
		func() interface{} {
			return viper.GetString("config_dir") + "certs/"
		})

	rootCfg.Key("cert_pem", verdeter.IsStr, "", "path of the TLS certificate (absolute or relative to the certificates directory)")
	rootCfg.SetDefault("cert_pem", verdeter.ExpandPath(viper.GetString("cert_dir"), "cert.pem"))
	rootCfg.SetNormalize("cert_pem", func(val interface{}) interface{} {
		return verdeter.ExpandPath(viper.GetString("cert_dir"), val)
	})

	rootCfg.Key("key_pem", verdeter.IsStr, "", "path of the TLS private key (absolute or relative to the certificates directory)")
	rootCfg.SetDefault("key_pem", verdeter.ExpandPath(viper.GetString("cert_dir"), "key.pem"))
	rootCfg.SetNormalize("key_pem", func(val interface{}) interface{} {
		return verdeter.ExpandPath(viper.GetString("cert_dir"), val)
	})

	rootCfg.Key("ca_cert_pem", verdeter.IsStr, "", "path of the CA certificate (absolute or relative to the certificates directory)")
	rootCfg.SetDefault("ca_cert_pem", verdeter.ExpandPath(viper.GetString("cert_dir"), "ca_cert.pem"))
	rootCfg.SetNormalize("ca_cert_pem", func(val interface{}) interface{} {
		return verdeter.ExpandPath(viper.GetString("cert_dir"), val)
	})

	rootCfg.Key("ca_key_pem", verdeter.IsStr, "", "path of the CA key (absolute or relative to the certificates directory)")
	rootCfg.SetDefault("ca_key_pem", verdeter.ExpandPath(viper.GetString("cert_dir"), "ca_key.pem"))
	rootCfg.SetNormalize("ca_key_pem", func(val interface{}) interface{} {
		return verdeter.ExpandPath(viper.GetString("cert_dir"), val)
	})
	rootCfg.Key("log_dir", verdeter.IsStr, "", "directory to store gandalf logfile")
	//rootCfg.SetDefault("log_dir", "/var/log/")
	rootCfg.SetDefault("log_dir", "/var/log/gandalf/")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	verdeter.InitConfig(rootCfg)
}
