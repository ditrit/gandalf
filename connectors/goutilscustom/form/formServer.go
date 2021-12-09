package form

import (
	"fmt"
	"log"
	"net/http"

	goclient "github.com/ditrit/gandalf/libraries/goclient"

	"github.com/go-chi/chi"
	goformit "github.com/kirves/go-form-it"
)

const (
	server_address = "localhost"
	server_port    = ":3011"
)

type FormServer struct {
	address    string
	port       string
	router     *chi.Mux
	controller *FormController
	form       *goformit.Form
	Rooturl    string
	Url        string
}

func NewFormServer(uuid string, payload FormPayload, clientGandalf *goclient.ClientGandalf) *FormServer {
	formServer := new(FormServer)
	formServer.address = server_address
	formServer.port = server_port
	formServer.Rooturl = server_address + server_port

	formServer.Url = ReturnHashURLS()
	fmt.Println("toto0")
	formServer.form = CreateFormWithUrl(formServer.Url, uuid, payload.Fields)
	fmt.Println("toto1")
	formServer.controller = NewFormController(formServer.form, clientGandalf)
	//formServer.router = GetRouter(formServer.controller)
	fmt.Println("toto2")
	formServer.router = GetRouterWithUrl(formServer.Url, formServer.controller)
	fmt.Println("toto3")

	return formServer
}

func (f FormServer) Run() {
	// Start the server
	log.Printf("Listening on localhost: %s", f.port)
	log.Println(http.ListenAndServe(f.Rooturl, f.router))
}
