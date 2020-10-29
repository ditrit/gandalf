package worker

import (
	"github.com/ditrit/gandalf/libraries/goclient/models"

	"github.com/ditrit/gandalf/connectors/go/functions"
	goclient "github.com/ditrit/gandalf/libraries/goclient"
	"github.com/ditrit/shoset/msg"
)

//Worker : Worker
type Worker struct {
	major           int64
	minor           int64
	commandes       []string
	clientGandalf   *goclient.ClientGandalf
	CommandesActive map[string]int
	EventsActive    map[string]int
	CommandesFuncs  map[string]func(clientGandalf *goclient.ClientGandalf, major, minor int64, command msg.Command) int
	EventsFuncs     map[string]func(clientGandalf *goclient.ClientGandalf, major, minor int64, event msg.Event) int
	Start           func() *goclient.ClientGandalf
	SendCommands    func(clientGandalf *goclient.ClientGandalf, major, minor int64, commandes []string)
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
func (w Worker) GetMajor() int64 {
	return w.major
}

//GetVersion : GetVersion
func (w Worker) GetMinor() int64 {
	return w.minor
}

//Run : Run
func (w Worker) Run() {
	w.clientGandalf = w.Start()

	w.SendCommands(w.clientGandalf, w.major, w.minor, w.commandes)

	for true {
		for key, function := range w.CommandesFuncs {
			//CREATE ITERATOR
			id := w.clientGandalf.CreateIteratorCommand()
			//WAIT COMMANDS
			go w.waitCommands(id, key, function)
		}
		for key, function := range w.EventsFuncs {
			//CREATE ITERATOR
			id := w.clientGandalf.CreateIteratorEvent()
			//WAIT COMMANDS
			go w.WaitEvents(id, key, function)
		}
	}
	/* done := make(chan bool)
	//START WORKER ADMIN
	fmt.Println("EXECUTE WORKER")
	w.Execute()
	<-done */
}

func (w Worker) waitCommands(id, commandName string, function func(clientGandalf *goclient.ClientGandalf, major, minor int64, command msg.Command) int) {
	command := w.clientGandalf.WaitCommand(commandName, id, w.major)
	w.CommandesActive[commandName]++
	go w.executeCommands(command, function)

}

func (w Worker) executeCommands(command msg.Command, function func(clientGandalf *goclient.ClientGandalf, major, minor int64, command msg.Command) int) {
	result := function(w.clientGandalf, w.major, w.minor, command)
	if result == 0 {
		w.clientGandalf.SendReply(command.GetCommand(), "SUCCES", command.GetUUID(), models.NewOptions("", ""))
	} else {
		w.clientGandalf.SendReply(command.GetCommand(), "FAIL", command.GetUUID(), models.NewOptions("", ""))
	}
	w.CommandesActive[command.GetCommand()]--
}

//TODO REVOIR
func (w Worker) WaitEvents(id, eventName string, function func(clientGandalf *goclient.ClientGandalf, major, minor int64, event msg.Event) int) {
	/* event := w.clientGandalf.WaitEvent(eventName, id, w.major)

	var mailPayload mail.MailPayload
	err := json.Unmarshal([]byte(event.GetPayload()), &mailPayload)

	if err != nil {
		w.EventsActive[eventName]++
		go w.ExecuteEvents(event, function)
	} */
}

//TODO REVOIR
func (w Worker) ExecuteEvents(event msg.Event, function func(clientGandalf *goclient.ClientGandalf, major, minor int64, event msg.Event) int) {
	/* 	result := function(w.clientGandalf, w.major, w.minor)
	   	if result == 0 {
	   		//w.clientGandalf.SendReply(command.GetCommand(), "SUCCES", command.GetUUID(), models.NewOptions("", ""))
	   	} else {
	   		//w.clientGandalf.SendReply(command.GetCommand(), "FAIL", command.GetUUID(), models.NewOptions("", ""))
	   	}
	   	w.EventsActive[event.GetEvent()]-- */
}
