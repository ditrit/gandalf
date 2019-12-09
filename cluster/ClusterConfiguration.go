package clustergandalf

import(
	 gonfig "github.com/tkanos/gonfig"
)

type ClusterConfiguration struct {
	clusterEventSendConnection    string
	clusterEventReceiveConnection string
	clusterEventCaptureConnection    string
	clusterCommandSendConnection    string
	clusterCommandReceiveConnection string
	clusterCommandCaptureConnection string
	identity                        string
}

func loadConfiguration(path string) (clusterConfiguration ClusterConfiguration, err error) {
	clusterConfiguration := ClusterConfiguration{}
	err := gonfig.GetConf(path, &clusterConfiguration)
	return
}