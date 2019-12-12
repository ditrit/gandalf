package workergandalf

import(
	 gonfig "github.com/tkanos/gonfig"
)

type WorkerConfiguration struct {
	senderCommandConnection string
	senderEventConnection string
	receiverCommandConnection string
	receiverEventConnection string
	identity string                          string
}

func loadConfiguration(path string) (workerConfiguration WorkerConfiguration, err error) {
	workerConfiguration := WorkerConfiguration{}
	err := gonfig.GetConf(path, &workerConfiguration)
	return
}