package models

type ConfigurationConnector struct {
	LogicalName         string //logical
	Tenant              string //logical
	BindAddress         string
	LinkAddress         string
	LogPath             string
	GRPCSocketDirectory string
	GRPCSocketBind      string //computed value = ConnectorType + Product + Hash
	WorkersPath         string
	Secret              string //logical
	ConnectorType       string //logical
	Product             string //logical
	WorkersUrl          string //logical
	AutoUpdateTime      string //logical
	AutoUpdate          bool
	MaxTimeout          int64            //logical
	Versions            []models.Version //logical
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
