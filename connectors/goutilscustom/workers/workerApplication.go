package workers

import (
	"log"
	"net/http"

	"github.com/ditrit/gandalf/connectors/goutilscustom/app"

	goclient "github.com/ditrit/gandalf/libraries/goclient"

	"github.com/go-chi/chi"
)

const (
	server_address = "localhost"
	server_port    = ":3010"
)

type WorkerApplication struct {
	address       string
	port          string
	rooturl       string
	router        *chi.Mux
	controller    *app.AppController
	url           string
	clientGandalf *goclient.ClientGandalf
	version       int64
}

func NewWorkerApplication(clientGandalf *goclient.ClientGandalf, version int64) *WorkerApplication {
	workerApplication := new(WorkerApplication)
	workerApplication.address = server_address
	workerApplication.port = server_port
	workerApplication.rooturl = server_address + server_port
	workerApplication.clientGandalf = clientGandalf
	workerApplication.version = version
	workerApplication.url = app.ReturnURLS()

	controllerUrl := "http://" + workerApplication.rooturl + workerApplication.url
	workerApplication.controller = app.NewAppController(controllerUrl, workerApplication.clientGandalf)
	workerApplication.router = app.GetRouterWithUrl(workerApplication.url, workerApplication.controller)

	return workerApplication
}

func (wa WorkerApplication) Run() {
	// Start the workerUpload
	log.Printf("Listening on localhost: %s", wa.port)
	log.Println(http.ListenAndServe(wa.rooturl, wa.router))
}

func (wa WorkerApplication) GetUrl() string {
	return "http://" + wa.rooturl + wa.url
}
