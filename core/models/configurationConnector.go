package models

type ConfigurationConnector struct {
	LogicalName     string
	GRPCBindAddress string
	ConnectorType   string
	Product         string
	WorkersUrl      string
	Workers         string
	AutoUpdateTime  string
	AutoUpdate      bool
	MaxTimeout      int64
	VersionsMajor   int64
	VersionsMinor   int64
}
