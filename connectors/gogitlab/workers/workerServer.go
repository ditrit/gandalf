package workers

import (
	"connectors/gogitlab/server"
	"fmt"
	"libraries/goclient"
	"libraries/goclient/models"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

const (
	server_address = "localhost"
	server_port    = ":3003"
)

type WorkerServer struct {
	address       string
	port          string
	rooturl       string
	router        *chi.Mux
	controller    *server.ServerController
	url           string
	clientGandalf *goclient.ClientGandalf
}

func NewWorkerServer(identity string, connections []string) *WorkerServer {
	workerServer := new(WorkerServer)
	workerServer.address = server_address
	workerServer.port = server_port
	workerServer.rooturl = server_address + server_port
	workerServer.clientGandalf = goclient.NewClientGandalf(identity, "", connections)

	workerServer.url = server.ReturnHashURLS()
	fmt.Println(workerServer.url)

	workerServer.controller = server.NewServerController(workerServer.clientGandalf)
	workerServer.router = server.GetRouterWithUrl(workerServer.url, workerServer.controller)

	return workerServer
}

func NewWorkerServerTEST() *WorkerServer {
	workerServer := new(WorkerServer)
	workerServer.address = server_address
	workerServer.port = server_port
	workerServer.rooturl = server_address + server_port

	workerServer.url = server.ReturnHashURLS()
	fmt.Println(workerServer.url)

	workerServer.controller = server.NewServerController(workerServer.clientGandalf)
	workerServer.router = server.GetRouterWithUrl(workerServer.url, workerServer.controller)

	return workerServer
}

func (ws WorkerServer) SendEventTest() {
	ws.clientGandalf.SendEvent("Gitlab", "HOOK", models.NewOptions("", "toto"))
}

func NewWorkerServerCustom(address, port, identity string, connections []string) *WorkerServer {
	workerServer := new(WorkerServer)
	workerServer.address = address
	workerServer.port = port
	workerServer.rooturl = address + port
	workerServer.clientGandalf = goclient.NewClientGandalf(identity, "", connections)

	workerServer.url = server.ReturnHashURLS()
	fmt.Println(workerServer.url)

	workerServer.controller = server.NewServerController(workerServer.clientGandalf)
	workerServer.router = server.GetRouterWithUrl(workerServer.url, workerServer.controller)

	return workerServer
}

func (ws WorkerServer) Run() {
	// Start the workerServer
	log.Printf("Listening on localhost: %s", ws.port)
	log.Println(http.ListenAndServe(ws.rooturl, ws.router))
}

func (ws WorkerServer) GetUrl() string {
	return "http://" + ws.rooturl + ws.url
}
