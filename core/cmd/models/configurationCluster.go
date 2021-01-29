package models

import (
	"github.com/spf13/viper"
)

type ConfigurationCluster struct {
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

func (cc ConfigurationCluster) GetLogPath() string {
	return viper.GetString("logicalName")
}

func (cc ConfigurationCluster) SetLogPath(logPath string) {
	viper.Set("", logPath)
}

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
	return viper.GetString("bind")
}

func (cc ConfigurationCluster) SetDatabaseBindAddress(databaseBindAddress string) {
	viper.Set("bind", databaseBindAddress)
}

func (cc ConfigurationCluster) GetDatabaseHttpAddress() string {
	return viper.GetString("logicalName")
}

func (cc ConfigurationCluster) SetDatabaseHttpAddress(databaseHttpAddress string) {
	viper.Set("", databaseHttpAddress)
}

func (cc ConfigurationCluster) GetSecret() string {
	return viper.GetString("secret")
}

func (cc ConfigurationCluster) SetSecret(secret string) {
	viper.Set("secret", secret)
}

func (cc ConfigurationCluster) GetMaxTimeout() int64 {
	return viper.Getint64("logicalName")
}

func (cc ConfigurationCluster) SetMaxTimeout(maxTimeout int64) {
	viper.Set("", maxTimeout)
}
