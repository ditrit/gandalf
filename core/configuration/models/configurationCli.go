package models

import "github.com/spf13/viper"

type ConfigurationCli struct {
}

func NewConfigurationCli() *ConfigurationCli {
	configurationCli := new(ConfigurationCli)
	return configurationCli
}

func (cc ConfigurationCli) GetAddress() string {
	return viper.GetString("bind")
}

func (cc ConfigurationCli) SetAddress(bindAddress string) {
	viper.Set("bind", bindAddress)
}

func (cc ConfigurationCli) GetAPIPort() int {
	return viper.GetInt("api_port")
}

func (cc ConfigurationCli) SetAPIPort(apiPort int) {
	viper.Set("api_port", apiPort)
}

func (cc ConfigurationCli) GetAPIBindAddress() string {
	return viper.GetString("bind") + ":" + viper.GetString("api_port")
}

func (cc ConfigurationCli) GetDatabaseMode() string {
	return viper.GetString("database_mode")
}

func (cc ConfigurationCli) SetDatabaseMode(databaseMode string) {
	viper.Set("database_mode", databaseMode)
}

func (cc ConfigurationCli) GetTenant() string {
	return viper.GetString("database_mode")
}

func (cc ConfigurationCli) SetTenant(tenant string) {
	viper.Set("database_mode", tenant)
}

func (cc ConfigurationCli) GetModel() string {
	return viper.GetString("model")
}

func (cc ConfigurationCli) SetModel(model string) {
	viper.Set("model", model)
}

func (cc ConfigurationCli) GetCommand() string {
	return viper.GetString("command")
}

func (cc ConfigurationCli) SetCommand(command string) {
	viper.Set("command", command)
}

func (cc ConfigurationCli) GetToken() string {
	return viper.GetString("token")
}

func (cc ConfigurationCli) SetToken(token string) {
	viper.Set("token", token)
}

func (cc ConfigurationCli) GetID() int {
	return viper.GetInt("id")
}

func (cc ConfigurationCli) SetID(id int) {
	viper.Set("id", id)
}

func (cc ConfigurationCli) GetValue() string {
	return viper.GetString("value")
}

func (cc ConfigurationCli) SetValue(value string) {
	viper.Set("value", value)
}
