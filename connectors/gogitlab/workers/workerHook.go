package workers

import (
	"connectors/gogitlab/client"
	"connectors/gogitlab/hook"
	"encoding/json"
	"fmt"
	"libraries/goclient"
	"libraries/goclient/models"
)

type WorkerHook struct {
	ClientGandalf *goclient.ClientGandalf
	ClientGilab   *client.ClientGitlab
	WorkerServer  *WorkerServer
}

func NewWorkerHook(identity, token string, connections []string) *WorkerHook {
	workerHook := new(WorkerHook)
	workerHook.ClientGandalf = goclient.NewClientGandalf(identity, "", connections)
	workerHook.ClientGilab = client.NewClient(token)

	return workerHook
}

func NewWorkerHookBasic(identity, login, password, rooturl string, connections []string) *WorkerHook {
	workerHook := new(WorkerHook)
	workerHook.ClientGandalf = goclient.NewClientGandalf(identity, "", connections)
	workerHook.ClientGilab = client.NewClientBasic(login, password, rooturl)

	return workerHook
}

func NewWorkerHookTEST(login, password, rooturl string) *WorkerHook {
	workerHook := new(WorkerHook)
	fmt.Println(client.NewClientBasic(login, password, rooturl))
	workerHook.ClientGilab = client.NewClientBasic(login, password, rooturl)

	return workerHook
}

func (r WorkerHook) Run() {
	done := make(chan bool)
	go r.AddHook()
	go r.DeleteHook()
	go r.ListHook()
	go r.TestHook()
	<-done
}

func (r WorkerHook) AddHook() {
	id := r.ClientGandalf.CreateIteratorCommand()
	for true {
		fmt.Println(id)
		command := r.ClientGandalf.WaitCommand("ADD_HOOK", id)
		fmt.Println(command)

		var addHookPayload hook.AddHookPayload
		err := json.Unmarshal([]byte(command.GetPayload()), &addHookPayload)

		if err == nil {
			result := hook.AddHook(r.ClientGilab, addHookPayload.Url, addHookPayload.PushEvents, addHookPayload.TagPushEvents, addHookPayload.MergeRequestsEvents, addHookPayload.RepositoryUpdateEvents, addHookPayload.EnableSSLVerification)
			if result {
				r.ClientGandalf.SendEvent(command.GetUUID(), "SUCCES", models.NewOptions("", "test"))
			} else {
				r.ClientGandalf.SendEvent(command.GetUUID(), "FAIL", models.NewOptions("", ""))
			}
		}
	}
}

func (r WorkerHook) AddHookTEST(urlserver, urlhook string, push, tagpush, merge, update, enable bool) {

	result := hook.AddHook(r.ClientGilab, urlhook, &push, &tagpush, &merge, &update, &enable)

	fmt.Println("result")
	fmt.Println(result)

}

func (r WorkerHook) DeleteHook() {
	id := r.ClientGandalf.CreateIteratorCommand()
	for true {
		fmt.Println(id)
		command := r.ClientGandalf.WaitCommand("DELETE_HOOK", id)
		fmt.Println(command)

		var deleteHookPayload hook.DeleteHookPayload
		err := json.Unmarshal([]byte(command.GetPayload()), &deleteHookPayload)

		if err == nil {
			result := hook.DeleteHook(r.ClientGilab, deleteHookPayload.Id)
			if result {
				r.ClientGandalf.SendEvent(command.GetUUID(), "SUCCES", models.NewOptions("", "test"))
			} else {
				r.ClientGandalf.SendEvent(command.GetUUID(), "FAIL", models.NewOptions("", ""))
			}
		}
	}
}

func (r WorkerHook) ListHook() {
	id := r.ClientGandalf.CreateIteratorCommand()
	for true {
		fmt.Println(id)
		command := r.ClientGandalf.WaitCommand("LIST_HOOK", id)
		fmt.Println(command)

		result := hook.ListHook(r.ClientGilab)
		if result {
			r.ClientGandalf.SendEvent(command.GetUUID(), "SUCCES", models.NewOptions("", "test"))
		} else {
			r.ClientGandalf.SendEvent(command.GetUUID(), "FAIL", models.NewOptions("", ""))
		}
	}
}

func (r WorkerHook) TestHook() {
	id := r.ClientGandalf.CreateIteratorCommand()
	for true {
		fmt.Println(id)
		command := r.ClientGandalf.WaitCommand("TEST_HOOK", id)
		fmt.Println(command)

		var testHookPayload hook.TestHookPayload
		err := json.Unmarshal([]byte(command.GetPayload()), &testHookPayload)

		if err == nil {
			result := hook.TestHook(r.ClientGilab, testHookPayload.Id)
			if result {
				r.ClientGandalf.SendEvent(command.GetUUID(), "SUCCES", models.NewOptions("", "test"))
			} else {
				r.ClientGandalf.SendEvent(command.GetUUID(), "FAIL", models.NewOptions("", ""))
			}
		}
	}
}
