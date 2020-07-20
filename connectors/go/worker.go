package worker

import (
	"github.com/ditrit/gandalf/connectors/go/functions"
	goclient "github.com/ditrit/gandalf/libraries/goclient"
)

//Worker : Worker
type Worker struct {
	version       int64
	commandes     []string
	clientGandalf *goclient.ClientGandalf

	Start        func() *goclient.ClientGandalf
	SendCommands func(clientGandalf *goclient.ClientGandalf, version int64, commandes []string)
	//Execute      func()
}

//NewWorker : NewWorker
func NewWorker(version int64, commandes []string) *Worker {
	worker := new(Worker)
	worker.version = version
	worker.commandes = commandes
	worker.Start = functions.Start
	worker.SendCommands = functions.SendCommands

	return worker
}

//GetClientGandalf : GetClientGandalf
func (w Worker) GetClientGandalf() *goclient.ClientGandalf {
	return w.clientGandalf
}

//GetVersion : GetVersion
func (w Worker) GetVersion() int64 {
	return w.version
}

//Run : Run
func (w Worker) Run() {
	w.clientGandalf = w.Start()

	w.SendCommands(w.clientGandalf, w.version, w.commandes)

	/* done := make(chan bool)
	//START WORKER ADMIN
	fmt.Println("EXECUTE WORKER")
	w.Execute()
	<-done */
}
