package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

const (
	server_address = "localhost"
	server_port    = ":3010"
)

type ServerAPI struct {
	address string
	port    string
	rooturl string
	router  *chi.Mux
}

func NewServerAPI(databasePath string) *ServerAPI {
	serverAPI := new(ServerAPI)
	serverAPI.address = server_address
	serverAPI.port = server_port
	serverAPI.rooturl = server_address + server_port

	serverAPI.router = GetRouter(databasePath)

	return serverAPI
}

func (sa ServerAPI) Run() {
	// Start the workerUpload
	log.Printf("Listening on localhost: %s", sa.port)
	log.Println(http.ListenAndServe(sa.rooturl, sa.router))
}
