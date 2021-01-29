package models

type ConfigurationCluster struct {
	LogicalName         string //logical
	BindAddress         string
	JoinAddress         string
	LogPath             string
	DatabasePath        string
	DatabaseName        string
	DatabaseBindAddress string
	DatabaseHttpAddress string
	Secret              string //logical
	MaxTimeout          int64  //logical
}

func NewConfigurationCluster(logicalName, bindAddress, joinAddress, logPath, databasePath, databaseName, databaseBindAddress, databaseHttpAddress, secret string, maxTimeout int64) *ConfigurationCluster {
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
	configurationCluster.MaxTimeout = maxTimeout

	return configurationCluster
}
