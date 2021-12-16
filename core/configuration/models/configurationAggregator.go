package models

import (
	"github.com/spf13/viper"
)

type ConfigurationAggregator struct {
}

func NewConfigurationAggregator() *ConfigurationAggregator {
	configurationAggregator := new(ConfigurationAggregator)
	return configurationAggregator
}

func (ca ConfigurationAggregator) GetOffset() int {
	return viper.GetInt("offset")
}

func (ca ConfigurationAggregator) GetLogicalName() string {
	return viper.GetString("lname")
}

func (ca ConfigurationAggregator) SetLogicalName(logicalName string) {
	viper.Set("lname", logicalName)
}

func (ca ConfigurationAggregator) GetTenant() string {
	return viper.GetString("tenant")
}

func (ca ConfigurationAggregator) SetTenant(tenant string) {
	viper.Set("tenant", tenant)

}

func (cc ConfigurationAggregator) GetAddress() string {
	return viper.GetString("bind")
}

func (cc ConfigurationAggregator) SetAddress(bindAddress string) {
	viper.Set("bind", bindAddress)
}

func (cc ConfigurationAggregator) GetPort() string {
	return viper.GetString("port")
}

func (cc ConfigurationAggregator) SetPort(bindAddress string) {
	viper.Set("port", bindAddress)
}

func (cc ConfigurationAggregator) GetBindAddress() string {
	return viper.GetString("bind") + ":" + viper.GetString("port")
}

func (ca ConfigurationAggregator) GetLinkAddress() string {
	return viper.GetString("cluster")
}

func (ca ConfigurationAggregator) SetLinkAddress(linkAddress string) {
	viper.Set("cluster", linkAddress)

}

func (ca ConfigurationAggregator) GetAPIPort() int {
	return viper.GetInt("api_port")
}

func (ca ConfigurationAggregator) SetAPIPort(apiPort int) {
	viper.Set("api_port", apiPort)
}

func (ca ConfigurationAggregator) GetAPIPath() string {
	return viper.GetString("api_path")
}

func (ca ConfigurationAggregator) SetAPIPath(apiPath string) {
	viper.Set("api_path", apiPath)
}

func (ca ConfigurationAggregator) GetAPIBindAddress() string {
	return viper.GetString("bind") + ":" + viper.GetString("api_port")
}

func (ca ConfigurationAggregator) GetSecret() string {
	return viper.GetString("secret")
}

func (ca ConfigurationAggregator) SetSecret(secret string) {
	viper.Set("secret", secret)

}

func (ca ConfigurationAggregator) GetMaxTimeout() int64 {
	return viper.GetInt64("max_timeout")
}

func (ca ConfigurationAggregator) SetMaxTimeout(maxTimeout int64) {
	viper.Set("max_timeout", maxTimeout)

}

func (ca ConfigurationAggregator) GetCertsPath() string {
	return viper.GetString("cert_dir")
}

func (ca ConfigurationAggregator) SetCertsPath(certsPath string) {
	viper.Set("cert_dir", certsPath)
}

func (ca ConfigurationAggregator) GetConfigPath() string {
	return viper.GetString("config_dir")
}

func (ca ConfigurationAggregator) SetConfigPath(configPath string) {
	viper.Set("config_dir", configPath)
}
