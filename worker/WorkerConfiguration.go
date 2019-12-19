package worker

import (
	gonfig "github.com/tkanos/gonfig"
)

type WorkerConfiguration struct {
	SenderCommandConnection   string
	SenderEventConnection     string
	ReceiverCommandConnection string
	ReceiverEventConnection   string
	Identity                  string
}

func LoadConfiguration(path string) (workerConfiguration *WorkerConfiguration, err error) {
	workerConfiguration = WorkerConfiguration{}
	err = gonfig.GetConf(path, &workerConfiguration)
	return
}
