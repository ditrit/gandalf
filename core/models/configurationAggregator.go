package models

type ConfigurationAggregator struct {
	LogicalName string
	Tenant      string
	BindAddress string
	LinkAddress string
	LogPath     string
	Secret      string
	MaxTimeout  int64
}

func NewConfigurationAggregator(logicalName, tenant, bindAddress, linkAddress, logPath, secret string, maxTimeout int64) *ConfigurationAggregator {
	configurationAggregator := new(ConfigurationAggregator)
	configurationAggregator.LogicalName = logicalName
	configurationAggregator.Tenant = tenant
	configurationAggregator.BindAddress = bindAddress
	configurationAggregator.LinkAddress = linkAddress
	configurationAggregator.LogPath = logPath
	configurationAggregator.Secret = secret
	configurationAggregator.MaxTimeout = maxTimeout

	return configurationAggregator
}
