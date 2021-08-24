package api

import (
	"net/http"

	"log"

	net "github.com/ditrit/shoset"

	"github.com/ditrit/gandalf/core/aggregator/database"

	"github.com/rs/cors"
)

// ServerAPI :
type ServerAPI struct {
	bindAddress string
	//router      *mux.Router
	handler http.Handler

	databaseConnection *database.DatabaseConnection
	shoset             *net.Shoset
	//gandalfDatabaseClient    *gorm.DB
	//mapTenantDatabaseClients map[string]*gorm.DB
}

// NewServerAPI :
func NewServerAPI(bindAddress string) *ServerAPI {
	serverAPI := new(ServerAPI)
	serverAPI.bindAddress = bindAddress
	//serverAPI.gandalfDatabaseClient = gandalfDatabaseClient
	//serverAPI.mapTenantDatabaseClients = mapTenantDatabaseClients

	//serverAPI.router = NewRouter()
	router := NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // NOT IN PRODUCTION !!!!!
		AllowCredentials: true,
	})

	serverAPI.handler = c.Handler(router)

	return serverAPI
}

// Run :
func (sa ServerAPI) Run() {
	// Start the workerUpload
	log.Println("Listening on localhost: " + sa.bindAddress)
	log.Println(http.ListenAndServe(sa.bindAddress, sa.handler))
}
