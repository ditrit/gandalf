package database

import (
	gonfig "github.com/tkanos/gonfig"
)

type DatabaseClusterConfiguration struct {
	DatabaseClusterConnections []string
	DatabaseClusterDirectory   string
}

func LoadConfiguration(path string) (databaseClusterConfiguration *DatabaseClusterConfiguration, err error) {
	databaseClusterConfiguration = new(DatabaseClusterConfiguration)
	err = gonfig.GetConf(path, databaseClusterConfiguration)
	return
}
