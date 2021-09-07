package upload

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ditrit/gandalf/libraries/goclient"

	"github.com/docker/docker/client"

	"github.com/gorilla/mux"
)

// ServerAPI :
type ServerUpload struct {
	router *mux.Router
}

// NewServerAPI :
func NewServerUpload(cli *client.Client, identity, timeout string, connections []string) *ServerUpload {
	serverUpload := new(ServerUpload)
	fmt.Println("toto")
	serverUpload.router = GetRouter(cli, identity, timeout, connections)
	fmt.Println("toto1")

	return serverUpload
}

// Run :
func (sa ServerUpload) Run(context map[string]interface{}, clientGandalf *goclient.ClientGandalf) {
	// Start the workerUpload
	log.Printf("Listening on localhost: %s", "localhost:8080")
	log.Println(http.ListenAndServe(":8080", sa.router))
}
