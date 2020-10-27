package workers

import (
	"encoding/json"
	"fmt"

	"github.com/ditrit/gandalf/connectors/goutilscustom/form"

	goclient "github.com/ditrit/gandalf/libraries/goclient"
	models "github.com/ditrit/gandalf/libraries/goclient/models"
)

type WorkerForm struct {
	clientGandalf    *goclient.ClientGandalf
	major            int64
	minor            int64
	clientFormServer *form.FormServer
}

func NewWorkerForm(clientGandalf *goclient.ClientGandalf, major, minor int64) *WorkerForm {
	workerForm := new(WorkerForm)
	workerForm.clientGandalf = clientGandalf
	workerForm.major = major
	workerForm.minor = minor

	return workerForm
}

func (r WorkerForm) Run() {
	done := make(chan bool)
	go r.CreateForm()
	<-done
}

func (r WorkerForm) CreateForm() {
	id := r.clientGandalf.CreateIteratorCommand()
	for true {
		command := r.clientGandalf.WaitCommand("CREATE_FORM", id, r.major)

		var formPayload form.FormPayload
		err := json.Unmarshal([]byte(command.GetPayload()), &formPayload)

		if err == nil {
			r.clientFormServer = form.NewFormServer(command.GetUUID(), formPayload, r.clientGandalf)
			go r.clientFormServer.Run()

			fmt.Println("SUCCES")
			r.clientGandalf.SendReply(command.GetCommand(), "SUCCES", command.GetUUID(), models.NewOptions("", "http://"+r.clientFormServer.Rooturl+r.clientFormServer.Url))
		} else {
			fmt.Println("FALSE")
			r.clientGandalf.SendReply(command.GetCommand(), "FAIL", command.GetUUID(), models.NewOptions("", ""))
		}
	}

}
