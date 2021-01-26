package models

type ConfigurationConnector struct {
	LogicalName         string
	Tenant              string
	BindAddress         string
	LinkAddress         string
	LogPath             string
	GRPCSocketDirectory string
	GRPCSocketBind      string
	WorkersPath         string
	Secret              string
	ConnectorType       string
	Product             string
	WorkersUrl          string
	AutoUpdateTime      string
	AutoUpdate          bool
	MaxTimeout          int64
	VersionsMajor       int8
	VersionsMinor       int8
}

func NewConfigurationConnector(logicalName, tenant, bindAddress, linkAdress, logPath, gRPCSocketDirectory, workersPath, secret, connectorType, product, workersUrl, autoUpdateTime string, autoUpdate bool, maxTimeout int64, versionsMajor, versionsMinor int8) *ConfigurationConnector {
	configurationConnector := new(ConfigurationConnector)
	configurationConnector.LogicalName = logicalName
	configurationConnector.Tenant = tenant
	configurationConnector.BindAddress = bindAddress
	configurationConnector.LinkAddress = linkAdress
	configurationConnector.LogPath = logPath
	configurationConnector.GRPCSocketDirectory = gRPCSocketDirectory
	configurationConnector.WorkersPath = workersPath
	configurationConnector.Secret = secret
	configurationConnector.ConnectorType = connectorType
	configurationConnector.Product = product
	configurationConnector.WorkersUrl = workersUrl
	configurationConnector.AutoUpdateTime = autoUpdateTime
	configurationConnector.AutoUpdate = autoUpdate
	configurationConnector.MaxTimeout = maxTimeout
	configurationConnector.VersionsMajor = versionsMajor
	configurationConnector.VersionsMinor = versionsMinor

	return configurationConnector
}
