package models

type ConfigurationAggregator struct {
	LogicalName string
	Tenant      string
	BindAddress string
	LinkAddress string
	LogPath     string
	Secret      string
}

func NewConfigurationAggregator(logicalName, tenant, bindAddress, linkAddress, logPath, secret string) *ConfigurationAggregator {
	configurationAggregator := new(ConfigurationAggregator)
	configurationAggregator.LogicalName = logicalName
	configurationAggregator.Tenant = tenant
	configurationAggregator.BindAddress = bindAddress
	configurationAggregator.LinkAddress = linkAddress
	configurationAggregator.LogPath = logPath
	configurationAggregator.Secret = secret

	return configurationAggregator
}
