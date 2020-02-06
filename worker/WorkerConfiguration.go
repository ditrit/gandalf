package worker

import (
	gonfig "github.com/tkanos/gonfig"
)

//WorkerConfiguration :
type WorkerConfiguration struct {
	SenderCommandConnection string
	SenderEventConnection   string
	WaiterCommandConnection string
	WaiterEventConnection   string
	Identity                string
}

//LoadConfiguration :
func LoadConfiguration(path string) (workerConfiguration *WorkerConfiguration, err error) {
	workerConfiguration = new(WorkerConfiguration)
	err = gonfig.GetConf(path, workerConfiguration)

	return
}
