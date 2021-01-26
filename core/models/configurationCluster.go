package models

type ConfigurationCluster struct {
	LogicalName         string
	BindAddress         string
	JoinAddress         string
	LogPath             string
	DatabasePath        string
	DatabaseName        string
	DatabaseBindAddress string
	DatabaseHttpAddress string
	Secret              string
}

func NewConfigurationCluster(logicalName, bindAddress, joinAddress, logPath, databasePath, databaseName, databaseBindAddress, databaseHttpAddress, secret string) *ConfigurationCluster {
	configurationCluster := new(ConfigurationCluster)
	configurationCluster.LogicalName = logicalName
	configurationCluster.BindAddress = bindAddress
	configurationCluster.JoinAddress = joinAddress
	configurationCluster.LogPath = logPath
	configurationCluster.DatabasePath = databasePath
	configurationCluster.DatabaseName = databaseName
	configurationCluster.DatabaseBindAddress = databaseBindAddress
	configurationCluster.DatabaseHttpAddress = databaseHttpAddress
	configurationCluster.Secret = secret

	return configurationCluster
}
