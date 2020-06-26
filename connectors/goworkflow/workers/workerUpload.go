package workers

import (
	"connectors/gandalf-core/connectors/goworkflow/upload"

	goclient "github.com/ditrit/gandalf/libraries/goclient"

	"log"
	"net/http"

	"github.com/go-chi/chi"
)

const (
	server_address = "localhost"
	server_port    = ":3004"
)

type WorkerUpload struct {
	address       string
	port          string
	rooturl       string
	router        *chi.Mux
	controller    *upload.UploadController
	url           string
	clientGandalf *goclient.ClientGandalf
}

func NewWorkerUpload(clientGandalf *goclient.ClientGandalf) *WorkerUpload {

	workerUpload := new(WorkerUpload)
	workerUpload.address = server_address
	workerUpload.port = server_port
	workerUpload.rooturl = server_address + server_port
	workerUpload.clientGandalf = clientGandalf
	workerUpload.url = upload.ReturnURLS()

	controllerUrl := "http://" + workerUpload.rooturl + workerUpload.url
	workerUpload.controller = upload.NewUploadController(controllerUrl, workerUpload.clientGandalf)
	workerUpload.router = upload.GetRouterWithUrl(workerUpload.url, workerUpload.controller)

	return workerUpload
}

func (ws WorkerUpload) Run() {
	// Start the workerUpload
	log.Printf("Listening on localhost: %s", ws.port)
	log.Println(http.ListenAndServe(ws.rooturl, ws.router))
}

func (ws WorkerUpload) GetUrl() string {
	return "http://" + ws.rooturl + ws.url
}
