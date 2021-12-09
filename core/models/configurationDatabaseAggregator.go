package models

type ConfigurationDatabaseAggregator struct {
	Tenant              string
	Password            string
	DatabaseBindAddress string
}

func NewConfigurationDatabaseAggregator(tenant, password, bindAddress string) *ConfigurationDatabaseAggregator {
	configurationDatabaseAggregator := new(ConfigurationDatabaseAggregator)
	configurationDatabaseAggregator.Tenant = tenant
	configurationDatabaseAggregator.Password = password
	configurationDatabaseAggregator.DatabaseBindAddress = bindAddress

	return configurationDatabaseAggregator
}
