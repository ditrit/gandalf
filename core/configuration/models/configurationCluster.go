package models

import (
	"github.com/ditrit/gandalf/core/models"

	"github.com/spf13/viper"
)

type ConfigurationCluster struct {
}

func NewConfigurationCluster() *ConfigurationCluster {
	configurationCluster := new(ConfigurationCluster)
	return configurationCluster
}

func (cc ConfigurationCluster) GetOffset() int {
	return viper.GetInt("offset")
}

func (cc ConfigurationCluster) GetLogicalName() string {
	return viper.GetString("lname")
}

func (cc ConfigurationCluster) SetLogicalName(logicalName string) {
	viper.Set("lname", logicalName)
}

func (cc ConfigurationCluster) GetAddress() string {
	return viper.GetString("bind")
}

func (cc ConfigurationCluster) SetAddress(bindAddress string) {
	viper.Set("bind", bindAddress)
}

func (cc ConfigurationCluster) GetPort() string {
	return viper.GetString("port")
}

func (cc ConfigurationCluster) SetPort(bindAddress string) {
	viper.Set("port", bindAddress)
}

func (cc ConfigurationCluster) GetBindAddress() string {
	return viper.GetString("bind") + ":" + viper.GetString("port")
}

func (cc ConfigurationCluster) GetJoinAddress() string {
	return viper.GetString("join")
}

func (cc ConfigurationCluster) SetJoinAddress(joinAddress string) {
	viper.Set("join", joinAddress)
}

func (cc ConfigurationCluster) GetAPIPort() int {
	return viper.GetInt("api_port")
}

func (cc ConfigurationCluster) SetAPIPort(apiPort int) {
	viper.Set("api_port", apiPort)
}

func (cc ConfigurationCluster) GetAPIBindAddress() string {
	return viper.GetString("bind") + ":" + viper.GetString("api_port")
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

func (cc ConfigurationCluster) GetDatabasePort() int {
	return viper.GetInt("db_port")
}

func (cc ConfigurationCluster) SetDatabasePort(databaseBindAddress int) {
	viper.Set("db_port", databaseBindAddress)
}

func (cc ConfigurationCluster) GetDatabaseBindAddress() string {
	return viper.GetString("bind") + ":" + viper.GetString("db_port")
}

func (cc ConfigurationCluster) GetDatabaseHttpPort() int {
	return viper.GetInt("db_http_port")
}

func (cc ConfigurationCluster) SetDatabaseHttpPort(databaseHttpAddress int) {
	viper.Set("db_http_port", databaseHttpAddress)
}

func (cc ConfigurationCluster) GetDatabaseHttpAddress() string {
	return viper.GetString("bind") + ":" + viper.GetString("db_http_port")
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

func (cc ConfigurationCluster) ConfigurationToDatabase() *models.ConfigurationLogicalCluster {
	configurationLogicalCluster := new(models.ConfigurationLogicalCluster)
	configurationLogicalCluster.LogicalName = cc.GetLogicalName()
	configurationLogicalCluster.Secret = cc.GetSecret()
	configurationLogicalCluster.MaxTimeout = cc.GetMaxTimeout()
	return configurationLogicalCluster
}

func (cc ConfigurationCluster) DatabaseToConfiguration(configurationLogicalCluster *models.ConfigurationLogicalCluster) {
	cc.SetLogicalName(configurationLogicalCluster.LogicalName)
	cc.SetSecret(configurationLogicalCluster.Secret)
	cc.SetMaxTimeout(configurationLogicalCluster.MaxTimeout)
}
