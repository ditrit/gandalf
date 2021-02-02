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

func (ca ConfigurationAggregator) GetBindAddress() string {
	return viper.GetString("bind")
}

func (ca ConfigurationAggregator) SetBindAddress(bindAdress string) {
	viper.Set("bind", bindAdress)

}

func (ca ConfigurationAggregator) GetLinkAddress() string {
	return viper.GetInt64("cluster")
}

func (ca ConfigurationAggregator) SetLinkAddress(linkAddress string) {
	viper.Set("cluster", linkAddress)

}

/* func (ca ConfigurationAggregator) GetLogPath() string {
	return viper.GetString("logicalName")
}

func (ca ConfigurationAggregator) SetLogPath(logPath string) {
	viper.Set("", logPath)

} */

func (ca ConfigurationAggregator) GetSecret() string {
	return viper.GetString("secret")
}

func (ca ConfigurationAggregator) SetSecret(secret string) {
	viper.Set("secret", secret)

}

func (ca ConfigurationAggregator) GetMaxTimeout() int64 {
	return viper.GetString("max_timeout")
}

func (ca ConfigurationAggregator) SetMaxTimeout(maxTimeout int64) {
	viper.Set("max_timeout", maxTimeout)

}
