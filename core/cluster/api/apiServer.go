package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
)

const (
	server_address = "localhost"
	server_port    = ":8443"
)

// ServerAPI :
type ServerAPI struct {
	address                  string
	port                     string
	rooturl                  string
	router                   *mux.Router
	gandalfDatabaseClient    *gorm.DB
	mapTenantDatabaseClients map[string]*gorm.DB
}

// NewServerAPI :
func NewServerAPI(databasePath, databaseBindAddr string, gandalfDatabaseClient *gorm.DB, mapTenantDatabaseClients map[string]*gorm.DB) *ServerAPI {
	serverAPI := new(ServerAPI)
	serverAPI.address = server_address
	serverAPI.port = server_port
	serverAPI.rooturl = server_address + server_port

	serverAPI.gandalfDatabaseClient = gandalfDatabaseClient
	serverAPI.mapTenantDatabaseClients = mapTenantDatabaseClients

	serverAPI.router = GetRouter(serverAPI.gandalfDatabaseClient, serverAPI.mapTenantDatabaseClients, databasePath, databaseBindAddr)

	return serverAPI
}

// Run :
func (sa ServerAPI) Run() {
	// Start the workerUpload
	log.Printf("Listening on localhost: %s", sa.port)
	log.Println(http.ListenAndServe(sa.rooturl, sa.router))
}
