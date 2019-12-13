package client

import(
	 gonfig "github.com/tkanos/gonfig"
)

type ClientConfiguration struct {
	senderCommandConnection    	string
	senderEventConnection 		string
	receiverCommandConnection   string
	receiverEventConnection 	string
	identity                    string
}

func loadConfiguration(path string) (clusterConfiguration ClusterConfiguration, err error) {
	clusterConfiguration := ClusterConfiguration{}
	err := gonfig.GetConf(path, &clusterConfiguration)
	return
}