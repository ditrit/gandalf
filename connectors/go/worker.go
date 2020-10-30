package worker

import (
	"encoding/json"
	"gandalf/connectors/goutilscustom/mail"

	"github.com/ditrit/gandalf/libraries/goclient/models"

	"github.com/ditrit/gandalf/connectors/go/functions"
	gomodels "github.com/ditrit/gandalf/connectors/go/models"
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
	EventsFuncs     map[gomodels.TopicEvent]func(clientGandalf *goclient.ClientGandalf, major, minor int64, event msg.Event) int
	Start           func() *goclient.ClientGandalf
	SendCommands    func(clientGandalf *goclient.ClientGandalf, major, minor int64, commandes []string)
	//Execute      func()
}

//NewWorker : NewWorker
func NewWorker(major, minor int64) *Worker {
	worker := new(Worker)
	worker.major = major
	worker.minor = minor
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

//GetVersion : GetVersion
func (w Worker) RegisterCommandsFuncs(command string, function func(clientGandalf *goclient.ClientGandalf, major, minor int64, command msg.Command) int) {
	w.commandes = append(w.commandes, command)
	w.CommandesFuncs[command] = function
}

//GetVersion : GetVersion
func (w Worker) RegisterEventsFuncs(topicevent gomodels.TopicEvent, function func(clientGandalf *goclient.ClientGandalf, major, minor int64, event msg.Event) int) {
	w.EventsFuncs[topicevent] = function
}

//Run : Run
func (w Worker) Run() {
	w.clientGandalf = w.Start()

	w.SendCommands(w.clientGandalf, w.major, w.minor, w.commandes)

	//TODO REVOIR CONDITION SORTIE
	for true {
		for key, function := range w.CommandesFuncs {
			id := w.clientGandalf.CreateIteratorCommand()

			go w.waitCommands(id, key, function)
		}
		for key, function := range w.EventsFuncs {
			id := w.clientGandalf.CreateIteratorEvent()

			go w.WaitEvents(id, key, function)
		}
	}
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

func (w Worker) WaitEvents(id string, topicEvent gomodels.TopicEvent, function func(clientGandalf *goclient.ClientGandalf, major, minor int64, event msg.Event) int) {
	event := w.clientGandalf.WaitEvent(topicEvent.Topic, topicEvent.Event, id)

	var mailPayload mail.MailPayload
	err := json.Unmarshal([]byte(event.GetPayload()), &mailPayload)

	if err != nil {
		w.EventsActive[topicEvent.Event]++
		go w.ExecuteEvents(event, function)
	}
}

func (w Worker) ExecuteEvents(event msg.Event, function func(clientGandalf *goclient.ClientGandalf, major, minor int64, event msg.Event) int) {
	result := function(w.clientGandalf, w.major, w.minor, event)
	if result == 0 {
		w.clientGandalf.SendReply(event.GetEvent(), "SUCCES", event.GetUUID(), models.NewOptions("", ""))
	} else {
		w.clientGandalf.SendReply(event.GetEvent(), "FAIL", event.GetUUID(), models.NewOptions("", ""))
	}
	w.EventsActive[event.GetEvent()]--
}
