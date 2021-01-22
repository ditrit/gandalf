package models

type ConfigurationInstanceAggregator struct {
	BindAddress string
	LinkAddress string
	LogPath     string
	Secret      string
}

func NewConfigurationInstanceAggregator(bindAddress, linkAddress, logPath, secret string) *ConfigurationInstanceAggregator {
	configurationInstanceAggregator := new(ConfigurationInstanceAggregator)
	configurationInstanceAggregator.BindAddress = bindAddress
	configurationInstanceAggregator.LinkAddress = linkAddress
	configurationInstanceAggregator.LogPath = logPath
	configurationInstanceAggregator.Secret = secret

	return configurationInstanceAggregator
}
