package api

import (
	"log"
	"net/http"

	"github.com/ditrit/gandalf/core/cluster/database"

	"github.com/gorilla/mux"
)

// ServerAPI :
type ServerAPI struct {
	bindAddress        string
	router             *mux.Router
	databaseConnection *database.DatabaseConnection
	//gandalfDatabaseClient    *gorm.DB
	//mapTenantDatabaseClients map[string]*gorm.DB
}

// NewServerAPI :
func NewServerAPI(bindAddress string, databaseConnection *database.DatabaseConnection) *ServerAPI {
	serverAPI := new(ServerAPI)
	serverAPI.bindAddress = bindAddress
	serverAPI.databaseConnection = databaseConnection
	//serverAPI.gandalfDatabaseClient = gandalfDatabaseClient
	//serverAPI.mapTenantDatabaseClients = mapTenantDatabaseClients

	serverAPI.router = GetRouter(serverAPI.databaseConnection)

	return serverAPI
}

// Run :
func (sa ServerAPI) Run() {
	// Start the workerUpload
	log.Printf("Listening on localhost: %s", sa.bindAddress)
	log.Println(http.ListenAndServe(sa.bindAddress, sa.router))
}
