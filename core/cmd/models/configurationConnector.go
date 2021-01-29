package models

import (
	"github.com/spf13/viper"
)

type ConfigurationConnector struct {
}

func (cc ConfigurationConnector) GetLogicalName() string {
	return viper.GetString("lname")
}

func (cc ConfigurationConnector) SetLogicalName(logicalName string) {
	viper.Set("lname", logicalName)
}

func (cc ConfigurationConnector) GetTenant() string {
	return viper.GetString("tenant")
}

func (cc ConfigurationConnector) SetTenant(tenant string) {
	viper.Set("tenant", tenant)
}

func (cc ConfigurationConnector) GetBindAddress() string {
	return viper.GetString("bind")
}

func (cc ConfigurationConnector) SetBindAddress(bindAddress string) {
	viper.Set("bind", bindAddress)
}

func (cc ConfigurationConnector) GetLinkAddress() string {
	return viper.GetString("aggregator")
}

func (cc ConfigurationConnector) SetLinkAddress(linkAddress string) {
	viper.Set("aggregator", linkAddress)
}

func (cc ConfigurationConnector) GetLogPath() string {
	return viper.GetString("logicalName")
}

func (cc ConfigurationConnector) SetLogPath(logPath string) {
	viper.Set("", logPath)
}

func (cc ConfigurationConnector) GetGRPCSocketDirectory() string {
	return viper.GetString("logicalName")
}

func (cc ConfigurationConnector) SetGRPCSocketDirectory(gRPCSocketDirectory string) {
	viper.Set("", gRPCSocketDirectory)
}

func (cc ConfigurationConnector) GetGRPCSocketBind() string {
	return viper.GetString("logicalName")
}

func (cc ConfigurationConnector) SetGRPCSocketBind(gRPCSocketBind string) {
	viper.Set("", gRPCSocketBind)
}

func (cc ConfigurationConnector) GetWorkersPath() string {
	return viper.GetString("workers")
}

func (cc ConfigurationConnector) SetWorkersPath(workersPath string) {
	viper.Set("workers", workersPath)
}

func (cc ConfigurationConnector) GetSecret() string {
	return viper.GetString("secret")
}

func (cc ConfigurationConnector) SetSecret(secret string) {
	viper.Set("secret", secret)
}

func (cc ConfigurationConnector) GetConnectorType() string {
	return viper.GetString("class")
}

func (cc ConfigurationConnector) SetConnectorType(connectorType string) {
	viper.Set("class", connectorType)
}

func (cc ConfigurationConnector) GetProduct() string {
	return viper.GetString("product")
}

func (cc ConfigurationConnector) SetProduct(product string) {
	viper.Set("product", product)
}

func (cc ConfigurationConnector) GetWorkersUrl() string {
	return viper.GetString("logicalName")
}

func (cc ConfigurationConnector) SetWorkersUrl(workersUrl string) {
	viper.Set("", workersUrl)
}

func (cc ConfigurationConnector) GetAutoUpdateTime() string {
	return viper.GetString("update_time")
}

func (cc ConfigurationConnector) SetAutoUpdateTime(autoUpdateTime string) {
	viper.Set("update_time", autoUpdateTime)
}

func (cc ConfigurationConnector) GetAutoUpdate() string {
	return viper.GetString("update_mode")
}

func (cc ConfigurationConnector) SetAutoUpdate(autoUpdate string) {
	viper.Set("update_mode", autoUpdate)
}

func (cc ConfigurationConnector) GetMaxTimeout() int64 {
	return viper.Getint64("max_timeout")
}

func (cc ConfigurationConnector) SetMaxTimeout(maxTimeout int64) {
	viper.Set("max_timeout", maxTimeout)
}

//TODO REVOIR
func (cc ConfigurationConnector) GetVersions() string {
	return viper.GetString("logicalName")
}

func (cc ConfigurationConnector) SetVersions(logicalName string) {
	viper.Set("", logicalName)
}
