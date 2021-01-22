package models

type ConfigurationInstanceConnector struct {
	BindAddress         string
	LinkAddress         string
	LogPath             string
	GRPCSocketDirectory string
	GRPCSocketBind      string
	WorkersPath         string
	Secret              string
}

func NewConfigurationInstanceConnector(bindAddress, linkAdress, logPath, gRPCSocketDirectory, workersPath, secret string) *ConfigurationInstanceConnector {
	configurationInstanceConnector := new(ConfigurationInstanceConnector)
	configurationInstanceConnector.BindAddress = bindAddress
	configurationInstanceConnector.LinkAddress = linkAdress
	configurationInstanceConnector.LogPath = logPath
	configurationInstanceConnector.GRPCSocketDirectory = gRPCSocketDirectory
	configurationInstanceConnector.WorkersPath = workersPath
	configurationInstanceConnector.Secret = secret

	return configurationInstanceConnector
}
