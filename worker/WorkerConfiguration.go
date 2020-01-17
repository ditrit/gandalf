package worker

import (
	gonfig "github.com/tkanos/gonfig"
)

type WorkerConfiguration struct {
	SenderCommandConnection string
	SenderEventConnection   string
	WaiterCommandConnection string
	WaiterEventConnection   string
	Identity                string
}

func LoadConfiguration(path string) (workerConfiguration *WorkerConfiguration, err error) {
	workerConfiguration = new(WorkerConfiguration)
	err = gonfig.GetConf(path, workerConfiguration)
	return
}
