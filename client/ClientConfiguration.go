package client

import(
	 gonfig "github.com/tkanos/gonfig"
)

type ClientConfiguration struct {
	SenderCommandConnection    	string
	SenderEventConnection 		string
	ReceiverCommandConnection   string
	ReceiverEventConnection 	string
	Identity                    string
}

func LoadConfiguration(path string) (clusterConfiguration *ClusterConfiguration, err error) {
	clusterConfiguration = new(ClusterConfiguration)
	err := gonfig.GetConf(path, &clusterConfiguration)
	return
}