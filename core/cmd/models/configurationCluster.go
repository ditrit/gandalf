package models

import (
	"github.com/spf13/viper"
)

type ConfigurationCluster struct {
}

func NewConfigurationCluster() *ConfigurationCluster {
	configurationCluster := new(ConfigurationCluster)
	return configurationCluster
}

func (cc ConfigurationCluster) GetLogicalName() string {
	return viper.GetString("lname")
}

func (cc ConfigurationCluster) SetLogicalName(logicalName string) {
	viper.Set("lname", logicalName)
}

func (cc ConfigurationCluster) GetBindAddress() string {
	return viper.GetString("logicalName")
}

func (cc ConfigurationCluster) SetBindAddress(bindAddress string) {
	viper.Set("", bindAddress)
}

func (cc ConfigurationCluster) GetJoinAddress() string {
	return viper.GetString("logicalName")
}

func (cc ConfigurationCluster) SetJoinAddress(joinAddress string) {
	viper.Set("", joinAddress)
}

/* func (cc ConfigurationCluster) GetLogPath() string {
	return viper.GetString("logicalName")
}

func (cc ConfigurationCluster) SetLogPath(logPath string) {
	viper.Set("", logPath)
} */

func (cc ConfigurationCluster) GetDatabasePath() string {
	return viper.GetString("db_path")
}

func (cc ConfigurationCluster) SetDatabasePath(databasePath string) {
	viper.Set("db_path", databasePath)
}

func (cc ConfigurationCluster) GetDatabaseName() string {
	return viper.GetString("db_nodename")
}

func (cc ConfigurationCluster) SetDatabaseName(databaseName string) {
	viper.Set("db_nodename", databaseName)
}

func (cc ConfigurationCluster) GetDatabaseBindAddress() string {
	return viper.GetString("db_bind")
}

func (cc ConfigurationCluster) SetDatabaseBindAddress(databaseBindAddress string) {
	viper.Set("db_bind", databaseBindAddress)
}

func (cc ConfigurationCluster) GetDatabaseHttpAddress() string {
	return viper.GetString("db_http_bind")
}

func (cc ConfigurationCluster) SetDatabaseHttpAddress(databaseHttpAddress string) {
	viper.Set("db_http_bind", databaseHttpAddress)
}

func (cc ConfigurationCluster) GetSecret() string {
	return viper.GetString("secret")
}

func (cc ConfigurationCluster) SetSecret(secret string) {
	viper.Set("secret", secret)
}

func (cc ConfigurationCluster) GetMaxTimeout() int64 {
	return viper.GetInt64("max_timeout")
}

func (cc ConfigurationCluster) SetMaxTimeout(maxTimeout int64) {
	viper.Set("max_timeout", maxTimeout)
}
