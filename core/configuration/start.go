// Package configuration manages commands and configuration
package configuration

import (
	"github.com/ditrit/gandalf/verdeter"
	"github.com/spf13/viper"
)

var startCfg = verdeter.NewConfigCmd(
	"start",

	"Launch gandalf as a server in 'cluster', 'aggregator' or 'connector' mode.",

	`'start' command allows to start the Gandalf binary as a componant instance of a Gandalf distributed system. 
	The component (depending of the used subcommand or the 'mode' config key) can be a :
	   - a member of a Gandalf cluster,
	   - a server instance of a Gandalf aggregator,
	   - a server instance of a Gandalf connector.
	`,

	func(cfg *verdeter.ConfigCmd, args []string) {
		mode := viper.GetString("mode")
		cfg.CallSubRun(mode)
	})

func init() {
	rootCfg.AddConfig(startCfg)

	startCfg.GKey("offset", verdeter.IsInt, "", "Offset used in case of multiple Gandals instances hosted on the same host")

	// flags common to all commands
	startCfg.GKey("lname", verdeter.IsStr, "l", "logical name (non empty value required)")
	startCfg.SetCheck("lname", verdeter.CheckNotEmpty)
	startCfg.SetNormalize("lname", verdeter.TrimToLower)
	startCfg.SetRequired("lname")

	startCfg.GKey("config_dir", verdeter.IsStr, "", "Path to the config directory")
	startCfg.SetNormalize("config_dir", func(val interface{}) interface{} {
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

	startCfg.GKey("config_file", verdeter.IsStr, "", "Path to the config file")

	startCfg.GKey("secret", verdeter.IsStr, "", "Path to the secret (absolute or relative to the configuration directory)")
	startCfg.SetCheck("secret", verdeter.CheckNotEmpty)

	startCfg.GKey("retry_time", verdeter.IsInt, "", "Time to wait in second between each error")
	startCfg.SetDefault("retry_time", 5)

	startCfg.GKey("retry_max", verdeter.IsInt, "", "Max number of retry in case of error")
	startCfg.SetDefault("retry_max", 5)

	startCfg.GKey("max_timeout", verdeter.IsInt, "", "maximum timeout of the connector")
	startCfg.SetDefault("max_timeout", 1000)

	startCfg.GKey("bind", verdeter.IsStr, "", "Address to bind (default is 127.0.0.1)")
	startCfg.SetDefault("bind", "127.0.0.1")
	startCfg.SetNormalize("bind", verdeter.TrimToLower)
	//If no offset use localhost else if local address is unique return it
	startCfg.SetComputedValue("bind",
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

	startCfg.GKey("port", verdeter.IsInt, "", "Port to bind (default is 9099 + offset if defined)")
	startCfg.SetCheck("port", verdeter.CheckTCPHighPort)
	startCfg.SetComputedValue("port",
		func() interface{} {
			return 9099 + verdeter.GetOffset()
		})

	startCfg.GKey("cert_dir", verdeter.IsStr, "", "path of the certificates directory (absolute or relative to the configuration directory)")
	startCfg.SetDefault("cert_dir", "/etc/gandalf/certs/")
	startCfg.SetComputedValue("cert_dir",
		func() interface{} {
			return viper.GetString("config_dir") + "certs/"
		})

	startCfg.GKey("cert_pem", verdeter.IsStr, "", "path of the TLS certificate (absolute or relative to the certificates directory)")
	startCfg.SetDefault("cert_pem", verdeter.ExpandPath(viper.GetString("cert_dir"), "cert.pem"))
	startCfg.SetNormalize("cert_pem", func(val interface{}) interface{} {
		return verdeter.ExpandPath(viper.GetString("cert_dir"), val)
	})

	startCfg.GKey("key_pem", verdeter.IsStr, "", "path of the TLS private key (absolute or relative to the certificates directory)")
	startCfg.SetDefault("key_pem", verdeter.ExpandPath(viper.GetString("cert_dir"), "key.pem"))
	startCfg.SetNormalize("key_pem", func(val interface{}) interface{} {
		return verdeter.ExpandPath(viper.GetString("cert_dir"), val)
	})

	startCfg.GKey("ca_cert_pem", verdeter.IsStr, "", "path of the CA certificate (absolute or relative to the certificates directory)")
	startCfg.SetDefault("ca_cert_pem", verdeter.ExpandPath(viper.GetString("cert_dir"), "ca_cert.pem"))
	startCfg.SetNormalize("ca_cert_pem", func(val interface{}) interface{} {
		return verdeter.ExpandPath(viper.GetString("cert_dir"), val)
	})

	startCfg.GKey("ca_key_pem", verdeter.IsStr, "", "path of the CA key (absolute or relative to the certificates directory)")
	startCfg.SetDefault("ca_key_pem", verdeter.ExpandPath(viper.GetString("cert_dir"), "ca_key.pem"))
	startCfg.SetNormalize("ca_key_pem", func(val interface{}) interface{} {
		return verdeter.ExpandPath(viper.GetString("cert_dir"), val)
	})
	startCfg.GKey("log_dir", verdeter.IsStr, "", "directory to store gandalf logfile")
	startCfg.SetDefault("log_dir", "/var/log/gandalf/")
}
