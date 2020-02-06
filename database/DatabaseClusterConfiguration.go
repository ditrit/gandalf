//Package database :
//File DatabaseClusterConfiguration.go
package database

import (
	gonfig "github.com/tkanos/gonfig"
)

//DatabaseClusterConfiguration :
type DatabaseClusterConfiguration struct {
	DatabaseClusterConnections []string
	DatabaseClusterDirectory   string
}

//LoadConfiguration :
func LoadConfiguration(path string) (databaseClusterConfiguration *DatabaseClusterConfiguration, err error) {
	databaseClusterConfiguration = new(DatabaseClusterConfiguration)
	err = gonfig.GetConf(path, databaseClusterConfiguration)

	return
}
