package models

import (
	"gandalf/core/models"

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
	return viper.GetInt64("max_timeout")
}

func (ca ConfigurationAggregator) SetMaxTimeout(maxTimeout int64) {
	viper.Set("max_timeout", maxTimeout)

}

func (ca ConfigurationAggregator) ConfigurationToDatabase() *models.ConfigurationLogicalAggregator {
	configurationLogicalAggregator := new(models.ConfigurationLogicalAggregator)
	configurationLogicalAggregator.LogicalName = ca.GetLogicalName()
	configurationLogicalAggregator.Tenant = ca.GetTenant()
	configurationLogicalAggregator.Secret = ca.GetSecret()
	configurationLogicalAggregator.MaxTimeout = ca.GetMaxTimeout()
	return configurationLogicalAggregator
}

func (ca ConfigurationAggregator) DatabaseToConfiguration(configurationLogicalAggregator *models.ConfigurationLogicalAggregator) {
	ca.SetLogicalName(configurationLogicalAggregator.LogicalName)
	ca.SetTenant(configurationLogicalAggregator.Tenant)
	ca.SetSecret(configurationLogicalAggregator.Secret)
	ca.SetMaxTimeout(configurationLogicalAggregator.MaxTimeout)
}
