package server

import (
	"fmt"
	"log"
	"net/http"

	goclient "github.com/ditrit/gandalf/libraries/goclient"

	"github.com/go-chi/chi"
)

const (
	server_address = "localhost"
	server_port    = ":3011"
)

type WorkflowServer struct {
	address    string
	port       string
	router     *chi.Mux
	controller *WorkflowController
	Rooturl    string
	Url        urls
	b          *Broker
}

func NewWorkflowServer(clientGandalf *goclient.ClientGandalf) *WorkflowServer {
	formServer := new(WorkflowServer)
	formServer.address = server_address
	formServer.port = server_port
	formServer.Rooturl = server_address + server_port

	formServer.b = &Broker{
		make(map[chan string]bool),
		make(chan (chan string)),
		make(chan (chan string)),
		make(chan string),
	}

	formServer.Url = ReturnURLS()
	fmt.Println("toto0")
	formServer.controller = NewWorkflowController(clientGandalf, formServer.b)
	//formServer.router = GetRouter(formServer.controller)
	fmt.Println("toto2")
	formServer.router = GetRouter(formServer.controller)
	fmt.Println("toto3")

	return formServer
}

func (f WorkflowServer) Run() {

	f.b.Start()
	http.Handle("/events/", f.b)

	// Start the server
	log.Printf("Listening on localhost: %s", f.port)
	log.Println(http.ListenAndServe(f.Rooturl, f.router))
}
