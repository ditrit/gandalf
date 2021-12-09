package utils

import (
	"github.com/ditrit/gandalf/core/aggregator/database"
	net "github.com/ditrit/shoset"
)

var (
	Shoset             *net.Shoset
	DatabaseConnection *database.DatabaseConnection
)

func InitAPIGlobals(s *net.Shoset, dbC *database.DatabaseConnection) {
	Shoset = s
	DatabaseConnection = dbC
}
