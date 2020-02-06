//Package cluster :
//File ClusterConfiguration.go
package cluster

import (
	gonfig "github.com/tkanos/gonfig"
)

//ClusterConfiguration :
type ClusterConfiguration struct {
	ClusterEventSendConnection            string
	ClusterEventReceiveConnection         string
	ClusterEventCaptureConnection         string
	ClusterCommandSendConnection          string
	ClusterCommandReceiveConnection       string
	ClusterCommandCaptureConnection       string
	WorkerCaptureCommandReceiveConnection string
	WorkerCaptureEventReceiveConnection   string
	DatabaseClusterConnections            []string
	Identity                              string
	Topics                                []string
}

//LoadConfiguration :
func LoadConfiguration(path string) (clusterConfiguration *ClusterConfiguration, err error) {
	clusterConfiguration = new(ClusterConfiguration)
	//clusterConfiguration = ClusterConfiguration{}
	err = gonfig.GetConf(path, clusterConfiguration)

	return
}
