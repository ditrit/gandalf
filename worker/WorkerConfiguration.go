package workergandalf

import(
	 gonfig "github.com/tkanos/gonfig"
)

type WorkerConfiguration struct {
	workerCommandReceiveConnection string
	workerEventReceiveConnection string
	identity string                          string
}

func loadConfiguration(path string) (workerConfiguration WorkerConfiguration, err error) {
	workerConfiguration := WorkerConfiguration{}
	err := gonfig.GetConf(path, &workerConfiguration)
	return
}