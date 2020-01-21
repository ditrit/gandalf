package database

import "fmt"

type DatabaseClusterGandalf struct {
	databaseStopChannel   chan int
	databaseConfiguration *DatabaseClusterConfiguration
	databaseCluster       *DatabaseCluster
}

func NewDatabaseClusterGandalf(path string) (databaseClusterGandalf *DatabaseClusterGandalf) {
	databaseClusterGandalf = new(DatabaseClusterGandalf)
	databaseClusterGandalf.databaseStopChannel = make(chan int)

	databaseClusterGandalf.databaseConfiguration, _ = LoadConfiguration(path)

	databaseClusterGandalf.databaseCluster = NewDatabaseCluster(databaseClusterGandalf.databaseConfiguration.DatabaseClusterDirectory, databaseClusterGandalf.databaseConfiguration.DatabaseClusterConnections)

	return
}

func (dc DatabaseClusterGandalf) Run() {
	dc.databaseCluster.Run()
	for {
		select {
		case <-dc.databaseStopChannel:
			fmt.Println("quit")
			break
		}
	}
}
