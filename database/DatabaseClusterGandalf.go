package database

import "fmt"

//DatabaseClusterGandalf :
type DatabaseClusterGandalf struct {
	databaseStopChannel   chan int
	databaseConfiguration *DatabaseClusterConfiguration
	databaseCluster       *DatabaseCluster
}

//NewDatabaseClusterGandalf :
func NewDatabaseClusterGandalf(path string) (databaseClusterGandalf *DatabaseClusterGandalf) {
	databaseClusterGandalf = new(DatabaseClusterGandalf)
	databaseClusterGandalf.databaseStopChannel = make(chan int)

	databaseClusterGandalf.databaseConfiguration, _ = LoadConfiguration(path)

	databaseClusterGandalf.databaseCluster = NewDatabaseCluster(databaseClusterGandalf.databaseConfiguration.DatabaseClusterDirectory, databaseClusterGandalf.databaseConfiguration.DatabaseClusterConnections)

	return
}

//Run :
func (dc DatabaseClusterGandalf) Run() {
	dc.databaseCluster.Run()
	<-dc.databaseStopChannel
	fmt.Println("quit")
}
