package worker

import (
	"github.com/ditrit/gandalf/connectors/go/functions"
	goclient "github.com/ditrit/gandalf/libraries/goclient"
)

//Worker : Worker
type Worker struct {
	major         int64
	minor         int64
	commandes     []string
	clientGandalf *goclient.ClientGandalf

	Start        func() *goclient.ClientGandalf
	SendCommands func(clientGandalf *goclient.ClientGandalf, major, minor int64, commandes []string)
	//Execute      func()
}

//NewWorker : NewWorker
func NewWorker(major, minor int64, commandes []string) *Worker {
	worker := new(Worker)
	worker.major = major
	worker.minor = minor
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
func (w Worker) GetMajor() int8 {
	return w.major
}

//GetVersion : GetVersion
func (w Worker) GetMinor() int8 {
	return w.minor
}

//Run : Run
func (w Worker) Run() {
	w.clientGandalf = w.Start()

	w.SendCommands(w.clientGandalf, w.major, w.minor, w.commandes)

	/* done := make(chan bool)
	//START WORKER ADMIN
	fmt.Println("EXECUTE WORKER")
	w.Execute()
	<-done */
}
