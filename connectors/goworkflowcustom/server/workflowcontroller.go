package server

import (
	"fmt"
	"html/template"
	"net/http"

	goclient "github.com/ditrit/gandalf/libraries/goclient"
	models "github.com/ditrit/gandalf/libraries/goclient/models"
)

type WorkflowController struct {
	clientGandalf *goclient.ClientGandalf
	b             *Broker
}

func NewWorkflowController(clientGandalf *goclient.ClientGandalf, b *Broker) *WorkflowController {
	workflowController := new(WorkflowController)
	workflowController.clientGandalf = clientGandalf
	workflowController.b = b

	return workflowController
}

func (fc WorkflowController) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("INDEX")

	tmpl := template.New("name")
	tmpl = template.Must(tmpl.ParseFiles("tmpl/layout.tmpl", "tmpl/content.tmpl"))
	//tmpl.Execute(w, token)
	tmpl.ExecuteTemplate(w, "layout", map[string]interface{}{"Title": "workflow"})
}

func (fc WorkflowController) SendCommand(w http.ResponseWriter, r *http.Request) {

	go func() {
		id := fc.clientGandalf.CreateIteratorEvent()

		fmt.Println("SEND COMMMAND CREATE_FORM")
		payload := `{"Fields":[{"Name":"ID","HtmlType":"TextField","Value":"Id"}]}`
		commandMessageUUID := fc.clientGandalf.SendCommand("Utils.CREATE_FORM", models.NewOptions("", payload))
		formUUID := commandMessageUUID.GetUUID()
		fmt.Println(formUUID)
		for true {
			//event := fc.clientGandalf.WaitReplyByEvent("CREATE_FORM", "STATE", formUUID, id)
			event := fc.clientGandalf.WaitReplyByTopic("CREATE_FORM", formUUID, id)
			fmt.Println(event)
			//TRAITEMENT PAYLOAD
			if event.GetEvent() == "SUCESS" {
				break
			}

			fc.b.messages <- event.GetPayload()
		}
	}()

	tmpl := template.New("name")
	tmpl = template.Must(tmpl.ParseFiles("tmpl/layout.tmpl", "tmpl/content.tmpl"))
	//tmpl.Execute(w, token)
	tmpl.ExecuteTemplate(w, "layout", map[string]interface{}{"Title": "workflow"})

}

func (fc WorkflowController) SendUpdate(w http.ResponseWriter, r *http.Request) {

	go func() {
		//id := fc.clientGandalf.CreateIteratorEvent()

		fmt.Println("SEND COMMMAND ADMIN_UPDATE")
		commandMessageUUIDupdate := fc.clientGandalf.SendAdminCommand("Utils.ADMIN_UPDATE", models.NewOptions("", `""`))
		updateUUID := commandMessageUUIDupdate.GetUUID()
		fmt.Println(updateUUID)
		//event := fc.clientGandalf.WaitReplyByEvent("ADMIN_UPDATE", "SUCCES", updateUUID, id)
		//fmt.Println(event)
	}()

	tmpl := template.New("name")
	tmpl = template.Must(tmpl.ParseFiles("tmpl/layout.tmpl", "tmpl/content.tmpl"))
	//tmpl.Execute(w, token)
	tmpl.ExecuteTemplate(w, "layout", map[string]interface{}{"Title": "workflow"})
}
