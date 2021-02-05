package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
)

// ServerAPI :
type ServerAPI struct {
	bindAddress              string
	router                   *mux.Router
	gandalfDatabaseClient    *gorm.DB
	mapTenantDatabaseClients map[string]*gorm.DB
}

// NewServerAPI :
func NewServerAPI(bindAddress, databasePath, certsPath, databaseBindAddress string, gandalfDatabaseClient *gorm.DB, mapTenantDatabaseClients map[string]*gorm.DB) *ServerAPI {
	serverAPI := new(ServerAPI)
	serverAPI.bindAddress = bindAddress

	serverAPI.gandalfDatabaseClient = gandalfDatabaseClient
	serverAPI.mapTenantDatabaseClients = mapTenantDatabaseClients

	serverAPI.router = GetRouter(serverAPI.gandalfDatabaseClient, serverAPI.mapTenantDatabaseClients, databasePath, certsPath, databaseBindAddress)

	return serverAPI
}

// Run :
func (sa ServerAPI) Run() {
	// Start the workerUpload
	log.Printf("Listening on localhost: %s", sa.bindAddress)
	log.Println(http.ListenAndServe(sa.bindAddress, sa.router))
}
