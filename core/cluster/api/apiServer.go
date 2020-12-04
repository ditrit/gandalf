package api

import (
	"log"
	"net/http"

	"github.com/ditrit/gandalf/core/cluster/database"
	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
)

const (
	server_address = "localhost"
	server_port    = ":3010"
)

// ServerAPI :
type ServerAPI struct {
	address         string
	port            string
	rooturl         string
	router          *mux.Router
	gandalfDatabase *gorm.DB
	mapDatabase     map[string]*gorm.DB
}

// NewServerAPI :
func NewServerAPI(databasePath string) *ServerAPI {
	serverAPI := new(ServerAPI)
	serverAPI.address = server_address
	serverAPI.port = server_port
	serverAPI.rooturl = server_address + server_port

	serverAPI.mapDatabase = make(map[string]*gorm.DB)
	gandalfDatabase, _ := database.NewGandalfDatabaseClient(databasePath, "gandalf")
	serverAPI.gandalfDatabase = gandalfDatabase

	serverAPI.router = GetRouter(gandalfDatabase, serverAPI.mapDatabase, databasePath)

	return serverAPI
}

// Run :
func (sa ServerAPI) Run() {
	// Start the workerUpload
	log.Printf("Listening on localhost: %s", sa.port)
	log.Println(http.ListenAndServe(sa.rooturl, sa.router))
}
