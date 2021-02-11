package models

import "github.com/spf13/viper"

type ConfigurationCli struct {
}

func NewConfigurationCli() *ConfigurationCli {
	configurationCli := new(ConfigurationCli)
	return configurationCli
}

func (cc ConfigurationCli) GetEndpoint() string {
	return viper.GetString("endpoint")
}

func (cc ConfigurationCli) SetEndpoint(endpoint string) {
	viper.Set("endpoint", endpoint)
}

func (cc ConfigurationCli) GetToken() string {
	return viper.GetString("token")
}

func (cc ConfigurationCli) SetToken(token string) {
	viper.Set("token", token)
}
