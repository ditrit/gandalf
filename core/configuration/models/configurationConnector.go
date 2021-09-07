package models

import (
	"strconv"
	"strings"

	"github.com/ditrit/gandalf/verdeter"

	"github.com/ditrit/gandalf/core/models"

	"github.com/spf13/viper"
)

type ConfigurationConnector struct {
	cfg *verdeter.ConfigCmd
}

func NewConfigurationConnector(cfg *verdeter.ConfigCmd) *ConfigurationConnector {
	configurationConnector := new(ConfigurationConnector)
	configurationConnector.cfg = cfg
	return configurationConnector
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

func (cc ConfigurationConnector) GetAddress() string {
	return viper.GetString("bind")
}

func (cc ConfigurationConnector) SetAddress(bindAddress string) {
	viper.Set("bind", bindAddress)
}

func (cc ConfigurationConnector) GetPort() string {
	return viper.GetString("port")
}

func (cc ConfigurationConnector) SetPort(bindAddress string) {
	viper.Set("port", bindAddress)
}

func (cc ConfigurationConnector) GetBindAddress() string {
	return viper.GetString("bind") + ":" + viper.GetString("port")
}

func (cc ConfigurationConnector) GetLinkAddress() string {
	return viper.GetString("aggregator")
}

func (cc ConfigurationConnector) SetLinkAddress(linkAddress string) {
	viper.Set("aggregator", linkAddress)
}

/* func (cc ConfigurationConnector) GetLogPath() string {
	return viper.GetString("logicalName")
}

func (cc ConfigurationConnector) SetLogPath(logPath string) {
	viper.Set("", logPath)
} */

func (cc ConfigurationConnector) GetGRPCSocketDirectory() string {
	return viper.GetString("grpc_dir")
}

func (cc ConfigurationConnector) SetGRPCSocketDirectory(gRPCSocketDirectory string) {
	viper.Set("grpc_dir", gRPCSocketDirectory)
}

func (cc ConfigurationConnector) GetGRPCSocketBind() string {
	return viper.GetString("grpc_bind")
}

func (cc ConfigurationConnector) SetGRPCSocketBind(gRPCSocketBind string) {
	viper.Set("grpc_bind", gRPCSocketBind)
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
	return viper.GetString("workers_url")
}

func (cc ConfigurationConnector) SetWorkersUrl(workersUrl string) {
	viper.Set("workers_url", workersUrl)
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
	return viper.GetInt64("max_timeout")
}

func (cc ConfigurationConnector) SetMaxTimeout(maxTimeout int64) {
	viper.Set("max_timeout", maxTimeout)
}

func (cc ConfigurationConnector) GetCertsPath() string {
	return viper.GetString("cert_dir")
}

func (cc ConfigurationConnector) SetCertsPath(certsPath string) {
	viper.Set("cert_dir", certsPath)
}

func (cc ConfigurationConnector) GetConfigPath() string {
	return viper.GetString("config_dir")
}

func (cc ConfigurationConnector) SetConfigPath(configPath string) {
	viper.Set("config_dir", configPath)
}

func (cc ConfigurationConnector) GetVersions() (versions []models.Version) {
	versionsSplit := strings.Split(viper.GetString("versions"), ",")
	for _, versionSplit := range versionsSplit {
		version := strings.Split(versionSplit, ".")
		versionMajor, err := strconv.ParseInt(version[0], 10, 8)
		if err == nil {
			versionMinor, err := strconv.ParseInt(version[1], 10, 8)
			if err == nil {
				versions = append(versions, models.Version{Major: int8(versionMajor), Minor: int8(versionMinor)})
			} else {
				//TODO ERROR
			}
		} else {
			//TODO ERROR
		}
	}
	return
}

func (cc ConfigurationConnector) GetVersionsString() string {
	return viper.GetString("versions")
}

func (cc ConfigurationConnector) SetVersionsString(versions string) {
	viper.Set("versions", versions)
}
