package worker

import (
	"flag"
	"strings"
	"time"

	"github.com/ditrit/gandalf/connectors/go/functions"
	gomodels "github.com/ditrit/gandalf/connectors/go/models"
	goclient "github.com/ditrit/gandalf/libraries/goclient"
	"github.com/ditrit/shoset/msg"
)

//Worker : Worker
type Worker struct {
	major             int64
	minor             int64
	identity          string
	timeout           string
	connections       []string
	clientGandalf     *goclient.ClientGandalf
	OngoingTreatments *gomodels.OngoingTreatments
	WorkerState       *gomodels.WorkerState
	CommandsFuncs     map[string]func(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int
	EventsFuncs       map[gomodels.TopicEvent]func(clientGandalf *goclient.ClientGandalf, major int64, event msg.Event) int
	//Start             func() *goclient.ClientGandalf
	Stop         func(clientGandalf *goclient.ClientGandalf, major, minor int64, workerState *gomodels.WorkerState)
	SendCommands func(clientGandalf *goclient.ClientGandalf, major, minor int64, commandes []string) bool
	//Execute      func()
}

//NewWorker : NewWorker
func NewWorker(major, minor int64) *Worker {
	worker := new(Worker)
	worker.major = major
	worker.minor = minor
	worker.CommandsFuncs = make(map[string]func(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int)
	worker.EventsFuncs = make(map[gomodels.TopicEvent]func(clientGandalf *goclient.ClientGandalf, major int64, event msg.Event) int)
	worker.OngoingTreatments = gomodels.NewOngoingTreatments()
	worker.WorkerState = gomodels.NewWorkerState()
	//worker.Start = functions.Start
	worker.Stop = functions.Stop
	worker.SendCommands = functions.SendCommands

	return worker
}

//GetClientGandalf : GetClientGandalf
func (w Worker) GetClientGandalf() *goclient.ClientGandalf {
	return w.clientGandalf
}

//GetMajor : GetMajor
func (w Worker) GetMajor() int64 {
	return w.major
}

//GetMinor : GetMinor
func (w Worker) GetMinor() int64 {
	return w.minor
}

//GetIdentity : GetIdentity
func (w Worker) GetIdentity() string {
	return w.identity
}

//GetTimeout : GetTimeout
func (w Worker) GetTimeout() string {
	return w.timeout
}

//GetConnections : GetConnections
func (w Worker) GetConnections() []string {
	return w.connections
}

//RegisterCommandsFuncs : RegisterCommandsFuncs
func (w Worker) RegisterCommandsFuncs(command string, function func(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int) {
	w.CommandsFuncs[command] = function
}

//RegisterEventsFuncs : RegisterEventsFuncs
func (w Worker) RegisterEventsFuncs(topicevent gomodels.TopicEvent, function func(clientGandalf *goclient.ClientGandalf, major int64, event msg.Event) int) {
	w.EventsFuncs[topicevent] = function
}

func (w *Worker) Start() {
	flag.Parse()
	args := flag.Args()
	w.identity = args[0]
	w.timeout = args[1]
	w.connections = strings.Split(args[2], ",")
	w.clientGandalf = goclient.NewClientGandalf(w.identity, w.timeout, w.connections)
	//return goclient.NewClientGandalf(args[0], args[1], strings.Split(args[2], ","))
}

//Run : Run
func (w Worker) Run() {
	w.Start()

	keys := make([]string, 0, len(w.CommandsFuncs))
	for k := range w.CommandsFuncs {
		keys = append(keys, k)
	}

	valid := w.SendCommands(w.clientGandalf, w.major, w.minor, keys)

	if valid {
		go w.Stop(w.clientGandalf, w.major, w.minor, w.WorkerState)

		for key, function := range w.CommandsFuncs {
			id := w.clientGandalf.CreateIteratorCommand()

			go w.waitCommands(id, key, function)
		}
		for key, function := range w.EventsFuncs {
			id := w.clientGandalf.CreateIteratorEvent()

			go w.waitEvents(id, key, function)
		}
		//TODO REVOIR CONDITION SORTIE
		for w.WorkerState.GetState() == 0 {
		}
		for w.OngoingTreatments.GetIndex() > 0 {
			time.Sleep(2 * time.Second)
		}
	} else {
		//SEND EVENT INVALID CONFIGURATION
	}
}

func (w Worker) waitCommands(id, commandName string, function func(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int) {

	for true {
		command := w.clientGandalf.WaitCommand(commandName, id, w.major)

		if w.WorkerState.GetState() == 0 {
			go w.executeCommands(command, function)
		} else {
			break
		}

	}
	for w.OngoingTreatments.GetIndex() > 0 {
		time.Sleep(2 * time.Second)
	}
}

func (w Worker) executeCommands(command msg.Command, function func(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int) {
	w.OngoingTreatments.IncrementOngoingTreatments()
	result := function(w.clientGandalf, w.major, command)
	if result == 0 {
		w.clientGandalf.SendReply(command.GetCommand(), "SUCCES", command.GetUUID(), map[string]string{})
	} else {
		w.clientGandalf.SendReply(command.GetCommand(), "FAIL", command.GetUUID(), map[string]string{})
	}
	w.OngoingTreatments.DecrementOngoingTreatments()
}

func (w Worker) waitEvents(id string, topicEvent gomodels.TopicEvent, function func(clientGandalf *goclient.ClientGandalf, major int64, event msg.Event) int) {
	for true {
		event := w.clientGandalf.WaitEvent(topicEvent.Topic, topicEvent.Event, id)
		if w.WorkerState.GetState() == 0 {
			go w.executeEvents(event, function)
		} else {
			break
		}

	}
	for w.OngoingTreatments.GetIndex() > 0 {
		time.Sleep(2 * time.Second)
	}
}

func (w Worker) executeEvents(event msg.Event, function func(clientGandalf *goclient.ClientGandalf, major int64, event msg.Event) int) {
	w.OngoingTreatments.IncrementOngoingTreatments()
	result := function(w.clientGandalf, w.major, event)
	if result == 0 {
		w.clientGandalf.SendReply(event.GetEvent(), "SUCCES", event.GetUUID(), map[string]string{})
	} else {
		w.clientGandalf.SendReply(event.GetEvent(), "FAIL", event.GetUUID(), map[string]string{})
	}
	w.OngoingTreatments.DecrementOngoingTreatments()
}
