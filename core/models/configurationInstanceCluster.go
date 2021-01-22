package models

type ConfigurationInstanceCluster struct {
	BindAddress         string
	JoinAddress         string
	LogPath             string
	DatabasePath        string
	DatabaseName        string
	DatabaseBindAddress string
	DatabaseHttpAddress string
	Secret              string
}

func NewConfigurationInstanceCluster(bindAddress, joinAddress, logPath, databasePath, databaseName, databaseBindAddress, databaseHttpAddress, secret string) *ConfigurationInstanceCluster {
	configurationInstanceCluster := new(ConfigurationInstanceCluster)
	configurationInstanceCluster.BindAddress = bindAddress
	configurationInstanceCluster.JoinAddress = joinAddress
	configurationInstanceCluster.LogPath = logPath
	configurationInstanceCluster.DatabasePath = databasePath
	configurationInstanceCluster.DatabaseName = databaseName
	configurationInstanceCluster.DatabaseBindAddress = databaseBindAddress
	configurationInstanceCluster.DatabaseHttpAddress = databaseHttpAddress
	configurationInstanceCluster.Secret = secret

	return configurationInstanceCluster
}
