package api

import (
	"net/http"

	"log"

	net "github.com/ditrit/shoset"

	"github.com/ditrit/gandalf/core/aggregator/database"

	"github.com/gorilla/mux"
)

// ServerAPI :
type ServerAPI struct {
	bindAddress        string
	router             *mux.Router
	databaseConnection *database.DatabaseConnection
	shoset             *net.Shoset
	//gandalfDatabaseClient    *gorm.DB
	//mapTenantDatabaseClients map[string]*gorm.DB
}

// NewServerAPI :
func NewServerAPI(bindAddress string, databaseConnection *database.DatabaseConnection, shoset *net.Shoset) *ServerAPI {
	serverAPI := new(ServerAPI)
	serverAPI.bindAddress = bindAddress
	serverAPI.databaseConnection = databaseConnection
	serverAPI.shoset = shoset
	//serverAPI.gandalfDatabaseClient = gandalfDatabaseClient
	//serverAPI.mapTenantDatabaseClients = mapTenantDatabaseClients

	serverAPI.router = GetRouter(serverAPI.databaseConnection, serverAPI.shoset)

	return serverAPI
}

// Run :
func (sa ServerAPI) Run() {
	// Start the workerUpload
	log.Println("Listening on localhost: " + sa.bindAddress)
	log.Println(http.ListenAndServe(sa.bindAddress, sa.router))
}
